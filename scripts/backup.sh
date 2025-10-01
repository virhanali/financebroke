#!/bin/bash

# FinanceBroke Backup Script
set -e

echo "💾 Creating backup of FinanceBroke..."

BACKUP_DIR="/opt/backups/financebroke"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="financebroke_backup_$DATE.tar.gz"

# Create backup directory
mkdir -p $BACKUP_DIR

# Backup database
echo "🗄️ Backing up database..."
docker exec financebroke-postgres pg_dump -U postgres financebroke_db > $BACKUP_DIR/database_$DATE.sql

# Backup environment file
echo "⚙️ Backing up configuration..."
cp .env $BACKUP_DIR/env_$DATE

# Backup SSL certificates
if [ -d "ssl" ]; then
    echo "🔒 Backing up SSL certificates..."
    cp -r ssl $BACKUP_DIR/ssl_$DATE
fi

# Create compressed backup
echo "📦 Creating compressed backup..."
tar -czf $BACKUP_DIR/$BACKUP_FILE \
    -C $BACKUP_DIR \
    database_$DATE.sql \
    env_$DATE \
    ssl_$Date

# Cleanup individual files
rm $BACKUP_DIR/database_$DATE.sql
rm $BACKUP_DIR/env_$DATE
rm -rf $BACKUP_DIR/ssl_$DATE

echo "✅ Backup completed: $BACKUP_DIR/$BACKUP_FILE"

# Keep only last 7 backups
echo "🧹 Cleaning up old backups..."
ls -t $BACKUP_DIR/financebroke_backup_*.tar.gz | tail -n +8 | xargs -r rm

echo "✅ Backup process completed!"