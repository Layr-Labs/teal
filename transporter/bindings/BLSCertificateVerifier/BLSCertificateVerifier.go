// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BLSCertificateVerifier

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSCertificateVerifierTypesBN254Certificate is an auto generated low-level Go binding around an user-defined struct.
type IBLSCertificateVerifierTypesBN254Certificate struct {
	ReferenceTimestamp uint32
	MessageHash        [32]byte
	Signature          BN254G1Point
	Apk                BN254G2Point
	NonSignerIndices   []uint32
	NonSignerWitnesses []IBLSCertificateVerifierTypesBN254OperatorInfoWitness
}

// IBLSCertificateVerifierTypesBN254OperatorInfoWitness is an auto generated low-level Go binding around an user-defined struct.
type IBLSCertificateVerifierTypesBN254OperatorInfoWitness struct {
	OperatorIndex      uint32
	OperatorInfoProofs []byte
	OperatorInfo       IBLSTableCalculatorTypesBN254OperatorInfo
}

// IBLSTableCalculatorTypesBN254OperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBLSTableCalculatorTypesBN254OperatorInfo struct {
	Pubkey  BN254G1Point
	Weights []*big.Int
}

// IBLSTableCalculatorTypesBN254OperatorSetInfo is an auto generated low-level Go binding around an user-defined struct.
type IBLSTableCalculatorTypesBN254OperatorSetInfo struct {
	NumOperators    *big.Int
	AggregatePubkey BN254G1Point
	TotalWeights    []*big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// BLSCertificateVerifierMetaData contains all meta data concerning the BLSCertificateVerifier contract.
var BLSCertificateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initialOperatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxOperatorTableStaleness\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ejectOperators\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"witnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProofs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestReferenceTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxOperatorTableStaleness\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorTableUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxOperatorTableStaleness\",\"inputs\":[{\"name\":\"_maxOperatorTableStaleness\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorTableUpdater\",\"inputs\":[{\"name\":\"_operatorTableUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorTable\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyCertificate\",\"inputs\":[{\"name\":\"certificate\",\"type\":\"tuple\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProofs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateNominal\",\"inputs\":[{\"name\":\"certificate\",\"type\":\"tuple\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProofs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}]}]},{\"name\":\"totalStakeNominalThresholds\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCertificateProportion\",\"inputs\":[{\"name\":\"certificate\",\"type\":\"tuple\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254Certificate\",\"components\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"nonSignerIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerWitnesses\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSCertificateVerifierTypes.BN254OperatorInfoWitness[]\",\"components\":[{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorInfoProofs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorInfo\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}]}]},{\"name\":\"totalStakeProportionThresholds\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxOperatorTableStalenessSet\",\"inputs\":[{\"name\":\"maxOperatorTableStaleness\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorTableUpdaterSet\",\"inputs\":[{\"name\":\"operatorTableUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TableUpdated\",\"inputs\":[{\"name\":\"referenceTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"pubkey\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"operatorInfoTreeRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CertVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyTableUpdater\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TableStale\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610d92380380610d9283398101604081905261002e91610116565b61003733610099565b835160018054602087015163ffffffff908116600160a01b9081026001600160c01b03199384166001600160a01b0396871617179093556002805491871690930291169286169290921791909117905561009081610099565b505050506101ba565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146100fe575f5ffd5b919050565b805163ffffffff811681146100fe575f5ffd5b5f5f5f5f84860360a081121561012a575f5ffd5b6040811215610137575f5ffd5b50604080519081016001600160401b038111828210171561016657634e487b7160e01b5f52604160045260245ffd5b604052610172866100e8565b815261018060208701610103565b60208201529350610193604086016100e8565b92506101a160608601610103565b91506101af608086016100e8565b905092959194509250565b610bcb806101c75f395ff3fe608060405234801561000f575f5ffd5b50600436106100e5575f3560e01c8063a973342f11610088578063d8d201fe11610063578063d8d201fe146101f5578063ecf43fc114610203578063f2fde38b14610216578063fa2c481214610229575f5ffd5b8063a973342f146101a5578063b697a30b146101b8578063c89fbdbe146101cf575f5ffd5b80637192a81c116100c35780637192a81c146101365780637514bc3e146101565780638da5cb5b14610169578063a6ff593614610179575f5ffd5b80633ae5bbe0146100e957806368d6e081146100fe578063715018a61461012e575b5f5ffd5b6100fc6100f73660046106a6565b61027d565b005b600254610111906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100fc6102ce565b6101496101443660046106ea565b6102e1565b6040516101259190610723565b6100fc610164366004610786565b610338565b5f546001600160a01b0316610111565b60025461019090600160c01b900463ffffffff1681565b60405163ffffffff9091168152602001610125565b6100fc6101b33660046107e6565b61038f565b60025461019090600160a01b900463ffffffff1681565b6101e56101dd366004610914565b600192915050565b6040519015158152602001610125565b6101e56101dd366004610a5d565b6100fc610211366004610ac0565b6103c1565b6100fc6102243660046106a6565b6104b7565b6040805180820182525f808252602091820152815180830183526001546001600160a01b03811680835263ffffffff600160a01b909204821692840192835284519081529151169181019190915201610125565b610285610535565b600280546001600160a01b0319166001600160a01b0383169081179091556040517f64c64c9415569fff409712d78176f348185168935b776890b4649e4cd64a25b4905f90a250565b6102d6610535565b6102df5f61058e565b565b6040805160018082528183019092526060915f9190602080830190803683370190505090506064815f8151811061031a5761031a610b81565b6001600160601b039092166020928302919091019091015292915050565b610340610535565b6002805463ffffffff60a01b1916600160a01b63ffffffff8416908102919091179091556040517f72f9e8d81faad2896db70482f1cf2e509e4716d07edf38a36a28d31110a035f1905f90a250565b6002546001600160a01b031633146103ba5760405163030c1b6b60e11b815260040160405180910390fd5b5050505050565b6002546001600160a01b031633146103ec5760405163030c1b6b60e11b815260040160405180910390fd5b63ffffffff83165f908152600360208181526040928390208551815581860151805160018301558201516002820155928501518051869493610433939085019201906105dd565b50505063ffffffff83165f818152600460209081526040918290208490556002805463ffffffff60c01b1916600160c01b8502179055848101518251938452805184830152015190820152606081018290527f349eb659faf73a073d293b6e4ee24fcfda5191a5770f5bc0a315bc45b5f69ad29060800160405180910390a1505050565b6104bf610535565b6001600160a01b0381166105295760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6105328161058e565b50565b5f546001600160a01b031633146102df5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610520565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054828255905f5260205f2090600101600290048101928215610682579160200282015f5b8382111561064d57835183826101000a8154816001600160601b0302191690836001600160601b031602179055509260200192600c01602081600b01049283019260010302610603565b80156106805782816101000a8154906001600160601b030219169055600c01602081600b0104928301926001030261064d565b505b5061068e929150610692565b5090565b5b8082111561068e575f8155600101610693565b5f602082840312156106b6575f5ffd5b81356001600160a01b03811681146106cc575f5ffd5b9392505050565b5f61014082840312156106e4575f5ffd5b50919050565b5f602082840312156106fa575f5ffd5b81356001600160401b0381111561070f575f5ffd5b61071b848285016106d3565b949350505050565b602080825282518282018190525f918401906040840190835b818110156107635783516001600160601b031683526020938401939092019160010161073c565b509095945050505050565b803563ffffffff81168114610781575f5ffd5b919050565b5f60208284031215610796575f5ffd5b6106cc8261076e565b5f5f83601f8401126107af575f5ffd5b5081356001600160401b038111156107c5575f5ffd5b6020830191508360208260051b85010111156107df575f5ffd5b9250929050565b5f5f5f5f5f606086880312156107fa575f5ffd5b6108038661076e565b945060208601356001600160401b0381111561081d575f5ffd5b6108298882890161079f565b90955093505060408601356001600160401b03811115610847575f5ffd5b6108538882890161079f565b969995985093965092949392505050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b038111828210171561089a5761089a610864565b60405290565b604080519081016001600160401b038111828210171561089a5761089a610864565b604051601f8201601f191681016001600160401b03811182821017156108ea576108ea610864565b604052919050565b5f6001600160401b0382111561090a5761090a610864565b5060051b60200190565b5f5f60408385031215610925575f5ffd5b82356001600160401b0381111561093a575f5ffd5b610946858286016106d3565b92505060208301356001600160401b03811115610961575f5ffd5b8301601f81018513610971575f5ffd5b803561098461097f826108f2565b6108c2565b8082825260208201915060208360051b8501019250878311156109a5575f5ffd5b6020840193505b828410156109d657833561ffff811681146109c5575f5ffd5b8252602093840193909101906109ac565b809450505050509250929050565b5f82601f8301126109f3575f5ffd5b8135610a0161097f826108f2565b8082825260208201915060208360051b860101925085831115610a22575f5ffd5b602085015b83811015610a535780356001600160601b0381168114610a45575f5ffd5b835260209283019201610a27565b5095945050505050565b5f5f60408385031215610a6e575f5ffd5b82356001600160401b03811115610a83575f5ffd5b610a8f858286016106d3565b92505060208301356001600160401b03811115610aaa575f5ffd5b610ab6858286016109e4565b9150509250929050565b5f5f5f60608486031215610ad2575f5ffd5b610adb8461076e565b925060208401356001600160401b03811115610af5575f5ffd5b84018086036080811215610b07575f5ffd5b610b0f610878565b823581526040601f1983011215610b24575f5ffd5b610b2c6108a0565b6020848101358252604085013581830152820152606083013591506001600160401b03821115610b5a575f5ffd5b610b66888385016109e4565b60408281019190915295989097509590940135949350505050565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212206dc0e44d6935e82eae9d08d3991f84a01efe7d8ab17671470eb14c20664d09c364736f6c634300081b0033",
}

// BLSCertificateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSCertificateVerifierMetaData.ABI instead.
var BLSCertificateVerifierABI = BLSCertificateVerifierMetaData.ABI

// BLSCertificateVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BLSCertificateVerifierMetaData.Bin instead.
var BLSCertificateVerifierBin = BLSCertificateVerifierMetaData.Bin

// DeployBLSCertificateVerifier deploys a new Ethereum contract, binding an instance of BLSCertificateVerifier to it.
func DeployBLSCertificateVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _initialOperatorSet OperatorSet, _operatorTableUpdater common.Address, _maxOperatorTableStaleness uint32, owner common.Address) (common.Address, *types.Transaction, *BLSCertificateVerifier, error) {
	parsed, err := BLSCertificateVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BLSCertificateVerifierBin), backend, _initialOperatorSet, _operatorTableUpdater, _maxOperatorTableStaleness, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BLSCertificateVerifier{BLSCertificateVerifierCaller: BLSCertificateVerifierCaller{contract: contract}, BLSCertificateVerifierTransactor: BLSCertificateVerifierTransactor{contract: contract}, BLSCertificateVerifierFilterer: BLSCertificateVerifierFilterer{contract: contract}}, nil
}

// BLSCertificateVerifier is an auto generated Go binding around an Ethereum contract.
type BLSCertificateVerifier struct {
	BLSCertificateVerifierCaller     // Read-only binding to the contract
	BLSCertificateVerifierTransactor // Write-only binding to the contract
	BLSCertificateVerifierFilterer   // Log filterer for contract events
}

// BLSCertificateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSCertificateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSCertificateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSCertificateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSCertificateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSCertificateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSCertificateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSCertificateVerifierSession struct {
	Contract     *BLSCertificateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BLSCertificateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSCertificateVerifierCallerSession struct {
	Contract *BLSCertificateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BLSCertificateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSCertificateVerifierTransactorSession struct {
	Contract     *BLSCertificateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BLSCertificateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSCertificateVerifierRaw struct {
	Contract *BLSCertificateVerifier // Generic contract binding to access the raw methods on
}

// BLSCertificateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSCertificateVerifierCallerRaw struct {
	Contract *BLSCertificateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BLSCertificateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSCertificateVerifierTransactorRaw struct {
	Contract *BLSCertificateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSCertificateVerifier creates a new instance of BLSCertificateVerifier, bound to a specific deployed contract.
func NewBLSCertificateVerifier(address common.Address, backend bind.ContractBackend) (*BLSCertificateVerifier, error) {
	contract, err := bindBLSCertificateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifier{BLSCertificateVerifierCaller: BLSCertificateVerifierCaller{contract: contract}, BLSCertificateVerifierTransactor: BLSCertificateVerifierTransactor{contract: contract}, BLSCertificateVerifierFilterer: BLSCertificateVerifierFilterer{contract: contract}}, nil
}

// NewBLSCertificateVerifierCaller creates a new read-only instance of BLSCertificateVerifier, bound to a specific deployed contract.
func NewBLSCertificateVerifierCaller(address common.Address, caller bind.ContractCaller) (*BLSCertificateVerifierCaller, error) {
	contract, err := bindBLSCertificateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierCaller{contract: contract}, nil
}

// NewBLSCertificateVerifierTransactor creates a new write-only instance of BLSCertificateVerifier, bound to a specific deployed contract.
func NewBLSCertificateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSCertificateVerifierTransactor, error) {
	contract, err := bindBLSCertificateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierTransactor{contract: contract}, nil
}

// NewBLSCertificateVerifierFilterer creates a new log filterer instance of BLSCertificateVerifier, bound to a specific deployed contract.
func NewBLSCertificateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSCertificateVerifierFilterer, error) {
	contract, err := bindBLSCertificateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierFilterer{contract: contract}, nil
}

// bindBLSCertificateVerifier binds a generic wrapper to an already deployed contract.
func bindBLSCertificateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSCertificateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSCertificateVerifier *BLSCertificateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSCertificateVerifier.Contract.BLSCertificateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSCertificateVerifier *BLSCertificateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.BLSCertificateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSCertificateVerifier *BLSCertificateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.BLSCertificateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSCertificateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.contract.Transact(opts, method, params...)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0xa6ff5936.
//
// Solidity: function latestReferenceTimestamp() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) LatestReferenceTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "latestReferenceTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0xa6ff5936.
//
// Solidity: function latestReferenceTimestamp() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) LatestReferenceTimestamp() (uint32, error) {
	return _BLSCertificateVerifier.Contract.LatestReferenceTimestamp(&_BLSCertificateVerifier.CallOpts)
}

// LatestReferenceTimestamp is a free data retrieval call binding the contract method 0xa6ff5936.
//
// Solidity: function latestReferenceTimestamp() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) LatestReferenceTimestamp() (uint32, error) {
	return _BLSCertificateVerifier.Contract.LatestReferenceTimestamp(&_BLSCertificateVerifier.CallOpts)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0xb697a30b.
//
// Solidity: function maxOperatorTableStaleness() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) MaxOperatorTableStaleness(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "maxOperatorTableStaleness")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0xb697a30b.
//
// Solidity: function maxOperatorTableStaleness() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) MaxOperatorTableStaleness() (uint32, error) {
	return _BLSCertificateVerifier.Contract.MaxOperatorTableStaleness(&_BLSCertificateVerifier.CallOpts)
}

// MaxOperatorTableStaleness is a free data retrieval call binding the contract method 0xb697a30b.
//
// Solidity: function maxOperatorTableStaleness() view returns(uint32)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) MaxOperatorTableStaleness() (uint32, error) {
	return _BLSCertificateVerifier.Contract.MaxOperatorTableStaleness(&_BLSCertificateVerifier.CallOpts)
}

// OperatorSet is a free data retrieval call binding the contract method 0xfa2c4812.
//
// Solidity: function operatorSet() view returns((address,uint32))
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) OperatorSet(opts *bind.CallOpts) (OperatorSet, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "operatorSet")

	if err != nil {
		return *new(OperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorSet)).(*OperatorSet)

	return out0, err

}

// OperatorSet is a free data retrieval call binding the contract method 0xfa2c4812.
//
// Solidity: function operatorSet() view returns((address,uint32))
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) OperatorSet() (OperatorSet, error) {
	return _BLSCertificateVerifier.Contract.OperatorSet(&_BLSCertificateVerifier.CallOpts)
}

// OperatorSet is a free data retrieval call binding the contract method 0xfa2c4812.
//
// Solidity: function operatorSet() view returns((address,uint32))
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) OperatorSet() (OperatorSet, error) {
	return _BLSCertificateVerifier.Contract.OperatorSet(&_BLSCertificateVerifier.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) OperatorTableUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "operatorTableUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) OperatorTableUpdater() (common.Address, error) {
	return _BLSCertificateVerifier.Contract.OperatorTableUpdater(&_BLSCertificateVerifier.CallOpts)
}

// OperatorTableUpdater is a free data retrieval call binding the contract method 0x68d6e081.
//
// Solidity: function operatorTableUpdater() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) OperatorTableUpdater() (common.Address, error) {
	return _BLSCertificateVerifier.Contract.OperatorTableUpdater(&_BLSCertificateVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) Owner() (common.Address, error) {
	return _BLSCertificateVerifier.Contract.Owner(&_BLSCertificateVerifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) Owner() (common.Address, error) {
	return _BLSCertificateVerifier.Contract.Owner(&_BLSCertificateVerifier.CallOpts)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x7192a81c.
//
// Solidity: function verifyCertificate((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate) view returns(uint96[])
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) VerifyCertificate(opts *bind.CallOpts, certificate IBLSCertificateVerifierTypesBN254Certificate) ([]*big.Int, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "verifyCertificate", certificate)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// VerifyCertificate is a free data retrieval call binding the contract method 0x7192a81c.
//
// Solidity: function verifyCertificate((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate) view returns(uint96[])
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) VerifyCertificate(certificate IBLSCertificateVerifierTypesBN254Certificate) ([]*big.Int, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificate(&_BLSCertificateVerifier.CallOpts, certificate)
}

// VerifyCertificate is a free data retrieval call binding the contract method 0x7192a81c.
//
// Solidity: function verifyCertificate((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate) view returns(uint96[])
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) VerifyCertificate(certificate IBLSCertificateVerifierTypesBN254Certificate) ([]*big.Int, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificate(&_BLSCertificateVerifier.CallOpts, certificate)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xd8d201fe.
//
// Solidity: function verifyCertificateNominal((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint96[] totalStakeNominalThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) VerifyCertificateNominal(opts *bind.CallOpts, certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (bool, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "verifyCertificateNominal", certificate, totalStakeNominalThresholds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xd8d201fe.
//
// Solidity: function verifyCertificateNominal((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint96[] totalStakeNominalThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) VerifyCertificateNominal(certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (bool, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificateNominal(&_BLSCertificateVerifier.CallOpts, certificate, totalStakeNominalThresholds)
}

// VerifyCertificateNominal is a free data retrieval call binding the contract method 0xd8d201fe.
//
// Solidity: function verifyCertificateNominal((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint96[] totalStakeNominalThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) VerifyCertificateNominal(certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeNominalThresholds []*big.Int) (bool, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificateNominal(&_BLSCertificateVerifier.CallOpts, certificate, totalStakeNominalThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc89fbdbe.
//
// Solidity: function verifyCertificateProportion((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint16[] totalStakeProportionThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierCaller) VerifyCertificateProportion(opts *bind.CallOpts, certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (bool, error) {
	var out []interface{}
	err := _BLSCertificateVerifier.contract.Call(opts, &out, "verifyCertificateProportion", certificate, totalStakeProportionThresholds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc89fbdbe.
//
// Solidity: function verifyCertificateProportion((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint16[] totalStakeProportionThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) VerifyCertificateProportion(certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (bool, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificateProportion(&_BLSCertificateVerifier.CallOpts, certificate, totalStakeProportionThresholds)
}

// VerifyCertificateProportion is a free data retrieval call binding the contract method 0xc89fbdbe.
//
// Solidity: function verifyCertificateProportion((uint32,bytes32,(uint256,uint256),(uint256[2],uint256[2]),uint32[],(uint32,bytes,((uint256,uint256),uint96[]))[]) certificate, uint16[] totalStakeProportionThresholds) view returns(bool)
func (_BLSCertificateVerifier *BLSCertificateVerifierCallerSession) VerifyCertificateProportion(certificate IBLSCertificateVerifierTypesBN254Certificate, totalStakeProportionThresholds []uint16) (bool, error) {
	return _BLSCertificateVerifier.Contract.VerifyCertificateProportion(&_BLSCertificateVerifier.CallOpts, certificate, totalStakeProportionThresholds)
}

// EjectOperators is a paid mutator transaction binding the contract method 0xa973342f.
//
// Solidity: function ejectOperators(uint32 referenceTimestamp, uint32[] operatorIndices, (uint32,bytes,((uint256,uint256),uint96[]))[] witnesses) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) EjectOperators(opts *bind.TransactOpts, referenceTimestamp uint32, operatorIndices []uint32, witnesses []IBLSCertificateVerifierTypesBN254OperatorInfoWitness) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "ejectOperators", referenceTimestamp, operatorIndices, witnesses)
}

// EjectOperators is a paid mutator transaction binding the contract method 0xa973342f.
//
// Solidity: function ejectOperators(uint32 referenceTimestamp, uint32[] operatorIndices, (uint32,bytes,((uint256,uint256),uint96[]))[] witnesses) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) EjectOperators(referenceTimestamp uint32, operatorIndices []uint32, witnesses []IBLSCertificateVerifierTypesBN254OperatorInfoWitness) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.EjectOperators(&_BLSCertificateVerifier.TransactOpts, referenceTimestamp, operatorIndices, witnesses)
}

// EjectOperators is a paid mutator transaction binding the contract method 0xa973342f.
//
// Solidity: function ejectOperators(uint32 referenceTimestamp, uint32[] operatorIndices, (uint32,bytes,((uint256,uint256),uint96[]))[] witnesses) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) EjectOperators(referenceTimestamp uint32, operatorIndices []uint32, witnesses []IBLSCertificateVerifierTypesBN254OperatorInfoWitness) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.EjectOperators(&_BLSCertificateVerifier.TransactOpts, referenceTimestamp, operatorIndices, witnesses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.RenounceOwnership(&_BLSCertificateVerifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.RenounceOwnership(&_BLSCertificateVerifier.TransactOpts)
}

// SetMaxOperatorTableStaleness is a paid mutator transaction binding the contract method 0x7514bc3e.
//
// Solidity: function setMaxOperatorTableStaleness(uint32 _maxOperatorTableStaleness) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) SetMaxOperatorTableStaleness(opts *bind.TransactOpts, _maxOperatorTableStaleness uint32) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "setMaxOperatorTableStaleness", _maxOperatorTableStaleness)
}

// SetMaxOperatorTableStaleness is a paid mutator transaction binding the contract method 0x7514bc3e.
//
// Solidity: function setMaxOperatorTableStaleness(uint32 _maxOperatorTableStaleness) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) SetMaxOperatorTableStaleness(_maxOperatorTableStaleness uint32) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.SetMaxOperatorTableStaleness(&_BLSCertificateVerifier.TransactOpts, _maxOperatorTableStaleness)
}

// SetMaxOperatorTableStaleness is a paid mutator transaction binding the contract method 0x7514bc3e.
//
// Solidity: function setMaxOperatorTableStaleness(uint32 _maxOperatorTableStaleness) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) SetMaxOperatorTableStaleness(_maxOperatorTableStaleness uint32) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.SetMaxOperatorTableStaleness(&_BLSCertificateVerifier.TransactOpts, _maxOperatorTableStaleness)
}

// SetOperatorTableUpdater is a paid mutator transaction binding the contract method 0x3ae5bbe0.
//
// Solidity: function setOperatorTableUpdater(address _operatorTableUpdater) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) SetOperatorTableUpdater(opts *bind.TransactOpts, _operatorTableUpdater common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "setOperatorTableUpdater", _operatorTableUpdater)
}

// SetOperatorTableUpdater is a paid mutator transaction binding the contract method 0x3ae5bbe0.
//
// Solidity: function setOperatorTableUpdater(address _operatorTableUpdater) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) SetOperatorTableUpdater(_operatorTableUpdater common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.SetOperatorTableUpdater(&_BLSCertificateVerifier.TransactOpts, _operatorTableUpdater)
}

// SetOperatorTableUpdater is a paid mutator transaction binding the contract method 0x3ae5bbe0.
//
// Solidity: function setOperatorTableUpdater(address _operatorTableUpdater) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) SetOperatorTableUpdater(_operatorTableUpdater common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.SetOperatorTableUpdater(&_BLSCertificateVerifier.TransactOpts, _operatorTableUpdater)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.TransferOwnership(&_BLSCertificateVerifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.TransferOwnership(&_BLSCertificateVerifier.TransactOpts, newOwner)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0xecf43fc1.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, (uint256,(uint256,uint256),uint96[]) operatorSetInfo, bytes32 operatorInfoTreeRoot) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactor) UpdateOperatorTable(opts *bind.TransactOpts, referenceTimestamp uint32, operatorSetInfo IBLSTableCalculatorTypesBN254OperatorSetInfo, operatorInfoTreeRoot [32]byte) (*types.Transaction, error) {
	return _BLSCertificateVerifier.contract.Transact(opts, "updateOperatorTable", referenceTimestamp, operatorSetInfo, operatorInfoTreeRoot)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0xecf43fc1.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, (uint256,(uint256,uint256),uint96[]) operatorSetInfo, bytes32 operatorInfoTreeRoot) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierSession) UpdateOperatorTable(referenceTimestamp uint32, operatorSetInfo IBLSTableCalculatorTypesBN254OperatorSetInfo, operatorInfoTreeRoot [32]byte) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.UpdateOperatorTable(&_BLSCertificateVerifier.TransactOpts, referenceTimestamp, operatorSetInfo, operatorInfoTreeRoot)
}

// UpdateOperatorTable is a paid mutator transaction binding the contract method 0xecf43fc1.
//
// Solidity: function updateOperatorTable(uint32 referenceTimestamp, (uint256,(uint256,uint256),uint96[]) operatorSetInfo, bytes32 operatorInfoTreeRoot) returns()
func (_BLSCertificateVerifier *BLSCertificateVerifierTransactorSession) UpdateOperatorTable(referenceTimestamp uint32, operatorSetInfo IBLSTableCalculatorTypesBN254OperatorSetInfo, operatorInfoTreeRoot [32]byte) (*types.Transaction, error) {
	return _BLSCertificateVerifier.Contract.UpdateOperatorTable(&_BLSCertificateVerifier.TransactOpts, referenceTimestamp, operatorSetInfo, operatorInfoTreeRoot)
}

// BLSCertificateVerifierMaxOperatorTableStalenessSetIterator is returned from FilterMaxOperatorTableStalenessSet and is used to iterate over the raw logs and unpacked data for MaxOperatorTableStalenessSet events raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierMaxOperatorTableStalenessSetIterator struct {
	Event *BLSCertificateVerifierMaxOperatorTableStalenessSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSCertificateVerifierMaxOperatorTableStalenessSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSCertificateVerifierMaxOperatorTableStalenessSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSCertificateVerifierMaxOperatorTableStalenessSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSCertificateVerifierMaxOperatorTableStalenessSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSCertificateVerifierMaxOperatorTableStalenessSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSCertificateVerifierMaxOperatorTableStalenessSet represents a MaxOperatorTableStalenessSet event raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierMaxOperatorTableStalenessSet struct {
	MaxOperatorTableStaleness uint32
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMaxOperatorTableStalenessSet is a free log retrieval operation binding the contract event 0x72f9e8d81faad2896db70482f1cf2e509e4716d07edf38a36a28d31110a035f1.
//
// Solidity: event MaxOperatorTableStalenessSet(uint32 indexed maxOperatorTableStaleness)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) FilterMaxOperatorTableStalenessSet(opts *bind.FilterOpts, maxOperatorTableStaleness []uint32) (*BLSCertificateVerifierMaxOperatorTableStalenessSetIterator, error) {

	var maxOperatorTableStalenessRule []interface{}
	for _, maxOperatorTableStalenessItem := range maxOperatorTableStaleness {
		maxOperatorTableStalenessRule = append(maxOperatorTableStalenessRule, maxOperatorTableStalenessItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.FilterLogs(opts, "MaxOperatorTableStalenessSet", maxOperatorTableStalenessRule)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierMaxOperatorTableStalenessSetIterator{contract: _BLSCertificateVerifier.contract, event: "MaxOperatorTableStalenessSet", logs: logs, sub: sub}, nil
}

// WatchMaxOperatorTableStalenessSet is a free log subscription operation binding the contract event 0x72f9e8d81faad2896db70482f1cf2e509e4716d07edf38a36a28d31110a035f1.
//
// Solidity: event MaxOperatorTableStalenessSet(uint32 indexed maxOperatorTableStaleness)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) WatchMaxOperatorTableStalenessSet(opts *bind.WatchOpts, sink chan<- *BLSCertificateVerifierMaxOperatorTableStalenessSet, maxOperatorTableStaleness []uint32) (event.Subscription, error) {

	var maxOperatorTableStalenessRule []interface{}
	for _, maxOperatorTableStalenessItem := range maxOperatorTableStaleness {
		maxOperatorTableStalenessRule = append(maxOperatorTableStalenessRule, maxOperatorTableStalenessItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.WatchLogs(opts, "MaxOperatorTableStalenessSet", maxOperatorTableStalenessRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSCertificateVerifierMaxOperatorTableStalenessSet)
				if err := _BLSCertificateVerifier.contract.UnpackLog(event, "MaxOperatorTableStalenessSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxOperatorTableStalenessSet is a log parse operation binding the contract event 0x72f9e8d81faad2896db70482f1cf2e509e4716d07edf38a36a28d31110a035f1.
//
// Solidity: event MaxOperatorTableStalenessSet(uint32 indexed maxOperatorTableStaleness)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) ParseMaxOperatorTableStalenessSet(log types.Log) (*BLSCertificateVerifierMaxOperatorTableStalenessSet, error) {
	event := new(BLSCertificateVerifierMaxOperatorTableStalenessSet)
	if err := _BLSCertificateVerifier.contract.UnpackLog(event, "MaxOperatorTableStalenessSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSCertificateVerifierOperatorTableUpdaterSetIterator is returned from FilterOperatorTableUpdaterSet and is used to iterate over the raw logs and unpacked data for OperatorTableUpdaterSet events raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierOperatorTableUpdaterSetIterator struct {
	Event *BLSCertificateVerifierOperatorTableUpdaterSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSCertificateVerifierOperatorTableUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSCertificateVerifierOperatorTableUpdaterSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSCertificateVerifierOperatorTableUpdaterSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSCertificateVerifierOperatorTableUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSCertificateVerifierOperatorTableUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSCertificateVerifierOperatorTableUpdaterSet represents a OperatorTableUpdaterSet event raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierOperatorTableUpdaterSet struct {
	OperatorTableUpdater common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterOperatorTableUpdaterSet is a free log retrieval operation binding the contract event 0x64c64c9415569fff409712d78176f348185168935b776890b4649e4cd64a25b4.
//
// Solidity: event OperatorTableUpdaterSet(address indexed operatorTableUpdater)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) FilterOperatorTableUpdaterSet(opts *bind.FilterOpts, operatorTableUpdater []common.Address) (*BLSCertificateVerifierOperatorTableUpdaterSetIterator, error) {

	var operatorTableUpdaterRule []interface{}
	for _, operatorTableUpdaterItem := range operatorTableUpdater {
		operatorTableUpdaterRule = append(operatorTableUpdaterRule, operatorTableUpdaterItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.FilterLogs(opts, "OperatorTableUpdaterSet", operatorTableUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierOperatorTableUpdaterSetIterator{contract: _BLSCertificateVerifier.contract, event: "OperatorTableUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchOperatorTableUpdaterSet is a free log subscription operation binding the contract event 0x64c64c9415569fff409712d78176f348185168935b776890b4649e4cd64a25b4.
//
// Solidity: event OperatorTableUpdaterSet(address indexed operatorTableUpdater)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) WatchOperatorTableUpdaterSet(opts *bind.WatchOpts, sink chan<- *BLSCertificateVerifierOperatorTableUpdaterSet, operatorTableUpdater []common.Address) (event.Subscription, error) {

	var operatorTableUpdaterRule []interface{}
	for _, operatorTableUpdaterItem := range operatorTableUpdater {
		operatorTableUpdaterRule = append(operatorTableUpdaterRule, operatorTableUpdaterItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.WatchLogs(opts, "OperatorTableUpdaterSet", operatorTableUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSCertificateVerifierOperatorTableUpdaterSet)
				if err := _BLSCertificateVerifier.contract.UnpackLog(event, "OperatorTableUpdaterSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorTableUpdaterSet is a log parse operation binding the contract event 0x64c64c9415569fff409712d78176f348185168935b776890b4649e4cd64a25b4.
//
// Solidity: event OperatorTableUpdaterSet(address indexed operatorTableUpdater)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) ParseOperatorTableUpdaterSet(log types.Log) (*BLSCertificateVerifierOperatorTableUpdaterSet, error) {
	event := new(BLSCertificateVerifierOperatorTableUpdaterSet)
	if err := _BLSCertificateVerifier.contract.UnpackLog(event, "OperatorTableUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSCertificateVerifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierOwnershipTransferredIterator struct {
	Event *BLSCertificateVerifierOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSCertificateVerifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSCertificateVerifierOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSCertificateVerifierOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSCertificateVerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSCertificateVerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSCertificateVerifierOwnershipTransferred represents a OwnershipTransferred event raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BLSCertificateVerifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierOwnershipTransferredIterator{contract: _BLSCertificateVerifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BLSCertificateVerifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BLSCertificateVerifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSCertificateVerifierOwnershipTransferred)
				if err := _BLSCertificateVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) ParseOwnershipTransferred(log types.Log) (*BLSCertificateVerifierOwnershipTransferred, error) {
	event := new(BLSCertificateVerifierOwnershipTransferred)
	if err := _BLSCertificateVerifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSCertificateVerifierTableUpdatedIterator is returned from FilterTableUpdated and is used to iterate over the raw logs and unpacked data for TableUpdated events raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierTableUpdatedIterator struct {
	Event *BLSCertificateVerifierTableUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSCertificateVerifierTableUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSCertificateVerifierTableUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSCertificateVerifierTableUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSCertificateVerifierTableUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSCertificateVerifierTableUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSCertificateVerifierTableUpdated represents a TableUpdated event raised by the BLSCertificateVerifier contract.
type BLSCertificateVerifierTableUpdated struct {
	ReferenceTimestamp   uint32
	Pubkey               BN254G1Point
	OperatorInfoTreeRoot [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTableUpdated is a free log retrieval operation binding the contract event 0x349eb659faf73a073d293b6e4ee24fcfda5191a5770f5bc0a315bc45b5f69ad2.
//
// Solidity: event TableUpdated(uint32 referenceTimestamp, (uint256,uint256) pubkey, bytes32 operatorInfoTreeRoot)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) FilterTableUpdated(opts *bind.FilterOpts) (*BLSCertificateVerifierTableUpdatedIterator, error) {

	logs, sub, err := _BLSCertificateVerifier.contract.FilterLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return &BLSCertificateVerifierTableUpdatedIterator{contract: _BLSCertificateVerifier.contract, event: "TableUpdated", logs: logs, sub: sub}, nil
}

// WatchTableUpdated is a free log subscription operation binding the contract event 0x349eb659faf73a073d293b6e4ee24fcfda5191a5770f5bc0a315bc45b5f69ad2.
//
// Solidity: event TableUpdated(uint32 referenceTimestamp, (uint256,uint256) pubkey, bytes32 operatorInfoTreeRoot)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) WatchTableUpdated(opts *bind.WatchOpts, sink chan<- *BLSCertificateVerifierTableUpdated) (event.Subscription, error) {

	logs, sub, err := _BLSCertificateVerifier.contract.WatchLogs(opts, "TableUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSCertificateVerifierTableUpdated)
				if err := _BLSCertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTableUpdated is a log parse operation binding the contract event 0x349eb659faf73a073d293b6e4ee24fcfda5191a5770f5bc0a315bc45b5f69ad2.
//
// Solidity: event TableUpdated(uint32 referenceTimestamp, (uint256,uint256) pubkey, bytes32 operatorInfoTreeRoot)
func (_BLSCertificateVerifier *BLSCertificateVerifierFilterer) ParseTableUpdated(log types.Log) (*BLSCertificateVerifierTableUpdated, error) {
	event := new(BLSCertificateVerifierTableUpdated)
	if err := _BLSCertificateVerifier.contract.UnpackLog(event, "TableUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
