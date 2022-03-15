// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proof

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
)

// ProofChainMetaData contains all meta data concerning the ProofChain contract.
var ProofChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"specimenLength\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"}],\"name\":\"BlockSpecimenProductionProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"reward\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"BlockSpecimenRewardAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenRewardAllocation\",\"type\":\"uint128\"}],\"name\":\"BlockSpecimenRewardChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"blockHeight\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"}],\"name\":\"BlockSpecimenSessionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"OperatorStartedRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"OperatorStoppedRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorumThreshold\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionQuorumChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUDITOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_SPECIMEN_PRODUCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"activeIDs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRole\",\"type\":\"bytes32\"}],\"name\":\"addRoleType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"commissionRate\",\"type\":\"uint128\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"definitiveSpecimenHash\",\"type\":\"bytes32\"}],\"name\":\"arbitrateBlockSpecimenSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenQuorumThresholdNumerator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenRewardAllocation\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenSessionDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"countOperatorsRoles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roleCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"disableValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoleTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roleTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPreapprovedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"isSufficientlyStakedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"newStakeAmount\",\"type\":\"uint128\"}],\"name\":\"setRequiredStakeForRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingContractAddress\",\"type\":\"address\"}],\"name\":\"setStakingInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"startOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"stopOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"specimenLength\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"}],\"name\":\"submitBlockSpecimenProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenReward\",\"type\":\"uint128\"}],\"name\":\"updateBlockSpecimenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"updateBlockSpecimenSessionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quorumThresholdNumerator\",\"type\":\"uint64\"}],\"name\":\"updateQuorumThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorIDs\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProofChainABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofChainMetaData.ABI instead.
var ProofChainABI = ProofChainMetaData.ABI

// ProofChain is an auto generated Go binding around an Ethereum contract.
type ProofChain struct {
	ProofChainCaller     // Read-only binding to the contract
	ProofChainTransactor // Write-only binding to the contract
	ProofChainFilterer   // Log filterer for contract events
}

// ProofChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofChainSession struct {
	Contract     *ProofChain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofChainCallerSession struct {
	Contract *ProofChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProofChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofChainTransactorSession struct {
	Contract     *ProofChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProofChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofChainRaw struct {
	Contract *ProofChain // Generic contract binding to access the raw methods on
}

// ProofChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofChainCallerRaw struct {
	Contract *ProofChainCaller // Generic read-only contract binding to access the raw methods on
}

// ProofChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofChainTransactorRaw struct {
	Contract *ProofChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofChain creates a new instance of ProofChain, bound to a specific deployed contract.
func NewProofChain(address common.Address, backend bind.ContractBackend) (*ProofChain, error) {
	contract, err := bindProofChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofChain{ProofChainCaller: ProofChainCaller{contract: contract}, ProofChainTransactor: ProofChainTransactor{contract: contract}, ProofChainFilterer: ProofChainFilterer{contract: contract}}, nil
}

// NewProofChainCaller creates a new read-only instance of ProofChain, bound to a specific deployed contract.
func NewProofChainCaller(address common.Address, caller bind.ContractCaller) (*ProofChainCaller, error) {
	contract, err := bindProofChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofChainCaller{contract: contract}, nil
}

// NewProofChainTransactor creates a new write-only instance of ProofChain, bound to a specific deployed contract.
func NewProofChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofChainTransactor, error) {
	contract, err := bindProofChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofChainTransactor{contract: contract}, nil
}

// NewProofChainFilterer creates a new log filterer instance of ProofChain, bound to a specific deployed contract.
func NewProofChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofChainFilterer, error) {
	contract, err := bindProofChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofChainFilterer{contract: contract}, nil
}

// bindProofChain binds a generic wrapper to an already deployed contract.
func bindProofChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofChain *ProofChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofChain.Contract.ProofChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofChain *ProofChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofChain.Contract.ProofChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofChain *ProofChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofChain.Contract.ProofChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofChain *ProofChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofChain *ProofChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofChain *ProofChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofChain.Contract.contract.Transact(opts, method, params...)
}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCaller) AUDITORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "AUDITOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainSession) AUDITORROLE() ([32]byte, error) {
	return _ProofChain.Contract.AUDITORROLE(&_ProofChain.CallOpts)
}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) AUDITORROLE() ([32]byte, error) {
	return _ProofChain.Contract.AUDITORROLE(&_ProofChain.CallOpts)
}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCaller) BLOCKSPECIMENPRODUCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "BLOCK_SPECIMEN_PRODUCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainSession) BLOCKSPECIMENPRODUCERROLE() ([32]byte, error) {
	return _ProofChain.Contract.BLOCKSPECIMENPRODUCERROLE(&_ProofChain.CallOpts)
}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) BLOCKSPECIMENPRODUCERROLE() ([32]byte, error) {
	return _ProofChain.Contract.BLOCKSPECIMENPRODUCERROLE(&_ProofChain.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainSession) GOVERNANCEROLE() ([32]byte, error) {
	return _ProofChain.Contract.GOVERNANCEROLE(&_ProofChain.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _ProofChain.Contract.GOVERNANCEROLE(&_ProofChain.CallOpts)
}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_ProofChain *ProofChainCaller) ActiveIDs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "activeIDs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_ProofChain *ProofChainSession) ActiveIDs(arg0 *big.Int) (bool, error) {
	return _ProofChain.Contract.ActiveIDs(&_ProofChain.CallOpts, arg0)
}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_ProofChain *ProofChainCallerSession) ActiveIDs(arg0 *big.Int) (bool, error) {
	return _ProofChain.Contract.ActiveIDs(&_ProofChain.CallOpts, arg0)
}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_ProofChain *ProofChainCaller) BlockSpecimenQuorumThresholdNumerator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "blockSpecimenQuorumThresholdNumerator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_ProofChain *ProofChainSession) BlockSpecimenQuorumThresholdNumerator() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenQuorumThresholdNumerator(&_ProofChain.CallOpts)
}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_ProofChain *ProofChainCallerSession) BlockSpecimenQuorumThresholdNumerator() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenQuorumThresholdNumerator(&_ProofChain.CallOpts)
}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_ProofChain *ProofChainCaller) BlockSpecimenRewardAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "blockSpecimenRewardAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_ProofChain *ProofChainSession) BlockSpecimenRewardAllocation() (*big.Int, error) {
	return _ProofChain.Contract.BlockSpecimenRewardAllocation(&_ProofChain.CallOpts)
}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_ProofChain *ProofChainCallerSession) BlockSpecimenRewardAllocation() (*big.Int, error) {
	return _ProofChain.Contract.BlockSpecimenRewardAllocation(&_ProofChain.CallOpts)
}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_ProofChain *ProofChainCaller) BlockSpecimenSessionDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "blockSpecimenSessionDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_ProofChain *ProofChainSession) BlockSpecimenSessionDuration() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenSessionDuration(&_ProofChain.CallOpts)
}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_ProofChain *ProofChainCallerSession) BlockSpecimenSessionDuration() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenSessionDuration(&_ProofChain.CallOpts)
}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_ProofChain *ProofChainCaller) CountOperatorsRoles(opts *bind.CallOpts, operatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "countOperatorsRoles", operatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_ProofChain *ProofChainSession) CountOperatorsRoles(operatorAddress common.Address) (*big.Int, error) {
	return _ProofChain.Contract.CountOperatorsRoles(&_ProofChain.CallOpts, operatorAddress)
}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_ProofChain *ProofChainCallerSession) CountOperatorsRoles(operatorAddress common.Address) (*big.Int, error) {
	return _ProofChain.Contract.CountOperatorsRoles(&_ProofChain.CallOpts, operatorAddress)
}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_ProofChain *ProofChainCaller) GetRoleTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "getRoleTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_ProofChain *ProofChainSession) GetRoleTypes() ([][32]byte, error) {
	return _ProofChain.Contract.GetRoleTypes(&_ProofChain.CallOpts)
}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_ProofChain *ProofChainCallerSession) GetRoleTypes() ([][32]byte, error) {
	return _ProofChain.Contract.GetRoleTypes(&_ProofChain.CallOpts)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_ProofChain *ProofChainCaller) IsPreapprovedForRole(opts *bind.CallOpts, roleName [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isPreapprovedForRole", roleName, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_ProofChain *ProofChainSession) IsPreapprovedForRole(roleName [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole(&_ProofChain.CallOpts, roleName, account)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsPreapprovedForRole(roleName [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole(&_ProofChain.CallOpts, roleName, account)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_ProofChain *ProofChainCaller) IsSufficientlyStakedForRole(opts *bind.CallOpts, roleName [32]byte, validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isSufficientlyStakedForRole", roleName, validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_ProofChain *ProofChainSession) IsSufficientlyStakedForRole(roleName [32]byte, validatorId *big.Int) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole(&_ProofChain.CallOpts, roleName, validatorId)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsSufficientlyStakedForRole(roleName [32]byte, validatorId *big.Int) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole(&_ProofChain.CallOpts, roleName, validatorId)
}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_ProofChain *ProofChainCaller) ValidatorIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "validatorIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_ProofChain *ProofChainSession) ValidatorIDs(arg0 common.Address) (*big.Int, error) {
	return _ProofChain.Contract.ValidatorIDs(&_ProofChain.CallOpts, arg0)
}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_ProofChain *ProofChainCallerSession) ValidatorIDs(arg0 common.Address) (*big.Int, error) {
	return _ProofChain.Contract.ValidatorIDs(&_ProofChain.CallOpts, arg0)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRole) returns()
func (_ProofChain *ProofChainTransactor) AddRoleType(opts *bind.TransactOpts, newRole [32]byte) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "addRoleType", newRole)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRole) returns()
func (_ProofChain *ProofChainSession) AddRoleType(newRole [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.AddRoleType(&_ProofChain.TransactOpts, newRole)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRole) returns()
func (_ProofChain *ProofChainTransactorSession) AddRoleType(newRole [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.AddRoleType(&_ProofChain.TransactOpts, newRole)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_ProofChain *ProofChainTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "addValidator", validator, commissionRate)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_ProofChain *ProofChainSession) AddValidator(validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.AddValidator(&_ProofChain.TransactOpts, validator, commissionRate)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_ProofChain *ProofChainTransactorSession) AddValidator(validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.AddValidator(&_ProofChain.TransactOpts, validator, commissionRate)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainTransactor) ArbitrateBlockSpecimenSession(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "arbitrateBlockSpecimenSession", chainId, blockHeight, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.ArbitrateBlockSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainTransactorSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.ArbitrateBlockSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight, definitiveSpecimenHash)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_ProofChain *ProofChainTransactor) DisableValidator(opts *bind.TransactOpts, validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "disableValidator", validatorId, blockNumber)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_ProofChain *ProofChainSession) DisableValidator(validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.DisableValidator(&_ProofChain.TransactOpts, validatorId, blockNumber)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_ProofChain *ProofChainTransactorSession) DisableValidator(validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.DisableValidator(&_ProofChain.TransactOpts, validatorId, blockNumber)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainTransactor) GrantRolePreapproval(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "grantRolePreapproval", role, account)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainSession) GrantRolePreapproval(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.GrantRolePreapproval(&_ProofChain.TransactOpts, role, account)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainTransactorSession) GrantRolePreapproval(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.GrantRolePreapproval(&_ProofChain.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_ProofChain *ProofChainTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "initialize", initialOwner, stakingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_ProofChain *ProofChainSession) Initialize(initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.Initialize(&_ProofChain.TransactOpts, initialOwner, stakingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_ProofChain *ProofChainTransactorSession) Initialize(initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.Initialize(&_ProofChain.TransactOpts, initialOwner, stakingContract)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_ProofChain *ProofChainTransactor) RemoveOperator(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "removeOperator", validatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_ProofChain *ProofChainSession) RemoveOperator(validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.RemoveOperator(&_ProofChain.TransactOpts, validatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_ProofChain *ProofChainTransactorSession) RemoveOperator(validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.RemoveOperator(&_ProofChain.TransactOpts, validatorId)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainTransactor) RevokeRolePreapproval(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "revokeRolePreapproval", role, account)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainSession) RevokeRolePreapproval(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.RevokeRolePreapproval(&_ProofChain.TransactOpts, role, account)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 role, address account) returns()
func (_ProofChain *ProofChainTransactorSession) RevokeRolePreapproval(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.RevokeRolePreapproval(&_ProofChain.TransactOpts, role, account)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_ProofChain *ProofChainTransactor) SetRequiredStakeForRole(opts *bind.TransactOpts, roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setRequiredStakeForRole", roleName, newStakeAmount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_ProofChain *ProofChainSession) SetRequiredStakeForRole(roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetRequiredStakeForRole(&_ProofChain.TransactOpts, roleName, newStakeAmount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_ProofChain *ProofChainTransactorSession) SetRequiredStakeForRole(roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetRequiredStakeForRole(&_ProofChain.TransactOpts, roleName, newStakeAmount)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_ProofChain *ProofChainTransactor) SetStakingInterface(opts *bind.TransactOpts, stakingContractAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setStakingInterface", stakingContractAddress)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_ProofChain *ProofChainSession) SetStakingInterface(stakingContractAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.SetStakingInterface(&_ProofChain.TransactOpts, stakingContractAddress)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_ProofChain *ProofChainTransactorSession) SetStakingInterface(stakingContractAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.SetStakingInterface(&_ProofChain.TransactOpts, stakingContractAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_ProofChain *ProofChainTransactor) StartOperatorRole(opts *bind.TransactOpts, roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "startOperatorRole", roleName, validatorId, operatorAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_ProofChain *ProofChainSession) StartOperatorRole(roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.StartOperatorRole(&_ProofChain.TransactOpts, roleName, validatorId, operatorAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_ProofChain *ProofChainTransactorSession) StartOperatorRole(roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.StartOperatorRole(&_ProofChain.TransactOpts, roleName, validatorId, operatorAddress)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_ProofChain *ProofChainTransactor) StopOperatorRole(opts *bind.TransactOpts, roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "stopOperatorRole", roleName, validatorId)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_ProofChain *ProofChainSession) StopOperatorRole(roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.StopOperatorRole(&_ProofChain.TransactOpts, roleName, validatorId)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_ProofChain *ProofChainTransactorSession) StopOperatorRole(roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.StopOperatorRole(&_ProofChain.TransactOpts, roleName, validatorId)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainTransactor) SubmitBlockSpecimenProof(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "submitBlockSpecimenProof", chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.Contract.SubmitBlockSpecimenProof(&_ProofChain.TransactOpts, chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainTransactorSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.Contract.SubmitBlockSpecimenProof(&_ProofChain.TransactOpts, chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// UpdateBlockSpecimenReward is a paid mutator transaction binding the contract method 0xec7d3882.
//
// Solidity: function updateBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainTransactor) UpdateBlockSpecimenReward(opts *bind.TransactOpts, newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "updateBlockSpecimenReward", newBlockSpecimenReward)
}

// UpdateBlockSpecimenReward is a paid mutator transaction binding the contract method 0xec7d3882.
//
// Solidity: function updateBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainSession) UpdateBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateBlockSpecimenReward(&_ProofChain.TransactOpts, newBlockSpecimenReward)
}

// UpdateBlockSpecimenReward is a paid mutator transaction binding the contract method 0xec7d3882.
//
// Solidity: function updateBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainTransactorSession) UpdateBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateBlockSpecimenReward(&_ProofChain.TransactOpts, newBlockSpecimenReward)
}

// UpdateBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x1ae83023.
//
// Solidity: function updateBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainTransactor) UpdateBlockSpecimenSessionDuration(opts *bind.TransactOpts, newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "updateBlockSpecimenSessionDuration", newSessionDuration)
}

// UpdateBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x1ae83023.
//
// Solidity: function updateBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainSession) UpdateBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateBlockSpecimenSessionDuration(&_ProofChain.TransactOpts, newSessionDuration)
}

// UpdateBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x1ae83023.
//
// Solidity: function updateBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainTransactorSession) UpdateBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateBlockSpecimenSessionDuration(&_ProofChain.TransactOpts, newSessionDuration)
}

// UpdateQuorumThreshold is a paid mutator transaction binding the contract method 0x6e8b4819.
//
// Solidity: function updateQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_ProofChain *ProofChainTransactor) UpdateQuorumThreshold(opts *bind.TransactOpts, quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "updateQuorumThreshold", quorumThresholdNumerator)
}

// UpdateQuorumThreshold is a paid mutator transaction binding the contract method 0x6e8b4819.
//
// Solidity: function updateQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_ProofChain *ProofChainSession) UpdateQuorumThreshold(quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateQuorumThreshold(&_ProofChain.TransactOpts, quorumThresholdNumerator)
}

// UpdateQuorumThreshold is a paid mutator transaction binding the contract method 0x6e8b4819.
//
// Solidity: function updateQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_ProofChain *ProofChainTransactorSession) UpdateQuorumThreshold(quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.UpdateQuorumThreshold(&_ProofChain.TransactOpts, quorumThresholdNumerator)
}

// ProofChainBlockSpecimenProductionProofSubmittedIterator is returned from FilterBlockSpecimenProductionProofSubmitted and is used to iterate over the raw logs and unpacked data for BlockSpecimenProductionProofSubmitted events raised by the ProofChain contract.
type ProofChainBlockSpecimenProductionProofSubmittedIterator struct {
	Event *ProofChainBlockSpecimenProductionProofSubmitted // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenProductionProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenProductionProofSubmitted)
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
		it.Event = new(ProofChainBlockSpecimenProductionProofSubmitted)
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
func (it *ProofChainBlockSpecimenProductionProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenProductionProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenProductionProofSubmitted represents a BlockSpecimenProductionProofSubmitted event raised by the ProofChain contract.
type ProofChainBlockSpecimenProductionProofSubmitted struct {
	ChainId        uint64
	BlockHeight    uint64
	SpecimenSize   uint64
	SpecimenLength uint64
	SpecimenHash   [32]byte
	StorageURL     string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenProductionProofSubmitted is a free log retrieval operation binding the contract event 0xd6167162e246a2d474fb881cb6dee451485884b64f01655f8f668ee415cbf9b2.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 indexed specimenHash, string storageURL)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenProductionProofSubmitted(opts *bind.FilterOpts, blockHeight []uint64, specimenHash [][32]byte) (*ProofChainBlockSpecimenProductionProofSubmittedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, specimenHashRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenProductionProofSubmittedIterator{contract: _ProofChain.contract, event: "BlockSpecimenProductionProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenProductionProofSubmitted is a free log subscription operation binding the contract event 0xd6167162e246a2d474fb881cb6dee451485884b64f01655f8f668ee415cbf9b2.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 indexed specimenHash, string storageURL)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenProductionProofSubmitted(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenProductionProofSubmitted, blockHeight []uint64, specimenHash [][32]byte) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, specimenHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenProductionProofSubmitted)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenProductionProofSubmitted", log); err != nil {
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

// ParseBlockSpecimenProductionProofSubmitted is a log parse operation binding the contract event 0xd6167162e246a2d474fb881cb6dee451485884b64f01655f8f668ee415cbf9b2.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 indexed specimenHash, string storageURL)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenProductionProofSubmitted(log types.Log) (*ProofChainBlockSpecimenProductionProofSubmitted, error) {
	event := new(ProofChainBlockSpecimenProductionProofSubmitted)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenProductionProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainBlockSpecimenRewardAwardedIterator is returned from FilterBlockSpecimenRewardAwarded and is used to iterate over the raw logs and unpacked data for BlockSpecimenRewardAwarded events raised by the ProofChain contract.
type ProofChainBlockSpecimenRewardAwardedIterator struct {
	Event *ProofChainBlockSpecimenRewardAwarded // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenRewardAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenRewardAwarded)
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
		it.Event = new(ProofChainBlockSpecimenRewardAwarded)
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
func (it *ProofChainBlockSpecimenRewardAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenRewardAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenRewardAwarded represents a BlockSpecimenRewardAwarded event raised by the ProofChain contract.
type ProofChainBlockSpecimenRewardAwarded struct {
	ChainId         uint64
	BlockHeight     uint64
	Reward          *big.Int
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenRewardAwarded is a free log retrieval operation binding the contract event 0x105c3d60f0e35c6731e5828358dcbf158e40f122d68691c098931d2f57b661d1.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 chainId, uint64 indexed blockHeight, uint128 reward, address indexed operatorAddress)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenRewardAwarded(opts *bind.FilterOpts, blockHeight []uint64, operatorAddress []common.Address) (*ProofChainBlockSpecimenRewardAwardedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenRewardAwarded", blockHeightRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenRewardAwardedIterator{contract: _ProofChain.contract, event: "BlockSpecimenRewardAwarded", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenRewardAwarded is a free log subscription operation binding the contract event 0x105c3d60f0e35c6731e5828358dcbf158e40f122d68691c098931d2f57b661d1.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 chainId, uint64 indexed blockHeight, uint128 reward, address indexed operatorAddress)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenRewardAwarded(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenRewardAwarded, blockHeight []uint64, operatorAddress []common.Address) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenRewardAwarded", blockHeightRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenRewardAwarded)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenRewardAwarded", log); err != nil {
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

// ParseBlockSpecimenRewardAwarded is a log parse operation binding the contract event 0x105c3d60f0e35c6731e5828358dcbf158e40f122d68691c098931d2f57b661d1.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 chainId, uint64 indexed blockHeight, uint128 reward, address indexed operatorAddress)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenRewardAwarded(log types.Log) (*ProofChainBlockSpecimenRewardAwarded, error) {
	event := new(ProofChainBlockSpecimenRewardAwarded)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenRewardAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainBlockSpecimenRewardChangedIterator is returned from FilterBlockSpecimenRewardChanged and is used to iterate over the raw logs and unpacked data for BlockSpecimenRewardChanged events raised by the ProofChain contract.
type ProofChainBlockSpecimenRewardChangedIterator struct {
	Event *ProofChainBlockSpecimenRewardChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenRewardChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenRewardChanged)
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
		it.Event = new(ProofChainBlockSpecimenRewardChanged)
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
func (it *ProofChainBlockSpecimenRewardChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenRewardChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenRewardChanged represents a BlockSpecimenRewardChanged event raised by the ProofChain contract.
type ProofChainBlockSpecimenRewardChanged struct {
	NewBlockSpecimenRewardAllocation *big.Int
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenRewardChanged is a free log retrieval operation binding the contract event 0x01eb821dd596243f2f8c5f6c7478e281b855ac12a9f4be2c486cb2778a0bb81e.
//
// Solidity: event BlockSpecimenRewardChanged(uint128 newBlockSpecimenRewardAllocation)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenRewardChanged(opts *bind.FilterOpts) (*ProofChainBlockSpecimenRewardChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenRewardChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenRewardChangedIterator{contract: _ProofChain.contract, event: "BlockSpecimenRewardChanged", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenRewardChanged is a free log subscription operation binding the contract event 0x01eb821dd596243f2f8c5f6c7478e281b855ac12a9f4be2c486cb2778a0bb81e.
//
// Solidity: event BlockSpecimenRewardChanged(uint128 newBlockSpecimenRewardAllocation)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenRewardChanged(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenRewardChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenRewardChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenRewardChanged)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenRewardChanged", log); err != nil {
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

// ParseBlockSpecimenRewardChanged is a log parse operation binding the contract event 0x01eb821dd596243f2f8c5f6c7478e281b855ac12a9f4be2c486cb2778a0bb81e.
//
// Solidity: event BlockSpecimenRewardChanged(uint128 newBlockSpecimenRewardAllocation)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenRewardChanged(log types.Log) (*ProofChainBlockSpecimenRewardChanged, error) {
	event := new(ProofChainBlockSpecimenRewardChanged)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenRewardChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainBlockSpecimenSessionFinalizedIterator is returned from FilterBlockSpecimenSessionFinalized and is used to iterate over the raw logs and unpacked data for BlockSpecimenSessionFinalized events raised by the ProofChain contract.
type ProofChainBlockSpecimenSessionFinalizedIterator struct {
	Event *ProofChainBlockSpecimenSessionFinalized // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenSessionFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenSessionFinalized)
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
		it.Event = new(ProofChainBlockSpecimenSessionFinalized)
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
func (it *ProofChainBlockSpecimenSessionFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenSessionFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenSessionFinalized represents a BlockSpecimenSessionFinalized event raised by the ProofChain contract.
type ProofChainBlockSpecimenSessionFinalized struct {
	BlockHeight *big.Int
	ProofHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenSessionFinalized is a free log retrieval operation binding the contract event 0xcbbbca2698029deaaa991c42053b23dd488b5caf488e5223da54a08a48d7c31e.
//
// Solidity: event BlockSpecimenSessionFinalized(uint128 indexed blockHeight, bytes32 indexed proofHash)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenSessionFinalized(opts *bind.FilterOpts, blockHeight []*big.Int, proofHash [][32]byte) (*ProofChainBlockSpecimenSessionFinalizedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var proofHashRule []interface{}
	for _, proofHashItem := range proofHash {
		proofHashRule = append(proofHashRule, proofHashItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenSessionFinalized", blockHeightRule, proofHashRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenSessionFinalizedIterator{contract: _ProofChain.contract, event: "BlockSpecimenSessionFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenSessionFinalized is a free log subscription operation binding the contract event 0xcbbbca2698029deaaa991c42053b23dd488b5caf488e5223da54a08a48d7c31e.
//
// Solidity: event BlockSpecimenSessionFinalized(uint128 indexed blockHeight, bytes32 indexed proofHash)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenSessionFinalized(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenSessionFinalized, blockHeight []*big.Int, proofHash [][32]byte) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var proofHashRule []interface{}
	for _, proofHashItem := range proofHash {
		proofHashRule = append(proofHashRule, proofHashItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenSessionFinalized", blockHeightRule, proofHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenSessionFinalized)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenSessionFinalized", log); err != nil {
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

// ParseBlockSpecimenSessionFinalized is a log parse operation binding the contract event 0xcbbbca2698029deaaa991c42053b23dd488b5caf488e5223da54a08a48d7c31e.
//
// Solidity: event BlockSpecimenSessionFinalized(uint128 indexed blockHeight, bytes32 indexed proofHash)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenSessionFinalized(log types.Log) (*ProofChainBlockSpecimenSessionFinalized, error) {
	event := new(ProofChainBlockSpecimenSessionFinalized)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenSessionFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorStartedRoleIterator is returned from FilterOperatorStartedRole and is used to iterate over the raw logs and unpacked data for OperatorStartedRole events raised by the ProofChain contract.
type ProofChainOperatorStartedRoleIterator struct {
	Event *ProofChainOperatorStartedRole // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorStartedRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorStartedRole)
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
		it.Event = new(ProofChainOperatorStartedRole)
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
func (it *ProofChainOperatorStartedRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorStartedRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorStartedRole represents a OperatorStartedRole event raised by the ProofChain contract.
type ProofChainOperatorStartedRole struct {
	Role            [32]byte
	OperatorAddress common.Address
	ValidatorId     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorStartedRole is a free log retrieval operation binding the contract event 0x77bc260bf2591a08387f4fb053a665992d20fb87bd485d6210d285ec7706527b.
//
// Solidity: event OperatorStartedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) FilterOperatorStartedRole(opts *bind.FilterOpts, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (*ProofChainOperatorStartedRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorStartedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorStartedRoleIterator{contract: _ProofChain.contract, event: "OperatorStartedRole", logs: logs, sub: sub}, nil
}

// WatchOperatorStartedRole is a free log subscription operation binding the contract event 0x77bc260bf2591a08387f4fb053a665992d20fb87bd485d6210d285ec7706527b.
//
// Solidity: event OperatorStartedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) WatchOperatorStartedRole(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorStartedRole, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorStartedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorStartedRole)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorStartedRole", log); err != nil {
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

// ParseOperatorStartedRole is a log parse operation binding the contract event 0x77bc260bf2591a08387f4fb053a665992d20fb87bd485d6210d285ec7706527b.
//
// Solidity: event OperatorStartedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) ParseOperatorStartedRole(log types.Log) (*ProofChainOperatorStartedRole, error) {
	event := new(ProofChainOperatorStartedRole)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorStartedRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorStoppedRoleIterator is returned from FilterOperatorStoppedRole and is used to iterate over the raw logs and unpacked data for OperatorStoppedRole events raised by the ProofChain contract.
type ProofChainOperatorStoppedRoleIterator struct {
	Event *ProofChainOperatorStoppedRole // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorStoppedRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorStoppedRole)
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
		it.Event = new(ProofChainOperatorStoppedRole)
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
func (it *ProofChainOperatorStoppedRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorStoppedRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorStoppedRole represents a OperatorStoppedRole event raised by the ProofChain contract.
type ProofChainOperatorStoppedRole struct {
	Role            [32]byte
	OperatorAddress common.Address
	ValidatorId     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorStoppedRole is a free log retrieval operation binding the contract event 0xfcf71fe4ab3a430dbd2e89438168cf16c179f2ed54e627d026157b3582761042.
//
// Solidity: event OperatorStoppedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) FilterOperatorStoppedRole(opts *bind.FilterOpts, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (*ProofChainOperatorStoppedRoleIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorStoppedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorStoppedRoleIterator{contract: _ProofChain.contract, event: "OperatorStoppedRole", logs: logs, sub: sub}, nil
}

// WatchOperatorStoppedRole is a free log subscription operation binding the contract event 0xfcf71fe4ab3a430dbd2e89438168cf16c179f2ed54e627d026157b3582761042.
//
// Solidity: event OperatorStoppedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) WatchOperatorStoppedRole(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorStoppedRole, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorStoppedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorStoppedRole)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorStoppedRole", log); err != nil {
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

// ParseOperatorStoppedRole is a log parse operation binding the contract event 0xfcf71fe4ab3a430dbd2e89438168cf16c179f2ed54e627d026157b3582761042.
//
// Solidity: event OperatorStoppedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_ProofChain *ProofChainFilterer) ParseOperatorStoppedRole(log types.Log) (*ProofChainOperatorStoppedRole, error) {
	event := new(ProofChainOperatorStoppedRole)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorStoppedRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ProofChain contract.
type ProofChainRoleAdminChangedIterator struct {
	Event *ProofChainRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainRoleAdminChanged)
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
		it.Event = new(ProofChainRoleAdminChanged)
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
func (it *ProofChainRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainRoleAdminChanged represents a RoleAdminChanged event raised by the ProofChain contract.
type ProofChainRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProofChain *ProofChainFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProofChainRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainRoleAdminChangedIterator{contract: _ProofChain.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProofChain *ProofChainFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProofChainRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainRoleAdminChanged)
				if err := _ProofChain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProofChain *ProofChainFilterer) ParseRoleAdminChanged(log types.Log) (*ProofChainRoleAdminChanged, error) {
	event := new(ProofChainRoleAdminChanged)
	if err := _ProofChain.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ProofChain contract.
type ProofChainRoleGrantedIterator struct {
	Event *ProofChainRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProofChainRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainRoleGranted)
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
		it.Event = new(ProofChainRoleGranted)
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
func (it *ProofChainRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainRoleGranted represents a RoleGranted event raised by the ProofChain contract.
type ProofChainRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProofChainRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainRoleGrantedIterator{contract: _ProofChain.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProofChainRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainRoleGranted)
				if err := _ProofChain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) ParseRoleGranted(log types.Log) (*ProofChainRoleGranted, error) {
	event := new(ProofChainRoleGranted)
	if err := _ProofChain.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ProofChain contract.
type ProofChainRoleRevokedIterator struct {
	Event *ProofChainRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProofChainRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainRoleRevoked)
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
		it.Event = new(ProofChainRoleRevoked)
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
func (it *ProofChainRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainRoleRevoked represents a RoleRevoked event raised by the ProofChain contract.
type ProofChainRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProofChainRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainRoleRevokedIterator{contract: _ProofChain.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProofChainRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainRoleRevoked)
				if err := _ProofChain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProofChain *ProofChainFilterer) ParseRoleRevoked(log types.Log) (*ProofChainRoleRevoked, error) {
	event := new(ProofChainRoleRevoked)
	if err := _ProofChain.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainSpecimenSessionDurationChangedIterator is returned from FilterSpecimenSessionDurationChanged and is used to iterate over the raw logs and unpacked data for SpecimenSessionDurationChanged events raised by the ProofChain contract.
type ProofChainSpecimenSessionDurationChangedIterator struct {
	Event *ProofChainSpecimenSessionDurationChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainSpecimenSessionDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainSpecimenSessionDurationChanged)
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
		it.Event = new(ProofChainSpecimenSessionDurationChanged)
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
func (it *ProofChainSpecimenSessionDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainSpecimenSessionDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainSpecimenSessionDurationChanged represents a SpecimenSessionDurationChanged event raised by the ProofChain contract.
type ProofChainSpecimenSessionDurationChanged struct {
	NewSessionDuration uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSpecimenSessionDurationChanged is a free log retrieval operation binding the contract event 0x94bc488f4d9a985dd5f9d11e8f0a614a62828888eb65b704a90fa852be937549.
//
// Solidity: event SpecimenSessionDurationChanged(uint64 newSessionDuration)
func (_ProofChain *ProofChainFilterer) FilterSpecimenSessionDurationChanged(opts *bind.FilterOpts) (*ProofChainSpecimenSessionDurationChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "SpecimenSessionDurationChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainSpecimenSessionDurationChangedIterator{contract: _ProofChain.contract, event: "SpecimenSessionDurationChanged", logs: logs, sub: sub}, nil
}

// WatchSpecimenSessionDurationChanged is a free log subscription operation binding the contract event 0x94bc488f4d9a985dd5f9d11e8f0a614a62828888eb65b704a90fa852be937549.
//
// Solidity: event SpecimenSessionDurationChanged(uint64 newSessionDuration)
func (_ProofChain *ProofChainFilterer) WatchSpecimenSessionDurationChanged(opts *bind.WatchOpts, sink chan<- *ProofChainSpecimenSessionDurationChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "SpecimenSessionDurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainSpecimenSessionDurationChanged)
				if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionDurationChanged", log); err != nil {
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

// ParseSpecimenSessionDurationChanged is a log parse operation binding the contract event 0x94bc488f4d9a985dd5f9d11e8f0a614a62828888eb65b704a90fa852be937549.
//
// Solidity: event SpecimenSessionDurationChanged(uint64 newSessionDuration)
func (_ProofChain *ProofChainFilterer) ParseSpecimenSessionDurationChanged(log types.Log) (*ProofChainSpecimenSessionDurationChanged, error) {
	event := new(ProofChainSpecimenSessionDurationChanged)
	if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainSpecimenSessionQuorumChangedIterator is returned from FilterSpecimenSessionQuorumChanged and is used to iterate over the raw logs and unpacked data for SpecimenSessionQuorumChanged events raised by the ProofChain contract.
type ProofChainSpecimenSessionQuorumChangedIterator struct {
	Event *ProofChainSpecimenSessionQuorumChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainSpecimenSessionQuorumChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainSpecimenSessionQuorumChanged)
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
		it.Event = new(ProofChainSpecimenSessionQuorumChanged)
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
func (it *ProofChainSpecimenSessionQuorumChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainSpecimenSessionQuorumChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainSpecimenSessionQuorumChanged represents a SpecimenSessionQuorumChanged event raised by the ProofChain contract.
type ProofChainSpecimenSessionQuorumChanged struct {
	NewQuorumThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSpecimenSessionQuorumChanged is a free log retrieval operation binding the contract event 0x7ab8150f5b613293e16e3a99396812661a51bad017fbe4204fb8faef1c315cb0.
//
// Solidity: event SpecimenSessionQuorumChanged(uint64 newQuorumThreshold)
func (_ProofChain *ProofChainFilterer) FilterSpecimenSessionQuorumChanged(opts *bind.FilterOpts) (*ProofChainSpecimenSessionQuorumChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "SpecimenSessionQuorumChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainSpecimenSessionQuorumChangedIterator{contract: _ProofChain.contract, event: "SpecimenSessionQuorumChanged", logs: logs, sub: sub}, nil
}

// WatchSpecimenSessionQuorumChanged is a free log subscription operation binding the contract event 0x7ab8150f5b613293e16e3a99396812661a51bad017fbe4204fb8faef1c315cb0.
//
// Solidity: event SpecimenSessionQuorumChanged(uint64 newQuorumThreshold)
func (_ProofChain *ProofChainFilterer) WatchSpecimenSessionQuorumChanged(opts *bind.WatchOpts, sink chan<- *ProofChainSpecimenSessionQuorumChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "SpecimenSessionQuorumChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainSpecimenSessionQuorumChanged)
				if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionQuorumChanged", log); err != nil {
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

// ParseSpecimenSessionQuorumChanged is a log parse operation binding the contract event 0x7ab8150f5b613293e16e3a99396812661a51bad017fbe4204fb8faef1c315cb0.
//
// Solidity: event SpecimenSessionQuorumChanged(uint64 newQuorumThreshold)
func (_ProofChain *ProofChainFilterer) ParseSpecimenSessionQuorumChanged(log types.Log) (*ProofChainSpecimenSessionQuorumChanged, error) {
	event := new(ProofChainSpecimenSessionQuorumChanged)
	if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionQuorumChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
