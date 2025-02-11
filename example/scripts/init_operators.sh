#! /bin/bash

# Default values
NUM_OPERATORS=1
FUNDS_PK=""
RPC_URL="http://0.0.0.0:8545"

# Parse named arguments
while [ $# -gt 0 ]; do
  case "$1" in
    --num-operators)
      NUM_OPERATORS="$2"
      shift 2
      ;;
    --funds-pk)
      FUNDS_PK="$2"
      shift 2
      ;;
    --rpc-url)
      RPC_URL="$2"
      shift 2
      ;;
    --help)
      echo "Usage: $0 --num-operators <num-operators> --funds-pk <funds-pk> --rpc-url <rpc-url>"
      exit 0
      ;;
    *)
      echo "Unknown parameter: $1"
      exit 1
      ;;
  esac
done

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "Creating $NUM_OPERATORS operators"
START_SOCKET=8080
if [ "$NUM_OPERATORS" -gt 0 ]; then
  for i in $(seq 0 $((NUM_OPERATORS-1))); do
    echo "Creating operator $i"
    "$SCRIPT_DIR"/initialize-operator.sh --id "$i" --funds-pk "$FUNDS_PK" --rpc-url "$RPC_URL" --socket "$START_SOCKET"
    START_SOCKET=$((START_SOCKET+1))
  done
else
  echo "No operators to create"
fi
