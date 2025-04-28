package node

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Layr-Labs/teal/example/utils"
	"github.com/Layr-Labs/teal/node/server"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	MaxDataSize   = 128000
	MaxGas        = 30000000
	MinBlockDepth = 100
	MaxBlockDepth = 10000
)

type EthCallNode struct {
	*server.BaseNode
	ethClient *ethclient.Client
}

func NewEthCallNode(nodeConfig server.Config, rpcUrl string) *EthCallNode {
	node := &EthCallNode{}
	node.BaseNode = server.NewBaseNode(nodeConfig, node)

	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}
	node.ethClient = ethClient

	return node
}

func (n *EthCallNode) GetResponse(nodeConfig server.Config, data []byte) ([]byte, error) {
	if len(data) < utils.MinDataSize {
		return nil, fmt.Errorf("data too short")
	}

	if len(data) > MaxDataSize {
		return nil, fmt.Errorf("data too long")
	}

	currBlockNumber, err := n.ethClient.BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get current block number: %w", err)
	}

	// next bytes are the call msg
	blockNumber, callMsg := utils.CallFromBytes(data)
	if blockNumber+MinBlockDepth > currBlockNumber || blockNumber+MaxBlockDepth < currBlockNumber {
		return nil, fmt.Errorf("block number out of range")
	}
	if callMsg.Gas > MaxGas {
		return nil, fmt.Errorf("gas too high")
	}

	returnData, err := n.ethClient.CallContract(context.Background(), callMsg, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	// summarise request and return data and return response!
	requestDataHash := crypto.Keccak256(data)
	returnDataHash := crypto.Keccak256(returnData)
	return crypto.Keccak256(append(requestDataHash, returnDataHash...)), nil
}
