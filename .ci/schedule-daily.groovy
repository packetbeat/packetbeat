@Library('apm@current') _

pipeline {
  agent none
  environment {
    NOTIFY_TO = credentials('notify-to')
    PIPELINE_LOG_LEVEL = 'INFO'
  }
  options {
    timeout(time: 1, unit: 'HOURS')
    buildDiscarder(logRotator(numToKeepStr: '20', artifactNumToKeepStr: '20'))
    timestamps()
    ansiColor('xterm')
    disableResume()
    durabilityHint('PERFORMANCE_OPTIMIZED')
  }
  triggers {
    cron('H H(2-3) * * 1-5')
  }
  stages {
    stage('Nighly beats builds') {
      steps {
        runBuild(quietPeriod: 0, branch: 'master')
        runBuild(quietPeriod: 2000, branch: '8.<minor>')
        runBuild(quietPeriod: 4000, branch: '7.<next-minor>')
      }
    }
  }
  post {
    cleanup {
      notifyBuildResult(prComment: false)
    }
  }
}

def runBuild(Map args = [:]) {
  def branch = args.branch
  // special macro to look for the latest minor version
  if (branch.contains('8.<minor>')) {
    branch = bumpUtils.getMajorMinor(bumpUtils.getCurrentMinorReleaseFor8())
  }
  if (branch.contains('7.<minor>')) {
    branch = bumpUtils.getMajorMinor(bumpUtils.getCurrentMinorReleaseFor7())
  }
  if (branch.contains('7.<next-minor>')) {
    branch = bumpUtils.getMajorMinor(bumpUtils.getNextMinorReleaseFor7())
  }
  build(quietPeriod: args.quietPeriod, job: "Beats/beats/${branch}", parameters: [booleanParam(name: 'macosTest', value: true)], wait: false, propagate: false)
}
