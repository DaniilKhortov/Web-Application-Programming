$ErrorActionPreference = "Stop"

$APP_NAME = "queue-app"
$SERVER = "user@your-server.com"
$DEPLOY_PATH = "/opt/$APP_NAME"
$SERVICE_NAME = "$APP_NAME.service"

Write-Host "Deploying $APP_NAME to $SERVER..."

# Виклик build
.\scripts\build.ps1

# Зупинка сервісу на сервері
ssh $SERVER "sudo systemctl stop $SERVICE_NAME || true"

# Копіюємо бінарник
scp .\dist\$APP_NAME-linux-amd64 $SERVER:$DEPLOY_PATH/$APP_NAME

# Копіюємо статичні файли
scp -r .\web $SERVER:$DEPLOY_PATH/

# Встановлення прав
ssh $SERVER "sudo chmod +x $DEPLOY_PATH/$APP_NAME"

# Старт сервісу
ssh $SERVER "sudo systemctl start $SERVICE_NAME"
ssh $SERVER "sudo systemctl status $SERVICE_NAME"

Write-Host "Deployment complete!"
