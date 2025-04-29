package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	v1 "github.com/Layr-Labs/teal/go/api/service/v1"
	"github.com/Layr-Labs/teal/go/node/service"
)

type Config struct {
	ServicePort int
	BlsKeyPair  *bls.KeyPair
}
type Certifier interface {
	GetResponse(config Config, data []byte) ([]byte, error)
}

type BaseNode struct {
	config    Config
	certifier Certifier
}

type Node interface {
	Certifier
	Start() error
}

// NewBaseNode creates a new base node implementation
func NewBaseNode(config Config, certifier Certifier) *BaseNode {
	return &BaseNode{
		config:    config,
		certifier: certifier,
	}
}

// Start implements the Node interface
func (n *BaseNode) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", n.config.ServicePort))
	if err != nil {
		return err
	}
	return n.StartWithListener(lis)
}

func (n *BaseNode) StartWithListener(lis net.Listener) error {
	grpcServer := grpc.NewServer()

	// Create a closure that captures the config for validation
	getResponse := func(data []byte) ([]byte, error) {
		return n.certifier.GetResponse(n.config, data)
	}

	v1.RegisterNodeServiceServer(grpcServer, service.NewCertifyingService(
		n.config.BlsKeyPair,
		getResponse,
	))

	reflection.Register(grpcServer)

	log.Printf("Starting server on port %d", n.config.ServicePort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve on port %d: %v", n.config.ServicePort, err)
		return err
	}
	return nil
}
