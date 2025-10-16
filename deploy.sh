#!/bin/bash
set -e

echo "🔨 Building binary for Linux..."
GOOS=linux GOARCH=amd64 go build -o group-management-bot

echo "📦 Binary size: $(ls -lh group-management-bot | awk '{print $5}')"

echo "🛑 Stopping bot service..."
ssh vultr_vpn "systemctl stop bot"

echo "📤 Uploading binary to VPS..."
scp group-management-bot vultr_vpn:/root/broBot/group-management-bot

echo "🔐 Setting permissions..."
ssh vultr_vpn "chmod +x /root/broBot/group-management-bot"

echo "🚀 Starting bot service..."
ssh vultr_vpn "systemctl start bot"

echo "⏳ Waiting for bot to start..."
sleep 2

echo "✅ Deployment complete! Service status:"
ssh vultr_vpn "systemctl status bot --no-pager -l"

echo ""
echo "📝 Recent logs:"
ssh vultr_vpn "journalctl -u bot -n 10 --no-pager"
