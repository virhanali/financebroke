#!/bin/bash

# FinanceBroke SSL Setup Script with Let's Encrypt
set -e

echo "🔒 Setting up SSL certificate for FinanceBroke..."

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root or with sudo"
    exit 1
fi

# Load environment variables
APP_DIR="/opt/financebroke"
cd $APP_DIR
source .env

DOMAIN="financebroke.virhanali.com"

echo "📋 Setting up SSL for domain: $DOMAIN"

# Get SSL certificate
echo "🔑 Obtaining SSL certificate..."
certbot certonly --standalone \
    --non-interactive \
    --agree-tos \
    --email "admin@${DOMAIN}" \
    -d "$DOMAIN"

# Copy certificates to app directory
echo "📁 Copying SSL certificates..."
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem $APP_DIR/ssl/cert.pem
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem $APP_DIR/ssl/key.pem
chown $SUDO_USER:$SUDO_USER $APP_DIR/ssl/*.pem

# Setup automatic renewal
echo "⏰ Setting up automatic renewal..."
(crontab -l 2>/dev/null; echo "0 12 * * * certbot renew --quiet && docker restart financebroke-nginx") | crontab -

# Start SSL enabled containers
echo "🚀 Starting with SSL enabled..."
cd $APP_DIR
sudo -u $SUDO_USER docker-compose --profile ssl up -d

echo "✅ SSL setup completed!"
echo ""
echo "🌐 Your app is available with HTTPS at:"
echo "   - HTTPS: https://$DOMAIN"
echo ""
echo "🔧 Certificate details:"
echo "   - Certificate: $APP_DIR/ssl/cert.pem"
echo "   - Private Key: $APP_DIR/ssl/key.pem"
echo ""
echo "📅 Auto-renewal: Scheduled daily at 12:00 PM"