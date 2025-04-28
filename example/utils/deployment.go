package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/ethereum/go-ethereum/common"
)

type EigenLayerDeployment struct {
	DelegationManager    common.Address `json:"delegationManager"`
	AvsDirectory         common.Address `json:"avsDirectory"`
	RewardsCoordinator   common.Address `json:"rewardsCoordinator"`
	PermissionController common.Address `json:"permissionController"`
	AllocationManager    common.Address `json:"allocationManager"`
}

type AVSDeployment struct {
	DeploymentBlock             uint64           `json:"deploymentBlock"`
	CertificateVerifier         common.Address   `json:"certificateVerifier"`
	TableCalculator             common.Address   `json:"tableCalculator"`
	SlashingRegistryCoordinator common.Address   `json:"slashingRegistryCoordinator"`
	OperatorStateRetriever      common.Address   `json:"operatorStateRetriever"`
	ServiceManager              common.Address   `json:"serviceManager"`
	Strategies                  []common.Address `json:"strategies"`
}

type CertificateVerifierDeployment struct {
	CertificateVerifier common.Address `json:"certificateVerifier"`
	ChainID             uint64         `json:"chainID"`
}

type CertificateVerifierDeployments struct {
	CrossChainDeployments []CertificateVerifierDeployment `json:"crossChainDeployments"`
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
		PermissionsControllerAddress: deployment.PermissionController,
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

func ReadCertificateVerifierDeployment(path string, destinationChainID *big.Int) (common.Address, error) {
	// Read the JSON file
	jsonFile, err := os.Open(path)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to open certificate verifier deployment file: %w", err)
	}
	defer jsonFile.Close()

	// Parse the JSON
	var deployments CertificateVerifierDeployments
	err = json.NewDecoder(jsonFile).Decode(&deployments)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to decode certificate verifier deployment: %w", err)
	}

	// Find the deployment for the target chain ID
	for _, deployment := range deployments.CrossChainDeployments {
		if deployment.ChainID == destinationChainID.Uint64() {
			return deployment.CertificateVerifier, nil
		}
	}

	return common.Address{}, fmt.Errorf("no certificate verifier deployment found for chain ID %d", destinationChainID)
}

func (d AVSDeployment) ToConfig() avsregistry.Config {
	return avsregistry.Config{
		RegistryCoordinatorAddress:    d.SlashingRegistryCoordinator,
		OperatorStateRetrieverAddress: d.OperatorStateRetriever,
		ServiceManagerAddress:         d.ServiceManager,
	}
}
