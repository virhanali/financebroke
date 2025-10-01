#!/bin/bash

# FinanceBroke Docker Deployment Script
set -e

echo "🚀 Deploying FinanceBroke with Docker..."

# Check if .env file exists
if [ ! -f ".env" ]; then
    echo "❌ .env file not found. Please copy .env.example to .env and configure it."
    exit 1
fi

# Load environment variables
source .env

echo "🐳 Building and starting containers..."

# Stop existing containers
echo "🛑 Stopping existing containers..."
docker-compose down

# Build new images
echo "🏗️ Building application image..."
docker-compose build

# Start containers
echo "🚀 Starting containers..."
docker-compose up -d

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
sleep 10

# Check if containers are running
echo "🔍 Checking container status..."
docker-compose ps

# Show logs
echo "📊 Showing recent logs..."
docker-compose logs --tail=50

echo "✅ Deployment completed!"
echo ""
echo "🌐 Your app is available at:"
echo "   - HTTP: http://financebroke.virhanali.com"
echo "   - API: http://financebroke.virhanali.com/api/v1"
echo ""
echo "📊 Check logs with: docker-compose logs -f"
echo "🛑 Stop with: docker-compose down"
echo "🔄 Update with: git pull && docker-compose up -d --build"