# Install protoc on Windows (manual step)
# This script downloads protoc win64 zip and adds to PATH for current session.
# Run as Administrator if you want to add permanently.

Write-Host "Downloading protoc win64..." -ForegroundColor Cyan
Invoke-WebRequest -Uri "https://github.com/protocolbuffers/protobuf/releases/download/v28.3/protoc-28.3-win64.zip" -OutFile "protoc-win64.zip"

Write-Host "Extracting to C:\protoc..." -ForegroundColor Cyan
if (!(Test-Path "C:\protoc")) { New-Item -ItemType Directory -Path "C:\protoc" }
Expand-Archive -Path "protoc-win64.zip" -DestinationPath "C:\protoc" -Force

Write-Host "Adding C:\protoc\bin to PATH for this session..." -ForegroundColor Cyan
$env:PATH += ";C:\protoc\bin"

Write-Host "Cleaning up zip..." -ForegroundColor Cyan
Remove-Item "protoc-win64.zip"

Write-Host "protoc installed. Verify: protoc --version" -ForegroundColor Green
protoc --version
