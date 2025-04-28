package main

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/testutils"
	tableCalculator "github.com/Layr-Labs/teal/transporter/bindings/BLSTableCalculator"

	"github.com/stretchr/testify/assert"
)

func TestEncodeOperatorInfoLeaf(t *testing.T) {
	// Create test values
	xPoint, _ := new(big.Int).SetString("5138697240077803445514669414784254799933862402946278134326199877546184124353", 10)
	yPoint, _ := new(big.Int).SetString("12587011617949543324467535889916856826666519601316494966427400843934921824601", 10)

	// Create a weight of 1e18
	weight := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 1e18
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

// TODO: Validate against impl
func TestCalculateOperatorInfoTree(t *testing.T) {
	// Define the test values
	// TODO: fuzz this, generate arbitrary size tree
	pubkey1X, _ := new(big.Int).SetString("12876932181527085302090685170570508842280292907932645173379329305332411180712", 10)
	pubkey1Y, _ := new(big.Int).SetString("7062153448257026379008314446710613533480546459680849608831392377655541567657", 10)
	weight1 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 1e18

	pubkey2X, _ := new(big.Int).SetString("4905309787131494400818314146123823043382392666010137491213667351654865077267", 10)
	pubkey2Y, _ := new(big.Int).SetString("19262514345326869704859660524305846818700960529405071279918676987606253782084", 10)
	weight2 := new(big.Int).Mul(
		new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
		big.NewInt(2),
	) // 2e18

	// Create test operator infos
	operatorInfos := []tableCalculator.IBLSTableCalculatorTypesBN254OperatorInfo{
		{
			Pubkey: tableCalculator.BN254G1Point{
				X: pubkey1X,
				Y: pubkey1Y,
			},
			Weights: []*big.Int{weight1},
		},
		{
			Pubkey: tableCalculator.BN254G1Point{
				X: pubkey2X,
				Y: pubkey2Y,
			},
			Weights: []*big.Int{weight2},
		},
	}

	// Calculate tree
	logger := testutils.GetTestLogger()
	operatorInfoTree, _, err := calculateOperatorInfoTree(logger, operatorInfos)

	assert.NoError(t, err)

	// Check that the leaves are present in the tree
	assert.Len(t, operatorInfoTree.Data, len(operatorInfos))
	for i, operatorInfo := range operatorInfos {
		assert.Equal(t, operatorInfoTree.Data[i], EncodeOperatorInfoLeaf(operatorInfo))
	}
}
