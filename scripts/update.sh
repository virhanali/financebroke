#!/bin/bash

# FinanceBroke Update Script
set -e

echo "🔄 Updating FinanceBroke..."

# Pull latest changes
echo "📥 Pulling latest changes..."
git pull origin master

# Deploy with Docker Compose
echo "🚀 Deploying updated version..."
docker-compose up -d --build

# Show status
echo "📊 Checking deployment status..."
docker-compose ps

echo "✅ Update completed!"
echo "📊 Check logs with: docker-compose logs -f"