#!/bin/bash

# Deployment script for Diaxel services

echo "Choose deployment option:"
echo "1) Deploy all services except AI (production)"
echo "2) Deploy all services including AI (development)"
read -p "Enter your choice (1 or 2): " choice

case $choice in
    1)
        echo "Deploying services without AI service..."
        docker-compose down
        docker-compose up --build -d
        echo "Services deployed without AI service!"
        echo "API Gateway: http://localhost:8081"
        echo "Frontend: http://localhost:3000"
        echo "Auth Service: http://localhost:8083"
        echo "Telegram Service: http://localhost:8084"
        echo "Database Service: localhost:50051"
        ;;
    2)
        echo "Deploying all services including AI service..."
        docker-compose -f docker-compose.dev.yml down
        docker-compose -f docker-compose.dev.yml up --build -d
        echo "All services deployed including AI service!"
        echo "API Gateway: http://localhost:8081"
        echo "Frontend: http://localhost:3000"
        echo "AI Service: http://localhost:8082"
        echo "Auth Service: http://localhost:8083"
        echo "Telegram Service: http://localhost:8084"
        echo "Database Service: localhost:50051"
        ;;
    *)
        echo "Invalid choice. Please run the script again."
        exit 1
        ;;
esac

echo "Checking service status..."
sleep 10
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
