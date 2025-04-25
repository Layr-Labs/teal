#! /bin/bash

set -e

# Default values
RPC_URL="http://0.0.0.0:8545"


# Parse named arguments
while [ $# -gt 0 ]; do
  case "$1" in
    --rpc-url)
      RPC_URL="$2"
      shift 2
      ;;
    --help)
      echo "Usage: $0 --rpc-url <rpc-url>"
      exit 0
      ;;
    *)
      echo "Unknown parameter: $1"
      exit 1
      ;;
  esac
done

if [ -z "$RPC_URL" ]; then
  echo "Error: --rpc-url is required"
  exit 1
fi

echo "Registering operator to AVS with RPC URL $RPC_URL"

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Build the nodes
PARENT_DIR=$SCRIPT_DIR/..

for file in "$SCRIPT_DIR"/operators/*.json; do
  echo "Starting node from $file"
  if [ -r "$file" ]; then
    BLS_PRIVATE_KEY=$(jq -r '.bls_private_key' "$file")
    ECDSA_PRIVATE_KEY=$(jq -r '.ecdsa_private_key' "$file")
    SOCKET=$(jq -r '.socket' "$file")
    echo "Registering operator to AVS with BLS private key $BLS_PRIVATE_KEY, ECDSA private key $ECDSA_PRIVATE_KEY, socket $SOCKET"
    go run "$SCRIPT_DIR"/register.go \
      --eth-url "$RPC_URL" \
      --eigenlayer-deployment-path "$PARENT_DIR"/contracts/script/input/testnet.json \
      --avs-deployment-path "$PARENT_DIR"/contracts/script/output/avs_deploy_output.json \
      --ecdsa-private-key "$ECDSA_PRIVATE_KEY" \
      --bls-private-key "$BLS_PRIVATE_KEY" \
      --socket "$SOCKET"
  else
    echo "File $file is not readable"
  fi
done