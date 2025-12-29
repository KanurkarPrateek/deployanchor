# Getting Started

This guide walks you through setting up DeployAnchor and deploying your first app.

## Prerequisites

### System Requirements

- **OS**: Linux (Ubuntu 20.04+, Debian 11+, CentOS 8+)
- **CPU**: 2+ cores
- **RAM**: 4GB+ (8GB recommended)
- **Disk**: 20GB+ free space
- **Network**: Static IP recommended

### Domain Setup

You need a domain where you can create DNS A records. For each app, you'll point a subdomain to your server IP.

## Installation

### 1. Install the CLI

```bash
curl -sSL https://deployanchor.dev/install | bash
```

Or manually:

```bash
wget https://github.com/deployanchor/deployanchor/releases/latest/download/anchor-linux-amd64
chmod +x anchor-linux-amd64
sudo mv anchor-linux-amd64 /usr/local/bin/anchor
```

### 2. Bootstrap the Cluster

```bash
sudo anchor cluster init
```

This installs and configures:

| Component | Purpose |
|-----------|---------|
| k3s | Lightweight Kubernetes |
| MetalLB | LoadBalancer for bare metal |
| Envoy Gateway | Gateway API routing |
| cert-manager | TLS certificates |

Output:
```
Installing k3s...
  Downloading k3s v1.29.0... done
  Starting k3s... done

Installing MetalLB...
  Configuring IP pool... done

Installing Envoy Gateway...
  Creating Gateway... done

Installing cert-manager...
  Creating ClusterIssuer... done

Cluster ready!
Your cluster IP: 192.168.1.200
```

### 3. Configure Your Domain

```bash
anchor config set domain.base example.com
```

### 4. Deploy Your First App

```bash
anchor deploy nginx:latest --name hello --subdomain hello
```

Output:
```
Deploying hello...
  Creating Deployment... done
  Creating Service... done
  Creating HTTPRoute... done

Your app is ready!
Add DNS record: hello.example.com -> 192.168.1.200
```

Now add the DNS A record, wait for propagation, and visit your app!

## Multi-Node Cluster

### On the first node (control plane):

```bash
sudo anchor cluster init --multi-node
```

Save the join token from the output.

### On worker nodes:

```bash
sudo anchor cluster join 192.168.1.100 --token <token>
```

## Common Operations

### Deploy with Options

```bash
# Custom port
anchor deploy myapi:v1 --name api --subdomain api --port 3000

# With environment variables
anchor deploy postgres:15 --name db --subdomain db \
  --env POSTGRES_USER=admin \
  --env POSTGRES_DB=myapp

# With secrets
anchor deploy myapp:v1 --name app --subdomain app \
  --secret API_KEY=sk-xxx

# With storage
anchor deploy postgres:15 --name db --subdomain db \
  --volume /var/lib/postgresql/data:10Gi

# Multiple replicas
anchor deploy myapp:v1 --name api --subdomain api --replicas 3
```

### Manage Apps

```bash
anchor list                          # List all apps
anchor status myapp                  # Detailed status
anchor logs myapp -f                 # Stream logs
anchor scale myapp --replicas 5      # Scale
anchor env set myapp KEY=value       # Set env var
anchor rollback myapp                # Rollback
anchor delete myapp                  # Delete
```

### Enable TLS

```bash
anchor tls enable myapp
```

### Start Web Dashboard

```bash
anchor server
# Open http://localhost:8080
```

## Configuration

Config file: `~/.deployanchor/config.yaml`

```yaml
cluster:
  kubeconfig: ~/.kube/config

server:
  port: 8080
  host: 127.0.0.1

domain:
  base: example.com

metallb:
  ip_range: "192.168.1.200-192.168.1.250"
```

## Troubleshooting

### Cluster not starting

```bash
sudo systemctl status k3s
sudo journalctl -u k3s -f
```

### App not accessible

```bash
anchor status myapp              # Check app status
kubectl get httproute -A         # Check routes
kubectl get svc -n envoy-gateway-system  # Check gateway
```

### Reset and start over

```bash
sudo anchor cluster reset --full
sudo anchor cluster init
```
