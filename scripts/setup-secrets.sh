#!/bin/bash

# Script to generate base64 encoded secrets for Kubernetes
set -e

echo "🔐 Setting up Kubernetes secrets..."

echo "Please enter the following values:"

# Get database password
read -s -p "Database password: " DB_PASSWORD
echo
DB_PASSWORD_B64=$(echo -n "$DB_PASSWORD" | base64 -w 0)

# Get JWT secret
read -s -p "JWT secret (long random string): " JWT_SECRET
echo
JWT_SECRET_B64=$(echo -n "$JWT_SECRET" | base64 -w 0)

# Get Telegram bot token
read -s -p "Telegram bot token: " TELEGRAM_TOKEN
echo
TELEGRAM_TOKEN_B64=$(echo -n "$TELEGRAM_TOKEN" | base64 -w 0)

# Get SMTP credentials
read -p "SMTP username: " SMTP_USERNAME
SMTP_USERNAME_B64=$(echo -n "$SMTP_USERNAME" | base64 -w 0)

read -s -p "SMTP password: " SMTP_PASSWORD
echo
SMTP_PASSWORD_B64=$(echo -n "$SMTP_PASSWORD" | base64 -w 0)

# Update secret.yaml file
echo "📝 Updating k8s/backend/secret.yaml..."
sed -i "s/<base64-encoded-db-password>/$DB_PASSWORD_B64/" k8s/backend/secret.yaml
sed -i "s/<base64-encoded-jwt-secret>/$JWT_SECRET_B64/" k8s/backend/secret.yaml
sed -i "s/<base64-encoded-telegram-token>/$TELEGRAM_TOKEN_B64/" k8s/backend/secret.yaml
sed -i "s/<base64-encoded-smtp-username>/$SMTP_USERNAME_B64/" k8s/backend/secret.yaml
sed -i "s/<base64-encoded-smtp-password>/$SMTP_PASSWORD_B64/" k8s/backend/secret.yaml

echo "✅ Secrets configured successfully!"
echo "📝 Secret file updated: k8s/backend/secret.yaml"