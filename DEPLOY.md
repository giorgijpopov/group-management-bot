# Deployment Guide

## Quick Deploy

To rebuild and deploy the bot to your VPS:

```bash
./deploy.sh
```

## What the script does:

1. 🔨 Builds binary for Linux x86_64
2. 🛑 Stops the bot service on VPS
3. 📤 Uploads binary via SCP
4. 🔐 Sets executable permissions
5. 🚀 Starts the bot service
6. ✅ Shows service status and recent logs

## Requirements

- SSH access configured as `vultr_vpn` in `~/.ssh/config`
- Go installed locally
- Bot service configured at `/etc/systemd/system/bot.service`
- Binary located at `/root/broBot/group-management-bot`

## Manual deployment

If you prefer to deploy manually:

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o group-management-bot

# Stop service
ssh vultr_vpn "systemctl stop bot"

# Upload
scp group-management-bot vultr_vpn:/root/broBot/group-management-bot

# Set permissions and start
ssh vultr_vpn "chmod +x /root/broBot/group-management-bot && systemctl start bot"

# Check status
ssh vultr_vpn "systemctl status bot"
```

## Check logs

```bash
ssh vultr_vpn "journalctl -u bot -f"
```

## Environment Variables

The bot requires these environment variables (set in the systemd service):

- `TBOT_SECRET` - Telegram Bot API token
- `TBOT_DADDY_ID` - Telegram user ID for error notifications
