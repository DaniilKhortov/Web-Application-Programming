# Скрипт крос-компіляції для Windows
$ErrorActionPreference = "Stop"

$APP_NAME = "queue-app"
$VERSION = $env:VERSION
if (-not $VERSION) { $VERSION = "1.0.0" }

$BUILD_TIME = Get-Date -Format "yyyy-MM-dd_HH:mm:ss"
try {
    $GIT_COMMIT = git rev-parse --short HEAD
} catch {
    $GIT_COMMIT = "nogit"
}

# Створення папки dist
if (-not (Test-Path -Path "dist")) {
    New-Item -ItemType Directory -Path "dist" | Out-Null
}

function Build($GOOS, $GOARCH, $EXT) {
    Write-Host "Building for $GOOS/$GOARCH"
    $env:CGO_ENABLED = "0"
    go build -ldflags "-s -w -X main.Version=$VERSION -X main.BuildTime=$BUILD_TIME -X main.GitCommit=$GIT_COMMIT" `
        -trimpath -o "dist\$APP_NAME-$GOOS-$GOARCH$EXT" ./cmd/web
}

# Крос-компіляція
Build "windows" "amd64" ".exe"
Build "linux" "amd64" ""
Build "darwin" "amd64" ""

# Вивід результату
Get-ChildItem dist | Select-Object Name, @{Name="Size(MB)";Expression={[math]::Round($_.Length/1MB,2)}}
Write-Host "Build finished. Files are in ./dist"
