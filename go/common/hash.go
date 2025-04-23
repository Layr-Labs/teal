package common

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256HashFn(response types.TaskResponse) (types.TaskResponseDigest, error) {
	responseBytes, ok := response.([]byte)
	if !ok {
		return types.TaskResponseDigest{}, fmt.Errorf("response is not a byte array")
	}

	return [32]byte(crypto.Keccak256(responseBytes)), nil
}
