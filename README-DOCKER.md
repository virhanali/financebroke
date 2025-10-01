# FinanceBroke Docker Deployment Guide

## 🚀 Quick Start

### 1. Setup VPS
```bash
# SSH into your VPS
ssh root@43.133.140.5

# Clone and run setup
git clone https://github.com/virhanali/financebroke.git
cd financebroke
sudo ./scripts/setup-vps.sh
```

### 2. Configure Environment
```bash
# Copy and edit environment file
cp .env.example .env
nano .env

# Fill in your configuration:
DB_PASSWORD=your_secure_db_password
JWT_SECRET=your_very_long_and_secure_jwt_secret
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_email_app_password
```

### 3. Deploy Application
```bash
./scripts/deploy.sh
```

### 4. Setup SSL (Optional but Recommended)
```bash
sudo ./scripts/setup-ssl.sh
```

## 📋 Prerequisites

- Ubuntu 20.04+ (VPS: 2GB RAM, 2 CPU cores)
- Domain name pointing to VPS IP
- Docker & Docker Compose (auto-installed by setup script)

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│      Nginx      │    │   Frontend      │    │    Backend      │
│  (Port 80/443)  │───▶│   (React)       │───▶│      Go         │
│                 │    │                 │    │   (Port 8080)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                        │
                                                        ▼
                                                ┌─────────────────┐
                                                │   PostgreSQL    │
                                                │   (Port 5432)   │
                                                └─────────────────┘
```

## 🔧 Configuration

### Environment Variables
Create `.env` file from `.env.example`:

```bash
# Database
DB_PASSWORD=your_secure_db_password

# JWT
JWT_SECRET=your_very_long_and_secure_jwt_secret

# Telegram
TELEGRAM_BOT_TOKEN=your_telegram_bot_token

# Email
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_email_app_password
FROM_EMAIL=noreply@financebroke.com
```

## 🌐 Access Points

After deployment:
- **Frontend:** http://financebroke.virhanali.com
- **API:** http://financebroke.virhanali.com/api/v1
- **HTTPS (with SSL):** https://financebroke.virhanali.com

## 🔄 Management Commands

### Update Application
```bash
./scripts/update.sh
```

### View Logs
```bash
docker-compose logs -f
```

### Stop Application
```bash
docker-compose down
```

### Backup Data
```bash
./scripts/backup.sh
```

### Check Status
```bash
docker-compose ps
```

## 🔍 Monitoring

### Database Access
```bash
# Connect to database
docker exec -it financebroke-postgres psql -U postgres -d financebroke_db

# View database logs
docker logs financebroke-postgres
```

### Application Logs
```bash
# View all logs
docker-compose logs

# Follow logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f app
docker-compose logs -f postgres
```

### Performance Monitoring
```bash
# Resource usage
docker stats

# System resources
htop
df -h
```

## 🔒 SSL/HTTPS Setup

The setup script includes automatic SSL with Let's Encrypt:

```bash
sudo ./scripts/setup-ssl.sh
```

This will:
- Obtain SSL certificate for financebroke.virhanali.com
- Configure nginx for HTTPS
- Setup automatic certificate renewal

## 🛠️ Troubleshooting

### Common Issues

1. **Docker build fails**
   ```bash
   docker-compose down
   docker system prune -f
   ./scripts/deploy.sh
   ```

2. **Database connection error**
   ```bash
   # Check database status
   docker-compose ps postgres

   # Check database logs
   docker logs financebroke-postgres

   # Restart database
   docker-compose restart postgres
   ```

3. **Application not accessible**
   ```bash
   # Check nginx status
   docker-compose ps

   # Check firewall
   sudo ufw status

   # Check port 80
   netstat -tlnp | grep :80
   ```

### Port Forwarding (for testing)
```bash
# Forward backend port
docker-compose port app 8080

# Forward database port
docker-compose port postgres 5432
```

## 📁 File Structure

```
financebroke/
├── scripts/
│   ├── setup-vps.sh     # Initial VPS setup
│   ├── deploy.sh        # Deploy application
│   ├── setup-ssl.sh     # Setup SSL certificates
│   ├── update.sh        # Update application
│   └── backup.sh        # Backup data
├── backend/             # Go backend application
├── frontend/            # React frontend
├── docker-compose.yml   # Docker Compose configuration
├── Dockerfile          # Multi-stage Docker build
├── nginx.conf          # Nginx configuration
├── nginx-ssl.conf      # Nginx SSL configuration
└── .env.example        # Environment variables template
```

## 🔧 Development

### Local Development
```bash
# Backend development
cd backend
go run cmd/server/main.go

# Frontend development
cd frontend
npm start
```

### Local Docker
```bash
# Build and run locally
docker-compose up -d

# With SSL (requires certificates)
docker-compose --profile ssl up -d
```

## 📞 Support

For issues:
1. Check logs: `docker-compose logs -f`
2. Verify configuration in `.env`
3. Check system resources: `htop`, `df -h`
4. Restart services: `docker-compose restart`

## 🎯 Performance Tips for 2GB VPS

- Monitor memory usage: `docker stats`
- Restart containers weekly: `docker-compose restart`
- Clean up unused Docker images: `docker system prune -f`
- Setup log rotation in `/etc/logrotate.d/`