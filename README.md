# DeployAnchor

**Kubernetes deployments as easy as Vercel.**

DeployAnchor is a self-hosted PaaS that sets up and manages a Kubernetes cluster for you. Deploy containers with a single command - just provide an image and subdomain.

```bash
anchor deploy nginx:latest --name myapp --subdomain myapp
```

## Features

- **One-Command Setup** - Bootstraps k3s, MetalLB, Gateway API automatically
- **One-Command Deploys** - No YAML, no kubectl, just `anchor deploy`
- **Web Dashboard** - Visual interface for managing apps, logs, and scaling
- **Self-Healing** - Kubernetes handles automatic recovery
- **Zero-Downtime Updates** - Rolling deployments out of the box
- **Works Anywhere** - On-prem servers, EC2, bare metal

## Quick Start

### Install

```bash
curl -sSL https://deployanchor.dev/install | bash
```

### Bootstrap Your Cluster

```bash
sudo anchor cluster init
```

This installs and configures:
- **k3s** - Lightweight Kubernetes
- **MetalLB** - Load balancer for bare metal
- **Envoy Gateway** - Gateway API implementation
- **cert-manager** - Automatic TLS certificates

### Set Your Domain

```bash
anchor config set domain.base example.com
```

### Deploy Your First App

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

### Manage Your Apps

```bash
anchor env set hello DATABASE_URL=postgres://...   # Set env vars
anchor scale hello --replicas 3                     # Scale up
anchor logs hello -f                                # Stream logs
anchor server                                       # Start web dashboard
```

## Documentation

- [Getting Started](docs/getting-started.md) - Full setup guide
- [CLI Reference](docs/cli-reference.md) - All commands
- [Architecture](docs/architecture.md) - How it works
- [API Reference](docs/api-reference.md) - REST API

## Requirements

- Linux server (Ubuntu 20.04+, Debian 11+)
- 2+ CPU cores, 4GB+ RAM
- Root access
- Domain with DNS access

## Tech Stack

| Component | Technology |
|-----------|------------|
| Backend | Go + Fiber + client-go |
| Frontend | React + Vite + shadcn/ui |
| Cluster | k3s + MetalLB + Envoy Gateway |

## License

MIT
# deployanchor
# deployanchor
