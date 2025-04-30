package main

import (
	"context"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	rpccalls "github.com/Layr-Labs/eigensdk-go/metrics/collectors/rpc_calls"
	avsservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/signerv2"
	"github.com/Layr-Labs/eigensdk-go/types"
	BLSCertificateVerifier "github.com/Layr-Labs/teal/transporter/bindings/BLSCertificateVerifier"
	"github.com/ethereum/go-ethereum"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/teal/aggregator"
	operatorrequester "github.com/Layr-Labs/teal/aggregator/operator_requester"
	"github.com/Layr-Labs/teal/common"
	crossChainUtils "github.com/Layr-Labs/teal/crossChainExample/utils"
	"github.com/Layr-Labs/teal/example/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/urfave/cli/v2"
)

var (
	strategyAddress = gethcommon.HexToAddress("0x7d704507b76571a51d9cae8addabbfd0ba0e63d3")
)

var (
	ReadRpcUrlFlag = cli.StringFlag{
		Name:     "source-rpc-url",
		Usage:    "WSS URL on source chain",
		Required: true,
	}
	WriteRpcUrlFlag = cli.StringFlag{
		Name:     "destination-rpc-url",
		Usage:    "The RPC URL to verify certificates on",
		Required: true,
	}
	CertVerifierDeploymentPathFlag = cli.StringFlag{
		Name:     "avs-cert-verifier-deployment-path",
		Usage:    "The path to the certificate verifier deployment file",
		Required: true,
	}
	OperatorSetIdFlag = cli.IntFlag{
		Name:     "operator-set-id",
		Usage:    "ID of the operator set to use",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "aggregator"
	app.Usage = "abc"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		&utils.EcdsaPrivateKeyFlag,
		&utils.AvsDeploymentPathFlag,
		&ReadRpcUrlFlag,
		&WriteRpcUrlFlag,
		&CertVerifierDeploymentPathFlag,
		&OperatorSetIdFlag,
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	ctx := context.Background()

	reg := prometheus.NewRegistry()
	rpcCallsCollector := rpccalls.NewCollector("exampleAvs", reg)

	readClient, err := eth.NewInstrumentedClient(c.String(ReadRpcUrlFlag.Name), rpcCallsCollector)
	if err != nil {
		panic(err)
	}

	writeClient, err := eth.NewInstrumentedClient(c.String(WriteRpcUrlFlag.Name), rpcCallsCollector)
	if err != nil {
		panic(err)
	}

	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelInfo})

	chainid, err := writeClient.ChainID(ctx)
	if err != nil {
		panic(err)
	}

	privateKeyString := c.String(utils.EcdsaPrivateKeyFlag.Name)
	// Print private key
	if privateKeyString[0:2] == "0x" {
		privateKeyString = privateKeyString[2:]
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		panic(err)
	}

	signerV2, addr, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainid)
	if err != nil {
		panic(err)
	}

	pkWallet, err := wallet.NewPrivateKeyWallet(writeClient, signerV2, addr, logger)
	if err != nil {
		panic(err)
	}

	txManager := txmgr.NewSimpleTxManager(
		pkWallet,
		writeClient,
		logger,
		addr,
	)

	avsDeployment, err := utils.ReadAVSDeployment(c.String(utils.AvsDeploymentPathFlag.Name))
	if err != nil {
		panic(err)
	}

	logger.Info("AVS deployment", "avsDeployment", avsDeployment.ToConfig())

	avsReader, err := avsregistry.NewReaderFromConfig(
		avsDeployment.ToConfig(),
		readClient,
		logger,
	)
	if err != nil {
		panic(err)
	}

	avsSubscriber, err := avsregistry.NewSubscriberFromConfig(
		avsDeployment.ToConfig(),
		readClient,
		logger,
	)
	if err != nil {
		panic(err)
	}

	operatorInfoService := operatorsinfo.NewOperatorsInfoServiceInMemory(
		ctx,
		avsSubscriber,
		avsReader,
		nil,
		operatorsinfo.Opts{
			StartBlock: big.NewInt(int64(avsDeployment.DeploymentBlock)),
		},
		logger,
	)

	avsRegistryService := avsservice.NewAvsRegistryServiceChainCaller(
		avsReader,
		operatorInfoService,
		logger,
	)

	blsAggService := blsagg.NewBlsAggregatorService(
		avsRegistryService,
		common.Keccak256HashFn,
		logger,
	)

	aggregator := aggregator.NewAggregatorService(
		logger,
		avsRegistryService,
		blsAggService,
		operatorrequester.NewOperatorRequester(logger),
	)

	certificateVerifierAddr, err := utils.ReadCertificateVerifierDeployment(c.String(CertVerifierDeploymentPathFlag.Name), chainid)
	if err != nil {
		panic(err)
	}

	certVerifier, err := BLSCertificateVerifier.NewBLSCertificateVerifier(
		certificateVerifierAddr,
		writeClient,
	)
	if err != nil {
		panic(err)
	}

	// TODO: update to a sane threshold, right now we assume 100% signing rate so we don't have to deal with merkle proofs
	threshold := big.NewInt(1e18)
	denominator := big.NewInt(1e18)
	quorumThreshold := types.QuorumThresholdPercentage(uint8(new(big.Int).Div(new(big.Int).Mul(threshold, big.NewInt(100)), denominator).Uint64()))

	quorumNumber := types.QuorumNum(0)
	requestNumber := uint32(0)

	// on a ticker every 30s, get the certificate
	for {
		func() {
			defer time.Sleep(5 * time.Second)
			requestNumber++

			currentBlockNumber, err := readClient.BlockNumber(ctx)
			if err != nil {
				logger.Error("Failed to get current block number", "error", err)
				return
			}
			referenceBlockNumber := uint32(currentBlockNumber - 5)

			callBlockNumber := currentBlockNumber - 150
			callMsg := ethereum.CallMsg{
				From:     gethcommon.HexToAddress("0x4242424242424242424242424242424242424242"),
				To:       &strategyAddress,
				Gas:      1000000,
				GasPrice: big.NewInt(10000),
				Value:    big.NewInt(0),
				// totalShares()
				Data: []byte{0x3a, 0x98, 0xef, 0x39},
			}
			request := utils.CallToBytes(uint64(callBlockNumber), callMsg)

			logger.Info("Requesting certificate of totalShares()", "callBlockNumber", callBlockNumber)

			resp, err := aggregator.GetCertificate(
				ctx,
				types.TaskIndex(requestNumber),
				uint32(referenceBlockNumber),
				quorumNumber,
				quorumThreshold,
				request,
				25*time.Second,
			)
			if err != nil {
				logger.Error("Failed to get certificate", "error", err)
				return
			}

			txOpts, err := txManager.GetNoSendTxOpts()
			if err != nil {
				logger.Error("Failed to get tx opts", "error", err)
				return
			}

			// TODO: query this against latest stake table update
			referenceTimestamp := uint32(1745965200)

			tx, err := certVerifier.VerifyCertificate(
				txOpts,
				BLSCertificateVerifier.IBLSCertificateVerifierTypesBN254Certificate{
					ReferenceTimestamp: referenceTimestamp,
					MessageHash:        resp.TaskResponseDigest,
					Signature:          crossChainUtils.ToBN254G1Point(resp.SignersAggSigG1.G1Point),
					Apk:                crossChainUtils.ToBN254G2Point(resp.SignersApkG2),
					NonSignerIndices:   []uint32{},
					NonSignerWitnesses: []BLSCertificateVerifier.IBLSCertificateVerifierTypesBN254OperatorInfoWitness{},
				},
			)
			if err != nil {
				logger.Error("Failed to assemble verify certificate tx", "error", err)
				return
			}

			receipt, err := txManager.Send(ctx, tx, true)
			if err != nil {
				logger.Error("Failed to send verify certificate tx", "tx", tx.Hash().Hex(), "error", err)
				return
			}

			logger.Info("Sent verify certificate tx", "tx", receipt.TxHash.Hex(), "requestNumber", requestNumber)
		}()
	}
}
