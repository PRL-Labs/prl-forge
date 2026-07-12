@echo off
title PRL Forge Launcher

echo ==========================================
echo         STARTING PRL FORGE
echo ==========================================
echo.

REM Forge Backend (8080)
start "Forge Backend" cmd /k "cd /d D:\Projects\prl-forge && go run cmd\server\main.go"

timeout /t 2 >nul

REM Website Backend (8081)
start "Website Backend" cmd /k "cd /d D:\Projects\prl-forge-web\backend && go run cmd\server\main.go"

timeout /t 2 >nul

REM Frontend (5173)
start "Frontend" cmd /k "cd /d D:\Projects\prl-forge-web\frontend && npm run dev"

echo.
echo ==========================================
echo     Everything is starting...
echo ==========================================
pause