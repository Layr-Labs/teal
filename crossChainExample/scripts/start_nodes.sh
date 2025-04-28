#!/bin/bash

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

declare -a PIDS=()

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Build the nodes
PARENT_DIR=$SCRIPT_DIR/..
go build -o "$PARENT_DIR"/bin/node "$PARENT_DIR"/node/cmd/main.go

for file in "$SCRIPT_DIR"/operators/*.json; do
  echo "Starting node from $file"
  if [ -r "$file" ]; then
    BLS_PRIVATE_KEY=$(jq -r '.bls_private_key' "$file")
    ECDSA_PRIVATE_KEY=$(jq -r '.ecdsa_private_key' "$file")
    SOCKET=$(jq -r '.socket' "$file")
    PORT=$(echo "$SOCKET" | cut -d ':' -f 2)
    echo "Starting node with BLS private key $BLS_PRIVATE_KEY, ECDSA private key $ECDSA_PRIVATE_KEY, socket $SOCKET, and port $PORT"
    "$PARENT_DIR"/bin/node --bls-private-key "$BLS_PRIVATE_KEY" --service-port "$PORT" --eth-url "$RPC_URL" & PIDS+=($!)
  else
    echo "File $file is not readable"
  fi
done

# Save PIDs to a file for later reference
echo "${PIDS[@]}" > /tmp/teal_pids.txt

# Function to check if all processes are running
check_processes() {
    for pid in "${PIDS[@]}"; do
        if ! kill -0 "$pid" 2>/dev/null; then
            echo "Process $pid has died"
            exit 1
        fi
    done
}

echo "Nodes started with PIDs: ${PIDS[*]}"

# Cleanup function
cleanup() {
    echo "Stopping all processes..."
    for pid in "${PIDS[@]}"; do
        kill "$pid" 2>/dev/null || true
    done
    rm -f /tmp/teal_pids.txt
}

trap cleanup EXIT INT TERM

# Keep script running and monitoring processes
while true; do
    check_processes
    sleep 5
done