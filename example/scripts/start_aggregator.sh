#! /bin/bash

set -e

# Default values
RPC_URL="http://0.0.0.0:8545"
PRIVATE_KEY=""


# Parse named arguments
while [ $# -gt 0 ]; do
  case "$1" in
    --rpc-url)
      RPC_URL="$2"
      shift 2
      ;;
    --private-key)
      PRIVATE_KEY="$2"
      shift 2
      ;;
    --help)
      echo "Usage: $0 --rpc-url <rpc-url> --private-key <private-key>"
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

if [ -z "$PRIVATE_KEY" ]; then
  echo "Error: --private-key is required"
  exit 1
fi

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Build the nodes
PARENT_DIR=$SCRIPT_DIR/..

go run "$PARENT_DIR"/aggregator/eth_call.go  \
  --eth-url "$RPC_URL" \
  --avs-deployment-path "$PARENT_DIR"/contracts/script/output/avs_deploy_output.json \
  --ecdsa-private-key "$PRIVATE_KEY"