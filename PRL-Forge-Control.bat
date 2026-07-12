@echo off
title PRL Forge Control Center
color 0A

:MENU
cls
echo.
echo ======================================================
echo                 PRL FORGE CONTROL CENTER
echo ======================================================
echo.
echo    [1] Start Everything
echo    [2] Stop Everything
echo    [3] Restart Everything
echo    [4] Open Website
echo    [5] Open Forge API
echo    [6] Exit
echo.
echo ======================================================
echo.

set /p choice=Select option:

if "%choice%"=="1" goto START
if "%choice%"=="2" goto STOP
if "%choice%"=="3" goto RESTART
if "%choice%"=="4" goto WEBSITE
if "%choice%"=="5" goto API
if "%choice%"=="6" exit

goto MENU

:START
cls
echo Starting PRL Forge...
echo.

echo [1/3] Forge Backend...
start "Forge Backend" cmd /k "cd /d D:\Projects\prl-forge && go run cmd\server\main.go"

timeout /t 2 >nul

echo [2/3] Website Backend...
start "Website Backend" cmd /k "cd /d D:\Projects\prl-forge-web\backend && go run cmd\server\main.go"

timeout /t 2 >nul

echo [3/3] Frontend...
start "Frontend" cmd /k "cd /d D:\Projects\prl-forge-web\frontend && npm run dev"

timeout /t 4 >nul

start http://localhost:5173

echo.
echo ==========================================
echo        PRL FORGE STARTED!
echo ==========================================
pause
goto MENU

:STOP
cls
echo Stopping PRL Forge...

taskkill /F /FI "WINDOWTITLE eq Forge Backend*" >nul 2>&1
taskkill /F /FI "WINDOWTITLE eq Website Backend*" >nul 2>&1
taskkill /F /FI "WINDOWTITLE eq Frontend*" >nul 2>&1

echo.
echo All services stopped.
pause
goto MENU

:RESTART
cls
echo Restarting...

taskkill /F /FI "WINDOWTITLE eq Forge Backend*" >nul 2>&1
taskkill /F /FI "WINDOWTITLE eq Website Backend*" >nul 2>&1
taskkill /F /FI "WINDOWTITLE eq Frontend*" >nul 2>&1

timeout /t 2 >nul

start "Forge Backend" cmd /k "cd /d D:\Projects\prl-forge && go run cmd\server\main.go"

timeout /t 2 >nul

start "Website Backend" cmd /k "cd /d D:\Projects\prl-forge-web\backend && go run cmd\server\main.go"

timeout /t 2 >nul

start "Frontend" cmd /k "cd /d D:\Projects\prl-forge-web\frontend && npm run dev"

timeout /t 4 >nul

start http://localhost:5173

echo.
echo Restart complete.
pause
goto MENU

:WEBSITE
start http://localhost:5173
goto MENU

:API
start http://localhost:8080/api/v1/dashboard
goto MENU