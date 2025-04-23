package utils

import (
	"github.com/urfave/cli/v2"
)

var (
	EthUrlFlag = cli.StringFlag{
		Name:     "eth-url",
		Usage:    "The URL of the Ethereum node",
		Value:    "",
		Required: true,
	}
	EigenlayerDeploymentPathFlag = cli.StringFlag{
		Name:     "eigenlayer-deployment-path",
		Usage:    "The path to the eigenlayer deployment",
		Value:    "",
		Required: true,
	}
	AvsDeploymentPathFlag = cli.StringFlag{
		Name:     "avs-deployment-path",
		Usage:    "The path to the avs deployment",
		Value:    "",
		Required: true,
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "The private key to use for the node",
		Value:    "",
		Required: true,
	}
)
