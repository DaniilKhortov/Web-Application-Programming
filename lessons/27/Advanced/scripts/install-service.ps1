$ErrorActionPreference = "Stop"

$APP_NAME = "queue-app"
$DEPLOY_PATH = "/opt/$APP_NAME"
$SERVICE_FILE = "/etc/systemd/system/$APP_NAME.service"

sudo tee $SERVICE_FILE > $null <<EOL
[Unit]
Description=$APP_NAME service
After=network.target

[Service]
Type=simple
ExecStart=$DEPLOY_PATH/$APP_NAME
Restart=on-failure
User=$(whoami)
WorkingDirectory=$DEPLOY_PATH

[Install]
WantedBy=multi-user.target
EOL

sudo systemctl daemon-reload
sudo systemctl enable $APP_NAME
sudo systemctl start $APP_NAME
sudo systemctl status $APP_NAME
