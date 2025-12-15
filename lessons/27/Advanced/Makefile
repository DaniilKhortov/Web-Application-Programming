APP_NAME=queue-app

.PHONY: build run test clean deploy install-service

build:
	@echo "Building application..."
	powershell -ExecutionPolicy Bypass -File scripts\build.ps1

run:
	@echo "Running locally..."
	go run ./cmd/web

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning..."
	rm -rf dist

deploy:
	@echo "Deploying..."
	powershell -ExecutionPolicy Bypass -File scripts\deploy.ps1

install-service:
	@echo "Installing systemd service..."
	powershell -ExecutionPolicy Bypass -File scripts\install-service.ps1
