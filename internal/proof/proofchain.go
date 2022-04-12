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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newMaxNumberOfHashesPer24H\",\"type\":\"uint64\"}],\"name\":\"BlockSpecimenMaxNumberOfHashesPer24HChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenMinSubmissionsRequired\",\"type\":\"uint128\"}],\"name\":\"BlockSpecimenMinSubmissionRequiredChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"submittedStake\",\"type\":\"uint128\"}],\"name\":\"BlockSpecimenProductionProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockhash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"specimenhash\",\"type\":\"bytes32\"}],\"name\":\"BlockSpecimenRewardAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenRewardAllocation\",\"type\":\"uint128\"}],\"name\":\"BlockSpecimenRewardChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"blockHeight\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"}],\"name\":\"BlockSpecimenSessionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newStakeRequirement\",\"type\":\"uint128\"}],\"name\":\"MinimumRequiredStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nthBlock\",\"type\":\"uint64\"}],\"name\":\"NthBlockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numberOfBlocks\",\"type\":\"uint64\"}],\"name\":\"NumberOfBlocksPer24HChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"QuorumNotReached\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"minSubmissions\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionMinSubmissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorumThreshold\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionQuorumChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterfaceAddress\",\"type\":\"address\"}],\"name\":\"StakingInterfaceChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUDITOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_SPECIMEN_PRODUCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"commissionRate\",\"type\":\"uint128\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"definitiveSpecimenHash\",\"type\":\"bytes32\"}],\"name\":\"arbitrateBlockSpecimenSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenQuorum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenRewardAllocation\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenSessionDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"disableOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"disableValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"enableOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"finalizeAndRewardSpecimenSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoleTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roleTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"isSufficientlyStakedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"maxNumberOfHashesPer24H\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionsRequired\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"nthBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorRoles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minSubmissions\",\"type\":\"uint64\"}],\"name\":\"setBlockSpecimenMinSubmissionRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenReward\",\"type\":\"uint128\"}],\"name\":\"setBlockSpecimenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"setBlockSpecimenSessionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSubmissions\",\"type\":\"uint64\"}],\"name\":\"setMaxNumberOfHashesPer24H\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"minSubmissions\",\"type\":\"uint64\"}],\"name\":\"setMinSubmissionsRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"n\",\"type\":\"uint64\"}],\"name\":\"setNthBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"numberOfBlocks\",\"type\":\"uint64\"}],\"name\":\"setNumberOfBlocksPer24H\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quorum\",\"type\":\"uint64\"}],\"name\":\"setQuorumThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"newStakeAmount\",\"type\":\"uint128\"}],\"name\":\"setRequiredStakeForRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingContractAddress\",\"type\":\"address\"}],\"name\":\"setStakingInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"}],\"name\":\"submitBlockSpecimenProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorIDs\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// BlockSpecimenQuorum is a free data retrieval call binding the contract method 0x96df32b5.
//
// Solidity: function blockSpecimenQuorum() view returns(uint64)
func (_ProofChain *ProofChainCaller) BlockSpecimenQuorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "blockSpecimenQuorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockSpecimenQuorum is a free data retrieval call binding the contract method 0x96df32b5.
//
// Solidity: function blockSpecimenQuorum() view returns(uint64)
func (_ProofChain *ProofChainSession) BlockSpecimenQuorum() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenQuorum(&_ProofChain.CallOpts)
}

// BlockSpecimenQuorum is a free data retrieval call binding the contract method 0x96df32b5.
//
// Solidity: function blockSpecimenQuorum() view returns(uint64)
func (_ProofChain *ProofChainCallerSession) BlockSpecimenQuorum() (uint64, error) {
	return _ProofChain.Contract.BlockSpecimenQuorum(&_ProofChain.CallOpts)
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

// GetOperators is a free data retrieval call binding the contract method 0xd3a8b2a8.
//
// Solidity: function getOperators(uint128 validatorId) view returns(address[])
func (_ProofChain *ProofChainCaller) GetOperators(opts *bind.CallOpts, validatorId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "getOperators", validatorId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0xd3a8b2a8.
//
// Solidity: function getOperators(uint128 validatorId) view returns(address[])
func (_ProofChain *ProofChainSession) GetOperators(validatorId *big.Int) ([]common.Address, error) {
	return _ProofChain.Contract.GetOperators(&_ProofChain.CallOpts, validatorId)
}

// GetOperators is a free data retrieval call binding the contract method 0xd3a8b2a8.
//
// Solidity: function getOperators(uint128 validatorId) view returns(address[])
func (_ProofChain *ProofChainCallerSession) GetOperators(validatorId *big.Int) ([]common.Address, error) {
	return _ProofChain.Contract.GetOperators(&_ProofChain.CallOpts, validatorId)
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

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address operator) view returns(bool)
func (_ProofChain *ProofChainCaller) IsEnabled(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isEnabled", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address operator) view returns(bool)
func (_ProofChain *ProofChainSession) IsEnabled(operator common.Address) (bool, error) {
	return _ProofChain.Contract.IsEnabled(&_ProofChain.CallOpts, operator)
}

// IsEnabled is a free data retrieval call binding the contract method 0x9015d371.
//
// Solidity: function isEnabled(address operator) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsEnabled(operator common.Address) (bool, error) {
	return _ProofChain.Contract.IsEnabled(&_ProofChain.CallOpts, operator)
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

// MaxNumberOfHashesPer24H is a free data retrieval call binding the contract method 0xf271f4dd.
//
// Solidity: function maxNumberOfHashesPer24H(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainCaller) MaxNumberOfHashesPer24H(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "maxNumberOfHashesPer24H", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxNumberOfHashesPer24H is a free data retrieval call binding the contract method 0xf271f4dd.
//
// Solidity: function maxNumberOfHashesPer24H(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainSession) MaxNumberOfHashesPer24H(arg0 uint64) (uint64, error) {
	return _ProofChain.Contract.MaxNumberOfHashesPer24H(&_ProofChain.CallOpts, arg0)
}

// MaxNumberOfHashesPer24H is a free data retrieval call binding the contract method 0xf271f4dd.
//
// Solidity: function maxNumberOfHashesPer24H(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainCallerSession) MaxNumberOfHashesPer24H(arg0 uint64) (uint64, error) {
	return _ProofChain.Contract.MaxNumberOfHashesPer24H(&_ProofChain.CallOpts, arg0)
}

// MinSubmissionsRequired is a free data retrieval call binding the contract method 0x6d261e00.
//
// Solidity: function minSubmissionsRequired() view returns(uint64)
func (_ProofChain *ProofChainCaller) MinSubmissionsRequired(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "minSubmissionsRequired")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinSubmissionsRequired is a free data retrieval call binding the contract method 0x6d261e00.
//
// Solidity: function minSubmissionsRequired() view returns(uint64)
func (_ProofChain *ProofChainSession) MinSubmissionsRequired() (uint64, error) {
	return _ProofChain.Contract.MinSubmissionsRequired(&_ProofChain.CallOpts)
}

// MinSubmissionsRequired is a free data retrieval call binding the contract method 0x6d261e00.
//
// Solidity: function minSubmissionsRequired() view returns(uint64)
func (_ProofChain *ProofChainCallerSession) MinSubmissionsRequired() (uint64, error) {
	return _ProofChain.Contract.MinSubmissionsRequired(&_ProofChain.CallOpts)
}

// NthBlock is a free data retrieval call binding the contract method 0xc7f99302.
//
// Solidity: function nthBlock(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainCaller) NthBlock(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "nthBlock", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NthBlock is a free data retrieval call binding the contract method 0xc7f99302.
//
// Solidity: function nthBlock(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainSession) NthBlock(arg0 uint64) (uint64, error) {
	return _ProofChain.Contract.NthBlock(&_ProofChain.CallOpts, arg0)
}

// NthBlock is a free data retrieval call binding the contract method 0xc7f99302.
//
// Solidity: function nthBlock(uint64 ) view returns(uint64)
func (_ProofChain *ProofChainCallerSession) NthBlock(arg0 uint64) (uint64, error) {
	return _ProofChain.Contract.NthBlock(&_ProofChain.CallOpts, arg0)
}

// OperatorRoles is a free data retrieval call binding the contract method 0x6ab9d8e8.
//
// Solidity: function operatorRoles(address ) view returns(bytes32)
func (_ProofChain *ProofChainCaller) OperatorRoles(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "operatorRoles", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorRoles is a free data retrieval call binding the contract method 0x6ab9d8e8.
//
// Solidity: function operatorRoles(address ) view returns(bytes32)
func (_ProofChain *ProofChainSession) OperatorRoles(arg0 common.Address) ([32]byte, error) {
	return _ProofChain.Contract.OperatorRoles(&_ProofChain.CallOpts, arg0)
}

// OperatorRoles is a free data retrieval call binding the contract method 0x6ab9d8e8.
//
// Solidity: function operatorRoles(address ) view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) OperatorRoles(arg0 common.Address) ([32]byte, error) {
	return _ProofChain.Contract.OperatorRoles(&_ProofChain.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProofChain *ProofChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProofChain *ProofChainSession) Owner() (common.Address, error) {
	return _ProofChain.Contract.Owner(&_ProofChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProofChain *ProofChainCallerSession) Owner() (common.Address, error) {
	return _ProofChain.Contract.Owner(&_ProofChain.CallOpts)
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

// AddOperator is a paid mutator transaction binding the contract method 0xd43e451d.
//
// Solidity: function addOperator(address operator, uint128 validatorId, bytes32 role) returns()
func (_ProofChain *ProofChainTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address, validatorId *big.Int, role [32]byte) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "addOperator", operator, validatorId, role)
}

// AddOperator is a paid mutator transaction binding the contract method 0xd43e451d.
//
// Solidity: function addOperator(address operator, uint128 validatorId, bytes32 role) returns()
func (_ProofChain *ProofChainSession) AddOperator(operator common.Address, validatorId *big.Int, role [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.AddOperator(&_ProofChain.TransactOpts, operator, validatorId, role)
}

// AddOperator is a paid mutator transaction binding the contract method 0xd43e451d.
//
// Solidity: function addOperator(address operator, uint128 validatorId, bytes32 role) returns()
func (_ProofChain *ProofChainTransactorSession) AddOperator(operator common.Address, validatorId *big.Int, role [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.AddOperator(&_ProofChain.TransactOpts, operator, validatorId, role)
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

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x4414beeb.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainTransactor) ArbitrateBlockSpecimenSession(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, blockHash [32]byte, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "arbitrateBlockSpecimenSession", chainId, blockHeight, blockHash, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x4414beeb.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, blockHash [32]byte, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.ArbitrateBlockSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight, blockHash, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x4414beeb.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 definitiveSpecimenHash) returns()
func (_ProofChain *ProofChainTransactorSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, blockHash [32]byte, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _ProofChain.Contract.ArbitrateBlockSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight, blockHash, definitiveSpecimenHash)
}

// DisableOperator is a paid mutator transaction binding the contract method 0xf56408ed.
//
// Solidity: function disableOperator(address operator) returns()
func (_ProofChain *ProofChainTransactor) DisableOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "disableOperator", operator)
}

// DisableOperator is a paid mutator transaction binding the contract method 0xf56408ed.
//
// Solidity: function disableOperator(address operator) returns()
func (_ProofChain *ProofChainSession) DisableOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.DisableOperator(&_ProofChain.TransactOpts, operator)
}

// DisableOperator is a paid mutator transaction binding the contract method 0xf56408ed.
//
// Solidity: function disableOperator(address operator) returns()
func (_ProofChain *ProofChainTransactorSession) DisableOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.DisableOperator(&_ProofChain.TransactOpts, operator)
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

// EnableOperator is a paid mutator transaction binding the contract method 0xdd307b99.
//
// Solidity: function enableOperator(address operator) returns()
func (_ProofChain *ProofChainTransactor) EnableOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "enableOperator", operator)
}

// EnableOperator is a paid mutator transaction binding the contract method 0xdd307b99.
//
// Solidity: function enableOperator(address operator) returns()
func (_ProofChain *ProofChainSession) EnableOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.EnableOperator(&_ProofChain.TransactOpts, operator)
}

// EnableOperator is a paid mutator transaction binding the contract method 0xdd307b99.
//
// Solidity: function enableOperator(address operator) returns()
func (_ProofChain *ProofChainTransactorSession) EnableOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.EnableOperator(&_ProofChain.TransactOpts, operator)
}

// FinalizeAndRewardSpecimenSession is a paid mutator transaction binding the contract method 0x8ecd30bc.
//
// Solidity: function finalizeAndRewardSpecimenSession(uint64 chainId, uint64 blockHeight) returns()
func (_ProofChain *ProofChainTransactor) FinalizeAndRewardSpecimenSession(opts *bind.TransactOpts, chainId uint64, blockHeight uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "finalizeAndRewardSpecimenSession", chainId, blockHeight)
}

// FinalizeAndRewardSpecimenSession is a paid mutator transaction binding the contract method 0x8ecd30bc.
//
// Solidity: function finalizeAndRewardSpecimenSession(uint64 chainId, uint64 blockHeight) returns()
func (_ProofChain *ProofChainSession) FinalizeAndRewardSpecimenSession(chainId uint64, blockHeight uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.FinalizeAndRewardSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight)
}

// FinalizeAndRewardSpecimenSession is a paid mutator transaction binding the contract method 0x8ecd30bc.
//
// Solidity: function finalizeAndRewardSpecimenSession(uint64 chainId, uint64 blockHeight) returns()
func (_ProofChain *ProofChainTransactorSession) FinalizeAndRewardSpecimenSession(chainId uint64, blockHeight uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.FinalizeAndRewardSpecimenSession(&_ProofChain.TransactOpts, chainId, blockHeight)
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

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ProofChain *ProofChainTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ProofChain *ProofChainSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.RemoveOperator(&_ProofChain.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_ProofChain *ProofChainTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.RemoveOperator(&_ProofChain.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProofChain *ProofChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProofChain *ProofChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProofChain.Contract.RenounceOwnership(&_ProofChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProofChain *ProofChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProofChain.Contract.RenounceOwnership(&_ProofChain.TransactOpts)
}

// SetBlockSpecimenMinSubmissionRequired is a paid mutator transaction binding the contract method 0x7432d8dd.
//
// Solidity: function setBlockSpecimenMinSubmissionRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainTransactor) SetBlockSpecimenMinSubmissionRequired(opts *bind.TransactOpts, minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setBlockSpecimenMinSubmissionRequired", minSubmissions)
}

// SetBlockSpecimenMinSubmissionRequired is a paid mutator transaction binding the contract method 0x7432d8dd.
//
// Solidity: function setBlockSpecimenMinSubmissionRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainSession) SetBlockSpecimenMinSubmissionRequired(minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenMinSubmissionRequired(&_ProofChain.TransactOpts, minSubmissions)
}

// SetBlockSpecimenMinSubmissionRequired is a paid mutator transaction binding the contract method 0x7432d8dd.
//
// Solidity: function setBlockSpecimenMinSubmissionRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainTransactorSession) SetBlockSpecimenMinSubmissionRequired(minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenMinSubmissionRequired(&_ProofChain.TransactOpts, minSubmissions)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainTransactor) SetBlockSpecimenReward(opts *bind.TransactOpts, newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setBlockSpecimenReward", newBlockSpecimenReward)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainSession) SetBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenReward(&_ProofChain.TransactOpts, newBlockSpecimenReward)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_ProofChain *ProofChainTransactorSession) SetBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenReward(&_ProofChain.TransactOpts, newBlockSpecimenReward)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainTransactor) SetBlockSpecimenSessionDuration(opts *bind.TransactOpts, newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setBlockSpecimenSessionDuration", newSessionDuration)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainSession) SetBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenSessionDuration(&_ProofChain.TransactOpts, newSessionDuration)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_ProofChain *ProofChainTransactorSession) SetBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetBlockSpecimenSessionDuration(&_ProofChain.TransactOpts, newSessionDuration)
}

// SetMaxNumberOfHashesPer24H is a paid mutator transaction binding the contract method 0xc700c43b.
//
// Solidity: function setMaxNumberOfHashesPer24H(uint64 chainId, uint64 maxSubmissions) returns()
func (_ProofChain *ProofChainTransactor) SetMaxNumberOfHashesPer24H(opts *bind.TransactOpts, chainId uint64, maxSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setMaxNumberOfHashesPer24H", chainId, maxSubmissions)
}

// SetMaxNumberOfHashesPer24H is a paid mutator transaction binding the contract method 0xc700c43b.
//
// Solidity: function setMaxNumberOfHashesPer24H(uint64 chainId, uint64 maxSubmissions) returns()
func (_ProofChain *ProofChainSession) SetMaxNumberOfHashesPer24H(chainId uint64, maxSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetMaxNumberOfHashesPer24H(&_ProofChain.TransactOpts, chainId, maxSubmissions)
}

// SetMaxNumberOfHashesPer24H is a paid mutator transaction binding the contract method 0xc700c43b.
//
// Solidity: function setMaxNumberOfHashesPer24H(uint64 chainId, uint64 maxSubmissions) returns()
func (_ProofChain *ProofChainTransactorSession) SetMaxNumberOfHashesPer24H(chainId uint64, maxSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetMaxNumberOfHashesPer24H(&_ProofChain.TransactOpts, chainId, maxSubmissions)
}

// SetMinSubmissionsRequired is a paid mutator transaction binding the contract method 0x93742b56.
//
// Solidity: function setMinSubmissionsRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainTransactor) SetMinSubmissionsRequired(opts *bind.TransactOpts, minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setMinSubmissionsRequired", minSubmissions)
}

// SetMinSubmissionsRequired is a paid mutator transaction binding the contract method 0x93742b56.
//
// Solidity: function setMinSubmissionsRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainSession) SetMinSubmissionsRequired(minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetMinSubmissionsRequired(&_ProofChain.TransactOpts, minSubmissions)
}

// SetMinSubmissionsRequired is a paid mutator transaction binding the contract method 0x93742b56.
//
// Solidity: function setMinSubmissionsRequired(uint64 minSubmissions) returns()
func (_ProofChain *ProofChainTransactorSession) SetMinSubmissionsRequired(minSubmissions uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetMinSubmissionsRequired(&_ProofChain.TransactOpts, minSubmissions)
}

// SetNthBlock is a paid mutator transaction binding the contract method 0xe3201409.
//
// Solidity: function setNthBlock(uint64 chainId, uint64 n) returns()
func (_ProofChain *ProofChainTransactor) SetNthBlock(opts *bind.TransactOpts, chainId uint64, n uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setNthBlock", chainId, n)
}

// SetNthBlock is a paid mutator transaction binding the contract method 0xe3201409.
//
// Solidity: function setNthBlock(uint64 chainId, uint64 n) returns()
func (_ProofChain *ProofChainSession) SetNthBlock(chainId uint64, n uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetNthBlock(&_ProofChain.TransactOpts, chainId, n)
}

// SetNthBlock is a paid mutator transaction binding the contract method 0xe3201409.
//
// Solidity: function setNthBlock(uint64 chainId, uint64 n) returns()
func (_ProofChain *ProofChainTransactorSession) SetNthBlock(chainId uint64, n uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetNthBlock(&_ProofChain.TransactOpts, chainId, n)
}

// SetNumberOfBlocksPer24H is a paid mutator transaction binding the contract method 0xf883465a.
//
// Solidity: function setNumberOfBlocksPer24H(uint64 numberOfBlocks) returns()
func (_ProofChain *ProofChainTransactor) SetNumberOfBlocksPer24H(opts *bind.TransactOpts, numberOfBlocks uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setNumberOfBlocksPer24H", numberOfBlocks)
}

// SetNumberOfBlocksPer24H is a paid mutator transaction binding the contract method 0xf883465a.
//
// Solidity: function setNumberOfBlocksPer24H(uint64 numberOfBlocks) returns()
func (_ProofChain *ProofChainSession) SetNumberOfBlocksPer24H(numberOfBlocks uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetNumberOfBlocksPer24H(&_ProofChain.TransactOpts, numberOfBlocks)
}

// SetNumberOfBlocksPer24H is a paid mutator transaction binding the contract method 0xf883465a.
//
// Solidity: function setNumberOfBlocksPer24H(uint64 numberOfBlocks) returns()
func (_ProofChain *ProofChainTransactorSession) SetNumberOfBlocksPer24H(numberOfBlocks uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetNumberOfBlocksPer24H(&_ProofChain.TransactOpts, numberOfBlocks)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorum) returns()
func (_ProofChain *ProofChainTransactor) SetQuorumThreshold(opts *bind.TransactOpts, quorum uint64) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setQuorumThreshold", quorum)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorum) returns()
func (_ProofChain *ProofChainSession) SetQuorumThreshold(quorum uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetQuorumThreshold(&_ProofChain.TransactOpts, quorum)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorum) returns()
func (_ProofChain *ProofChainTransactorSession) SetQuorumThreshold(quorum uint64) (*types.Transaction, error) {
	return _ProofChain.Contract.SetQuorumThreshold(&_ProofChain.TransactOpts, quorum)
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

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0x151fd8f3.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainTransactor) SubmitBlockSpecimenProof(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, blockHash [32]byte, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "submitBlockSpecimenProof", chainId, blockHeight, blockHash, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0x151fd8f3.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, blockHash [32]byte, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.Contract.SubmitBlockSpecimenProof(&_ProofChain.TransactOpts, chainId, blockHeight, blockHash, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0x151fd8f3.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, bytes32 blockHash, bytes32 specimenHash, string storageURL) returns()
func (_ProofChain *ProofChainTransactorSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, blockHash [32]byte, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _ProofChain.Contract.SubmitBlockSpecimenProof(&_ProofChain.TransactOpts, chainId, blockHeight, blockHash, specimenHash, storageURL)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProofChain *ProofChainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProofChain *ProofChainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.TransferOwnership(&_ProofChain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProofChain *ProofChainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.TransferOwnership(&_ProofChain.TransactOpts, newOwner)
}

// ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator is returned from FilterBlockSpecimenMaxNumberOfHashesPer24HChanged and is used to iterate over the raw logs and unpacked data for BlockSpecimenMaxNumberOfHashesPer24HChanged events raised by the ProofChain contract.
type ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator struct {
	Event *ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged)
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
		it.Event = new(ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged)
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
func (it *ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged represents a BlockSpecimenMaxNumberOfHashesPer24HChanged event raised by the ProofChain contract.
type ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged struct {
	ChainId                    uint64
	NewMaxNumberOfHashesPer24H uint64
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenMaxNumberOfHashesPer24HChanged is a free log retrieval operation binding the contract event 0xcd51a09527f1fdcb091e787cf5c06a462c010eaef61de7b7c3904916d95a4abd.
//
// Solidity: event BlockSpecimenMaxNumberOfHashesPer24HChanged(uint64 indexed chainId, uint64 newMaxNumberOfHashesPer24H)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenMaxNumberOfHashesPer24HChanged(opts *bind.FilterOpts, chainId []uint64) (*ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenMaxNumberOfHashesPer24HChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenMaxNumberOfHashesPer24HChangedIterator{contract: _ProofChain.contract, event: "BlockSpecimenMaxNumberOfHashesPer24HChanged", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenMaxNumberOfHashesPer24HChanged is a free log subscription operation binding the contract event 0xcd51a09527f1fdcb091e787cf5c06a462c010eaef61de7b7c3904916d95a4abd.
//
// Solidity: event BlockSpecimenMaxNumberOfHashesPer24HChanged(uint64 indexed chainId, uint64 newMaxNumberOfHashesPer24H)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenMaxNumberOfHashesPer24HChanged(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenMaxNumberOfHashesPer24HChanged", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenMaxNumberOfHashesPer24HChanged", log); err != nil {
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

// ParseBlockSpecimenMaxNumberOfHashesPer24HChanged is a log parse operation binding the contract event 0xcd51a09527f1fdcb091e787cf5c06a462c010eaef61de7b7c3904916d95a4abd.
//
// Solidity: event BlockSpecimenMaxNumberOfHashesPer24HChanged(uint64 indexed chainId, uint64 newMaxNumberOfHashesPer24H)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenMaxNumberOfHashesPer24HChanged(log types.Log) (*ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged, error) {
	event := new(ProofChainBlockSpecimenMaxNumberOfHashesPer24HChanged)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenMaxNumberOfHashesPer24HChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator is returned from FilterBlockSpecimenMinSubmissionRequiredChanged and is used to iterate over the raw logs and unpacked data for BlockSpecimenMinSubmissionRequiredChanged events raised by the ProofChain contract.
type ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator struct {
	Event *ProofChainBlockSpecimenMinSubmissionRequiredChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenMinSubmissionRequiredChanged)
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
		it.Event = new(ProofChainBlockSpecimenMinSubmissionRequiredChanged)
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
func (it *ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenMinSubmissionRequiredChanged represents a BlockSpecimenMinSubmissionRequiredChanged event raised by the ProofChain contract.
type ProofChainBlockSpecimenMinSubmissionRequiredChanged struct {
	NewBlockSpecimenMinSubmissionsRequired *big.Int
	Raw                                    types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenMinSubmissionRequiredChanged is a free log retrieval operation binding the contract event 0xe2c2e6446883e57e5ecbdcf7b0a4e67007bcdb9191111ae6589d05a053ad8617.
//
// Solidity: event BlockSpecimenMinSubmissionRequiredChanged(uint128 newBlockSpecimenMinSubmissionsRequired)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenMinSubmissionRequiredChanged(opts *bind.FilterOpts) (*ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenMinSubmissionRequiredChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenMinSubmissionRequiredChangedIterator{contract: _ProofChain.contract, event: "BlockSpecimenMinSubmissionRequiredChanged", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenMinSubmissionRequiredChanged is a free log subscription operation binding the contract event 0xe2c2e6446883e57e5ecbdcf7b0a4e67007bcdb9191111ae6589d05a053ad8617.
//
// Solidity: event BlockSpecimenMinSubmissionRequiredChanged(uint128 newBlockSpecimenMinSubmissionsRequired)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenMinSubmissionRequiredChanged(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenMinSubmissionRequiredChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenMinSubmissionRequiredChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenMinSubmissionRequiredChanged)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenMinSubmissionRequiredChanged", log); err != nil {
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

// ParseBlockSpecimenMinSubmissionRequiredChanged is a log parse operation binding the contract event 0xe2c2e6446883e57e5ecbdcf7b0a4e67007bcdb9191111ae6589d05a053ad8617.
//
// Solidity: event BlockSpecimenMinSubmissionRequiredChanged(uint128 newBlockSpecimenMinSubmissionsRequired)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenMinSubmissionRequiredChanged(log types.Log) (*ProofChainBlockSpecimenMinSubmissionRequiredChanged, error) {
	event := new(ProofChainBlockSpecimenMinSubmissionRequiredChanged)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenMinSubmissionRequiredChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	BlockHash      [32]byte
	SpecimenHash   [32]byte
	StorageURL     string
	SubmittedStake *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenProductionProofSubmitted is a free log retrieval operation binding the contract event 0x57b0cb34d2ff9ed661f8b3c684aaee6cbf0bda5da02f4044205556817fa8e76c.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, bytes32 indexed blockHash, bytes32 indexed specimenHash, string storageURL, uint128 submittedStake)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenProductionProofSubmitted(opts *bind.FilterOpts, blockHeight []uint64, blockHash [][32]byte, specimenHash [][32]byte) (*ProofChainBlockSpecimenProductionProofSubmittedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, blockHashRule, specimenHashRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenProductionProofSubmittedIterator{contract: _ProofChain.contract, event: "BlockSpecimenProductionProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenProductionProofSubmitted is a free log subscription operation binding the contract event 0x57b0cb34d2ff9ed661f8b3c684aaee6cbf0bda5da02f4044205556817fa8e76c.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, bytes32 indexed blockHash, bytes32 indexed specimenHash, string storageURL, uint128 submittedStake)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenProductionProofSubmitted(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenProductionProofSubmitted, blockHeight []uint64, blockHash [][32]byte, specimenHash [][32]byte) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, blockHashRule, specimenHashRule)
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

// ParseBlockSpecimenProductionProofSubmitted is a log parse operation binding the contract event 0x57b0cb34d2ff9ed661f8b3c684aaee6cbf0bda5da02f4044205556817fa8e76c.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, bytes32 indexed blockHash, bytes32 indexed specimenHash, string storageURL, uint128 submittedStake)
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
	ChainId      uint64
	BlockHeight  uint64
	Blockhash    [32]byte
	Specimenhash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenRewardAwarded is a free log retrieval operation binding the contract event 0xf05ac779af1ec75a7b2fbe9415b33a67c00294a121786f7ce2eb3f92e4a6424a.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 indexed chainId, uint64 indexed blockHeight, bytes32 indexed blockhash, bytes32 specimenhash)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenRewardAwarded(opts *bind.FilterOpts, chainId []uint64, blockHeight []uint64, blockhash [][32]byte) (*ProofChainBlockSpecimenRewardAwardedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var blockhashRule []interface{}
	for _, blockhashItem := range blockhash {
		blockhashRule = append(blockhashRule, blockhashItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenRewardAwarded", chainIdRule, blockHeightRule, blockhashRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenRewardAwardedIterator{contract: _ProofChain.contract, event: "BlockSpecimenRewardAwarded", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenRewardAwarded is a free log subscription operation binding the contract event 0xf05ac779af1ec75a7b2fbe9415b33a67c00294a121786f7ce2eb3f92e4a6424a.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 indexed chainId, uint64 indexed blockHeight, bytes32 indexed blockhash, bytes32 specimenhash)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenRewardAwarded(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenRewardAwarded, chainId []uint64, blockHeight []uint64, blockhash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var blockhashRule []interface{}
	for _, blockhashItem := range blockhash {
		blockhashRule = append(blockhashRule, blockhashItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenRewardAwarded", chainIdRule, blockHeightRule, blockhashRule)
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

// ParseBlockSpecimenRewardAwarded is a log parse operation binding the contract event 0xf05ac779af1ec75a7b2fbe9415b33a67c00294a121786f7ce2eb3f92e4a6424a.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 indexed chainId, uint64 indexed blockHeight, bytes32 indexed blockhash, bytes32 specimenhash)
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

// ProofChainMinimumRequiredStakeChangedIterator is returned from FilterMinimumRequiredStakeChanged and is used to iterate over the raw logs and unpacked data for MinimumRequiredStakeChanged events raised by the ProofChain contract.
type ProofChainMinimumRequiredStakeChangedIterator struct {
	Event *ProofChainMinimumRequiredStakeChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainMinimumRequiredStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainMinimumRequiredStakeChanged)
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
		it.Event = new(ProofChainMinimumRequiredStakeChanged)
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
func (it *ProofChainMinimumRequiredStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainMinimumRequiredStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainMinimumRequiredStakeChanged represents a MinimumRequiredStakeChanged event raised by the ProofChain contract.
type ProofChainMinimumRequiredStakeChanged struct {
	Role                [32]byte
	NewStakeRequirement *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinimumRequiredStakeChanged is a free log retrieval operation binding the contract event 0xcc5adc82271e3da3beed19bdd358519f24712369aa0cd14ec87e36a0eaa8efaa.
//
// Solidity: event MinimumRequiredStakeChanged(bytes32 indexed role, uint128 newStakeRequirement)
func (_ProofChain *ProofChainFilterer) FilterMinimumRequiredStakeChanged(opts *bind.FilterOpts, role [][32]byte) (*ProofChainMinimumRequiredStakeChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "MinimumRequiredStakeChanged", roleRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainMinimumRequiredStakeChangedIterator{contract: _ProofChain.contract, event: "MinimumRequiredStakeChanged", logs: logs, sub: sub}, nil
}

// WatchMinimumRequiredStakeChanged is a free log subscription operation binding the contract event 0xcc5adc82271e3da3beed19bdd358519f24712369aa0cd14ec87e36a0eaa8efaa.
//
// Solidity: event MinimumRequiredStakeChanged(bytes32 indexed role, uint128 newStakeRequirement)
func (_ProofChain *ProofChainFilterer) WatchMinimumRequiredStakeChanged(opts *bind.WatchOpts, sink chan<- *ProofChainMinimumRequiredStakeChanged, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "MinimumRequiredStakeChanged", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainMinimumRequiredStakeChanged)
				if err := _ProofChain.contract.UnpackLog(event, "MinimumRequiredStakeChanged", log); err != nil {
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

// ParseMinimumRequiredStakeChanged is a log parse operation binding the contract event 0xcc5adc82271e3da3beed19bdd358519f24712369aa0cd14ec87e36a0eaa8efaa.
//
// Solidity: event MinimumRequiredStakeChanged(bytes32 indexed role, uint128 newStakeRequirement)
func (_ProofChain *ProofChainFilterer) ParseMinimumRequiredStakeChanged(log types.Log) (*ProofChainMinimumRequiredStakeChanged, error) {
	event := new(ProofChainMinimumRequiredStakeChanged)
	if err := _ProofChain.contract.UnpackLog(event, "MinimumRequiredStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainNthBlockChangedIterator is returned from FilterNthBlockChanged and is used to iterate over the raw logs and unpacked data for NthBlockChanged events raised by the ProofChain contract.
type ProofChainNthBlockChangedIterator struct {
	Event *ProofChainNthBlockChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainNthBlockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainNthBlockChanged)
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
		it.Event = new(ProofChainNthBlockChanged)
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
func (it *ProofChainNthBlockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainNthBlockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainNthBlockChanged represents a NthBlockChanged event raised by the ProofChain contract.
type ProofChainNthBlockChanged struct {
	ChainId  uint64
	NthBlock uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNthBlockChanged is a free log retrieval operation binding the contract event 0xbbfa9310306e8a8485d109f8be6b0a808473ce55d2e94b8ca3447c9ddb2854b4.
//
// Solidity: event NthBlockChanged(uint64 indexed chainId, uint64 indexed nthBlock)
func (_ProofChain *ProofChainFilterer) FilterNthBlockChanged(opts *bind.FilterOpts, chainId []uint64, nthBlock []uint64) (*ProofChainNthBlockChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var nthBlockRule []interface{}
	for _, nthBlockItem := range nthBlock {
		nthBlockRule = append(nthBlockRule, nthBlockItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "NthBlockChanged", chainIdRule, nthBlockRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainNthBlockChangedIterator{contract: _ProofChain.contract, event: "NthBlockChanged", logs: logs, sub: sub}, nil
}

// WatchNthBlockChanged is a free log subscription operation binding the contract event 0xbbfa9310306e8a8485d109f8be6b0a808473ce55d2e94b8ca3447c9ddb2854b4.
//
// Solidity: event NthBlockChanged(uint64 indexed chainId, uint64 indexed nthBlock)
func (_ProofChain *ProofChainFilterer) WatchNthBlockChanged(opts *bind.WatchOpts, sink chan<- *ProofChainNthBlockChanged, chainId []uint64, nthBlock []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var nthBlockRule []interface{}
	for _, nthBlockItem := range nthBlock {
		nthBlockRule = append(nthBlockRule, nthBlockItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "NthBlockChanged", chainIdRule, nthBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainNthBlockChanged)
				if err := _ProofChain.contract.UnpackLog(event, "NthBlockChanged", log); err != nil {
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

// ParseNthBlockChanged is a log parse operation binding the contract event 0xbbfa9310306e8a8485d109f8be6b0a808473ce55d2e94b8ca3447c9ddb2854b4.
//
// Solidity: event NthBlockChanged(uint64 indexed chainId, uint64 indexed nthBlock)
func (_ProofChain *ProofChainFilterer) ParseNthBlockChanged(log types.Log) (*ProofChainNthBlockChanged, error) {
	event := new(ProofChainNthBlockChanged)
	if err := _ProofChain.contract.UnpackLog(event, "NthBlockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainNumberOfBlocksPer24HChangedIterator is returned from FilterNumberOfBlocksPer24HChanged and is used to iterate over the raw logs and unpacked data for NumberOfBlocksPer24HChanged events raised by the ProofChain contract.
type ProofChainNumberOfBlocksPer24HChangedIterator struct {
	Event *ProofChainNumberOfBlocksPer24HChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainNumberOfBlocksPer24HChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainNumberOfBlocksPer24HChanged)
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
		it.Event = new(ProofChainNumberOfBlocksPer24HChanged)
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
func (it *ProofChainNumberOfBlocksPer24HChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainNumberOfBlocksPer24HChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainNumberOfBlocksPer24HChanged represents a NumberOfBlocksPer24HChanged event raised by the ProofChain contract.
type ProofChainNumberOfBlocksPer24HChanged struct {
	NumberOfBlocks uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNumberOfBlocksPer24HChanged is a free log retrieval operation binding the contract event 0xd110b97bc9e6210ae4a53402f5a2835c92dbe790fc6b679468af199a13756d98.
//
// Solidity: event NumberOfBlocksPer24HChanged(uint64 numberOfBlocks)
func (_ProofChain *ProofChainFilterer) FilterNumberOfBlocksPer24HChanged(opts *bind.FilterOpts) (*ProofChainNumberOfBlocksPer24HChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "NumberOfBlocksPer24HChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainNumberOfBlocksPer24HChangedIterator{contract: _ProofChain.contract, event: "NumberOfBlocksPer24HChanged", logs: logs, sub: sub}, nil
}

// WatchNumberOfBlocksPer24HChanged is a free log subscription operation binding the contract event 0xd110b97bc9e6210ae4a53402f5a2835c92dbe790fc6b679468af199a13756d98.
//
// Solidity: event NumberOfBlocksPer24HChanged(uint64 numberOfBlocks)
func (_ProofChain *ProofChainFilterer) WatchNumberOfBlocksPer24HChanged(opts *bind.WatchOpts, sink chan<- *ProofChainNumberOfBlocksPer24HChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "NumberOfBlocksPer24HChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainNumberOfBlocksPer24HChanged)
				if err := _ProofChain.contract.UnpackLog(event, "NumberOfBlocksPer24HChanged", log); err != nil {
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

// ParseNumberOfBlocksPer24HChanged is a log parse operation binding the contract event 0xd110b97bc9e6210ae4a53402f5a2835c92dbe790fc6b679468af199a13756d98.
//
// Solidity: event NumberOfBlocksPer24HChanged(uint64 numberOfBlocks)
func (_ProofChain *ProofChainFilterer) ParseNumberOfBlocksPer24HChanged(log types.Log) (*ProofChainNumberOfBlocksPer24HChanged, error) {
	event := new(ProofChainNumberOfBlocksPer24HChanged)
	if err := _ProofChain.contract.UnpackLog(event, "NumberOfBlocksPer24HChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the ProofChain contract.
type ProofChainOperatorAddedIterator struct {
	Event *ProofChainOperatorAdded // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorAdded)
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
		it.Event = new(ProofChainOperatorAdded)
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
func (it *ProofChainOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorAdded represents a OperatorAdded event raised by the ProofChain contract.
type ProofChainOperatorAdded struct {
	Operator    common.Address
	ValidatorId *big.Int
	Role        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0x797ca55fc7be0f65c71f10996f7a16f801094f8ae3811874afc5a39730772a42.
//
// Solidity: event OperatorAdded(address operator, uint128 validatorId, bytes32 role)
func (_ProofChain *ProofChainFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*ProofChainOperatorAddedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorAddedIterator{contract: _ProofChain.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0x797ca55fc7be0f65c71f10996f7a16f801094f8ae3811874afc5a39730772a42.
//
// Solidity: event OperatorAdded(address operator, uint128 validatorId, bytes32 role)
func (_ProofChain *ProofChainFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorAdded)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0x797ca55fc7be0f65c71f10996f7a16f801094f8ae3811874afc5a39730772a42.
//
// Solidity: event OperatorAdded(address operator, uint128 validatorId, bytes32 role)
func (_ProofChain *ProofChainFilterer) ParseOperatorAdded(log types.Log) (*ProofChainOperatorAdded, error) {
	event := new(ProofChainOperatorAdded)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorDisabledIterator is returned from FilterOperatorDisabled and is used to iterate over the raw logs and unpacked data for OperatorDisabled events raised by the ProofChain contract.
type ProofChainOperatorDisabledIterator struct {
	Event *ProofChainOperatorDisabled // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorDisabled)
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
		it.Event = new(ProofChainOperatorDisabled)
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
func (it *ProofChainOperatorDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorDisabled represents a OperatorDisabled event raised by the ProofChain contract.
type ProofChainOperatorDisabled struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDisabled is a free log retrieval operation binding the contract event 0x23cd406c7cafe6d88c3f1c1cc16e438745a4236aec25906be2046ca16c36bd1e.
//
// Solidity: event OperatorDisabled(address operator)
func (_ProofChain *ProofChainFilterer) FilterOperatorDisabled(opts *bind.FilterOpts) (*ProofChainOperatorDisabledIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorDisabled")
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorDisabledIterator{contract: _ProofChain.contract, event: "OperatorDisabled", logs: logs, sub: sub}, nil
}

// WatchOperatorDisabled is a free log subscription operation binding the contract event 0x23cd406c7cafe6d88c3f1c1cc16e438745a4236aec25906be2046ca16c36bd1e.
//
// Solidity: event OperatorDisabled(address operator)
func (_ProofChain *ProofChainFilterer) WatchOperatorDisabled(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorDisabled) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorDisabled)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorDisabled", log); err != nil {
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

// ParseOperatorDisabled is a log parse operation binding the contract event 0x23cd406c7cafe6d88c3f1c1cc16e438745a4236aec25906be2046ca16c36bd1e.
//
// Solidity: event OperatorDisabled(address operator)
func (_ProofChain *ProofChainFilterer) ParseOperatorDisabled(log types.Log) (*ProofChainOperatorDisabled, error) {
	event := new(ProofChainOperatorDisabled)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorEnabledIterator is returned from FilterOperatorEnabled and is used to iterate over the raw logs and unpacked data for OperatorEnabled events raised by the ProofChain contract.
type ProofChainOperatorEnabledIterator struct {
	Event *ProofChainOperatorEnabled // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorEnabled)
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
		it.Event = new(ProofChainOperatorEnabled)
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
func (it *ProofChainOperatorEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorEnabled represents a OperatorEnabled event raised by the ProofChain contract.
type ProofChainOperatorEnabled struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorEnabled is a free log retrieval operation binding the contract event 0x9e532d260bd7dde07708a6b1f7c64042546243d79bac23514cd74fcfc1a01fe4.
//
// Solidity: event OperatorEnabled(address operator)
func (_ProofChain *ProofChainFilterer) FilterOperatorEnabled(opts *bind.FilterOpts) (*ProofChainOperatorEnabledIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorEnabled")
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorEnabledIterator{contract: _ProofChain.contract, event: "OperatorEnabled", logs: logs, sub: sub}, nil
}

// WatchOperatorEnabled is a free log subscription operation binding the contract event 0x9e532d260bd7dde07708a6b1f7c64042546243d79bac23514cd74fcfc1a01fe4.
//
// Solidity: event OperatorEnabled(address operator)
func (_ProofChain *ProofChainFilterer) WatchOperatorEnabled(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorEnabled) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorEnabled)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorEnabled", log); err != nil {
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

// ParseOperatorEnabled is a log parse operation binding the contract event 0x9e532d260bd7dde07708a6b1f7c64042546243d79bac23514cd74fcfc1a01fe4.
//
// Solidity: event OperatorEnabled(address operator)
func (_ProofChain *ProofChainFilterer) ParseOperatorEnabled(log types.Log) (*ProofChainOperatorEnabled, error) {
	event := new(ProofChainOperatorEnabled)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the ProofChain contract.
type ProofChainOperatorRemovedIterator struct {
	Event *ProofChainOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ProofChainOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOperatorRemoved)
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
		it.Event = new(ProofChainOperatorRemoved)
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
func (it *ProofChainOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOperatorRemoved represents a OperatorRemoved event raised by the ProofChain contract.
type ProofChainOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_ProofChain *ProofChainFilterer) FilterOperatorRemoved(opts *bind.FilterOpts) (*ProofChainOperatorRemovedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ProofChainOperatorRemovedIterator{contract: _ProofChain.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_ProofChain *ProofChainFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ProofChainOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOperatorRemoved)
				if err := _ProofChain.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address operator)
func (_ProofChain *ProofChainFilterer) ParseOperatorRemoved(log types.Log) (*ProofChainOperatorRemoved, error) {
	event := new(ProofChainOperatorRemoved)
	if err := _ProofChain.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProofChain contract.
type ProofChainOwnershipTransferredIterator struct {
	Event *ProofChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProofChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainOwnershipTransferred)
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
		it.Event = new(ProofChainOwnershipTransferred)
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
func (it *ProofChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainOwnershipTransferred represents a OwnershipTransferred event raised by the ProofChain contract.
type ProofChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProofChain *ProofChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProofChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainOwnershipTransferredIterator{contract: _ProofChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProofChain *ProofChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProofChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainOwnershipTransferred)
				if err := _ProofChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProofChain *ProofChainFilterer) ParseOwnershipTransferred(log types.Log) (*ProofChainOwnershipTransferred, error) {
	event := new(ProofChainOwnershipTransferred)
	if err := _ProofChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofChainQuorumNotReachedIterator is returned from FilterQuorumNotReached and is used to iterate over the raw logs and unpacked data for QuorumNotReached events raised by the ProofChain contract.
type ProofChainQuorumNotReachedIterator struct {
	Event *ProofChainQuorumNotReached // Event containing the contract specifics and raw log

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
func (it *ProofChainQuorumNotReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainQuorumNotReached)
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
		it.Event = new(ProofChainQuorumNotReached)
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
func (it *ProofChainQuorumNotReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainQuorumNotReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainQuorumNotReached represents a QuorumNotReached event raised by the ProofChain contract.
type ProofChainQuorumNotReached struct {
	ChainId     uint64
	BlockHeight uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQuorumNotReached is a free log retrieval operation binding the contract event 0x398fd8f638a7242217f011fd0720a06747f7a85b7d28d7276684b841baea4021.
//
// Solidity: event QuorumNotReached(uint64 indexed chainId, uint64 blockHeight)
func (_ProofChain *ProofChainFilterer) FilterQuorumNotReached(opts *bind.FilterOpts, chainId []uint64) (*ProofChainQuorumNotReachedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "QuorumNotReached", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofChainQuorumNotReachedIterator{contract: _ProofChain.contract, event: "QuorumNotReached", logs: logs, sub: sub}, nil
}

// WatchQuorumNotReached is a free log subscription operation binding the contract event 0x398fd8f638a7242217f011fd0720a06747f7a85b7d28d7276684b841baea4021.
//
// Solidity: event QuorumNotReached(uint64 indexed chainId, uint64 blockHeight)
func (_ProofChain *ProofChainFilterer) WatchQuorumNotReached(opts *bind.WatchOpts, sink chan<- *ProofChainQuorumNotReached, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "QuorumNotReached", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainQuorumNotReached)
				if err := _ProofChain.contract.UnpackLog(event, "QuorumNotReached", log); err != nil {
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

// ParseQuorumNotReached is a log parse operation binding the contract event 0x398fd8f638a7242217f011fd0720a06747f7a85b7d28d7276684b841baea4021.
//
// Solidity: event QuorumNotReached(uint64 indexed chainId, uint64 blockHeight)
func (_ProofChain *ProofChainFilterer) ParseQuorumNotReached(log types.Log) (*ProofChainQuorumNotReached, error) {
	event := new(ProofChainQuorumNotReached)
	if err := _ProofChain.contract.UnpackLog(event, "QuorumNotReached", log); err != nil {
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

// ProofChainSpecimenSessionMinSubmissionChangedIterator is returned from FilterSpecimenSessionMinSubmissionChanged and is used to iterate over the raw logs and unpacked data for SpecimenSessionMinSubmissionChanged events raised by the ProofChain contract.
type ProofChainSpecimenSessionMinSubmissionChangedIterator struct {
	Event *ProofChainSpecimenSessionMinSubmissionChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainSpecimenSessionMinSubmissionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainSpecimenSessionMinSubmissionChanged)
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
		it.Event = new(ProofChainSpecimenSessionMinSubmissionChanged)
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
func (it *ProofChainSpecimenSessionMinSubmissionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainSpecimenSessionMinSubmissionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainSpecimenSessionMinSubmissionChanged represents a SpecimenSessionMinSubmissionChanged event raised by the ProofChain contract.
type ProofChainSpecimenSessionMinSubmissionChanged struct {
	MinSubmissions uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSpecimenSessionMinSubmissionChanged is a free log retrieval operation binding the contract event 0x28312bbddd51eea4439db773218c441a4057f6ed285c642a569f1dcdba1cc047.
//
// Solidity: event SpecimenSessionMinSubmissionChanged(uint64 minSubmissions)
func (_ProofChain *ProofChainFilterer) FilterSpecimenSessionMinSubmissionChanged(opts *bind.FilterOpts) (*ProofChainSpecimenSessionMinSubmissionChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "SpecimenSessionMinSubmissionChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainSpecimenSessionMinSubmissionChangedIterator{contract: _ProofChain.contract, event: "SpecimenSessionMinSubmissionChanged", logs: logs, sub: sub}, nil
}

// WatchSpecimenSessionMinSubmissionChanged is a free log subscription operation binding the contract event 0x28312bbddd51eea4439db773218c441a4057f6ed285c642a569f1dcdba1cc047.
//
// Solidity: event SpecimenSessionMinSubmissionChanged(uint64 minSubmissions)
func (_ProofChain *ProofChainFilterer) WatchSpecimenSessionMinSubmissionChanged(opts *bind.WatchOpts, sink chan<- *ProofChainSpecimenSessionMinSubmissionChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "SpecimenSessionMinSubmissionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainSpecimenSessionMinSubmissionChanged)
				if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionMinSubmissionChanged", log); err != nil {
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

// ParseSpecimenSessionMinSubmissionChanged is a log parse operation binding the contract event 0x28312bbddd51eea4439db773218c441a4057f6ed285c642a569f1dcdba1cc047.
//
// Solidity: event SpecimenSessionMinSubmissionChanged(uint64 minSubmissions)
func (_ProofChain *ProofChainFilterer) ParseSpecimenSessionMinSubmissionChanged(log types.Log) (*ProofChainSpecimenSessionMinSubmissionChanged, error) {
	event := new(ProofChainSpecimenSessionMinSubmissionChanged)
	if err := _ProofChain.contract.UnpackLog(event, "SpecimenSessionMinSubmissionChanged", log); err != nil {
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

// ProofChainStakingInterfaceChangedIterator is returned from FilterStakingInterfaceChanged and is used to iterate over the raw logs and unpacked data for StakingInterfaceChanged events raised by the ProofChain contract.
type ProofChainStakingInterfaceChangedIterator struct {
	Event *ProofChainStakingInterfaceChanged // Event containing the contract specifics and raw log

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
func (it *ProofChainStakingInterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainStakingInterfaceChanged)
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
		it.Event = new(ProofChainStakingInterfaceChanged)
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
func (it *ProofChainStakingInterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainStakingInterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainStakingInterfaceChanged represents a StakingInterfaceChanged event raised by the ProofChain contract.
type ProofChainStakingInterfaceChanged struct {
	NewInterfaceAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStakingInterfaceChanged is a free log retrieval operation binding the contract event 0x70016f37fc9a299f674d1e3083a27743406649810887ed947a79884b064d2de9.
//
// Solidity: event StakingInterfaceChanged(address newInterfaceAddress)
func (_ProofChain *ProofChainFilterer) FilterStakingInterfaceChanged(opts *bind.FilterOpts) (*ProofChainStakingInterfaceChangedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "StakingInterfaceChanged")
	if err != nil {
		return nil, err
	}
	return &ProofChainStakingInterfaceChangedIterator{contract: _ProofChain.contract, event: "StakingInterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchStakingInterfaceChanged is a free log subscription operation binding the contract event 0x70016f37fc9a299f674d1e3083a27743406649810887ed947a79884b064d2de9.
//
// Solidity: event StakingInterfaceChanged(address newInterfaceAddress)
func (_ProofChain *ProofChainFilterer) WatchStakingInterfaceChanged(opts *bind.WatchOpts, sink chan<- *ProofChainStakingInterfaceChanged) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "StakingInterfaceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainStakingInterfaceChanged)
				if err := _ProofChain.contract.UnpackLog(event, "StakingInterfaceChanged", log); err != nil {
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

// ParseStakingInterfaceChanged is a log parse operation binding the contract event 0x70016f37fc9a299f674d1e3083a27743406649810887ed947a79884b064d2de9.
//
// Solidity: event StakingInterfaceChanged(address newInterfaceAddress)
func (_ProofChain *ProofChainFilterer) ParseStakingInterfaceChanged(log types.Log) (*ProofChainStakingInterfaceChanged, error) {
	event := new(ProofChainStakingInterfaceChanged)
	if err := _ProofChain.contract.UnpackLog(event, "StakingInterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
