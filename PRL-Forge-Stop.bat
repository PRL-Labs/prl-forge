@echo off
title Stop PRL Forge

echo Stopping PRL Forge...

taskkill /F /FI "WINDOWTITLE eq Forge Backend*"
taskkill /F /FI "WINDOWTITLE eq Website Backend*"
taskkill /F /FI "WINDOWTITLE eq Frontend*"

echo.
echo Done!
timeout /t 2 >nul