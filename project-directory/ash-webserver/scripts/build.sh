#!/bin/bash

echo "Building Ash Web Server..."

# Generate SSL certificates
echo "Generating SSL certificates..."
./scripts/generate-ssl.sh

# Build and start services
echo "Building Docker images..."
docker-compose build --no-cache

echo "Build completed successfully!"
