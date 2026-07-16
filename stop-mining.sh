#!/bin/bash

echo "=============================="
echo "Stopping PRL Forge..."
systemctl stop forge

echo "Stopping SRBMiner..."
pkill -f SRBMiner-MULTI
pkill -f SRBMiner

sleep 2

echo
echo "Running processes:"
pgrep -a forge
pgrep -a SRBMiner

echo
echo "? Everything stopped."
echo "=============================="