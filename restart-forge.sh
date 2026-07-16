#!/bin/bash

echo "=================================="
echo " PRL Forge Restart"
echo "=================================="

echo "[1/4] Stopping Forge..."
systemctl stop forge

echo "[2/4] Building..."
cd /opt/prl-forge || exit 1

go build -tags zkpow -o forge ./cmd/forge
if [ $? -ne 0 ]; then
    echo "❌ Build failed!"
    exit 1
fi

echo "[3/4] Starting Forge..."
systemctl start forge

sleep 2

echo "[4/4] Status:"
systemctl --no-pager --full status forge

echo
echo "✅ Forge restarted successfully."
echo
echo "⚠️  Now restart SRBMiner from HiveOS."
echo "=================================="
