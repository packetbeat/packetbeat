#!/usr/bin/env python3
import platform
import re
import subprocess
import sys
import tarfile

PATH = 'x-pack/agentbeat/build/distributions'


def log(msg):
    sys.stdout.write(f'{msg}\n')
    sys.stdout.flush()


def log_err(msg):
    sys.stderr.write(f'{msg}\n')
    sys.stderr.flush()


def get_os() -> str:
    return platform.system().lower()


def get_arch() -> str:
    arch = platform.machine().lower()

    if arch == 'amd64':
        return 'x86_64'
    else:
        if get_os() == 'darwin':
            return 'aarch64'
        else:
            return arch


def get_artifact_extension(agent_os) -> str:
    if agent_os == 'windows':
        return 'zip'
    else:
        return 'tar.gz'


def get_artifact_pattern() -> str:
    agent_os = get_os()
    agent_arch = get_arch()
    extension = get_artifact_extension(agent_os)
    print('Artifact params: ' + agent_os + ' ' + agent_arch + ' ' + extension)
    return f'{PATH}/agentbeat-*-{agent_os}-{agent_arch}.{extension}'


def download_agentbeat(pattern, path) -> str:
    try:
        subprocess.run(
            ['buildkite-agent', 'artifact', 'download', pattern, '.',
             '--step', 'agentbeat-package-linux'],
            check=True, stdout=sys.stdout, stderr=sys.stderr, text=True)

    except subprocess.CalledProcessError:
        exit(1)

    return get_filename(path)


def get_filename(path) -> str:
    try:
        out = subprocess.run(
            ['ls', '-p', path],
            check=True, capture_output=True, text=True)
        print("--- ls -p: " + out.stdout)
        return out.stdout.strip()
    except subprocess.CalledProcessError:
        exit(1)


def extract_agentbeat(filename):
    filepath = PATH + '/' + filename

    if filepath.endswith('.zip'):
        unzip_agentbeat(filepath)
    else:
        untar_agentbeat(filepath)


def unzip_agentbeat(filepath):
    try:
        subprocess.run(
            ['unzip', '-qq', filepath],
            check=True, stdout=sys.stdout, stderr=sys.stderr, text=True)
    except subprocess.CalledProcessError as e:
        log_err(e)
        exit(1)


def untar_agentbeat(filepath):
    try:
        subprocess.run(
            ['tar', '-xvf', filepath],
            check=True, stdout=sys.stdout, stderr=sys.stderr, text=True)
    except subprocess.CalledProcessError as e:
        log_err(e)
        exit(1)


def get_path_to_executable(filepath) -> str:
    pattern = r'(.*)(?=\.zip|.tar\.gz)'
    match = re.match(pattern, filepath)
    if match:
        path = f'../../{match.group(1)}/agentbeat'
        return path
    else:
        log_err("No agentbeat executable found")
        exit(1)

artifact_pattern = get_artifact_pattern()
archive = download_agentbeat(artifact_pattern, PATH)
extract_agentbeat(archive)
log(get_path_to_executable(archive))
