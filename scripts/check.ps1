Write-Host ""
Write-Host "========================================"
Write-Host "        PRL Forge Health Check"
Write-Host "========================================"
Write-Host ""

Write-Host "Git Status..."
git status

Write-Host ""
Write-Host "Building..."
go build ./...

if ($LASTEXITCODE -ne 0) {
    Write-Host ""
    Write-Host "BUILD FAILED"
    exit
}

Write-Host ""
Write-Host "Running gofmt..."
gofmt -w .

Write-Host ""
Write-Host "Done."