package main

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/teal/example/utils"
	certificateVerifier "github.com/Layr-Labs/teal/transporter/bindings/BLSCertificateVerifier"
	tableCalculator "github.com/Layr-Labs/teal/transporter/bindings/BLSTableCalculator"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"
	"github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/keccak256"
)

var (
	ReadRpcUrlFlag = cli.StringFlag{
		Name:     "read-rpc-url",
		Usage:    "The RPC URL for reading from the TableCalculator",
		Required: true,
	}
	WriteRpcUrlFlag = cli.StringFlag{
		Name:     "write-rpc-url",
		Usage:    "The RPC URL for writing to the CertificateVerifier",
		Required: true,
	}
	TableCalculatorAddrFlag = cli.StringFlag{
		Name:     "table-calculator-addr",
		Usage:    "Address of the BLSTableCalculator contract",
		Required: true,
	}
	CertificateVerifierAddrFlag = cli.StringFlag{
		Name:     "certificate-verifier-addr",
		Usage:    "Address of the BLSCertificateVerifier contract",
		Required: true,
	}
	AvsAddressFlag = cli.StringFlag{
		Name:     "avs-address",
		Usage:    "Address of the AVS",
		Required: true,
	}
	OperatorSetIdFlag = cli.IntFlag{
		Name:     "operator-set-id",
		Usage:    "ID of the operator set to use",
		Required: true,
	}
)

// TODO: make common package??
func convertOperatorSetInfo(from tableCalculator.IBLSTableCalculatorTypesBN254OperatorSetInfo) certificateVerifier.IBLSTableCalculatorTypesBN254OperatorSetInfo {
	// Create a new G1Point with the same values
	g1Point := certificateVerifier.BN254G1Point{
		X: new(big.Int).Set(from.AggregatePubkey.X),
		Y: new(big.Int).Set(from.AggregatePubkey.Y),
	}

	// Copy the total weights
	totalWeights := make([]*big.Int, len(from.TotalWeights))
	for i, weight := range from.TotalWeights {
		totalWeights[i] = new(big.Int).Set(weight)
	}

	// Create and return the converted struct
	return certificateVerifier.IBLSTableCalculatorTypesBN254OperatorSetInfo{
		NumOperators:    new(big.Int).Set(from.NumOperators),
		AggregatePubkey: g1Point,
		TotalWeights:    totalWeights,
	}
}

// Calculates operator table from the table calculator
func calculateOperatorTable(
	ctx context.Context,
	logger logging.Logger,
	tableCalc *tableCalculator.BLSTableCalculator,
	operatorSet tableCalculator.OperatorSet,
	blockNumber *big.Int,
) (tableCalculator.IBLSTableCalculatorTypesBN254OperatorSetInfo, error) {
	logger.Info("Calculating operator table", "operatorSetID", operatorSet.Id, "blockNumber", blockNumber)

	operatorSetInfo, err := tableCalc.CalculateOperatorTable(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, operatorSet)

	if err != nil {
		logger.Error("Failed to calculate operator table", "error", err)
		return operatorSetInfo, err
	}

	logger.Info("Operator table calculated", "numOperators", operatorSetInfo.NumOperators)
	return operatorSetInfo, nil
}

// Gets operator infos from the table calculator
func getOperatorInfos(
	ctx context.Context,
	logger logging.Logger,
	tableCalc *tableCalculator.BLSTableCalculator,
	operatorSet tableCalculator.OperatorSet,
	blockNumber *big.Int,
) ([]tableCalculator.IBLSTableCalculatorTypesBN254OperatorInfo, error) {
	logger.Info("Getting operator infos", "operatorSetID", operatorSet.Id, "blockNumber", blockNumber)

	operatorInfos, err := tableCalc.GetOperatorInfos(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, operatorSet)

	if err != nil {
		logger.Error("Failed to get operator infos", "error", err)
		return nil, err
	}

	logger.Info("Retrieved operator infos", "count", len(operatorInfos))
	return operatorInfos, nil
}

// Calculates merkle root from operator infos
func calculateOperatorInfoTree(
	logger logging.Logger,
	operatorInfos []tableCalculator.IBLSTableCalculatorTypesBN254OperatorInfo,
) (*merkletree.MerkleTree, [32]byte, error) {
	// Generate leaf nodes
	operatorInfoLeaves := make([][]byte, 0)

	for _, operatorInfo := range operatorInfos {
		operatorInfoLeaves = append(
			operatorInfoLeaves,
			EncodeOperatorInfoLeaf(operatorInfo),
		)
	}

	operatorInfoTree, err := merkletree.NewTree(
		merkletree.WithData(operatorInfoLeaves),
		merkletree.WithHashType(keccak256.New()),
	)
	if err != nil {
		logger.Error("Failed to create merkle tree", "error", err)
		return nil, [32]byte{}, err
	}

	// Get the root and convert to [32]byte
	rootBytes := operatorInfoTree.Root()
	logger.Info("Merkle root", "root", gethcommon.Bytes2Hex(rootBytes))
	var merkleRoot [32]byte
	copy(merkleRoot[:], rootBytes)

	return operatorInfoTree, merkleRoot, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "transporter"
	app.Usage = "does the transporting"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		&ReadRpcUrlFlag,
		&WriteRpcUrlFlag,
		&utils.EcdsaPrivateKeyFlag,
		&TableCalculatorAddrFlag,
		&CertificateVerifierAddrFlag,
		&AvsAddressFlag,
		&OperatorSetIdFlag,
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// TODO: turn this into a process that can be run in the background at a given interval
func start(c *cli.Context) error {
	ctx := context.Background()

	// Metrics/Logging
	logger := logging.NewTextSLogger(os.Stdout, nil)

	reg := prometheus.NewRegistry()
	rpcReadCallsCollector := rpccalls.NewCollector("transporterRead", reg)
	rpcWriteCallsCollector := rpccalls.NewCollector("transporterWrite", reg)

	// Initialize RPCs
	readClient, err := eth.NewInstrumentedClient(c.String(ReadRpcUrlFlag.Name), rpcReadCallsCollector)
	if err != nil {
		return err
	}

	writeClient, err := eth.NewInstrumentedClient(c.String(WriteRpcUrlFlag.Name), rpcWriteCallsCollector)
	if err != nil {
		return err
	}

	// Setup writes
	chainID, err := writeClient.ChainID(ctx)
	if err != nil {
		return err
	}

	privateKeyString := c.String(utils.EcdsaPrivateKeyFlag.Name)
	if privateKeyString[0:2] == "0x" {
		privateKeyString = privateKeyString[2:]
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return err
	}

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainID)
	if err != nil {
		return err
	}

	// Create wallet and transaction manager
	pkWallet, err := wallet.NewPrivateKeyWallet(writeClient, signerV2, addr, logger)
	if err != nil {
		return err
	}

	txManager := txmgr.NewSimpleTxManager(
		pkWallet,
		writeClient,
		logger,
		addr,
	)

	// Initialize addresses & operatorSet Info
	tableCalculatorAddr := gethcommon.HexToAddress(c.String(TableCalculatorAddrFlag.Name))
	certificateVerifierAddr := gethcommon.HexToAddress(c.String(CertificateVerifierAddrFlag.Name))

	avsAddress := gethcommon.HexToAddress(c.String(AvsAddressFlag.Name))

	operatorSetID := c.Int(OperatorSetIdFlag.Name)
	operatorSetIDUint32 := uint32(operatorSetID)

	operatorSet := tableCalculator.OperatorSet{
		Avs: avsAddress,
		Id:  operatorSetIDUint32,
	}

	tableCalc, err := tableCalculator.NewBLSTableCalculator(tableCalculatorAddr, readClient)
	if err != nil {
		return err
	}

	certVerifier, err := certificateVerifier.NewBLSCertificateVerifier(certificateVerifierAddr, writeClient)
	if err != nil {
		return err
	}

	// Update table
	err = updateOperatorTable(ctx, logger, readClient, tableCalc, certVerifier, txManager, operatorSet)
	if err != nil {
		logger.Error("Failed to update operator table", "error", err)
		return err
	}

	logger.Info("Successfully completed operator table update")
	return nil
}

func updateOperatorTable(
	ctx context.Context,
	logger logging.Logger,
	ethClient ethereum.ChainReader,
	tableCalc *tableCalculator.BLSTableCalculator,
	certVerifier *certificateVerifier.BLSCertificateVerifier,
	txManager *txmgr.SimpleTxManager,
	operatorSet tableCalculator.OperatorSet,
) error {
	// Get the operator table
	logger.Info("Reading operator table from calculator", "operatorSetID", operatorSet.Id)

	// Get the current block information
	header, err := ethClient.HeaderByNumber(ctx, nil) // nil means latest block
	if err != nil {
		logger.Error("Failed to get latest block header", "error", err)
		return err
	}

	currentBlockNumber := header.Number
	blockTimestamp := time.Unix(int64(header.Time), 0)

	logger.Info("Querying operator table",
		"blockNumber", currentBlockNumber,
		"blockTimestamp", blockTimestamp)

	// Calculate operator table
	operatorSetInfo, err := calculateOperatorTable(ctx, logger, tableCalc, operatorSet, currentBlockNumber)
	if err != nil {
		return err
	}

	// Get operator infos
	operatorInfos, err := getOperatorInfos(ctx, logger, tableCalc, operatorSet, currentBlockNumber)
	if err != nil {
		logger.Error("Failed to get operator infos", "error", err)
		return err
	}

	// Calculate operator info tree
	_, merkleRoot, err := calculateOperatorInfoTree(logger, operatorInfos)
	if err != nil {
		logger.Error("Failed to calculate operator infos merkle root", "error", err)
		return err
	}

	// Prepare transaction to update operator table
	txOpts, err := txManager.GetNoSendTxOpts()
	if err != nil {
		logger.Error("Failed to get tx opts", "error", err)
		return err
	}

	// Convert the operatorSetInfo to the correct type
	convertedInfo := convertOperatorSetInfo(operatorSetInfo)

	// Call updateOperatorTable
	// TODO: Should also be adding the blockNumber
	logger.Info("Updating operator table on verifier",
		"timestamp", blockTimestamp,
		"merkleRoot", gethcommon.Bytes2Hex(merkleRoot[:]))
	tx, err := certVerifier.UpdateOperatorTable(
		txOpts,
		uint32(blockTimestamp.Unix()),
		convertedInfo,
		merkleRoot,
	)
	if err != nil {
		logger.Error("Failed to update operator table", "error", err)
		return err
	}

	// Send the transaction and wait for receipt
	receipt, err := txManager.Send(ctx, tx, true)
	if err != nil {
		logger.Error("Failed to send update operator table tx", "tx", tx.Hash().Hex(), "error", err)
		return err
	}

	logger.Info("Successfully updated operator table", "txHash", receipt.TxHash.Hex())
	return nil
}

// encodes an operator info as a leaf node using Solidity's abi.encodePacked
// Format matches abi.encodePacked(G1Point.X, G1Point.Y, weights)
func EncodeOperatorInfoLeaf(operatorInfo tableCalculator.IBLSTableCalculatorTypesBN254OperatorInfo) []byte {
	// Get minimal byte representation of X and Y coordinates
	xBytes := operatorInfo.Pubkey.X.Bytes()
	yBytes := operatorInfo.Pubkey.Y.Bytes()

	result := make([]byte, 0, len(xBytes)+len(yBytes))
	result = append(result, xBytes...)
	result = append(result, yBytes...)

	// Encode each weight with minimal bytes
	for _, weight := range operatorInfo.Weights {
		// We pad to 32 bytes since abi.encodePacked doesn't pack arrays
		weightBytes := weight.FillBytes(make([]byte, 32))
		result = append(result, weightBytes...)
	}

	return result
}
