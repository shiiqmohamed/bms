moha@jaale:~/projects/bms$ cat > Makefile << 'EOF'
.PHONY: build run clean db-setup help

APP_NAME=bms

help:
	@echo "Available commands:"
	@echo "  make build     - Build the application"
	@echo "  make run       - Run the application"
	@echo "  make clean     - Clean build files"
	@echo "  make db-setup  - Setup database"

build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) cmd/main.go

run:
	@echo "Running $(APP_NAME)..."
	go run cmd/main.go

clean:
	@echo "Cleaning..."
	rm -rf bin/

db-setup:
	@echo "Setting up database..."
	chmod +x scripts/db-setup.sh
	./scripts/db-setup.sh

dev:
	@echo "Starting development server..."
	go run cmd/main.go
EOF