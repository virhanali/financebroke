#!/bin/bash

# Finance App Kubernetes Deployment Script
set -e

echo "🚀 Deploying Finance App to Kubernetes..."

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl not found. Please run setup-vps.sh first"
    exit 1
fi

# Change to project directory
cd "$(dirname "$0")/.."

# Build and push Docker image
echo "🐳 Building Docker image..."
docker build -f Dockerfile.k8s -t finance-app:latest .

# Load image into k3s
echo "📦 Loading image into k3s..."
docker save finance-app:latest | sudo k3s ctr images import -

# Apply Kubernetes manifests
echo "☸️ Applying Kubernetes manifests..."

# Create namespace
kubectl apply -f k8s/namespace.yaml

# Create secrets (you need to update this file first!)
echo "⚠️ Make sure to update k8s/backend/secret.yaml with your actual secrets!"
kubectl apply -f k8s/backend/secret.yaml

# Apply ConfigMap
kubectl apply -f k8s/backend/configmap.yaml

# Deploy database
kubectl apply -f k8s/database/pvc.yaml
kubectl apply -f k8s/database/deployment.yaml

# Wait for database to be ready
echo "⏳ Waiting for database to be ready..."
kubectl wait --for=condition=ready pod -l app=finance-app,component=database -n finance-app --timeout=300s

# Deploy backend
kubectl apply -f k8s/backend/deployment.yaml

# Wait for backend to be ready
echo "⏳ Waiting for backend to be ready..."
kubectl wait --for=condition=ready pod -l app=finance-app,component=backend -n finance-app --timeout=300s

# Apply ingress
kubectl apply -f k8s/ingress/ingress.yaml

# Get status
echo "📊 Deployment status:"
kubectl get pods -n finance-app
kubectl get services -n finance-app
kubectl get ingress -n finance-app

echo "✅ Deployment completed!"
echo "🌐 Your app should be available at your configured domain"
echo "📊 Check logs with: kubectl logs -f deployment/finance-app-backend -n finance-app"