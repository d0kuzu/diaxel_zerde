@echo off
echo ============================================
echo   Diaxel Zerde - Deployment Script
echo ============================================
echo.

REM Ensure microservices-net network exists (required before any service starts)
echo [1/3] Checking Docker network "microservices-net"...
docker network inspect microservices-net >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo      Network not found. Creating "microservices-net"...
    docker network create microservices-net
    if %ERRORLEVEL% NEQ 0 (
        echo [ERROR] Failed to create Docker network. Is Docker running?
        pause
        exit /b 1
    )
    echo      Network created successfully.
) else (
    echo      Network already exists. OK.
)
echo.

echo [2/3] Choose deployment option:
echo       1) Production  - all services (postgres, database, auth, ai, api-gateway)
echo       2) Stop all    - bring everything down
echo.
set /p choice="Enter your choice (1 or 2): "

if "%choice%"=="1" (
    echo.
    echo [3/3] Deploying all services from root...
    echo       Each service uses its own Dockerfile and .env
    echo.
    docker-compose down
    docker-compose up --build -d
    if %ERRORLEVEL% NEQ 0 (
        echo [ERROR] Deployment failed. Check the logs above.
        pause
        exit /b 1
    )
    echo.
    echo ============================================
    echo   All services deployed successfully!
    echo ============================================
    echo   Database Service : localhost:50051
    echo   Auth Service     : http://localhost:8082
    echo   AI Service       : http://localhost:8081
    echo   API Gateway      : http://localhost:8085
    echo ============================================
    goto status
)

if "%choice%"=="2" (
    echo.
    echo Stopping all services...
    docker-compose down
    echo Done. All services stopped.
    goto end
)

echo Invalid choice. Please run the script again.
exit /b 1

:status
echo.
echo Checking service status (waiting 10s for containers to start)...
timeout /t 10 /nobreak >nul
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

:end
echo.
pause
