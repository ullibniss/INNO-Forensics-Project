#!/bin/sh

TOKEN_FILE="/vault/file/vault-root-token"

INIT_STATUS=$(vault status -format=json 2>/dev/null)
INITIALIZED=$(echo "$INIT_STATUS" | grep -o '"initialized":[^,}]*' | cut -d: -f2 | tr -d ' ')

if [ "$INITIALIZED" != "true" ]; then
  echo "Initializing Vault..."
  INIT_OUTPUT=$(vault operator init -key-shares=1 -key-threshold=1 -format=json)
  if [ $? -ne 0 ]; then
    echo "ERROR: vault operator init failed"
    exit 1
  fi

  UNSEAL_KEY=$(echo "$INIT_OUTPUT" | grep -o '"unseal_keys_b64":\["[^"]*"' | cut -d'"' -f4)
  ROOT_TOKEN=$(echo "$INIT_OUTPUT" | grep -o '"root_token":"[^"]*"' | cut -d'"' -f4)

  echo "$ROOT_TOKEN" > "$TOKEN_FILE"
  echo "Root token saved to $TOKEN_FILE"

  vault operator unseal "$UNSEAL_KEY"
  if [ $? -ne 0 ]; then
    echo "ERROR: unseal failed"
    exit 1
  fi

  echo "Vault initialized and unsealed."
else
  SEALED=$(echo "$INIT_STATUS" | grep -o '"sealed":[^,}]*' | cut -d: -f2 | tr -d ' ')
  if [ "$SEALED" = "true" ]; then
    echo "ERROR: Vault is sealed but already initialized. Manual unseal required."
    exit 1
  fi
  echo "Vault already initialized and unsealed. Nothing to do."
fi
