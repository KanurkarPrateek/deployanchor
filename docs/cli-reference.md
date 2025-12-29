# CLI Reference

Complete reference for the `anchor` command-line interface.

## Global Flags

```
--config string       Config file (default: ~/.deployanchor/config.yaml)
--verbose, -v         Enable verbose output
--json                Output in JSON format
--help, -h            Show help
```

---

## Cluster Commands

### anchor cluster init

Bootstrap a new Kubernetes cluster.

```bash
sudo anchor cluster init [flags]
```

**Flags:**
```
--multi-node          Enable multi-node mode
--ip-range string     MetalLB IP range (default: auto-detect)
```

**Examples:**
```bash
# Single-node cluster
sudo anchor cluster init

# Multi-node cluster
sudo anchor cluster init --multi-node

# Custom IP range
sudo anchor cluster init --ip-range "10.0.0.100-10.0.0.150"
```

---

### anchor cluster join

Join a worker node to the cluster.

```bash
sudo anchor cluster join <master-ip> --token <token>
```

---

### anchor cluster status

Show cluster health.

```bash
anchor cluster status
```

---

### anchor cluster reset

Remove DeployAnchor and optionally k3s.

```bash
sudo anchor cluster reset [flags]
```

**Flags:**
```
--full    Also uninstall k3s
--force   Skip confirmation
```

---

## App Commands

### anchor deploy

Deploy a container image.

```bash
anchor deploy <image> --name <name> --subdomain <subdomain> [flags]
```

**Flags:**
```
--name string         App name (required)
--subdomain string    Subdomain (required)
--port int            Container port (default: 80)
--replicas int        Replicas (default: 1)
--env strings         Environment variables (KEY=value)
--secret strings      Secrets (KEY=value)
--volume strings      Volumes (path:size)
--cpu string          CPU limit (default: 500m)
--memory string       Memory limit (default: 256Mi)
```

**Examples:**
```bash
anchor deploy nginx:latest --name myapp --subdomain myapp

anchor deploy myapi:v1 --name api --subdomain api --port 3000

anchor deploy postgres:15 --name db --subdomain db \
  --env POSTGRES_USER=admin \
  --volume /var/lib/postgresql/data:10Gi
```

---

### anchor list

List all apps.

```bash
anchor list
```

---

### anchor status

Show app details.

```bash
anchor status <app-name>
```

---

### anchor logs

Stream app logs.

```bash
anchor logs <app-name> [flags]
```

**Flags:**
```
--follow, -f    Stream continuously
--tail int      Lines to show (default: 100)
```

---

### anchor scale

Scale replicas.

```bash
anchor scale <app-name> --replicas <n>
```

---

### anchor autoscale

Configure autoscaling.

```bash
anchor autoscale <app-name> --min <n> --max <n> --cpu <percent>
```

---

### anchor env

Manage environment variables.

```bash
anchor env set <app> KEY=value
anchor env unset <app> KEY
anchor env list <app>
```

---

### anchor secret

Manage secrets.

```bash
anchor secret set <app> KEY=value
anchor secret unset <app> KEY
anchor secret list <app>
```

---

### anchor history

Show deployment history.

```bash
anchor history <app-name>
```

---

### anchor rollback

Rollback to previous version.

```bash
anchor rollback <app-name> [--revision <n>]
```

---

### anchor delete

Delete an app.

```bash
anchor delete <app-name> [--force]
```

---

### anchor tls

Manage TLS.

```bash
anchor tls enable <app-name>
anchor tls disable <app-name>
anchor tls status <app-name>
```

---

## Server Commands

### anchor server

Start API server and web dashboard.

```bash
anchor server [flags]
```

**Flags:**
```
--port int       Port (default: 8080)
--host string    Host (default: 127.0.0.1)
```

---

## Config Commands

### anchor config

Manage configuration.

```bash
anchor config list
anchor config get <key>
anchor config set <key> <value>
```

**Examples:**
```bash
anchor config set domain.base example.com
anchor config set metallb.ip_range "10.0.0.100-10.0.0.150"
```

---

### anchor version

Show version.

```bash
anchor version
```

---

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | Cluster connection failed |
| 4 | Resource not found |
| 5 | Deployment failed |
