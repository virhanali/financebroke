#!/bin/bash

# Cleanup FinanceBroke App from Kubernetes
set -e

echo "🧹 Cleaning up FinanceBroke App from Kubernetes..."

# Delete all resources in reverse order
echo "🗑️ Deleting ingress..."
kubectl delete -f k8s/ingress/ingress.yaml --ignore-not-found=true

echo "🗑️ Deleting backend..."
kubectl delete -f k8s/backend/deployment.yaml --ignore-not-found=true

echo "🗑️ Deleting database..."
kubectl delete -f k8s/database/deployment.yaml --ignore-not-found=true
kubectl delete -f k8s/database/pvc.yaml --ignore-not-found=true

echo "🗑️ Deleting config and secrets..."
kubectl delete -f k8s/backend/configmap.yaml --ignore-not-found=true
kubectl delete -f k8s/backend/secret.yaml --ignore-not-found=true

echo "🗑️ Deleting namespace..."
kubectl delete -f k8s/namespace.yaml --ignore-not-found=true

echo "✅ Cleanup completed!"