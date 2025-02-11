#!/bin/bash

# ./acquire_and_deposit_steth.sh $ETH_RPC_URL $PRIVATE_KEY  ../contracts/script/input/testnet.json 0.1ether

ETH_RPC_URL=$1
PRIVATE_KEY=$2
EIGENLAYER_DEPLOYMENT_PATH=$3
STETH_STRATEGY=0x7d704507b76571a51d9cae8addabbfd0ba0e63d3
STETH_ADDRESS=0x3F1c547b21f65e10480dE3ad8E19fAAC46C95034
AMOUNT_TO_DEPOSIT=$4

OPERATOR_ADDRESS=$(cast wallet address "$PRIVATE_KEY")

cast send --rpc-url "$ETH_RPC_URL" --private-key "$PRIVATE_KEY" "$STETH_ADDRESS" --value="$AMOUNT_TO_DEPOSIT"
sleep 5


# set +e to avoid exiting the script if the server returns null response
set +e
balance=$(cast call --rpc-url "$ETH_RPC_URL" "$STETH_ADDRESS" "balanceOf(address)(uint256)" "$OPERATOR_ADDRESS")
echo "StETH Balance: $balance"
set -e

strategyManager=$(jq -r '.strategyManager' "$EIGENLAYER_DEPLOYMENT_PATH")

cast send --rpc-url "$ETH_RPC_URL" --private-key "$PRIVATE_KEY" "$STETH_ADDRESS" "approve(address,uint256)" "$strategyManager" "$balance"
echo "Approved Strategy Manager"
sleep 5

cast send --rpc-url "$ETH_RPC_URL" --private-key "$PRIVATE_KEY" "$strategyManager" "depositIntoStrategy(address,address,uint256)" "$STETH_STRATEGY" "$STETH_ADDRESS" "$balance"
echo "Deposited StETH"
