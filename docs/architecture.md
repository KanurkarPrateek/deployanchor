# Architecture

How DeployAnchor works under the hood.

## Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        DeployAnchor                              │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐    ┌─────────────────────────────────────────┐ │
│  │   CLI       │    │           Web Dashboard                  │ │
│  │  (anchor)   │    │                                          │ │
│  └──────┬──────┘    └────────────────┬────────────────────────┘ │
│         │                            │                           │
│         └────────────┬───────────────┘                           │
│                      ▼                                           │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                  Go API Server                             │  │
│  │                                                             │  │
│  │  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────┐  │  │
│  │  │ App Service │ │ Domain Svc  │ │ Cluster Bootstrap   │  │  │
│  │  └─────────────┘ └─────────────┘ └─────────────────────┘  │  │
│  │                          │                                  │  │
│  │                          ▼                                  │  │
│  │  ┌─────────────────────────────────────────────────────┐   │  │
│  │  │              K8s Client (client-go)                  │   │  │
│  │  └─────────────────────────────────────────────────────┘   │  │
│  └───────────────────────────────────────────────────────────┘  │
│                              │                                   │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                    SQLite Database                         │  │
│  └───────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Kubernetes (k3s)                              │
│                                                                  │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────────────────┐ │
│  │   MetalLB    │ │    Envoy     │ │     cert-manager         │ │
│  │              │ │   Gateway    │ │                          │ │
│  └──────────────┘ └──────────────┘ └──────────────────────────┘ │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │                    User Apps                              │   │
│  │  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐         │   │
│  │  │ Deployment  │ │ Deployment  │ │ Deployment  │         │   │
│  │  │ Service     │ │ Service     │ │ Service     │         │   │
│  │  │ HTTPRoute   │ │ HTTPRoute   │ │ HTTPRoute   │         │   │
│  │  └─────────────┘ └─────────────┘ └─────────────┘         │   │
│  └──────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

## Components

### CLI

Built with Cobra. Communicates with the API server or directly with K8s.

### API Server

Built with Fiber. Provides REST API and serves the web dashboard.

### Database

SQLite stores app metadata, deployment history, and configuration.

### Cluster Components

| Component | Purpose |
|-----------|---------|
| k3s | Lightweight Kubernetes |
| MetalLB | LoadBalancer for bare metal |
| Envoy Gateway | Gateway API implementation |
| cert-manager | TLS certificate automation |

## Deployment Flow

```
anchor deploy nginx:latest --name myapp --subdomain myapp
                    │
                    ▼
            ┌───────────────┐
            │ Validate      │
            │ - Image OK?   │
            │ - Name free?  │
            └───────┬───────┘
                    │
                    ▼
            ┌───────────────┐
            │ Store in DB   │
            │ status=deploy │
            └───────┬───────┘
                    │
                    ▼
            ┌───────────────┐
            │ Create K8s    │
            │ - Deployment  │
            │ - Service     │
            │ - HTTPRoute   │
            └───────┬───────┘
                    │
                    ▼
            ┌───────────────┐
            │ Watch rollout │
            │ Wait for pods │
            └───────┬───────┘
                    │
                    ▼
            ┌───────────────┐
            │ Update DB     │
            │ status=running│
            └───────────────┘
```

## K8s Resources Created

For each app, DeployAnchor creates:

### Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
  namespace: deployanchor
spec:
  replicas: 1
  template:
    spec:
      containers:
      - name: main
        image: nginx:latest
        ports:
        - containerPort: 80
```

### Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myapp
spec:
  type: ClusterIP
  ports:
  - port: 80
```

### HTTPRoute (Gateway API)

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: myapp
spec:
  parentRefs:
  - name: main-gateway
  hostnames:
  - "myapp.example.com"
  rules:
  - backendRefs:
    - name: myapp
      port: 80
```

## Gateway API

We use Gateway API instead of Ingress because:

- More expressive routing
- Better role separation
- Future-proof (Ingress successor)
- Works across implementations

Default implementation: **Envoy Gateway**

## MetalLB

Provides LoadBalancer support for bare metal. Uses Layer 2 mode:

```yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
spec:
  addresses:
  - 192.168.1.200-192.168.1.250
```

Traffic flow: Client → MetalLB IP → Envoy Gateway → App

## Database Schema

```sql
CREATE TABLE apps (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    image TEXT NOT NULL,
    subdomain TEXT UNIQUE NOT NULL,
    port INTEGER DEFAULT 80,
    replicas INTEGER DEFAULT 1,
    status TEXT DEFAULT 'pending',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE env_vars (
    app_id TEXT,
    key TEXT,
    value TEXT,
    is_secret BOOLEAN
);

CREATE TABLE deployments (
    app_id TEXT,
    revision INTEGER,
    image TEXT,
    created_at TIMESTAMP
);
```

## Project Structure

```
deployanchor/
├── cmd/anchor/main.go       # CLI entry point
├── internal/
│   ├── api/                 # HTTP handlers
│   ├── app/                 # Deployment logic
│   ├── cluster/             # Bootstrap logic
│   ├── k8s/                 # K8s client
│   ├── models/              # Data models
│   └── store/               # Database
└── web/                     # React dashboard
```
