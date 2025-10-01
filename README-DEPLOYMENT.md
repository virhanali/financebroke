# Finance App Kubernetes Deployment Guide

## 🚀 Quick Start

### 1. Setup VPS
```bash
# SSH into your VPS
ssh root@your-vps-ip

# Clone and run setup
git clone https://github.com/virhanali/finance-broke.git
cd finance-broke
./scripts/setup-vps.sh
```

### 2. Configure Secrets
```bash
# After setup completes, configure your secrets
./scripts/setup-secrets.sh
```

### 3. Update Domain
Edit `k8s/ingress/ingress.yaml`:
```yaml
# Replace your-domain.com with your actual domain
host: your-domain.com
```

### 4. Deploy
```bash
./scripts/deploy-k8s.sh
```

## 📋 Prerequisites

- Ubuntu 20.04+ or CentOS 8+
- At least 2GB RAM, 2 CPU cores
- Domain name (optional, for SSL)
- Docker Hub account (optional, for remote registry)

## 🔧 Configuration

### Environment Variables
Update these values in the secrets:
- Database password
- JWT secret
- Telegram bot token
- SMTP credentials

### Domain Setup
1. Point your domain A record to VPS IP
2. Update `k8s/ingress/ingress.yaml` with your domain
3. Cert-Manager will automatically provision SSL certificate

## 📊 Monitoring

### Check Deployment Status
```bash
kubectl get pods -n finance-app
kubectl get services -n finance-app
kubectl get ingress -n finance-app
```

### View Logs
```bash
# Backend logs
kubectl logs -f deployment/finance-app-backend -n finance-app

# Database logs
kubectl logs -f deployment/postgres -n finance-app
```

### Access Services
```bash
# Get shell access to backend
kubectl exec -it deployment/finance-app-backend -n finance-app -- /bin/sh

# Access database
kubectl exec -it deployment/postgres -n finance-app -- psql -U postgres -d finance_app
```

## 🔄 Updates

### Update Application
```bash
git pull origin master
./scripts/deploy-k8s.sh
```

### Rollback Deployment
```bash
kubectl rollout undo deployment/finance-app-backend -n finance-app
```

## 🧹 Cleanup

Remove entire application:
```bash
./scripts/cleanup.sh
```

## 🔍 Troubleshooting

### Common Issues

1. **Pod not starting**
   ```bash
   kubectl describe pod <pod-name> -n finance-app
   ```

2. **Database connection issues**
   ```bash
   kubectl logs deployment/postgres -n finance-app
   ```

3. **Ingress not working**
   ```bash
   kubectl logs -n ingress-nginx deployment/ingress-nginx-controller
   ```

### Port Forwarding (for testing)
```bash
# Forward backend port
kubectl port-forward service/finance-app-service 8080:80 -n finance-app

# Forward database port
kubectl port-forward service/postgres-service 5432:5432 -n finance-app
```

## 📁 File Structure

```
finance-app/
├── k8s/
│   ├── backend/
│   │   ├── configmap.yaml
│   │   ├── secret.yaml
│   │   └── deployment.yaml
│   ├── database/
│   │   ├── pvc.yaml
│   │   └── deployment.yaml
│   ├── ingress/
│   │   └── ingress.yaml
│   └── namespace.yaml
├── scripts/
│   ├── setup-vps.sh
│   ├── deploy-k8s.sh
│   ├── setup-secrets.sh
│   └── cleanup.sh
└── Dockerfile.k8s
```