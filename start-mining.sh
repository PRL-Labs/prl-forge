#!/bin/bash

echo "=============================="
echo "Starting PRL Forge..."

systemctl start forge

sleep 3

echo "Starting SRBMiner..."

cd /opt/srbminer || exit 1
nohup ./SRBMiner-MULTI --algorithm pearlhash --config Config.ini >/dev/null 2>&1 &

sleep 3

echo
echo "Running processes:"
pgrep -a forge
pgrep -a SRBMiner

echo
echo "Dashboard:"
echo "http://pool.prlforge.com"

echo
echo "? Mining started."
echo "=============================="