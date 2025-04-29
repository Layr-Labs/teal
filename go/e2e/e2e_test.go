package e2e_test

import (
	"context"
	"log/slog"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/teal/go/aggregator"
	operatorrequester "github.com/Layr-Labs/teal/go/aggregator/operator_requester"
	"github.com/Layr-Labs/teal/go/common"
	"github.com/Layr-Labs/teal/go/node/server"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	e2e "github.com/Layr-Labs/teal/go/e2e"
)

var (
	ANVIL_FIRST_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

func TestIntegrationBlsAgg(t *testing.T) {

	tasksTimeToExpiry := 10 * time.Second

	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)
	t.Run("1 quorums 1 operator", func(t *testing.T) {
		// read input from JSON if available, otherwise use default values
		var defaultInput = struct {
			QuorumNumbers              types.QuorumNums                 `json:"quorum_numbers"`
			QuorumThresholdPercentages types.QuorumThresholdPercentages `json:"quorum_threshold_percentages"`
			BlsPrivKey                 string                           `json:"bls_key"`
		}{
			QuorumNumbers:              types.QuorumNums{0},
			QuorumThresholdPercentages: types.QuorumThresholdPercentages{100},
			BlsPrivKey:                 "0x1",
		}
		testData := testutils.NewTestData(defaultInput)

		// define operator ecdsa and bls private keys
		ecdsaPrivKey, err := crypto.HexToECDSA(ANVIL_FIRST_PRIVATE_KEY)
		require.NoError(t, err)
		blsPrivKeyHex := testData.Input.BlsPrivKey
		blsKeyPair := newBlsKeyPairPanics(blsPrivKeyHex)
		// operatorId := types.OperatorIdFromG1Pubkey(blsKeyPair.GetPubKeyG1())

		// create avs clients to interact with contracts deployed on anvil
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
		avsClients, err := clients.BuildAll(clients.BuildAllConfig{
			EthHttpUrl:                 anvilHttpEndpoint,
			EthWsUrl:                   anvilWsEndpoint, // not used so doesn't matter that we pass an http url
			RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
			OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
			AvsName:                    "avs",
			PromMetricsIpPortAddress:   "localhost:9090",
			ServiceManagerAddress:      contractAddrs.ServiceManager.String(),
		}, ecdsaPrivKey, logger)
		require.NoError(t, err)
		avsWriter := avsClients.AvsRegistryChainWriter
		// avsServiceManager, err := avssm.NewContractMockAvsServiceManager(contractAddrs.ServiceManager, ethHttpClient)
		require.NoError(t, err)

		// create aggregation service
		operatorsInfoService := operatorsinfo.NewOperatorsInfoServiceInMemory(
			context.TODO(),
			avsClients.AvsRegistryChainSubscriber,
			avsClients.AvsRegistryChainReader,
			nil,
			operatorsinfo.Opts{},
			logger,
		)
		avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(
			avsClients.AvsRegistryChainReader,
			operatorsInfoService,
			logger,
		)
		blsAggService := blsagg.NewBlsAggregatorService(avsRegistryService, common.Keccak256HashFn, logger)

		// register operator
		quorumNumbers := testData.Input.QuorumNumbers
		_, err = avsWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivKey,
			blsKeyPair,
			quorumNumbers,
			"localhost:8080",
			true,
		)
		require.NoError(t, err)

		evenLovingNode := e2e.NewEvenLovingNode(server.Config{
			ServicePort: 8080,
			BlsKeyPair:  blsKeyPair,
		})
		go evenLovingNode.Start()

		// create the task related parameters: RBN, quorumThresholdPercentages, taskIndex and taskResponse
		curBlockNum, err := ethHttpClient.BlockNumber(context.Background())
		require.NoError(t, err)
		referenceBlockNumber := uint32(curBlockNum)
		// need to advance chain by 1 block because of the check in signatureChecker where RBN must be < current block
		// number
		testutils.AdvanceChainByNBlocksExecInContainer(context.TODO(), 1, anvilC)
		taskIndex := types.TaskIndex(0)
		// taskResponse := mockTaskResponse{123} // Initialize with appropriate data
		quorumThresholdPercentages := testData.Input.QuorumThresholdPercentages

		aggregator := aggregator.NewAggregatorService(
			logger,
			avsRegistryService,
			blsAggService,
			operatorrequester.NewOperatorRequester(logger),
		)

		_, err = aggregator.GetCertificate(
			context.Background(),
			taskIndex,
			referenceBlockNumber,
			quorumNumbers[0],
			quorumThresholdPercentages[0],
			big.NewInt(69420).Bytes(),
			tasksTimeToExpiry,
		)
		require.NoError(t, err)
	})
}

func newBlsKeyPairPanics(hexKey string) *bls.KeyPair {
	keypair, err := bls.NewKeyPairFromString(hexKey)
	if err != nil {
		panic(err)
	}
	return keypair
}
