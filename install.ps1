#Requires -Version 5.1
$ErrorActionPreference = "Stop"

$Url = "https://github.com/Akhil373/typr/releases/download/v1.0.0/typr.exe"
$BinName = "typr.exe"
$InstallDir = Join-Path $HOME "bin"
$BinPath = Join-Path $InstallDir $BinName

Write-Host "Installing $BinName to $InstallDir..."

New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

$TempFile = Join-Path $env:TEMP "$BinName.tmp"
try {
    Invoke-WebRequest -Uri $Url -OutFile $TempFile -UseBasicParsing
    Move-Item -Force -Path $TempFile -Destination $BinPath
} finally {
    if (Test-Path $TempFile) { Remove-Item $TempFile -Force }
}

Write-Host "Installed to $BinPath"

# Check PATH
$UserPath = [Environment]::GetEnvironmentVariable('Path', 'User')

if ($UserPath -split ";" -notcontains $InstallDir) {
    Write-Host "Adding $InstallDir to user PATH..."
    [Environment]::SetEnvironmentVariable('Path', "$InstallDir;$UserPath", 'User')
    $env:Path = "$InstallDir;$env:Path"
}
Write-Host "You're all set! Run: typr"
