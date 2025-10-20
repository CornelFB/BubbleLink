# docker_run.ps1
<#
.SYNOPSIS
Run Docker containers for frontend or backend.

.PARAMETER f
Run the frontend container (frontend:latest)

.PARAMETER b
Run the backend container (backend:latest)

.EXAMPLE
.\docker_run.ps1 -f
#>

param (
    [switch]$f,
    [switch]$b
)

function Show-Usage {
    Write-Host "Usage: .\docker_run.ps1 [-f | -b]"
    Write-Host "  -f    Run the frontend container (frontend:latest)"
    Write-Host "  -b    Run the backend container (backend:latest)"
    exit 1
}

# If no parameter is provided
if (-not ($f -or $b)) {
    Show-Usage
}

if ($f) {
    Write-Host "Running frontend container..."
    docker run -it --rm -p 8081:80 frontend:latest
}

if ($b) {
    Write-Host "Running backend container..."
    docker run -it --rm -p 3000:3000 backend:latest
}
