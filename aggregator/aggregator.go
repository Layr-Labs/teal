package aggregator

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/types"
	operatorrequester "github.com/Layr-Labs/teal/aggregator/operator_requester"
)

type AggregatorService struct {
	logger            logging.Logger
	avsRegistryReader avsregistry.AvsRegistryService
	blsAggService     blsagg.BlsAggregationService
	operatorRequester operatorrequester.OperatorRequester

	mu sync.Mutex
}

func NewAggregatorService(
	logger logging.Logger,
	avsRegistryReader avsregistry.AvsRegistryService,
	blsAggService blsagg.BlsAggregationService,
	operatorRequester operatorrequester.OperatorRequester,
) *AggregatorService {
	return &AggregatorService{
		logger:            logger,
		avsRegistryReader: avsRegistryReader,
		blsAggService:     blsAggService,
		operatorRequester: operatorRequester,
	}
}

// GetCertificate sends a task to all registered nodes and aggregates their responses
// Only works for single quorum for simplicity
func (s *AggregatorService) GetCertificate(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumber types.QuorumNum,
	quorumThresholdPercentage types.QuorumThresholdPercentage,
	data []byte,
	timeToExpiry time.Duration,
) (*blsagg.BlsAggregationServiceResponse, error) {
	// Only allow one task at a time
	s.mu.Lock()
	defer s.mu.Unlock()

	quorumNumbers := types.QuorumNums{quorumNumber}
	quorumThresholdPercentages := types.QuorumThresholdPercentages{quorumThresholdPercentage}

	// Initialize task in BLS aggregation service
	err := s.blsAggService.InitializeNewTaskWithWindow(
		taskIndex,
		taskCreatedBlock,
		quorumNumbers,
		quorumThresholdPercentages,
		timeToExpiry,
		1*time.Second,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize task: %w", err)
	}

	// Get operators from registry
	operators, err := s.avsRegistryReader.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, taskCreatedBlock)
	if err != nil {
		return nil, fmt.Errorf("failed to get operators: %w", err)
	}
	s.logger.Info("Number of operators registered", "numOperators", len(operators))

	// Send task to all operators in parallel
	for operatorId, operator := range operators {
		go func(operatorId types.OperatorId, operator types.OperatorAvsState) {
			s.logger.Info("Requesting certification from operator", "operatorId", operatorId, "socket", operator.OperatorInfo.Socket)
			// Create connection for this operator
			resp, err := s.operatorRequester.RequestCertification(ctx, operator, taskIndex, data)
			if err != nil {
				return
			}

			signature := &bls.Signature{G1Point: bls.NewG1Point(big.NewInt(0), big.NewInt(0))}
			_, err = signature.SetBytes(resp.Signature)
			if err != nil {
				s.logger.Error("Failed to unmarshal signature",
					"operatorId", operatorId,
					"error", err)
				return
			}

			s.logger.Info("Received signature from operator", "operatorId", operatorId)

			// Process signature from node
			err = s.blsAggService.ProcessNewSignature(
				ctx,
				taskIndex,
				types.TaskResponse(resp.Data),
				signature,
				operatorId,
			)
			if err != nil {
				s.logger.Error("Failed to process signature",
					"operatorId", operatorId,
					"error", err)
				return
			}
			s.logger.Info("Processed signature from operator", "operatorId", operatorId)
		}(operatorId, operator)
	}

	// Wait for aggregated response
	select {
	case resp := <-s.blsAggService.GetResponseChannel():
		if resp.Err != nil {
			return nil, fmt.Errorf("aggregation failed: %w", resp.Err)
		}
		return &resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
