# Teal Example

This example will take you end to end through deploying a minimal AVS that uses BLS aggregation to verify operator signatures on a SIMPLE TASK.

## Step 0: Setup Environment

```
export ETH_RPC_URL=<HOLSKY_RPC_URL>
export DESTINATION_RPC_URL=<DESTINATION_RPC_URL>
export PRIVATE_KEY=<PRIVATE_KEY_WITH_SOME_ETH_IN_IT>
export ETHERSCAN_API_KEY=<ETHERSCAN_API_KEY>
```

# Step 1: Deploy the AVS

For demo purposes, we'll deploy an AVS to EigenLayer's **Holesky Preprod** with a single quorum that just pays attention to the StETH strategy.

```
cd crossChainExample/contracts
forge script --rpc-url $ETH_RPC_URL --private-key $PRIVATE_KEY --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast script/DeployCrossChainAVS.s.sol --sig "run(string,uint256,address[])" -- ./script/input/preprod.json 200 "[0x7d704507b76571a51d9cae8addabbfd0ba0e63d3]"
```
See that `example/contracts/script/output/avs_deploy_output.json` has been created.

Only change the strategies if you know what you're doing.

Next, deploy the Certificate Verifier on the destination chain

```
forge script --rpc-url $DESTINATION_RPC_URL --private-key $PRIVATE_KEY --verify --etherscan-api-key $ETHERSCAN_API_KEY --broadcast script/DeployCertificateVerifier.sol --sig "run(string,uint32,uint32)" -- ./script/output/avs_deploy_output.json 0 86400
```

# Step 2: Setup Operators

First, let's create an operator and give it some stETH.
```
cd crossChainExample/scripts
./init_operators.sh --num-operators 3 --rpc-url $ETH_RPC_URL --funds-pk $PRIVATE_KEY
```
This will mint and deposit 0.1 ether of stETH into EigenLayer on behalf of the operator.

Now, let's register the operator with EigenLayer and the AVS.

```
./register_operator_avs.sh --rpc-url $ETH_RPC_URL
```

This will register your operator with the AVS and EigenLayer. 

## Step 3: Calculate the operator table and transport to the certificate verifier

Step 1 already deployed an `OperatorTableCalculator` on the source chain and a `CertificateVerifier` on the destination chain. 

Let's read the stakes and transport them. We use the original `PRIVATE_KEY` as the EOA to transport. Make sure it is funded with gas on the destination chain too. 

```
cd crossChainExample/scripts
./transport_stakes.sh --source-rpc-url $ETH_RPC_URL --destination-rpc-url $DESTINATION_RPC_URL --ecdsa-private-key $PRIVATE_KEY --operator-set-id 0
```

# Step 4: Run the Operators

Run the operator first.
```
./start_nodes.sh --rpc-url $ETH_RPC_URL
```

# Step 5: Run the aggregator and begin verifying certificates 
In a seperate terminal, run the aggregator.
```
./start_aggregator.sh --rpc-url $HOLESKY_WSS_URL --private-key $PRIVATE_KEY
```
