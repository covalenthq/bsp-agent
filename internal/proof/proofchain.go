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

// ProofMetaData contains all meta data concerning the Proof contract.
var ProofMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"specimenLength\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"}],\"name\":\"BlockSpecimenProductionProofSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"reward\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"BlockSpecimenRewardAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenRewardAllocation\",\"type\":\"uint128\"}],\"name\":\"BlockSpecimenRewardChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"blockHeight\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proofHash\",\"type\":\"bytes32\"}],\"name\":\"BlockSpecimenSessionFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newStakeRequirement\",\"type\":\"uint128\"}],\"name\":\"MinimumRequiredStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"OperatorStartedRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"OperatorStoppedRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorumThreshold\",\"type\":\"uint64\"}],\"name\":\"SpecimenSessionQuorumChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterfaceAddress\",\"type\":\"address\"}],\"name\":\"StakingInterfaceChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUDITOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_SPECIMEN_PRODUCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"activeIDs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoleName\",\"type\":\"bytes32\"}],\"name\":\"addRoleType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"commissionRate\",\"type\":\"uint128\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"definitiveSpecimenHash\",\"type\":\"bytes32\"}],\"name\":\"arbitrateBlockSpecimenSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenQuorumThresholdNumerator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenRewardAllocation\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockSpecimenSessionDuration\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"countOperatorsRoles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"roleCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"disableValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoleTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"roleTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPreapprovedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"isSufficientlyStakedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newBlockSpecimenReward\",\"type\":\"uint128\"}],\"name\":\"setBlockSpecimenReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSessionDuration\",\"type\":\"uint64\"}],\"name\":\"setBlockSpecimenSessionDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"quorumThresholdNumerator\",\"type\":\"uint64\"}],\"name\":\"setQuorumThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"newStakeAmount\",\"type\":\"uint128\"}],\"name\":\"setRequiredStakeForRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingContractAddress\",\"type\":\"address\"}],\"name\":\"setStakingInterface\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"startOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roleName\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"validatorId\",\"type\":\"uint128\"}],\"name\":\"stopOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"specimenLength\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"specimenHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"storageURL\",\"type\":\"string\"}],\"name\":\"submitBlockSpecimenProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorIDs\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProofABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofMetaData.ABI instead.
var ProofABI = ProofMetaData.ABI

// Proof is an auto generated Go binding around an Ethereum contract.
type Proof struct {
	ProofCaller     // Read-only binding to the contract
	ProofTransactor // Write-only binding to the contract
	ProofFilterer   // Log filterer for contract events
}

// ProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSession struct {
	Contract     *Proof            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofCallerSession struct {
	Contract *ProofCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofTransactorSession struct {
	Contract     *ProofTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofRaw struct {
	Contract *Proof // Generic contract binding to access the raw methods on
}

// ProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofCallerRaw struct {
	Contract *ProofCaller // Generic read-only contract binding to access the raw methods on
}

// ProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofTransactorRaw struct {
	Contract *ProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProof creates a new instance of Proof, bound to a specific deployed contract.
func NewProof(address common.Address, backend bind.ContractBackend) (*Proof, error) {
	contract, err := bindProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// NewProofCaller creates a new read-only instance of Proof, bound to a specific deployed contract.
func NewProofCaller(address common.Address, caller bind.ContractCaller) (*ProofCaller, error) {
	contract, err := bindProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofCaller{contract: contract}, nil
}

// NewProofTransactor creates a new write-only instance of Proof, bound to a specific deployed contract.
func NewProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofTransactor, error) {
	contract, err := bindProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofTransactor{contract: contract}, nil
}

// NewProofFilterer creates a new log filterer instance of Proof, bound to a specific deployed contract.
func NewProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofFilterer, error) {
	contract, err := bindProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofFilterer{contract: contract}, nil
}

// bindProof binds a generic wrapper to an already deployed contract.
func bindProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.ProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transact(opts, method, params...)
}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_Proof *ProofCaller) AUDITORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "AUDITOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_Proof *ProofSession) AUDITORROLE() ([32]byte, error) {
	return _Proof.Contract.AUDITORROLE(&_Proof.CallOpts)
}

// AUDITORROLE is a free data retrieval call binding the contract method 0x6e1d616e.
//
// Solidity: function AUDITOR_ROLE() view returns(bytes32)
func (_Proof *ProofCallerSession) AUDITORROLE() ([32]byte, error) {
	return _Proof.Contract.AUDITORROLE(&_Proof.CallOpts)
}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_Proof *ProofCaller) BLOCKSPECIMENPRODUCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "BLOCK_SPECIMEN_PRODUCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_Proof *ProofSession) BLOCKSPECIMENPRODUCERROLE() ([32]byte, error) {
	return _Proof.Contract.BLOCKSPECIMENPRODUCERROLE(&_Proof.CallOpts)
}

// BLOCKSPECIMENPRODUCERROLE is a free data retrieval call binding the contract method 0x9c49d8ee.
//
// Solidity: function BLOCK_SPECIMEN_PRODUCER_ROLE() view returns(bytes32)
func (_Proof *ProofCallerSession) BLOCKSPECIMENPRODUCERROLE() ([32]byte, error) {
	return _Proof.Contract.BLOCKSPECIMENPRODUCERROLE(&_Proof.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Proof *ProofCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Proof *ProofSession) GOVERNANCEROLE() ([32]byte, error) {
	return _Proof.Contract.GOVERNANCEROLE(&_Proof.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Proof *ProofCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _Proof.Contract.GOVERNANCEROLE(&_Proof.CallOpts)
}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_Proof *ProofCaller) ActiveIDs(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "activeIDs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_Proof *ProofSession) ActiveIDs(arg0 *big.Int) (bool, error) {
	return _Proof.Contract.ActiveIDs(&_Proof.CallOpts, arg0)
}

// ActiveIDs is a free data retrieval call binding the contract method 0x67476118.
//
// Solidity: function activeIDs(uint128 ) view returns(bool)
func (_Proof *ProofCallerSession) ActiveIDs(arg0 *big.Int) (bool, error) {
	return _Proof.Contract.ActiveIDs(&_Proof.CallOpts, arg0)
}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_Proof *ProofCaller) BlockSpecimenQuorumThresholdNumerator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "blockSpecimenQuorumThresholdNumerator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_Proof *ProofSession) BlockSpecimenQuorumThresholdNumerator() (uint64, error) {
	return _Proof.Contract.BlockSpecimenQuorumThresholdNumerator(&_Proof.CallOpts)
}

// BlockSpecimenQuorumThresholdNumerator is a free data retrieval call binding the contract method 0x0a5d288f.
//
// Solidity: function blockSpecimenQuorumThresholdNumerator() view returns(uint64)
func (_Proof *ProofCallerSession) BlockSpecimenQuorumThresholdNumerator() (uint64, error) {
	return _Proof.Contract.BlockSpecimenQuorumThresholdNumerator(&_Proof.CallOpts)
}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_Proof *ProofCaller) BlockSpecimenRewardAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "blockSpecimenRewardAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_Proof *ProofSession) BlockSpecimenRewardAllocation() (*big.Int, error) {
	return _Proof.Contract.BlockSpecimenRewardAllocation(&_Proof.CallOpts)
}

// BlockSpecimenRewardAllocation is a free data retrieval call binding the contract method 0xb143d7db.
//
// Solidity: function blockSpecimenRewardAllocation() view returns(uint128)
func (_Proof *ProofCallerSession) BlockSpecimenRewardAllocation() (*big.Int, error) {
	return _Proof.Contract.BlockSpecimenRewardAllocation(&_Proof.CallOpts)
}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_Proof *ProofCaller) BlockSpecimenSessionDuration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "blockSpecimenSessionDuration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_Proof *ProofSession) BlockSpecimenSessionDuration() (uint64, error) {
	return _Proof.Contract.BlockSpecimenSessionDuration(&_Proof.CallOpts)
}

// BlockSpecimenSessionDuration is a free data retrieval call binding the contract method 0x24dc223d.
//
// Solidity: function blockSpecimenSessionDuration() view returns(uint64)
func (_Proof *ProofCallerSession) BlockSpecimenSessionDuration() (uint64, error) {
	return _Proof.Contract.BlockSpecimenSessionDuration(&_Proof.CallOpts)
}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_Proof *ProofCaller) CountOperatorsRoles(opts *bind.CallOpts, operatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "countOperatorsRoles", operatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_Proof *ProofSession) CountOperatorsRoles(operatorAddress common.Address) (*big.Int, error) {
	return _Proof.Contract.CountOperatorsRoles(&_Proof.CallOpts, operatorAddress)
}

// CountOperatorsRoles is a free data retrieval call binding the contract method 0x14ec4a00.
//
// Solidity: function countOperatorsRoles(address operatorAddress) view returns(uint256 roleCount)
func (_Proof *ProofCallerSession) CountOperatorsRoles(operatorAddress common.Address) (*big.Int, error) {
	return _Proof.Contract.CountOperatorsRoles(&_Proof.CallOpts, operatorAddress)
}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_Proof *ProofCaller) GetRoleTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "getRoleTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_Proof *ProofSession) GetRoleTypes() ([][32]byte, error) {
	return _Proof.Contract.GetRoleTypes(&_Proof.CallOpts)
}

// GetRoleTypes is a free data retrieval call binding the contract method 0x96c61cb2.
//
// Solidity: function getRoleTypes() view returns(bytes32[] roleTypes)
func (_Proof *ProofCallerSession) GetRoleTypes() ([][32]byte, error) {
	return _Proof.Contract.GetRoleTypes(&_Proof.CallOpts)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_Proof *ProofCaller) IsPreapprovedForRole(opts *bind.CallOpts, roleName [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "isPreapprovedForRole", roleName, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_Proof *ProofSession) IsPreapprovedForRole(roleName [32]byte, account common.Address) (bool, error) {
	return _Proof.Contract.IsPreapprovedForRole(&_Proof.CallOpts, roleName, account)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 roleName, address account) view returns(bool)
func (_Proof *ProofCallerSession) IsPreapprovedForRole(roleName [32]byte, account common.Address) (bool, error) {
	return _Proof.Contract.IsPreapprovedForRole(&_Proof.CallOpts, roleName, account)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_Proof *ProofCaller) IsSufficientlyStakedForRole(opts *bind.CallOpts, roleName [32]byte, validatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "isSufficientlyStakedForRole", roleName, validatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_Proof *ProofSession) IsSufficientlyStakedForRole(roleName [32]byte, validatorId *big.Int) (bool, error) {
	return _Proof.Contract.IsSufficientlyStakedForRole(&_Proof.CallOpts, roleName, validatorId)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x74c595ed.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 roleName, uint128 validatorId) view returns(bool)
func (_Proof *ProofCallerSession) IsSufficientlyStakedForRole(roleName [32]byte, validatorId *big.Int) (bool, error) {
	return _Proof.Contract.IsSufficientlyStakedForRole(&_Proof.CallOpts, roleName, validatorId)
}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_Proof *ProofCaller) ValidatorIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Proof.contract.Call(opts, &out, "validatorIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_Proof *ProofSession) ValidatorIDs(arg0 common.Address) (*big.Int, error) {
	return _Proof.Contract.ValidatorIDs(&_Proof.CallOpts, arg0)
}

// ValidatorIDs is a free data retrieval call binding the contract method 0x0d92f4ed.
//
// Solidity: function validatorIDs(address ) view returns(uint128)
func (_Proof *ProofCallerSession) ValidatorIDs(arg0 common.Address) (*big.Int, error) {
	return _Proof.Contract.ValidatorIDs(&_Proof.CallOpts, arg0)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRoleName) returns()
func (_Proof *ProofTransactor) AddRoleType(opts *bind.TransactOpts, newRoleName [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "addRoleType", newRoleName)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRoleName) returns()
func (_Proof *ProofSession) AddRoleType(newRoleName [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.AddRoleType(&_Proof.TransactOpts, newRoleName)
}

// AddRoleType is a paid mutator transaction binding the contract method 0xcdb1c112.
//
// Solidity: function addRoleType(bytes32 newRoleName) returns()
func (_Proof *ProofTransactorSession) AddRoleType(newRoleName [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.AddRoleType(&_Proof.TransactOpts, newRoleName)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_Proof *ProofTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "addValidator", validator, commissionRate)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_Proof *ProofSession) AddValidator(validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.AddValidator(&_Proof.TransactOpts, validator, commissionRate)
}

// AddValidator is a paid mutator transaction binding the contract method 0xa2e7e441.
//
// Solidity: function addValidator(address validator, uint128 commissionRate) returns()
func (_Proof *ProofTransactorSession) AddValidator(validator common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.AddValidator(&_Proof.TransactOpts, validator, commissionRate)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_Proof *ProofTransactor) ArbitrateBlockSpecimenSession(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "arbitrateBlockSpecimenSession", chainId, blockHeight, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_Proof *ProofSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.ArbitrateBlockSpecimenSession(&_Proof.TransactOpts, chainId, blockHeight, definitiveSpecimenHash)
}

// ArbitrateBlockSpecimenSession is a paid mutator transaction binding the contract method 0x6a667fdd.
//
// Solidity: function arbitrateBlockSpecimenSession(uint64 chainId, uint64 blockHeight, bytes32 definitiveSpecimenHash) returns()
func (_Proof *ProofTransactorSession) ArbitrateBlockSpecimenSession(chainId uint64, blockHeight uint64, definitiveSpecimenHash [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.ArbitrateBlockSpecimenSession(&_Proof.TransactOpts, chainId, blockHeight, definitiveSpecimenHash)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_Proof *ProofTransactor) DisableValidator(opts *bind.TransactOpts, validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "disableValidator", validatorId, blockNumber)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_Proof *ProofSession) DisableValidator(validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.DisableValidator(&_Proof.TransactOpts, validatorId, blockNumber)
}

// DisableValidator is a paid mutator transaction binding the contract method 0xad9e91ee.
//
// Solidity: function disableValidator(uint128 validatorId, uint256 blockNumber) returns()
func (_Proof *ProofTransactorSession) DisableValidator(validatorId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.DisableValidator(&_Proof.TransactOpts, validatorId, blockNumber)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofTransactor) GrantRolePreapproval(opts *bind.TransactOpts, roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "grantRolePreapproval", roleName, account)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofSession) GrantRolePreapproval(roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.Contract.GrantRolePreapproval(&_Proof.TransactOpts, roleName, account)
}

// GrantRolePreapproval is a paid mutator transaction binding the contract method 0x5ff2bfa0.
//
// Solidity: function grantRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofTransactorSession) GrantRolePreapproval(roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.Contract.GrantRolePreapproval(&_Proof.TransactOpts, roleName, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_Proof *ProofTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "initialize", initialOwner, stakingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_Proof *ProofSession) Initialize(initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _Proof.Contract.Initialize(&_Proof.TransactOpts, initialOwner, stakingContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address stakingContract) returns()
func (_Proof *ProofTransactorSession) Initialize(initialOwner common.Address, stakingContract common.Address) (*types.Transaction, error) {
	return _Proof.Contract.Initialize(&_Proof.TransactOpts, initialOwner, stakingContract)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_Proof *ProofTransactor) RemoveOperator(opts *bind.TransactOpts, validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "removeOperator", validatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_Proof *ProofSession) RemoveOperator(validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.RemoveOperator(&_Proof.TransactOpts, validatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x97d35676.
//
// Solidity: function removeOperator(uint128 validatorId) returns()
func (_Proof *ProofTransactorSession) RemoveOperator(validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.RemoveOperator(&_Proof.TransactOpts, validatorId)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofTransactor) RevokeRolePreapproval(opts *bind.TransactOpts, roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "revokeRolePreapproval", roleName, account)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofSession) RevokeRolePreapproval(roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.Contract.RevokeRolePreapproval(&_Proof.TransactOpts, roleName, account)
}

// RevokeRolePreapproval is a paid mutator transaction binding the contract method 0x75d37c73.
//
// Solidity: function revokeRolePreapproval(bytes32 roleName, address account) returns()
func (_Proof *ProofTransactorSession) RevokeRolePreapproval(roleName [32]byte, account common.Address) (*types.Transaction, error) {
	return _Proof.Contract.RevokeRolePreapproval(&_Proof.TransactOpts, roleName, account)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_Proof *ProofTransactor) SetBlockSpecimenReward(opts *bind.TransactOpts, newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "setBlockSpecimenReward", newBlockSpecimenReward)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_Proof *ProofSession) SetBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.SetBlockSpecimenReward(&_Proof.TransactOpts, newBlockSpecimenReward)
}

// SetBlockSpecimenReward is a paid mutator transaction binding the contract method 0x89587f05.
//
// Solidity: function setBlockSpecimenReward(uint128 newBlockSpecimenReward) returns()
func (_Proof *ProofTransactorSession) SetBlockSpecimenReward(newBlockSpecimenReward *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.SetBlockSpecimenReward(&_Proof.TransactOpts, newBlockSpecimenReward)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_Proof *ProofTransactor) SetBlockSpecimenSessionDuration(opts *bind.TransactOpts, newSessionDuration uint64) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "setBlockSpecimenSessionDuration", newSessionDuration)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_Proof *ProofSession) SetBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _Proof.Contract.SetBlockSpecimenSessionDuration(&_Proof.TransactOpts, newSessionDuration)
}

// SetBlockSpecimenSessionDuration is a paid mutator transaction binding the contract method 0x67d07ad7.
//
// Solidity: function setBlockSpecimenSessionDuration(uint64 newSessionDuration) returns()
func (_Proof *ProofTransactorSession) SetBlockSpecimenSessionDuration(newSessionDuration uint64) (*types.Transaction, error) {
	return _Proof.Contract.SetBlockSpecimenSessionDuration(&_Proof.TransactOpts, newSessionDuration)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_Proof *ProofTransactor) SetQuorumThreshold(opts *bind.TransactOpts, quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "setQuorumThreshold", quorumThresholdNumerator)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_Proof *ProofSession) SetQuorumThreshold(quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _Proof.Contract.SetQuorumThreshold(&_Proof.TransactOpts, quorumThresholdNumerator)
}

// SetQuorumThreshold is a paid mutator transaction binding the contract method 0x7442f33a.
//
// Solidity: function setQuorumThreshold(uint64 quorumThresholdNumerator) returns()
func (_Proof *ProofTransactorSession) SetQuorumThreshold(quorumThresholdNumerator uint64) (*types.Transaction, error) {
	return _Proof.Contract.SetQuorumThreshold(&_Proof.TransactOpts, quorumThresholdNumerator)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_Proof *ProofTransactor) SetRequiredStakeForRole(opts *bind.TransactOpts, roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "setRequiredStakeForRole", roleName, newStakeAmount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_Proof *ProofSession) SetRequiredStakeForRole(roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.SetRequiredStakeForRole(&_Proof.TransactOpts, roleName, newStakeAmount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x4a40372a.
//
// Solidity: function setRequiredStakeForRole(bytes32 roleName, uint128 newStakeAmount) returns()
func (_Proof *ProofTransactorSession) SetRequiredStakeForRole(roleName [32]byte, newStakeAmount *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.SetRequiredStakeForRole(&_Proof.TransactOpts, roleName, newStakeAmount)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_Proof *ProofTransactor) SetStakingInterface(opts *bind.TransactOpts, stakingContractAddress common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "setStakingInterface", stakingContractAddress)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_Proof *ProofSession) SetStakingInterface(stakingContractAddress common.Address) (*types.Transaction, error) {
	return _Proof.Contract.SetStakingInterface(&_Proof.TransactOpts, stakingContractAddress)
}

// SetStakingInterface is a paid mutator transaction binding the contract method 0x3646aded.
//
// Solidity: function setStakingInterface(address stakingContractAddress) returns()
func (_Proof *ProofTransactorSession) SetStakingInterface(stakingContractAddress common.Address) (*types.Transaction, error) {
	return _Proof.Contract.SetStakingInterface(&_Proof.TransactOpts, stakingContractAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_Proof *ProofTransactor) StartOperatorRole(opts *bind.TransactOpts, roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "startOperatorRole", roleName, validatorId, operatorAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_Proof *ProofSession) StartOperatorRole(roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Proof.Contract.StartOperatorRole(&_Proof.TransactOpts, roleName, validatorId, operatorAddress)
}

// StartOperatorRole is a paid mutator transaction binding the contract method 0x79589ee6.
//
// Solidity: function startOperatorRole(bytes32 roleName, uint128 validatorId, address operatorAddress) returns()
func (_Proof *ProofTransactorSession) StartOperatorRole(roleName [32]byte, validatorId *big.Int, operatorAddress common.Address) (*types.Transaction, error) {
	return _Proof.Contract.StartOperatorRole(&_Proof.TransactOpts, roleName, validatorId, operatorAddress)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_Proof *ProofTransactor) StopOperatorRole(opts *bind.TransactOpts, roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "stopOperatorRole", roleName, validatorId)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_Proof *ProofSession) StopOperatorRole(roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.StopOperatorRole(&_Proof.TransactOpts, roleName, validatorId)
}

// StopOperatorRole is a paid mutator transaction binding the contract method 0x7e283822.
//
// Solidity: function stopOperatorRole(bytes32 roleName, uint128 validatorId) returns()
func (_Proof *ProofTransactorSession) StopOperatorRole(roleName [32]byte, validatorId *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.StopOperatorRole(&_Proof.TransactOpts, roleName, validatorId)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_Proof *ProofTransactor) SubmitBlockSpecimenProof(opts *bind.TransactOpts, chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "submitBlockSpecimenProof", chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_Proof *ProofSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _Proof.Contract.SubmitBlockSpecimenProof(&_Proof.TransactOpts, chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// SubmitBlockSpecimenProof is a paid mutator transaction binding the contract method 0xa5b14168.
//
// Solidity: function submitBlockSpecimenProof(uint64 chainId, uint64 blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 specimenHash, string storageURL) returns()
func (_Proof *ProofTransactorSession) SubmitBlockSpecimenProof(chainId uint64, blockHeight uint64, specimenSize uint64, specimenLength uint64, specimenHash [32]byte, storageURL string) (*types.Transaction, error) {
	return _Proof.Contract.SubmitBlockSpecimenProof(&_Proof.TransactOpts, chainId, blockHeight, specimenSize, specimenLength, specimenHash, storageURL)
}

// ProofBlockSpecimenProductionProofSubmittedIterator is returned from FilterBlockSpecimenProductionProofSubmitted and is used to iterate over the raw logs and unpacked data for BlockSpecimenProductionProofSubmitted events raised by the Proof contract.
type ProofBlockSpecimenProductionProofSubmittedIterator struct {
	Event *ProofBlockSpecimenProductionProofSubmitted // Event containing the contract specifics and raw log

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
func (it *ProofBlockSpecimenProductionProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofBlockSpecimenProductionProofSubmitted)
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
		it.Event = new(ProofBlockSpecimenProductionProofSubmitted)
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
func (it *ProofBlockSpecimenProductionProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofBlockSpecimenProductionProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofBlockSpecimenProductionProofSubmitted represents a BlockSpecimenProductionProofSubmitted event raised by the Proof contract.
type ProofBlockSpecimenProductionProofSubmitted struct {
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
func (_Proof *ProofFilterer) FilterBlockSpecimenProductionProofSubmitted(opts *bind.FilterOpts, blockHeight []uint64, specimenHash [][32]byte) (*ProofBlockSpecimenProductionProofSubmittedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _Proof.contract.FilterLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, specimenHashRule)
	if err != nil {
		return nil, err
	}
	return &ProofBlockSpecimenProductionProofSubmittedIterator{contract: _Proof.contract, event: "BlockSpecimenProductionProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenProductionProofSubmitted is a free log subscription operation binding the contract event 0xd6167162e246a2d474fb881cb6dee451485884b64f01655f8f668ee415cbf9b2.
//
// Solidity: event BlockSpecimenProductionProofSubmitted(uint64 chainId, uint64 indexed blockHeight, uint64 specimenSize, uint64 specimenLength, bytes32 indexed specimenHash, string storageURL)
func (_Proof *ProofFilterer) WatchBlockSpecimenProductionProofSubmitted(opts *bind.WatchOpts, sink chan<- *ProofBlockSpecimenProductionProofSubmitted, blockHeight []uint64, specimenHash [][32]byte) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var specimenHashRule []interface{}
	for _, specimenHashItem := range specimenHash {
		specimenHashRule = append(specimenHashRule, specimenHashItem)
	}

	logs, sub, err := _Proof.contract.WatchLogs(opts, "BlockSpecimenProductionProofSubmitted", blockHeightRule, specimenHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofBlockSpecimenProductionProofSubmitted)
				if err := _Proof.contract.UnpackLog(event, "BlockSpecimenProductionProofSubmitted", log); err != nil {
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
func (_Proof *ProofFilterer) ParseBlockSpecimenProductionProofSubmitted(log types.Log) (*ProofBlockSpecimenProductionProofSubmitted, error) {
	event := new(ProofBlockSpecimenProductionProofSubmitted)
	if err := _Proof.contract.UnpackLog(event, "BlockSpecimenProductionProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofBlockSpecimenRewardAwardedIterator is returned from FilterBlockSpecimenRewardAwarded and is used to iterate over the raw logs and unpacked data for BlockSpecimenRewardAwarded events raised by the Proof contract.
type ProofBlockSpecimenRewardAwardedIterator struct {
	Event *ProofBlockSpecimenRewardAwarded // Event containing the contract specifics and raw log

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
func (it *ProofBlockSpecimenRewardAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofBlockSpecimenRewardAwarded)
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
		it.Event = new(ProofBlockSpecimenRewardAwarded)
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
func (it *ProofBlockSpecimenRewardAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofBlockSpecimenRewardAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofBlockSpecimenRewardAwarded represents a BlockSpecimenRewardAwarded event raised by the Proof contract.
type ProofBlockSpecimenRewardAwarded struct {
	ChainId         uint64
	BlockHeight     uint64
	Reward          *big.Int
	OperatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenRewardAwarded is a free log retrieval operation binding the contract event 0x105c3d60f0e35c6731e5828358dcbf158e40f122d68691c098931d2f57b661d1.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 chainId, uint64 indexed blockHeight, uint128 reward, address indexed operatorAddress)
func (_Proof *ProofFilterer) FilterBlockSpecimenRewardAwarded(opts *bind.FilterOpts, blockHeight []uint64, operatorAddress []common.Address) (*ProofBlockSpecimenRewardAwardedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Proof.contract.FilterLogs(opts, "BlockSpecimenRewardAwarded", blockHeightRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProofBlockSpecimenRewardAwardedIterator{contract: _Proof.contract, event: "BlockSpecimenRewardAwarded", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenRewardAwarded is a free log subscription operation binding the contract event 0x105c3d60f0e35c6731e5828358dcbf158e40f122d68691c098931d2f57b661d1.
//
// Solidity: event BlockSpecimenRewardAwarded(uint64 chainId, uint64 indexed blockHeight, uint128 reward, address indexed operatorAddress)
func (_Proof *ProofFilterer) WatchBlockSpecimenRewardAwarded(opts *bind.WatchOpts, sink chan<- *ProofBlockSpecimenRewardAwarded, blockHeight []uint64, operatorAddress []common.Address) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}

	var operatorAddressRule []interface{}
	for _, operatorAddressItem := range operatorAddress {
		operatorAddressRule = append(operatorAddressRule, operatorAddressItem)
	}

	logs, sub, err := _Proof.contract.WatchLogs(opts, "BlockSpecimenRewardAwarded", blockHeightRule, operatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofBlockSpecimenRewardAwarded)
				if err := _Proof.contract.UnpackLog(event, "BlockSpecimenRewardAwarded", log); err != nil {
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
func (_Proof *ProofFilterer) ParseBlockSpecimenRewardAwarded(log types.Log) (*ProofBlockSpecimenRewardAwarded, error) {
	event := new(ProofBlockSpecimenRewardAwarded)
	if err := _Proof.contract.UnpackLog(event, "BlockSpecimenRewardAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofBlockSpecimenRewardChangedIterator is returned from FilterBlockSpecimenRewardChanged and is used to iterate over the raw logs and unpacked data for BlockSpecimenRewardChanged events raised by the Proof contract.
type ProofBlockSpecimenRewardChangedIterator struct {
	Event *ProofBlockSpecimenRewardChanged // Event containing the contract specifics and raw log

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
func (it *ProofBlockSpecimenRewardChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofBlockSpecimenRewardChanged)
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
		it.Event = new(ProofBlockSpecimenRewardChanged)
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
func (it *ProofBlockSpecimenRewardChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofBlockSpecimenRewardChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofBlockSpecimenRewardChanged represents a BlockSpecimenRewardChanged event raised by the Proof contract.
type ProofBlockSpecimenRewardChanged struct {
	NewBlockSpecimenRewardAllocation *big.Int
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenRewardChanged is a free log retrieval operation binding the contract event 0x01eb821dd596243f2f8c5f6c7478e281b855ac12a9f4be2c486cb2778a0bb81e.
//
// Solidity: event BlockSpecimenRewardChanged(uint128 newBlockSpecimenRewardAllocation)
func (_Proof *ProofFilterer) FilterBlockSpecimenRewardChanged(opts *bind.FilterOpts) (*ProofBlockSpecimenRewardChangedIterator, error) {

	logs, sub, err := _Proof.contract.FilterLogs(opts, "BlockSpecimenRewardChanged")
	if err != nil {
		return nil, err
	}
	return &ProofBlockSpecimenRewardChangedIterator{contract: _Proof.contract, event: "BlockSpecimenRewardChanged", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenRewardChanged is a free log subscription operation binding the contract event 0x01eb821dd596243f2f8c5f6c7478e281b855ac12a9f4be2c486cb2778a0bb81e.
//
// Solidity: event BlockSpecimenRewardChanged(uint128 newBlockSpecimenRewardAllocation)
func (_Proof *ProofFilterer) WatchBlockSpecimenRewardChanged(opts *bind.WatchOpts, sink chan<- *ProofBlockSpecimenRewardChanged) (event.Subscription, error) {

	logs, sub, err := _Proof.contract.WatchLogs(opts, "BlockSpecimenRewardChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofBlockSpecimenRewardChanged)
				if err := _Proof.contract.UnpackLog(event, "BlockSpecimenRewardChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseBlockSpecimenRewardChanged(log types.Log) (*ProofBlockSpecimenRewardChanged, error) {
	event := new(ProofBlockSpecimenRewardChanged)
	if err := _Proof.contract.UnpackLog(event, "BlockSpecimenRewardChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofBlockSpecimenSessionFinalizedIterator is returned from FilterBlockSpecimenSessionFinalized and is used to iterate over the raw logs and unpacked data for BlockSpecimenSessionFinalized events raised by the Proof contract.
type ProofBlockSpecimenSessionFinalizedIterator struct {
	Event *ProofBlockSpecimenSessionFinalized // Event containing the contract specifics and raw log

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
func (it *ProofBlockSpecimenSessionFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofBlockSpecimenSessionFinalized)
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
		it.Event = new(ProofBlockSpecimenSessionFinalized)
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
func (it *ProofBlockSpecimenSessionFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofBlockSpecimenSessionFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofBlockSpecimenSessionFinalized represents a BlockSpecimenSessionFinalized event raised by the Proof contract.
type ProofBlockSpecimenSessionFinalized struct {
	BlockHeight *big.Int
	ProofHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenSessionFinalized is a free log retrieval operation binding the contract event 0xcbbbca2698029deaaa991c42053b23dd488b5caf488e5223da54a08a48d7c31e.
//
// Solidity: event BlockSpecimenSessionFinalized(uint128 indexed blockHeight, bytes32 indexed proofHash)
func (_Proof *ProofFilterer) FilterBlockSpecimenSessionFinalized(opts *bind.FilterOpts, blockHeight []*big.Int, proofHash [][32]byte) (*ProofBlockSpecimenSessionFinalizedIterator, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var proofHashRule []interface{}
	for _, proofHashItem := range proofHash {
		proofHashRule = append(proofHashRule, proofHashItem)
	}

	logs, sub, err := _Proof.contract.FilterLogs(opts, "BlockSpecimenSessionFinalized", blockHeightRule, proofHashRule)
	if err != nil {
		return nil, err
	}
	return &ProofBlockSpecimenSessionFinalizedIterator{contract: _Proof.contract, event: "BlockSpecimenSessionFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenSessionFinalized is a free log subscription operation binding the contract event 0xcbbbca2698029deaaa991c42053b23dd488b5caf488e5223da54a08a48d7c31e.
//
// Solidity: event BlockSpecimenSessionFinalized(uint128 indexed blockHeight, bytes32 indexed proofHash)
func (_Proof *ProofFilterer) WatchBlockSpecimenSessionFinalized(opts *bind.WatchOpts, sink chan<- *ProofBlockSpecimenSessionFinalized, blockHeight []*big.Int, proofHash [][32]byte) (event.Subscription, error) {

	var blockHeightRule []interface{}
	for _, blockHeightItem := range blockHeight {
		blockHeightRule = append(blockHeightRule, blockHeightItem)
	}
	var proofHashRule []interface{}
	for _, proofHashItem := range proofHash {
		proofHashRule = append(proofHashRule, proofHashItem)
	}

	logs, sub, err := _Proof.contract.WatchLogs(opts, "BlockSpecimenSessionFinalized", blockHeightRule, proofHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofBlockSpecimenSessionFinalized)
				if err := _Proof.contract.UnpackLog(event, "BlockSpecimenSessionFinalized", log); err != nil {
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
func (_Proof *ProofFilterer) ParseBlockSpecimenSessionFinalized(log types.Log) (*ProofBlockSpecimenSessionFinalized, error) {
	event := new(ProofBlockSpecimenSessionFinalized)
	if err := _Proof.contract.UnpackLog(event, "BlockSpecimenSessionFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofMinimumRequiredStakeChangedIterator is returned from FilterMinimumRequiredStakeChanged and is used to iterate over the raw logs and unpacked data for MinimumRequiredStakeChanged events raised by the Proof contract.
type ProofMinimumRequiredStakeChangedIterator struct {
	Event *ProofMinimumRequiredStakeChanged // Event containing the contract specifics and raw log

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
func (it *ProofMinimumRequiredStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofMinimumRequiredStakeChanged)
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
		it.Event = new(ProofMinimumRequiredStakeChanged)
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
func (it *ProofMinimumRequiredStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofMinimumRequiredStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofMinimumRequiredStakeChanged represents a MinimumRequiredStakeChanged event raised by the Proof contract.
type ProofMinimumRequiredStakeChanged struct {
	Role                [32]byte
	NewStakeRequirement *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinimumRequiredStakeChanged is a free log retrieval operation binding the contract event 0xcc5adc82271e3da3beed19bdd358519f24712369aa0cd14ec87e36a0eaa8efaa.
//
// Solidity: event MinimumRequiredStakeChanged(bytes32 indexed role, uint128 newStakeRequirement)
func (_Proof *ProofFilterer) FilterMinimumRequiredStakeChanged(opts *bind.FilterOpts, role [][32]byte) (*ProofMinimumRequiredStakeChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Proof.contract.FilterLogs(opts, "MinimumRequiredStakeChanged", roleRule)
	if err != nil {
		return nil, err
	}
	return &ProofMinimumRequiredStakeChangedIterator{contract: _Proof.contract, event: "MinimumRequiredStakeChanged", logs: logs, sub: sub}, nil
}

// WatchMinimumRequiredStakeChanged is a free log subscription operation binding the contract event 0xcc5adc82271e3da3beed19bdd358519f24712369aa0cd14ec87e36a0eaa8efaa.
//
// Solidity: event MinimumRequiredStakeChanged(bytes32 indexed role, uint128 newStakeRequirement)
func (_Proof *ProofFilterer) WatchMinimumRequiredStakeChanged(opts *bind.WatchOpts, sink chan<- *ProofMinimumRequiredStakeChanged, role [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _Proof.contract.WatchLogs(opts, "MinimumRequiredStakeChanged", roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofMinimumRequiredStakeChanged)
				if err := _Proof.contract.UnpackLog(event, "MinimumRequiredStakeChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseMinimumRequiredStakeChanged(log types.Log) (*ProofMinimumRequiredStakeChanged, error) {
	event := new(ProofMinimumRequiredStakeChanged)
	if err := _Proof.contract.UnpackLog(event, "MinimumRequiredStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofOperatorStartedRoleIterator is returned from FilterOperatorStartedRole and is used to iterate over the raw logs and unpacked data for OperatorStartedRole events raised by the Proof contract.
type ProofOperatorStartedRoleIterator struct {
	Event *ProofOperatorStartedRole // Event containing the contract specifics and raw log

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
func (it *ProofOperatorStartedRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOperatorStartedRole)
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
		it.Event = new(ProofOperatorStartedRole)
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
func (it *ProofOperatorStartedRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOperatorStartedRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOperatorStartedRole represents a OperatorStartedRole event raised by the Proof contract.
type ProofOperatorStartedRole struct {
	Role            [32]byte
	OperatorAddress common.Address
	ValidatorId     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorStartedRole is a free log retrieval operation binding the contract event 0x77bc260bf2591a08387f4fb053a665992d20fb87bd485d6210d285ec7706527b.
//
// Solidity: event OperatorStartedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_Proof *ProofFilterer) FilterOperatorStartedRole(opts *bind.FilterOpts, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (*ProofOperatorStartedRoleIterator, error) {

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

	logs, sub, err := _Proof.contract.FilterLogs(opts, "OperatorStartedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofOperatorStartedRoleIterator{contract: _Proof.contract, event: "OperatorStartedRole", logs: logs, sub: sub}, nil
}

// WatchOperatorStartedRole is a free log subscription operation binding the contract event 0x77bc260bf2591a08387f4fb053a665992d20fb87bd485d6210d285ec7706527b.
//
// Solidity: event OperatorStartedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_Proof *ProofFilterer) WatchOperatorStartedRole(opts *bind.WatchOpts, sink chan<- *ProofOperatorStartedRole, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Proof.contract.WatchLogs(opts, "OperatorStartedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOperatorStartedRole)
				if err := _Proof.contract.UnpackLog(event, "OperatorStartedRole", log); err != nil {
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
func (_Proof *ProofFilterer) ParseOperatorStartedRole(log types.Log) (*ProofOperatorStartedRole, error) {
	event := new(ProofOperatorStartedRole)
	if err := _Proof.contract.UnpackLog(event, "OperatorStartedRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofOperatorStoppedRoleIterator is returned from FilterOperatorStoppedRole and is used to iterate over the raw logs and unpacked data for OperatorStoppedRole events raised by the Proof contract.
type ProofOperatorStoppedRoleIterator struct {
	Event *ProofOperatorStoppedRole // Event containing the contract specifics and raw log

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
func (it *ProofOperatorStoppedRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofOperatorStoppedRole)
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
		it.Event = new(ProofOperatorStoppedRole)
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
func (it *ProofOperatorStoppedRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofOperatorStoppedRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofOperatorStoppedRole represents a OperatorStoppedRole event raised by the Proof contract.
type ProofOperatorStoppedRole struct {
	Role            [32]byte
	OperatorAddress common.Address
	ValidatorId     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorStoppedRole is a free log retrieval operation binding the contract event 0xfcf71fe4ab3a430dbd2e89438168cf16c179f2ed54e627d026157b3582761042.
//
// Solidity: event OperatorStoppedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_Proof *ProofFilterer) FilterOperatorStoppedRole(opts *bind.FilterOpts, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (*ProofOperatorStoppedRoleIterator, error) {

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

	logs, sub, err := _Proof.contract.FilterLogs(opts, "OperatorStoppedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ProofOperatorStoppedRoleIterator{contract: _Proof.contract, event: "OperatorStoppedRole", logs: logs, sub: sub}, nil
}

// WatchOperatorStoppedRole is a free log subscription operation binding the contract event 0xfcf71fe4ab3a430dbd2e89438168cf16c179f2ed54e627d026157b3582761042.
//
// Solidity: event OperatorStoppedRole(bytes32 indexed role, address indexed operatorAddress, uint128 indexed validatorId)
func (_Proof *ProofFilterer) WatchOperatorStoppedRole(opts *bind.WatchOpts, sink chan<- *ProofOperatorStoppedRole, role [][32]byte, operatorAddress []common.Address, validatorId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Proof.contract.WatchLogs(opts, "OperatorStoppedRole", roleRule, operatorAddressRule, validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofOperatorStoppedRole)
				if err := _Proof.contract.UnpackLog(event, "OperatorStoppedRole", log); err != nil {
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
func (_Proof *ProofFilterer) ParseOperatorStoppedRole(log types.Log) (*ProofOperatorStoppedRole, error) {
	event := new(ProofOperatorStoppedRole)
	if err := _Proof.contract.UnpackLog(event, "OperatorStoppedRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Proof contract.
type ProofRoleAdminChangedIterator struct {
	Event *ProofRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProofRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofRoleAdminChanged)
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
		it.Event = new(ProofRoleAdminChanged)
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
func (it *ProofRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofRoleAdminChanged represents a RoleAdminChanged event raised by the Proof contract.
type ProofRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Proof *ProofFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProofRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Proof.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProofRoleAdminChangedIterator{contract: _Proof.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Proof *ProofFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProofRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Proof.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofRoleAdminChanged)
				if err := _Proof.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseRoleAdminChanged(log types.Log) (*ProofRoleAdminChanged, error) {
	event := new(ProofRoleAdminChanged)
	if err := _Proof.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Proof contract.
type ProofRoleGrantedIterator struct {
	Event *ProofRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProofRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofRoleGranted)
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
		it.Event = new(ProofRoleGranted)
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
func (it *ProofRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofRoleGranted represents a RoleGranted event raised by the Proof contract.
type ProofRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Proof *ProofFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProofRoleGrantedIterator, error) {

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

	logs, sub, err := _Proof.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProofRoleGrantedIterator{contract: _Proof.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Proof *ProofFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProofRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Proof.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofRoleGranted)
				if err := _Proof.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Proof *ProofFilterer) ParseRoleGranted(log types.Log) (*ProofRoleGranted, error) {
	event := new(ProofRoleGranted)
	if err := _Proof.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Proof contract.
type ProofRoleRevokedIterator struct {
	Event *ProofRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProofRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofRoleRevoked)
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
		it.Event = new(ProofRoleRevoked)
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
func (it *ProofRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofRoleRevoked represents a RoleRevoked event raised by the Proof contract.
type ProofRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Proof *ProofFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProofRoleRevokedIterator, error) {

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

	logs, sub, err := _Proof.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProofRoleRevokedIterator{contract: _Proof.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Proof *ProofFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProofRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Proof.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofRoleRevoked)
				if err := _Proof.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Proof *ProofFilterer) ParseRoleRevoked(log types.Log) (*ProofRoleRevoked, error) {
	event := new(ProofRoleRevoked)
	if err := _Proof.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofSpecimenSessionDurationChangedIterator is returned from FilterSpecimenSessionDurationChanged and is used to iterate over the raw logs and unpacked data for SpecimenSessionDurationChanged events raised by the Proof contract.
type ProofSpecimenSessionDurationChangedIterator struct {
	Event *ProofSpecimenSessionDurationChanged // Event containing the contract specifics and raw log

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
func (it *ProofSpecimenSessionDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofSpecimenSessionDurationChanged)
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
		it.Event = new(ProofSpecimenSessionDurationChanged)
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
func (it *ProofSpecimenSessionDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofSpecimenSessionDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofSpecimenSessionDurationChanged represents a SpecimenSessionDurationChanged event raised by the Proof contract.
type ProofSpecimenSessionDurationChanged struct {
	NewSessionDuration uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSpecimenSessionDurationChanged is a free log retrieval operation binding the contract event 0x94bc488f4d9a985dd5f9d11e8f0a614a62828888eb65b704a90fa852be937549.
//
// Solidity: event SpecimenSessionDurationChanged(uint64 newSessionDuration)
func (_Proof *ProofFilterer) FilterSpecimenSessionDurationChanged(opts *bind.FilterOpts) (*ProofSpecimenSessionDurationChangedIterator, error) {

	logs, sub, err := _Proof.contract.FilterLogs(opts, "SpecimenSessionDurationChanged")
	if err != nil {
		return nil, err
	}
	return &ProofSpecimenSessionDurationChangedIterator{contract: _Proof.contract, event: "SpecimenSessionDurationChanged", logs: logs, sub: sub}, nil
}

// WatchSpecimenSessionDurationChanged is a free log subscription operation binding the contract event 0x94bc488f4d9a985dd5f9d11e8f0a614a62828888eb65b704a90fa852be937549.
//
// Solidity: event SpecimenSessionDurationChanged(uint64 newSessionDuration)
func (_Proof *ProofFilterer) WatchSpecimenSessionDurationChanged(opts *bind.WatchOpts, sink chan<- *ProofSpecimenSessionDurationChanged) (event.Subscription, error) {

	logs, sub, err := _Proof.contract.WatchLogs(opts, "SpecimenSessionDurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofSpecimenSessionDurationChanged)
				if err := _Proof.contract.UnpackLog(event, "SpecimenSessionDurationChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseSpecimenSessionDurationChanged(log types.Log) (*ProofSpecimenSessionDurationChanged, error) {
	event := new(ProofSpecimenSessionDurationChanged)
	if err := _Proof.contract.UnpackLog(event, "SpecimenSessionDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofSpecimenSessionQuorumChangedIterator is returned from FilterSpecimenSessionQuorumChanged and is used to iterate over the raw logs and unpacked data for SpecimenSessionQuorumChanged events raised by the Proof contract.
type ProofSpecimenSessionQuorumChangedIterator struct {
	Event *ProofSpecimenSessionQuorumChanged // Event containing the contract specifics and raw log

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
func (it *ProofSpecimenSessionQuorumChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofSpecimenSessionQuorumChanged)
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
		it.Event = new(ProofSpecimenSessionQuorumChanged)
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
func (it *ProofSpecimenSessionQuorumChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofSpecimenSessionQuorumChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofSpecimenSessionQuorumChanged represents a SpecimenSessionQuorumChanged event raised by the Proof contract.
type ProofSpecimenSessionQuorumChanged struct {
	NewQuorumThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSpecimenSessionQuorumChanged is a free log retrieval operation binding the contract event 0x7ab8150f5b613293e16e3a99396812661a51bad017fbe4204fb8faef1c315cb0.
//
// Solidity: event SpecimenSessionQuorumChanged(uint64 newQuorumThreshold)
func (_Proof *ProofFilterer) FilterSpecimenSessionQuorumChanged(opts *bind.FilterOpts) (*ProofSpecimenSessionQuorumChangedIterator, error) {

	logs, sub, err := _Proof.contract.FilterLogs(opts, "SpecimenSessionQuorumChanged")
	if err != nil {
		return nil, err
	}
	return &ProofSpecimenSessionQuorumChangedIterator{contract: _Proof.contract, event: "SpecimenSessionQuorumChanged", logs: logs, sub: sub}, nil
}

// WatchSpecimenSessionQuorumChanged is a free log subscription operation binding the contract event 0x7ab8150f5b613293e16e3a99396812661a51bad017fbe4204fb8faef1c315cb0.
//
// Solidity: event SpecimenSessionQuorumChanged(uint64 newQuorumThreshold)
func (_Proof *ProofFilterer) WatchSpecimenSessionQuorumChanged(opts *bind.WatchOpts, sink chan<- *ProofSpecimenSessionQuorumChanged) (event.Subscription, error) {

	logs, sub, err := _Proof.contract.WatchLogs(opts, "SpecimenSessionQuorumChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofSpecimenSessionQuorumChanged)
				if err := _Proof.contract.UnpackLog(event, "SpecimenSessionQuorumChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseSpecimenSessionQuorumChanged(log types.Log) (*ProofSpecimenSessionQuorumChanged, error) {
	event := new(ProofSpecimenSessionQuorumChanged)
	if err := _Proof.contract.UnpackLog(event, "SpecimenSessionQuorumChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofStakingInterfaceChangedIterator is returned from FilterStakingInterfaceChanged and is used to iterate over the raw logs and unpacked data for StakingInterfaceChanged events raised by the Proof contract.
type ProofStakingInterfaceChangedIterator struct {
	Event *ProofStakingInterfaceChanged // Event containing the contract specifics and raw log

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
func (it *ProofStakingInterfaceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofStakingInterfaceChanged)
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
		it.Event = new(ProofStakingInterfaceChanged)
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
func (it *ProofStakingInterfaceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofStakingInterfaceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofStakingInterfaceChanged represents a StakingInterfaceChanged event raised by the Proof contract.
type ProofStakingInterfaceChanged struct {
	NewInterfaceAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStakingInterfaceChanged is a free log retrieval operation binding the contract event 0x70016f37fc9a299f674d1e3083a27743406649810887ed947a79884b064d2de9.
//
// Solidity: event StakingInterfaceChanged(address newInterfaceAddress)
func (_Proof *ProofFilterer) FilterStakingInterfaceChanged(opts *bind.FilterOpts) (*ProofStakingInterfaceChangedIterator, error) {

	logs, sub, err := _Proof.contract.FilterLogs(opts, "StakingInterfaceChanged")
	if err != nil {
		return nil, err
	}
	return &ProofStakingInterfaceChangedIterator{contract: _Proof.contract, event: "StakingInterfaceChanged", logs: logs, sub: sub}, nil
}

// WatchStakingInterfaceChanged is a free log subscription operation binding the contract event 0x70016f37fc9a299f674d1e3083a27743406649810887ed947a79884b064d2de9.
//
// Solidity: event StakingInterfaceChanged(address newInterfaceAddress)
func (_Proof *ProofFilterer) WatchStakingInterfaceChanged(opts *bind.WatchOpts, sink chan<- *ProofStakingInterfaceChanged) (event.Subscription, error) {

	logs, sub, err := _Proof.contract.WatchLogs(opts, "StakingInterfaceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofStakingInterfaceChanged)
				if err := _Proof.contract.UnpackLog(event, "StakingInterfaceChanged", log); err != nil {
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
func (_Proof *ProofFilterer) ParseStakingInterfaceChanged(log types.Log) (*ProofStakingInterfaceChanged, error) {
	event := new(ProofStakingInterfaceChanged)
	if err := _Proof.contract.UnpackLog(event, "StakingInterfaceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
