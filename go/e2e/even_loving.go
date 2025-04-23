package e2e

import (
	"errors"
	"math/big"

	"github.com/Layr-Labs/teal/go/node/server"
)

type EvenLovingNode struct {
	*server.BaseNode
}

func NewEvenLovingNode(config server.Config) *EvenLovingNode {
	node := &EvenLovingNode{}
	node.BaseNode = server.NewBaseNode(config, node)
	return node
}

func (n *EvenLovingNode) GetResponse(config server.Config, data []byte) ([]byte, error) {
	// convert data to bigInt
	dataInt := new(big.Int).SetBytes(data)

	// Can use both config and data for validation
	if dataInt.Mod(dataInt, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return data, nil
	}

	// Custom validation logic using config parameters
	return []byte{}, errors.New("invalid data")
}
