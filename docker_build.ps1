# docker_build.ps1
<#
.SYNOPSIS
Build Docker containers for frontend or backend.

.PARAMETER f
Build the frontend container (frontend:latest)

.PARAMETER b
Build the backend container (backend:latest)

.EXAMPLE
.\docker_build.ps1 -f
#>

param (
    [switch]$f,
    [switch]$b
)

function Show-Usage {
    Write-Host "Usage: .\docker_build.ps1 [-f | -b]"
    Write-Host "  -f    Build the frontend container (frontend:latest)"
    Write-Host "  -b    Build the backend container (backend:latest)"
    exit 1
}

# If no parameter is provided
if (-not ($f -or $b)) {
    Show-Usage
}

if ($f) {
    Write-Host "Building frontend container..."
    docker build -f Dockerfile.frontend -t frontend:latest .
}

if ($b) {
    Write-Host "Building backend container..."
    docker build -f Dockerfile.backend -t backend:latest .
}
