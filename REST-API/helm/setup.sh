#!/bin/bash
set -e

HELM_CHART_DIR="."
RELEASE_NAME="rest-api-release"
NAMESPACE="rest-api-namespace"

# ------------------------------
# 1️⃣ Add Helm repos for subcharts
# ------------------------------
echo "Adding Helm repos..."
cd ./rest-api-helm
helm dependency update

# ------------------------------
# 2️⃣ Deploy REST API Helm chart (with Vault & ESO subcharts)
# ------------------------------
echo "Deploying Helm chart with Vault & ESO subcharts..."
helm upgrade --install $RELEASE_NAME $HELM_CHART_DIR \
  --namespace $NAMESPACE \
  --create-namespace \
  --wait \
  -f $HELM_CHART_DIR/values.yaml

# ------------------------------
# 3️⃣ Port-forward Vault for CLI (optional)
# ------------------------------
echo "Port-forwarding Vault for CLI access..."
kubectl port-forward -n $NAMESPACE svc/${RELEASE_NAME}-vault 8200:8200 &
PF_VAULT=$!
sleep 5

export VAULT_ADDR="http://127.0.0.1:8200"
export VAULT_TOKEN="root"

echo "Vault CLI ready:"
echo "  export VAULT_ADDR=$VAULT_ADDR"
echo "  export VAULT_TOKEN=$VAULT_TOKEN"

# ------------------------------
# 4️⃣ Initialize Vault secrets if not templated
# ------------------------------
echo "Writing initial secrets to Vault..."
vault secrets enable -path=secret -version=2 kv || echo "KV v2 already enabled"

vault kv put secret/restapi \
  user=postgres \
  password=postgres \
  database=restapi \
  database_url='postgres://postgres:postgres@postgres-service:5432/restapi?sslmode=disable'

echo "Vault secrets created."

# ------------------------------
# 5️⃣ Wait for ESO to sync secrets
# ------------------------------
echo "Waiting for External Secrets Operator to sync secrets..."
kubectl wait --for=condition=available --timeout=60s deployment/external-secrets -n external-secrets || true
sleep 10

# ------------------------------
# 6️⃣ Port-forward REST API for local access
# ------------------------------
echo "Port-forwarding REST API service to localhost:8080..."
kubectl port-forward -n $NAMESPACE svc/${RELEASE_NAME}-service 8080:8080 &
PF_APP=$!

echo "✅ Helm-based setup complete!"
echo "Application is accessible at http://localhost:8080"

# ------------------------------
# 🔹 Cleanup
# ------------------------------
cleanup() {
  echo "Cleaning up port-forwards..."
  kill $PF_VAULT $PF_APP || true
}
trap cleanup EXIT
