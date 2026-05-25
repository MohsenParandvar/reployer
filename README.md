# Reployer

Lightweight Docker Compose deployment and auto-redeploy tool written in Go.

Reployer automatically checks for new container image versions and redeploys your Docker Compose services with minimal configuration.

---

## Features

- Automatic redeploy on image updates
- Lightweight and simple
- Daemon mode support
- YAML-based configuration
- Docker Compose integration
- Manual service update support

---

# Quick Start

## Create Configuration

Create a `config.yml` file:

```yaml
interval_seconds: 10

services:
  - name: nginx
    image: nginx
    deployer: compose
    update_policy: update

    spec:
      file: configs/docker-compose-example.yml
```

---

## Run in Daemon Mode

Daemon mode continuously checks for new image versions in the registry.

```bash
./reployer -daemon -config config.yml
```

When a new image is detected, Reployer automatically:

1. Pulls the latest image
2. Updates the Docker Compose deployment
3. Restarts the container

---

## Run in Update Mode

Update a specific service manually:

```bash
./reployer -config config.yml -update -service nginx
```

This command updates the service using the current image tag.

---

## Change Image Tag and Update

You can also change the image tag before updating:

```bash
./reployer -config config.yml -update -service nginx -tag latest
```

This command:

1. Changes the image tag
2. Pulls the new image
3. Redeploys the container

---

# Example Workflow

```text
Docker Registry
       ↓
Reployer checks image updates
       ↓
New image detected
       ↓
Docker Compose pulls latest image
       ↓
Container redeployed
```

---

# CLI Flags

| Flag | Description |
|---|---|
| `-config` | Configuration YAML file path |
| `-daemon` | Run in daemon mode for automatic image updates |
| `-update` | Update a specific service |
| `-service` | Service name defined in config file |
| `-tag` | Change the image tag before update |

---

# Philosophy

Reployer focuses on simplicity.

- No Kubernetes
- No heavy orchestration
- No unnecessary dependencies
- Just simple Docker Compose deployments

---

# License

MIT