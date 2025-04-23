package operatorrequester

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	pb "github.com/Layr-Labs/teal/go/api/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OperatorRequester interface {
	RequestCertification(ctx context.Context, operator types.OperatorAvsState, taskIndex types.TaskIndex, requestData []byte) (*pb.CertifyResponse, error)
}

type operatorRequester struct {
	logger logging.Logger
}

func NewOperatorRequester(logger logging.Logger) OperatorRequester {
	return &operatorRequester{
		logger: logger,
	}
}

func (or *operatorRequester) RequestCertification(ctx context.Context, operator types.OperatorAvsState, taskIndex types.TaskIndex, requestData []byte) (*pb.CertifyResponse, error) {
	conn, err := grpc.NewClient(
		operator.OperatorInfo.Socket.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		or.logger.Error("Failed to connect to operator",
			"operatorId", operator.OperatorId,
			"socket", operator.OperatorInfo.Socket,
			"error", err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewNodeServiceClient(conn)

	// Send task to node
	resp, err := client.Certify(ctx, &pb.CertifyRequest{
		TaskIndex: uint32(taskIndex),
		Data:      requestData,
	})
	if err != nil {
		or.logger.Error("Failed to send task to node",
			"operatorId", operator.OperatorId,
			"error", err)
		return nil, err
	}

	return resp, nil
}
