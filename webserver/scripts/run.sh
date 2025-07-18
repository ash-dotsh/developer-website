#!/bin/bash

echo "Starting Ash Web Server..."

# Start services
docker-compose up -d

echo "Services started successfully!"
echo ""
echo "Web server is running at:"
echo "  HTTP:  http://localhost"
echo "  HTTPS: https://localhost"
echo ""
echo "API endpoints:"
echo "  Health: https://localhost/health"
echo "  Users:  https://localhost/api/users"
echo ""
echo "To stop services, run: docker-compose down"
