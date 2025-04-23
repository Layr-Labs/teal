package aggregator_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/teal/go/aggregator"
	mockOperatorRequester "github.com/Layr-Labs/teal/go/aggregator/operator_requester/mocks"
	pb "github.com/Layr-Labs/teal/go/api/service/v1"
	"github.com/Layr-Labs/teal/go/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAggregatorService(t *testing.T) {
	ctrl := gomock.NewController(t)

	fakeOperatorRequester := mockOperatorRequester.NewMockOperatorRequester(ctrl)

	t.Run("successful certification with single operator", func(t *testing.T) {
		ctx := context.Background()

		testOperator1 := types.TestOperator{
			OperatorId: types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{
				0: big.NewInt(100),
			},
			BlsKeypair: newBlsKeyPairPanics("0x1"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumber := types.QuorumNum(0)
		quorumThresholdPercentage := types.QuorumThresholdPercentage(100)
		requestData := []byte("test 1")

		logger := testutils.GetTestLogger()
		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		operators, _ := fakeAvsRegistryService.GetOperatorsAvsStateAtBlock(ctx, types.QuorumNums{quorumNumber}, blockNum)

		for _, operator := range operators {
			responseData := []byte("test 1")
			taskResponseDigest, _ := common.Keccak256HashFn(responseData)
			fakeOperatorRequester.EXPECT().RequestCertification(ctx, operator, taskIndex, requestData).Return(&pb.CertifyResponse{
				Signature: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).Marshal(),
				Data:      responseData,
			}, nil)
		}

		blsAggService := blsagg.NewBlsAggregatorService(fakeAvsRegistryService, common.Keccak256HashFn, logger)

		// Create aggregator service
		aggregatorService := aggregator.NewAggregatorService(
			logger,
			fakeAvsRegistryService,
			blsAggService,
			fakeOperatorRequester,
		)

		resp, err := aggregatorService.GetCertificate(
			ctx,
			taskIndex,
			blockNum,
			quorumNumber,
			quorumThresholdPercentage,
			requestData,
			1*time.Second,
		)

		// Verify results
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, taskIndex, resp.TaskIndex)
	})

}

func newBlsKeyPairPanics(hexKey string) *bls.KeyPair {
	keypair, err := bls.NewKeyPairFromString(hexKey)
	if err != nil {
		panic(err)
	}
	return keypair
}
