#!/bin/bash

echo "🚀 Setting up BMS Project..."

# Check Go version
echo "📦 Checking Go version..."
go version

# Clean up old modules
echo "🧹 Cleaning up..."
rm -f go.sum
rm -rf vendor/

# Download dependencies
echo "📥 Downloading dependencies..."
go mod download

# Verify dependencies
echo "✅ Verifying dependencies..."
go mod verify

# Create .env file if not exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file..."
    cp .env.example .env
fi

# Make scripts executable
echo "⚙️  Setting up scripts..."
chmod +x scripts/*.sh

echo ""
echo "🎉 Setup completed!"
echo ""
echo "Next steps:"
echo "1. Edit .env file with your database credentials"
echo "2. Run: ./scripts/db-setup.sh"
echo "3. Run: go run cmd/main.go"
echo ""
echo "📡 API will be available at: http://localhost:8080"