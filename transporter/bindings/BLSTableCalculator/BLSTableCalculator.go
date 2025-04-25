// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BLSTableCalculator

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

// BLSTableCalculatorMetaData contains all meta data concerning the BLSTableCalculator contract.
var BLSTableCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorTable\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operatorSetInfo\",\"type\":\"tuple\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorSetInfo\",\"components\":[{\"name\":\"numOperators\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"aggregatePubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalWeights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorInfos\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIBLSTableCalculatorTypes.BN254OperatorInfo[]\",\"components\":[{\"name\":\"pubkey\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"weights\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorWeights\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"weights\",\"type\":\"uint96[][]\",\"internalType\":\"uint96[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateOperatorSet\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]}]",
}

// BLSTableCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSTableCalculatorMetaData.ABI instead.
var BLSTableCalculatorABI = BLSTableCalculatorMetaData.ABI

// BLSTableCalculator is an auto generated Go binding around an Ethereum contract.
type BLSTableCalculator struct {
	BLSTableCalculatorCaller     // Read-only binding to the contract
	BLSTableCalculatorTransactor // Write-only binding to the contract
	BLSTableCalculatorFilterer   // Log filterer for contract events
}

// BLSTableCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSTableCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSTableCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSTableCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSTableCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSTableCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSTableCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSTableCalculatorSession struct {
	Contract     *BLSTableCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BLSTableCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSTableCalculatorCallerSession struct {
	Contract *BLSTableCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BLSTableCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSTableCalculatorTransactorSession struct {
	Contract     *BLSTableCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BLSTableCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSTableCalculatorRaw struct {
	Contract *BLSTableCalculator // Generic contract binding to access the raw methods on
}

// BLSTableCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSTableCalculatorCallerRaw struct {
	Contract *BLSTableCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// BLSTableCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSTableCalculatorTransactorRaw struct {
	Contract *BLSTableCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSTableCalculator creates a new instance of BLSTableCalculator, bound to a specific deployed contract.
func NewBLSTableCalculator(address common.Address, backend bind.ContractBackend) (*BLSTableCalculator, error) {
	contract, err := bindBLSTableCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSTableCalculator{BLSTableCalculatorCaller: BLSTableCalculatorCaller{contract: contract}, BLSTableCalculatorTransactor: BLSTableCalculatorTransactor{contract: contract}, BLSTableCalculatorFilterer: BLSTableCalculatorFilterer{contract: contract}}, nil
}

// NewBLSTableCalculatorCaller creates a new read-only instance of BLSTableCalculator, bound to a specific deployed contract.
func NewBLSTableCalculatorCaller(address common.Address, caller bind.ContractCaller) (*BLSTableCalculatorCaller, error) {
	contract, err := bindBLSTableCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSTableCalculatorCaller{contract: contract}, nil
}

// NewBLSTableCalculatorTransactor creates a new write-only instance of BLSTableCalculator, bound to a specific deployed contract.
func NewBLSTableCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSTableCalculatorTransactor, error) {
	contract, err := bindBLSTableCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSTableCalculatorTransactor{contract: contract}, nil
}

// NewBLSTableCalculatorFilterer creates a new log filterer instance of BLSTableCalculator, bound to a specific deployed contract.
func NewBLSTableCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSTableCalculatorFilterer, error) {
	contract, err := bindBLSTableCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSTableCalculatorFilterer{contract: contract}, nil
}

// bindBLSTableCalculator binds a generic wrapper to an already deployed contract.
func bindBLSTableCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSTableCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSTableCalculator *BLSTableCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSTableCalculator.Contract.BLSTableCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSTableCalculator *BLSTableCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSTableCalculator.Contract.BLSTableCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSTableCalculator *BLSTableCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSTableCalculator.Contract.BLSTableCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSTableCalculator *BLSTableCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSTableCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSTableCalculator *BLSTableCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSTableCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSTableCalculator *BLSTableCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSTableCalculator.Contract.contract.Transact(opts, method, params...)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BLSTableCalculator *BLSTableCalculatorCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSTableCalculator.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BLSTableCalculator *BLSTableCalculatorSession) BlsApkRegistry() (common.Address, error) {
	return _BLSTableCalculator.Contract.BlsApkRegistry(&_BLSTableCalculator.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_BLSTableCalculator *BLSTableCalculatorCallerSession) BlsApkRegistry() (common.Address, error) {
	return _BLSTableCalculator.Contract.BlsApkRegistry(&_BLSTableCalculator.CallOpts)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((uint256,(uint256,uint256),uint96[]) operatorSetInfo)
func (_BLSTableCalculator *BLSTableCalculatorCaller) CalculateOperatorTable(opts *bind.CallOpts, operatorSet OperatorSet) (IBLSTableCalculatorTypesBN254OperatorSetInfo, error) {
	var out []interface{}
	err := _BLSTableCalculator.contract.Call(opts, &out, "calculateOperatorTable", operatorSet)

	if err != nil {
		return *new(IBLSTableCalculatorTypesBN254OperatorSetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSTableCalculatorTypesBN254OperatorSetInfo)).(*IBLSTableCalculatorTypesBN254OperatorSetInfo)

	return out0, err

}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((uint256,(uint256,uint256),uint96[]) operatorSetInfo)
func (_BLSTableCalculator *BLSTableCalculatorSession) CalculateOperatorTable(operatorSet OperatorSet) (IBLSTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BLSTableCalculator.Contract.CalculateOperatorTable(&_BLSTableCalculator.CallOpts, operatorSet)
}

// CalculateOperatorTable is a free data retrieval call binding the contract method 0x124c87e0.
//
// Solidity: function calculateOperatorTable((address,uint32) operatorSet) view returns((uint256,(uint256,uint256),uint96[]) operatorSetInfo)
func (_BLSTableCalculator *BLSTableCalculatorCallerSession) CalculateOperatorTable(operatorSet OperatorSet) (IBLSTableCalculatorTypesBN254OperatorSetInfo, error) {
	return _BLSTableCalculator.Contract.CalculateOperatorTable(&_BLSTableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint96[])[])
func (_BLSTableCalculator *BLSTableCalculatorCaller) GetOperatorInfos(opts *bind.CallOpts, operatorSet OperatorSet) ([]IBLSTableCalculatorTypesBN254OperatorInfo, error) {
	var out []interface{}
	err := _BLSTableCalculator.contract.Call(opts, &out, "getOperatorInfos", operatorSet)

	if err != nil {
		return *new([]IBLSTableCalculatorTypesBN254OperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBLSTableCalculatorTypesBN254OperatorInfo)).(*[]IBLSTableCalculatorTypesBN254OperatorInfo)

	return out0, err

}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint96[])[])
func (_BLSTableCalculator *BLSTableCalculatorSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBLSTableCalculatorTypesBN254OperatorInfo, error) {
	return _BLSTableCalculator.Contract.GetOperatorInfos(&_BLSTableCalculator.CallOpts, operatorSet)
}

// GetOperatorInfos is a free data retrieval call binding the contract method 0xcf2d90ef.
//
// Solidity: function getOperatorInfos((address,uint32) operatorSet) view returns(((uint256,uint256),uint96[])[])
func (_BLSTableCalculator *BLSTableCalculatorCallerSession) GetOperatorInfos(operatorSet OperatorSet) ([]IBLSTableCalculatorTypesBN254OperatorInfo, error) {
	return _BLSTableCalculator.Contract.GetOperatorInfos(&_BLSTableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint96[][] weights)
func (_BLSTableCalculator *BLSTableCalculatorCaller) GetOperatorWeights(opts *bind.CallOpts, operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	var out []interface{}
	err := _BLSTableCalculator.contract.Call(opts, &out, "getOperatorWeights", operatorSet)

	outstruct := new(struct {
		Operators []common.Address
		Weights   [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint96[][] weights)
func (_BLSTableCalculator *BLSTableCalculatorSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BLSTableCalculator.Contract.GetOperatorWeights(&_BLSTableCalculator.CallOpts, operatorSet)
}

// GetOperatorWeights is a free data retrieval call binding the contract method 0x71ca71d9.
//
// Solidity: function getOperatorWeights((address,uint32) operatorSet) view returns(address[] operators, uint96[][] weights)
func (_BLSTableCalculator *BLSTableCalculatorCallerSession) GetOperatorWeights(operatorSet OperatorSet) (struct {
	Operators []common.Address
	Weights   [][]*big.Int
}, error) {
	return _BLSTableCalculator.Contract.GetOperatorWeights(&_BLSTableCalculator.CallOpts, operatorSet)
}

// ValidateOperatorSet is a free data retrieval call binding the contract method 0xf0363949.
//
// Solidity: function validateOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_BLSTableCalculator *BLSTableCalculatorCaller) ValidateOperatorSet(opts *bind.CallOpts, operatorSet OperatorSet) (bool, error) {
	var out []interface{}
	err := _BLSTableCalculator.contract.Call(opts, &out, "validateOperatorSet", operatorSet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOperatorSet is a free data retrieval call binding the contract method 0xf0363949.
//
// Solidity: function validateOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_BLSTableCalculator *BLSTableCalculatorSession) ValidateOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _BLSTableCalculator.Contract.ValidateOperatorSet(&_BLSTableCalculator.CallOpts, operatorSet)
}

// ValidateOperatorSet is a free data retrieval call binding the contract method 0xf0363949.
//
// Solidity: function validateOperatorSet((address,uint32) operatorSet) view returns(bool)
func (_BLSTableCalculator *BLSTableCalculatorCallerSession) ValidateOperatorSet(operatorSet OperatorSet) (bool, error) {
	return _BLSTableCalculator.Contract.ValidateOperatorSet(&_BLSTableCalculator.CallOpts, operatorSet)
}
