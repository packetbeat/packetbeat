function Exec {
    [CmdletBinding()]
    param(
        [Parameter(Mandatory = $true)]
        [scriptblock]$cmd,
        [string]$errorMessage = ($msgs.error_bad_command -f $cmd)
    )

    try {
        $global:lastexitcode = 0
        & $cmd
        if ($lastexitcode -ne 0) {
            throw $errorMessage
        }
    }
    catch [Exception] {
        throw $_
    }
}

# Setup Go.
$env:GOPATH = $env:WORKSPACE
$env:PATH = "$env:GOPATH\bin;C:\tools\mingw64\bin;$env:PATH"
& gvm --format=powershell $(Get-Content .go-version) | Invoke-Expression

# Write cached magefile binaries to workspace to ensure
# each run starts from a clean slate.
$env:MAGEFILE_CACHE = "$env:WORKSPACE\.magefile"

# Configure testing parameters.
$env:TEST_COVERAGE = "true"
$env:RACE_DETECTOR = "true"

# Install mage from vendor.
exec { go install github.com/elastic/beats/vendor/github.com/magefile/mage } "mage install FAILURE"

if (Test-Path "$env:beat\magefile.go") {
    cd "$env:beat"
} else {
    echo "$env:beat\magefile.go does not exist"
    New-Item -ItemType directory -Path build | Out-Null
    New-Item -Name build\TEST-empty.out -ItemType File | Out-Null
    exit
}

if (Test-Path "build") { Remove-Item -Recurse -Force build }
New-Item -ItemType directory -Path build\coverage | Out-Null
New-Item -ItemType directory -Path build\system-tests | Out-Null
New-Item -ItemType directory -Path build\system-tests\run | Out-Null

echo "Building fields.yml"
exec { mage fields } "mage fields FAILURE"

echo "Building $env:beat"
exec { mage build } "Build FAILURE"

echo "Unit testing $env:beat"
exec { mage goTestUnit } "mage goTestUnit FAILURE"

echo "System testing $env:beat"
# Get a CSV list of package names.
$packages = $(go list ./... | select-string -Pattern "/vendor/" -NotMatch | select-string -Pattern "/scripts/cmd/" -NotMatch)
$packages = ($packages|group|Select -ExpandProperty Name) -join ","
exec { go test -race -c -cover -covermode=atomic -coverpkg $packages } "go test -race -cover FAILURE"

if (Test-Path "tests\system") {
    # Installation of python 3 to be removed once it is installed in the base image
    echo "Installing python 3"
    Invoke-WebRequest -Uri "https://www.python.org/ftp/python/3.7.6/python-3.7.6-embed-amd64.zip" | Expand-Archive -DestinationPath "build/python" -Force
    $currentDir = (Get-Item -Path ".\").FullName
    $env:PATH = "$currentDir\build\python;$env:PATH"

    echo "Running python tests"
    exec { mage pythonUnitTest } "System test FAILURE"
}
