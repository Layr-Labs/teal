package main

import (
	"log"
	"os"

	"github.com/Layr-Labs/teal/example/node"
	"github.com/Layr-Labs/teal/example/utils"
	"github.com/Layr-Labs/teal/node/server"
	"github.com/urfave/cli/v2"
)

var (
	ServicePortFlag = cli.IntFlag{
		Name:  "service-port",
		Usage: "The port to serve the service on",
		Value: 8080,
	}
	BlsPrivateKeyFlag = cli.StringFlag{
		Name:     "bls-private-key",
		Usage:    "The private key to use for the node",
		Value:    "",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "eth_call_node"
	app.Usage = "xyz"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		&utils.EthUrlFlag,
		&ServicePortFlag,
		&BlsPrivateKeyFlag,
	}

	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	keyPair := utils.NewBlsKeyPairPanics(c.String(BlsPrivateKeyFlag.Name))

	cfg := server.Config{
		ServicePort: c.Int(ServicePortFlag.Name),
		BlsKeyPair:  keyPair,
	}

	node := node.NewEthCallNode(cfg, c.String(utils.EthUrlFlag.Name))
	if err := node.Start(); err != nil {
		log.Fatal(err)
	}
	return nil
}
