#!/bin/bash

echo "============================================"
echo "  Diaxel Zerde - Deployment Script"
echo "============================================"
echo ""

# Ensure microservices-net network exists (required before any service starts)
echo "[1/3] Checking Docker network 'microservices-net'..."
if ! docker network inspect microservices-net > /dev/null 2>&1; then
    echo "     Network not found. Creating 'microservices-net'..."
    docker network create microservices-net
    if [ $? -ne 0 ]; then
        echo "[ERROR] Failed to create Docker network. Is Docker running?"
        exit 1
    fi
    echo "     Network created successfully."
else
    echo "     Network already exists. OK."
fi
echo ""

echo "[2/3] Choose deployment option:"
echo "      1) Production  - all services (postgres, database, auth, ai, api-gateway)"
echo "      2) Stop all    - bring everything down"
echo ""
read -p "Enter your choice (1 or 2): " choice

case $choice in
    1)
        echo ""
        echo "[3/3] Deploying all services from root..."
        echo "      Each service uses its own Dockerfile and .env"
        echo ""
        docker compose down
        docker compose up -d --build
        if [ $? -ne 0 ]; then
            echo "[ERROR] Deployment failed. Check the logs above."
            exit 1
        fi
        echo ""
        echo "============================================"
        echo "  All services deployed successfully!"
        echo "============================================"
        echo "  Database Service : localhost:50051"
        echo "  Auth Service     : http://localhost:8082"
        echo "  AI Service       : http://localhost:8081"
        echo "  API Gateway      : http://localhost:8085"
        echo "============================================"
        echo ""
        echo "Checking service status (waiting 10s for containers to start)..."
        sleep 10
        docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
        ;;
    2)
        echo ""
        echo "Stopping all services..."
        docker compose down
        echo "Done. All services stopped."
        ;;
    *)
        echo "Invalid choice. Please run the script again."
        exit 1
        ;;
esac
