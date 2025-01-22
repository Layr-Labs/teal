package utils

import (
	"encoding/json"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/ethereum/go-ethereum/common"
)

type EigenLayerDeployment struct {
	DelegationManager     common.Address `json:"delegationManager"`
	AvsDirectory          common.Address `json:"avsDirectory"`
	RewardsCoordinator    common.Address `json:"rewardsCoordinator"`
	PermissionsController common.Address `json:"permissionsController"`
}

type AVSDeployment struct {
	DeploymentBlock        uint64         `json:"deploymentBlock"`
	CertificateVerifier    common.Address `json:"certificateVerifier"`
	RegistryCoordinator    common.Address `json:"registryCoordinator"`
	OperatorStateRetriever common.Address `json:"operatorStateRetriever"`
}

func ReadEigenlayerDeployment(path string) (elcontracts.Config, error) {
	// read the json file
	jsonFile, err := os.Open(path)
	if err != nil {
		return elcontracts.Config{}, err
	}
	defer jsonFile.Close()

	// parse the json file
	var deployment EigenLayerDeployment
	err = json.NewDecoder(jsonFile).Decode(&deployment)
	if err != nil {
		return elcontracts.Config{}, err
	}
	return elcontracts.Config{
		DelegationManagerAddress:     deployment.DelegationManager,
		AvsDirectoryAddress:          deployment.AvsDirectory,
		RewardsCoordinatorAddress:    deployment.RewardsCoordinator,
		PermissionsControllerAddress: deployment.PermissionsController,
	}, nil
}

func ReadAVSDeployment(path string) (AVSDeployment, error) {
	// read the json file
	jsonFile, err := os.Open(path)
	if err != nil {
		return AVSDeployment{}, err
	}
	defer jsonFile.Close()

	// parse the json file
	var deployment AVSDeployment
	err = json.NewDecoder(jsonFile).Decode(&deployment)
	if err != nil {
		return AVSDeployment{}, err
	}
	return deployment, nil
}

func (d AVSDeployment) ToConfig() avsregistry.Config {
	return avsregistry.Config{
		RegistryCoordinatorAddress:    d.RegistryCoordinator,
		OperatorStateRetrieverAddress: d.OperatorStateRetriever,
	}
}
