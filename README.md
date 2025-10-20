# BubbleLink
Scaneaza coduri QR de la diverse atractii turistice/ momumente din orasul tau dar nu numai. Imparataseste ti experienta cu altii utilizatori si ajuta i si pe altii sa descopere noul lor loc favorit.

#**Aplicatia foloseste urmatoarele tehnologii**:
- **Project Strucure** : Fantastic Coffee
- **Database**: SQLite
- **Backend**: GO
- **API**: OpenAPI with swagger documentation
- **Frontend**: Vue JS
- **Containerization** : Docker 
- **Compatibility**: Premade docker scripts for both Linux and Windows

## Instructiuni de instalare

### Instalare Go

Distributii Linux bazate pe Arch : `sudo pacman -S go`
Distributii Linux bazate pe Debian: `sudo apt install golang-go`

Windows: Go to https://go.dev/dl/. Download and run the .msi installer.

### Instalare VueJs si yarn

Distributii Linux bazate pe Arch : `sudo pacman -Syu nodejs npm yarn`
Distributii Linux bazate pe Debian: `sudo npm install -g corepack` , `corepack enable` , `corepack prepare yarn@stable --activate`

Windows: Go to https://classic.yarnpkg.com/lang/en/docs/install/

### Docker

Pentru Docker se pot folosi instructiunile dupa site https://docs.docker.com/engine/install/

## Build & Run

Se poate rula local folosind npm si go, dar recomand utilizarea Dockerfile-urilor.

### Linux 

Build: ` ./docker_build.sh -b` pentru backend,  `./docker_build.sh -f` pentru frontend. 
Run: ` ./docker_run.sh -b` pentru backend,  `./docker_run.sh -f` pentru frontend. 

### Windows

Build: ` ./docker_build.ps1 -b` pentru backend,  `./docker_build.ps1 -f` pentru frontend. 
Run: ` ./docker_run.ps1 -b` pentru backend,  `./docker_run.ps1 -f` pentru frontend. 
