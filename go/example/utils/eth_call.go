package utils

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	MinDataSize = 96
)

func CallFromBytes(data []byte) (uint64, ethereum.CallMsg) {
	blockNumber := binary.BigEndian.Uint64(data[:8])
	from := common.BytesToAddress(data[8:28])
	to := common.BytesToAddress(data[28:48])
	gas := binary.BigEndian.Uint64(data[48:56])
	gasPrice := new(big.Int).SetBytes(data[56:64])
	value := new(big.Int).SetBytes(data[64:96])
	calldata := data[96:]

	return blockNumber, ethereum.CallMsg{
		From:     from,
		To:       &to,
		Gas:      gas,
		GasPrice: gasPrice,
		Value:    value,
		Data:     calldata,
	}
}

func CallToBytes(blockNumber uint64, callMsg ethereum.CallMsg) []byte {
	data := make([]byte, 96+len(callMsg.Data))
	binary.BigEndian.PutUint64(data[:8], blockNumber)
	copy(data[8:28], callMsg.From.Bytes())
	copy(data[28:48], callMsg.To.Bytes())
	binary.BigEndian.PutUint64(data[48:56], callMsg.Gas)
	copy(data[56:64], callMsg.GasPrice.Bytes())
	copy(data[64:96], callMsg.Value.Bytes())
	copy(data[96:], callMsg.Data)
	return data
}
