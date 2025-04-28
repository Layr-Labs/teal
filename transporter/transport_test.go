package main

import (
	"encoding/hex"
	"math/big"
	"testing"

	tableCalculator "github.com/Layr-Labs/teal/transporter/bindings/BLSTableCalculator"
	"github.com/stretchr/testify/assert"
)

func TestEncodeOperatorInfoLeaf(t *testing.T) {
	// Create test values
	xPoint, _ := new(big.Int).SetString("5138697240077803445514669414784254799933862402946278134326199877546184124353", 10)
	yPoint, _ := new(big.Int).SetString("12587011617949543324467535889916856826666519601316494966427400843934921824601", 10)

	// Create a weight of 1e18
	weight := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 10^18
	weights := []*big.Int{weight}

	// Create an operator info with the test values
	operatorInfo := tableCalculator.IBLSTableCalculatorTypesBN254OperatorInfo{
		Pubkey: tableCalculator.BN254G1Point{
			X: xPoint,
			Y: yPoint,
		},
		Weights: weights,
	}

	// Encode the operator info
	encoded := EncodeOperatorInfoLeaf(operatorInfo)

	// Convert to hex string for comparison
	encodedHex := "0x" + hex.EncodeToString(encoded)

	// Expected encoded value
	expectedHex := "0x0b5c664a181b916702c1acbbe4617a673cf9fffe733ff902353e315e11f9a7c11bd3ff011ee7d19e9d2ceef4475aac85313e8f31e4b8c66114d32e4ef0ee99590000000000000000000000000000000000000000000000000de0b6b3a7640000"

	// Print the encoded value and expected value for debugging
	t.Logf("Encoded operator info: %s", encodedHex)
	t.Logf("Expected value: %s", expectedHex)

	// Compare the encoded value with the expected value
	assert.Equal(t, expectedHex, encodedHex, "Encoding should match the pre-computed value")
}
