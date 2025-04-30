#! /bin/bash

set -e

# Default values
RPC_URL="http://0.0.0.0:8545"
PRIVATE_KEY=""


# Parse named arguments
while [ $# -gt 0 ]; do
  case "$1" in
    --source-rpc-url)
      SOURCE_RPC_URL="$2"
      shift 2
      ;;
    --destination-rpc-url)
      DESTINATION_RPC_URL="$2"
      shift 2
      ;;
    --operator-set-id)
      OPERATOR_SET_ID="$2"
      shift 2
      ;;
    --ecdsa-private-key)
      PRIVATE_KEY="$2"
      shift 2
      ;;
    --help)
      echo "Usage: $0 --source-rpc-url <source-rpc-url> --destination-rpc-url <destination-rpc-url> --operator-set-id <operator-set-id>"
      exit 0
      ;;
    *)
      echo "Unknown parameter: $1"
      exit 1
      ;;
  esac
done

if [ -z "$SOURCE_RPC_URL" ]; then
  echo "Error: --source-rpc-url is required"
  exit 1
fi

if [ -z "$DESTINATION_RPC_URL" ]; then
  echo "Error: --destination-rpc-url is required"
  exit 1
fi

if [ -z "$PRIVATE_KEY" ]; then
  echo "Error: --private-key is required"
  exit 1
fi

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Build the nodes
PARENT_DIR=$SCRIPT_DIR/..

go run "$PARENT_DIR"/aggregator/eth_call.go  \
  --source-rpc-url "$SOURCE_RPC_URL" \
  --destination-rpc-url "$DESTINATION_RPC_URL" \
  --ecdsa-private-key "$PRIVATE_KEY" \
  --avs-deployment-path "$PARENT_DIR"/contracts/script/output/avs_deploy_output.json \
  --avs-cert-verifier-deployment-path "$PARENT_DIR"/contracts/script/output/certificate_verifier_deploy_output.json \
  --operator-set-id "$OPERATOR_SET_ID" 