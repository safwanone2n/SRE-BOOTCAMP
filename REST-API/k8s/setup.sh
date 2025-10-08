#!/bin/bash
set -e

# ------------------------------
# 1️⃣ Install External Secrets Operator via Helm
# ------------------------------
echo "Installing External Secrets Operator..."
helm repo add external-secrets https://charts.external-secrets.io
helm repo update

helm upgrade --install external-secrets external-secrets/external-secrets \
  --namespace external-secrets \
  --create-namespace

echo "Waiting 15s for ESO pods to be ready..."
sleep 15

# ------------------------------
# 2️⃣ Deploy Vault
# ------------------------------
echo "Deploying Vault..."
cd ./base
kubectl apply -f vault.yml
echo "Waiting 15s for Vault pod to be ready..."
kubectl wait --for=condition=available --timeout=60s deployment/vault -n vault-ns

# ------------------------------
# 3️⃣ Set env vars for Vault CLI
# ------------------------------
export VAULT_ADDR="http://127.0.0.1:8200"
export VAULT_TOKEN="root"

# Port-forward Vault for local CLI
kubectl port-forward -n vault-ns svc/vault 8200:8200 &
PF_VAULT=$!
sleep 5

# ------------------------------
# 4️⃣ Configure Vault KV v2 and write secrets
# ------------------------------
vault secrets enable -path=secret -version=2 kv || echo "KV v2 already enabled"

vault kv put secret/restapi \
  user=postgres \
  password=postgres \
  database=restapi \
  database_url='postgres://postgres:postgres@postgres-service:5432/restapi?sslmode=disable'

echo "Vault secrets created."

# ------------------------------
# 5️⃣ Apply SecretStore / ClusterSecretStore
# ------------------------------
echo "Applying secret_store.yml..."
cd ../external-secrets
kubectl apply -f secret_store.yml

# ------------------------------
# 6️⃣ Apply ExternalSecret
# ------------------------------
echo "Applying external_secret.yml..."
kubectl apply -f external_secret.yml

# ------------------------------
# 7️⃣ Deploy DB
# ------------------------------
cd ../base
echo "Applying db.yml..."
kubectl apply -f db.yml

# ------------------------------
# 8️⃣ Deploy Application
# ------------------------------
echo "Applying app.yml..."
kubectl apply -f app.yml

# ------------------------------
# 9️⃣ Port-forward Application to localhost:8080
# ------------------------------
echo "Port-forwarding rest-api-service to localhost:8080..."
kubectl port-forward -n rest-api-namespace svc/rest-api-service 8080:8080 &
PF_APP=$!

echo "✅ Setup complete!"
echo "Vault CLI is accessible with:"
echo "  export VAULT_ADDR=http://127.0.0.1:8200"
echo "  export VAULT_TOKEN=root"
echo "Application is accessible at http://localhost:8080"

------------------------------
 🔹 Cleanup function on exit
------------------------------
cleanup() {
  echo "Cleaning up port-forwards..."
  kill $PF_VAULT $PF_APP || true
}
trap cleanup EXIT
