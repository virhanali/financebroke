#!/bin/bash

# FinanceBroke VPS Setup Script for Docker
set -e

echo "🚀 Setting up FinanceBroke on VPS with Docker..."

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root or with sudo"
    exit 1
fi

# Update system
echo "📦 Updating system packages..."
apt update && apt upgrade -y

# Install required packages
echo "🔧 Installing required packages..."
apt install -y curl wget git git htop

# Install Docker
echo "🐳 Installing Docker..."
if ! command -v docker &> /dev/null; then
    curl -fsSL https://get.docker.com -o get-docker.sh
    sh get-docker.sh
    systemctl enable docker
    systemctl start docker
    usermod -aG docker $SUDO_USER
    rm get-docker.sh
fi

# Install Docker Compose
echo "🔧 Installing Docker Compose..."
if ! command -v docker-compose &> /dev/null; then
    curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
fi

# Install Certbot for SSL
echo "🔒 Installing Certbot for SSL..."
apt install -y certbot

# Setup firewall
echo "🔥 Setting up firewall..."
ufw --force reset
ufw default deny incoming
ufw default allow outgoing
ufw allow ssh
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

# Create application directory
echo "📁 Creating application directory..."
APP_DIR="/opt/financebroke"
mkdir -p $APP_DIR
chown $SUDO_USER:$SUDO_USER $APP_DIR

# Clone or update repository
if [ -d "$APP_DIR/.git" ]; then
    echo "📁 Updating existing repository..."
    cd $APP_DIR
    sudo -u $SUDO_USER git pull origin master
else
    echo "📁 Cloning repository..."
    sudo -u $SUDO_USER git clone https://github.com/virhanali/financebroke.git $APP_DIR
    cd $APP_DIR
fi

# Create environment file
echo "⚙️ Setting up environment configuration..."
if [ ! -f "$APP_DIR/.env" ]; then
    sudo -u $SUDO_USER cp $APP_DIR/.env.example $APP_DIR/.env
    echo "📝 Please edit $APP_DIR/.env with your configuration"
fi

# Create SSL directory
echo "🔒 Creating SSL directory..."
mkdir -p $APP_DIR/ssl
chown $SUDO_USER:$SUDO_USER $APP_DIR/ssl

echo "✅ VPS setup completed!"
echo ""
echo "🔧 Next steps:"
echo "1. Edit $APP_DIR/.env with your configuration"
echo "2. Setup SSL: ./scripts/setup-ssl.sh"
echo "3. Deploy: ./scripts/deploy.sh"
echo ""
echo "📱 Your app will be available at:"
echo "   - HTTP: http://financebroke.virhanali.com"
echo "   - HTTPS: https://financebroke.virhanali.com (after SSL setup)"

# Log out and back in for docker group to take effect
echo ""
echo "⚠️  Log out and back in for Docker group to take effect"