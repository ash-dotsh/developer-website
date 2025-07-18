#!/bin/bash

echo "Setting up Ash Web Server..."

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running. Please start Docker Desktop."
    exit 1
fi

# Create go.sum file
cd app
go mod tidy
cd ..

# Make scripts executable
chmod +x scripts/*.sh

# Build the project
./scripts/build.sh

echo ""
echo "Setup completed successfully!"
echo ""
echo "To start the server, run: ./scripts/run.sh"
echo "To stop the server, run: ./scripts/stop.sh"
