APP_NAME=queue-app

.PHONY: build run test clean

# Використовуємо powershell -ExecutionPolicy Bypass для запуску скрипта
build:
	@echo "Building application..."
	powershell -ExecutionPolicy Bypass -File scripts/build.ps1

run:
	@echo "Running locally..."
	go run ./cmd/web

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning..."
	rm -rf dist
