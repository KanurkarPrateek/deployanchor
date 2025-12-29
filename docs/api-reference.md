# API Reference

REST API for DeployAnchor. Start the server with `anchor server`.

Base URL: `http://localhost:8080/api`

## Apps

### List Apps

```
GET /api/apps
```

**Response:**
```json
{
  "apps": [
    {
      "id": "abc123",
      "name": "myapp",
      "image": "nginx:latest",
      "subdomain": "myapp",
      "port": 80,
      "replicas": 2,
      "status": "running",
      "created_at": "2024-01-15T10:00:00Z"
    }
  ]
}
```

---

### Get App

```
GET /api/apps/:name
```

**Response:**
```json
{
  "id": "abc123",
  "name": "myapp",
  "image": "nginx:latest",
  "subdomain": "myapp",
  "port": 80,
  "replicas": 2,
  "status": "running",
  "env_vars": {
    "LOG_LEVEL": "info"
  },
  "secrets": ["API_KEY"],
  "created_at": "2024-01-15T10:00:00Z"
}
```

---

### Create App

```
POST /api/apps
```

**Request:**
```json
{
  "name": "myapp",
  "image": "nginx:latest",
  "subdomain": "myapp",
  "port": 80,
  "replicas": 1,
  "env_vars": {
    "LOG_LEVEL": "info"
  },
  "secrets": {
    "API_KEY": "sk-xxx"
  }
}
```

**Response:**
```json
{
  "id": "abc123",
  "name": "myapp",
  "status": "deploying"
}
```

---

### Update App

```
PUT /api/apps/:name
```

**Request:**
```json
{
  "image": "nginx:1.25",
  "replicas": 3
}
```

---

### Delete App

```
DELETE /api/apps/:name
```

---

### Scale App

```
POST /api/apps/:name/scale
```

**Request:**
```json
{
  "replicas": 5
}
```

---

### Get App Logs

```
GET /api/apps/:name/logs?tail=100
```

**Response:**
```json
{
  "logs": [
    {
      "timestamp": "2024-01-15T10:00:00Z",
      "message": "Server started on port 80"
    }
  ]
}
```

---

### Stream Logs (WebSocket)

```
WS /api/apps/:name/logs/stream
```

Messages are JSON:
```json
{"timestamp": "...", "message": "..."}
```

---

### Get Deployment History

```
GET /api/apps/:name/history
```

**Response:**
```json
{
  "revisions": [
    {
      "revision": 3,
      "image": "nginx:1.25",
      "created_at": "2024-01-15T10:00:00Z",
      "current": true
    },
    {
      "revision": 2,
      "image": "nginx:1.24",
      "created_at": "2024-01-14T10:00:00Z"
    }
  ]
}
```

---

### Rollback

```
POST /api/apps/:name/rollback
```

**Request:**
```json
{
  "revision": 2
}
```

---

## Environment Variables

### Set Env Var

```
POST /api/apps/:name/env
```

**Request:**
```json
{
  "key": "LOG_LEVEL",
  "value": "debug"
}
```

---

### Delete Env Var

```
DELETE /api/apps/:name/env/:key
```

---

## Secrets

### Set Secret

```
POST /api/apps/:name/secrets
```

**Request:**
```json
{
  "key": "API_KEY",
  "value": "sk-xxx"
}
```

---

### Delete Secret

```
DELETE /api/apps/:name/secrets/:key
```

---

## TLS

### Enable TLS

```
POST /api/apps/:name/tls/enable
```

---

### Disable TLS

```
POST /api/apps/:name/tls/disable
```

---

### TLS Status

```
GET /api/apps/:name/tls
```

**Response:**
```json
{
  "enabled": true,
  "issuer": "letsencrypt-prod",
  "expires_at": "2024-04-15T00:00:00Z"
}
```

---

## Cluster

### Cluster Status

```
GET /api/cluster/status
```

**Response:**
```json
{
  "status": "healthy",
  "nodes": [
    {
      "name": "node-1",
      "status": "Ready",
      "roles": ["control-plane"]
    }
  ],
  "components": {
    "metallb": "running",
    "envoy-gateway": "running",
    "cert-manager": "running"
  }
}
```

---

## Config

### Get Config

```
GET /api/config
```

**Response:**
```json
{
  "domain": {
    "base": "example.com"
  },
  "metallb": {
    "ip_range": "192.168.1.200-192.168.1.250"
  }
}
```

---

### Update Config

```
PUT /api/config
```

**Request:**
```json
{
  "domain": {
    "base": "newdomain.com"
  }
}
```

---

## Error Responses

All errors return:

```json
{
  "error": "App not found",
  "code": "NOT_FOUND"
}
```

**Error Codes:**

| Code | HTTP Status | Description |
|------|-------------|-------------|
| BAD_REQUEST | 400 | Invalid request |
| NOT_FOUND | 404 | Resource not found |
| CONFLICT | 409 | Name already exists |
| INTERNAL | 500 | Server error |
