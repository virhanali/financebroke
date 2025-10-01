#!/bin/bash

# Finance App VPS Setup Script for Kubernetes
set -e

echo "🚀 Setting up Finance App on VPS with Kubernetes..."

# Update system
echo "📦 Updating system packages..."
sudo apt update && sudo apt upgrade -y

# Install required packages
echo "🔧 Installing required packages..."
sudo apt install -y curl wget git apt-transport-https ca-certificates gnupg lsb-release

# Install Docker
echo "🐳 Installing Docker..."
if ! command -v docker &> /dev/null; then
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt update
    sudo apt install -y docker-ce docker-ce-cli containerd.io
    sudo usermod -aG docker $USER
fi

# Install Kubernetes components
echo "☸️ Installing Kubernetes components..."
# Disable swap
sudo swapoff -a
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

# Install kubelet, kubeadm, kubectl
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.28/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

sudo apt update
sudo apt install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

# Configure kernel modules for Kubernetes
echo "🔧 Configuring kernel modules..."
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# Configure sysctl
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF

sudo sysctl --system

# Install k3s (lightweight Kubernetes)
echo "🚀 Installing k3s..."
if ! command -v k3s &> /dev/null; then
    curl -sfL https://get.k3s.io | sh -s - --write-kubeconfig-mode 644
fi

# Wait for k3s to be ready
echo "⏳ Waiting for k3s to be ready..."
sleep 30

# Setup kubectl
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown $USER:$USER ~/.kube/config

# Install Helm
echo "📦 Installing Helm..."
if ! command -v helm &> /dev/null; then
    curl https://get.helm.sh/helm-v3.13.0-linux-amd64.tar.gz | tar xz
    sudo mv linux-amd64/helm /usr/local/bin/
    rm -rf linux-amd64
fi

# Install Ingress Controller
echo "🌐 Installing NGINX Ingress Controller..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml

# Install Cert-Manager for SSL
echo "🔒 Installing Cert-Manager..."
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.13.2 \
  --set installCRDs=true

# Clone repository
REPO_DIR="/opt/finance-app"
if [ -d "$REPO_DIR" ]; then
    echo "📁 Updating existing repository..."
    cd $REPO_DIR
    git pull origin master
else
    echo "📁 Cloning repository..."
    sudo mkdir -p $REPO_DIR
    sudo chown $USER:$USER $REPO_DIR
    git clone https://github.com/virhanali/finance-broke.git $REPO_DIR
    cd $REPO_DIR
fi

echo "✅ VPS setup completed!"
echo "🔧 Next steps:"
echo "1. Update the secrets in k8s/backend/secret.yaml"
echo "2. Update your domain in k8s/ingress/ingress.yaml"
echo "3. Run: ./scripts/deploy-k8s.sh"