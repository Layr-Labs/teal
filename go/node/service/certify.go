package service

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/Layr-Labs/teal/go/api/service/v1"
)

type CertifyingService struct {
	keyPair     *bls.KeyPair
	getResponse func(data []byte) ([]byte, error)

	v1.UnsafeNodeServiceServer
}

func NewCertifyingService(
	kp *bls.KeyPair,
	getResponse func(data []byte) ([]byte, error),
) *CertifyingService {
	return &CertifyingService{
		keyPair:     kp,
		getResponse: getResponse,
	}
}

func (s *CertifyingService) Certify(ctx context.Context, req *v1.CertifyRequest) (*v1.CertifyResponse, error) {
	response, err := s.getResponse(req.Data)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "data is invalid: %v", err)
	}

	digest := crypto.Keccak256(response)
	digestBytes := [32]byte(digest)

	signature := s.keyPair.SignMessage(digestBytes)
	signatureBytes := signature.Marshal()

	return &v1.CertifyResponse{Signature: signatureBytes[:], Data: response}, nil
}
