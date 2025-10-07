#!/usr/bin/env bash
set -euo pipefail

echo "==> Updating package lists"
sudo apt-get update -y

echo "==> Installing Docker"
sudo apt-get install -y docker.io curl

echo "==> Installing Docker Compose v2"
DOCKER_COMPOSE_VERSION="v2.29.2"
sudo mkdir -p /usr/libexec/docker/cli-plugins
sudo curl -L "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/libexec/docker/cli-plugins/docker-compose
sudo chmod +x /usr/libexec/docker/cli-plugins/docker-compose

echo "==> Verifying Docker Compose version"
docker compose version

echo "==> Adding vagrant user to docker group"
sudo usermod -aG docker vagrant

echo "==> Enabling and starting Docker service"
sudo systemctl enable --now docker

echo "==> Starting Docker Compose stack"
cd /vagrant


docker compose up -d

# echo "==> Running database migrations using a temporary Go container"
docker compose run --rm tools bash -c "\
  go install github.com/jackc/tern/v2@latest && \
  cd db/migrations && \
  tern migrate \
"

echo "==> Provisioning complete! App should be available at http://192.168.56.10:8080"
