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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seq\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"extractWorker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainHeightPos\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainHeightLen\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"specimenHash\",\"type\":\"uint256\"}],\"name\":\"BlockSpecimenPublicationProofAppended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCK_SPECIMEN_PRODUCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRequiredStakeForRole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setRequiredStakeForRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getStakedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setStakedBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPreapprovedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isPreapprovedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSufficientlyStakedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"isSufficientlyStakedForRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRolePreapproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainHeightPos\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainHeightLen\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"specimenSize\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"specimenHash\",\"type\":\"uint256\"}],\"name\":\"proveBlockSpecimenProduced\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProofChain.Contract.DEFAULTADMINROLE(&_ProofChain.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProofChain.Contract.DEFAULTADMINROLE(&_ProofChain.CallOpts)
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

// STAKINGORACLEROLE is a free data retrieval call binding the contract method 0xa0f1dbf8.
//
// Solidity: function STAKING_ORACLE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCaller) STAKINGORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "STAKING_ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGORACLEROLE is a free data retrieval call binding the contract method 0xa0f1dbf8.
//
// Solidity: function STAKING_ORACLE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainSession) STAKINGORACLEROLE() ([32]byte, error) {
	return _ProofChain.Contract.STAKINGORACLEROLE(&_ProofChain.CallOpts)
}

// STAKINGORACLEROLE is a free data retrieval call binding the contract method 0xa0f1dbf8.
//
// Solidity: function STAKING_ORACLE_ROLE() view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) STAKINGORACLEROLE() ([32]byte, error) {
	return _ProofChain.Contract.STAKINGORACLEROLE(&_ProofChain.CallOpts)
}

// GetRequiredStakeForRole is a free data retrieval call binding the contract method 0x5efdfe76.
//
// Solidity: function getRequiredStakeForRole(bytes32 role) view returns(uint256)
func (_ProofChain *ProofChainCaller) GetRequiredStakeForRole(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "getRequiredStakeForRole", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredStakeForRole is a free data retrieval call binding the contract method 0x5efdfe76.
//
// Solidity: function getRequiredStakeForRole(bytes32 role) view returns(uint256)
func (_ProofChain *ProofChainSession) GetRequiredStakeForRole(role [32]byte) (*big.Int, error) {
	return _ProofChain.Contract.GetRequiredStakeForRole(&_ProofChain.CallOpts, role)
}

// GetRequiredStakeForRole is a free data retrieval call binding the contract method 0x5efdfe76.
//
// Solidity: function getRequiredStakeForRole(bytes32 role) view returns(uint256)
func (_ProofChain *ProofChainCallerSession) GetRequiredStakeForRole(role [32]byte) (*big.Int, error) {
	return _ProofChain.Contract.GetRequiredStakeForRole(&_ProofChain.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProofChain *ProofChainCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProofChain *ProofChainSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProofChain.Contract.GetRoleAdmin(&_ProofChain.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProofChain *ProofChainCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProofChain.Contract.GetRoleAdmin(&_ProofChain.CallOpts, role)
}

// GetStakedBalance is a free data retrieval call binding the contract method 0x3a02a42d.
//
// Solidity: function getStakedBalance(address addr) view returns(uint256)
func (_ProofChain *ProofChainCaller) GetStakedBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "getStakedBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedBalance is a free data retrieval call binding the contract method 0x3a02a42d.
//
// Solidity: function getStakedBalance(address addr) view returns(uint256)
func (_ProofChain *ProofChainSession) GetStakedBalance(addr common.Address) (*big.Int, error) {
	return _ProofChain.Contract.GetStakedBalance(&_ProofChain.CallOpts, addr)
}

// GetStakedBalance is a free data retrieval call binding the contract method 0x3a02a42d.
//
// Solidity: function getStakedBalance(address addr) view returns(uint256)
func (_ProofChain *ProofChainCallerSession) GetStakedBalance(addr common.Address) (*big.Int, error) {
	return _ProofChain.Contract.GetStakedBalance(&_ProofChain.CallOpts, addr)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.HasRole(&_ProofChain.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.HasRole(&_ProofChain.CallOpts, role, account)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCaller) IsPreapprovedForRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isPreapprovedForRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainSession) IsPreapprovedForRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole(&_ProofChain.CallOpts, role, account)
}

// IsPreapprovedForRole is a free data retrieval call binding the contract method 0x6133dbf4.
//
// Solidity: function isPreapprovedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsPreapprovedForRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole(&_ProofChain.CallOpts, role, account)
}

// IsPreapprovedForRole0 is a free data retrieval call binding the contract method 0xf83f2523.
//
// Solidity: function isPreapprovedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainCaller) IsPreapprovedForRole0(opts *bind.CallOpts, role [32]byte) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isPreapprovedForRole0", role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPreapprovedForRole0 is a free data retrieval call binding the contract method 0xf83f2523.
//
// Solidity: function isPreapprovedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainSession) IsPreapprovedForRole0(role [32]byte) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole0(&_ProofChain.CallOpts, role)
}

// IsPreapprovedForRole0 is a free data retrieval call binding the contract method 0xf83f2523.
//
// Solidity: function isPreapprovedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsPreapprovedForRole0(role [32]byte) (bool, error) {
	return _ProofChain.Contract.IsPreapprovedForRole0(&_ProofChain.CallOpts, role)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x0d03ae25.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCaller) IsSufficientlyStakedForRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isSufficientlyStakedForRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x0d03ae25.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainSession) IsSufficientlyStakedForRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole(&_ProofChain.CallOpts, role, account)
}

// IsSufficientlyStakedForRole is a free data retrieval call binding the contract method 0x0d03ae25.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role, address account) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsSufficientlyStakedForRole(role [32]byte, account common.Address) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole(&_ProofChain.CallOpts, role, account)
}

// IsSufficientlyStakedForRole0 is a free data retrieval call binding the contract method 0x3fe9c720.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainCaller) IsSufficientlyStakedForRole0(opts *bind.CallOpts, role [32]byte) (bool, error) {
	var out []interface{}
	err := _ProofChain.contract.Call(opts, &out, "isSufficientlyStakedForRole0", role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSufficientlyStakedForRole0 is a free data retrieval call binding the contract method 0x3fe9c720.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainSession) IsSufficientlyStakedForRole0(role [32]byte) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole0(&_ProofChain.CallOpts, role)
}

// IsSufficientlyStakedForRole0 is a free data retrieval call binding the contract method 0x3fe9c720.
//
// Solidity: function isSufficientlyStakedForRole(bytes32 role) view returns(bool)
func (_ProofChain *ProofChainCallerSession) IsSufficientlyStakedForRole0(role [32]byte) (bool, error) {
	return _ProofChain.Contract.IsSufficientlyStakedForRole0(&_ProofChain.CallOpts, role)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ProofChain *ProofChainTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ProofChain *ProofChainSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.Initialize(&_ProofChain.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ProofChain *ProofChainTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ProofChain.Contract.Initialize(&_ProofChain.TransactOpts, initialOwner)
}

// ProveBlockSpecimenProduced is a paid mutator transaction binding the contract method 0x91ede81b.
//
// Solidity: function proveBlockSpecimenProduced(uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash) returns()
func (_ProofChain *ProofChainTransactor) ProveBlockSpecimenProduced(opts *bind.TransactOpts, chainID uint64, chainHeightPos uint64, chainHeightLen uint64, specimenSize uint64, specimenHash *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "proveBlockSpecimenProduced", chainID, chainHeightPos, chainHeightLen, specimenSize, specimenHash)
}

// ProveBlockSpecimenProduced is a paid mutator transaction binding the contract method 0x91ede81b.
//
// Solidity: function proveBlockSpecimenProduced(uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash) returns()
func (_ProofChain *ProofChainSession) ProveBlockSpecimenProduced(chainID uint64, chainHeightPos uint64, chainHeightLen uint64, specimenSize uint64, specimenHash *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.ProveBlockSpecimenProduced(&_ProofChain.TransactOpts, chainID, chainHeightPos, chainHeightLen, specimenSize, specimenHash)
}

// ProveBlockSpecimenProduced is a paid mutator transaction binding the contract method 0x91ede81b.
//
// Solidity: function proveBlockSpecimenProduced(uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash) returns()
func (_ProofChain *ProofChainTransactorSession) ProveBlockSpecimenProduced(chainID uint64, chainHeightPos uint64, chainHeightLen uint64, specimenSize uint64, specimenHash *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.ProveBlockSpecimenProduced(&_ProofChain.TransactOpts, chainID, chainHeightPos, chainHeightLen, specimenSize, specimenHash)
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

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x57f2a494.
//
// Solidity: function setRequiredStakeForRole(bytes32 role, uint256 amount) returns()
func (_ProofChain *ProofChainTransactor) SetRequiredStakeForRole(opts *bind.TransactOpts, role [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setRequiredStakeForRole", role, amount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x57f2a494.
//
// Solidity: function setRequiredStakeForRole(bytes32 role, uint256 amount) returns()
func (_ProofChain *ProofChainSession) SetRequiredStakeForRole(role [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetRequiredStakeForRole(&_ProofChain.TransactOpts, role, amount)
}

// SetRequiredStakeForRole is a paid mutator transaction binding the contract method 0x57f2a494.
//
// Solidity: function setRequiredStakeForRole(bytes32 role, uint256 amount) returns()
func (_ProofChain *ProofChainTransactorSession) SetRequiredStakeForRole(role [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetRequiredStakeForRole(&_ProofChain.TransactOpts, role, amount)
}

// SetStakedBalance is a paid mutator transaction binding the contract method 0xfbb62efe.
//
// Solidity: function setStakedBalance(address addr, uint256 amount) returns()
func (_ProofChain *ProofChainTransactor) SetStakedBalance(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.contract.Transact(opts, "setStakedBalance", addr, amount)
}

// SetStakedBalance is a paid mutator transaction binding the contract method 0xfbb62efe.
//
// Solidity: function setStakedBalance(address addr, uint256 amount) returns()
func (_ProofChain *ProofChainSession) SetStakedBalance(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetStakedBalance(&_ProofChain.TransactOpts, addr, amount)
}

// SetStakedBalance is a paid mutator transaction binding the contract method 0xfbb62efe.
//
// Solidity: function setStakedBalance(address addr, uint256 amount) returns()
func (_ProofChain *ProofChainTransactorSession) SetStakedBalance(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProofChain.Contract.SetStakedBalance(&_ProofChain.TransactOpts, addr, amount)
}

// ProofChainBlockSpecimenPublicationProofAppendedIterator is returned from FilterBlockSpecimenPublicationProofAppended and is used to iterate over the raw logs and unpacked data for BlockSpecimenPublicationProofAppended events raised by the ProofChain contract.
type ProofChainBlockSpecimenPublicationProofAppendedIterator struct {
	Event *ProofChainBlockSpecimenPublicationProofAppended // Event containing the contract specifics and raw log

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
func (it *ProofChainBlockSpecimenPublicationProofAppendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofChainBlockSpecimenPublicationProofAppended)
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
		it.Event = new(ProofChainBlockSpecimenPublicationProofAppended)
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
func (it *ProofChainBlockSpecimenPublicationProofAppendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofChainBlockSpecimenPublicationProofAppendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofChainBlockSpecimenPublicationProofAppended represents a BlockSpecimenPublicationProofAppended event raised by the ProofChain contract.
type ProofChainBlockSpecimenPublicationProofAppended struct {
	Seq            uint64
	ExtractWorker  common.Address
	ChainID        uint64
	ChainHeightPos uint64
	ChainHeightLen uint64
	SpecimenSize   uint64
	SpecimenHash   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockSpecimenPublicationProofAppended is a free log retrieval operation binding the contract event 0x3632b6df680763cbf71da60bfdbd60c965f5d33db6b5ca24acc84378cf86a4c3.
//
// Solidity: event BlockSpecimenPublicationProofAppended(uint64 seq, address extractWorker, uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash)
func (_ProofChain *ProofChainFilterer) FilterBlockSpecimenPublicationProofAppended(opts *bind.FilterOpts) (*ProofChainBlockSpecimenPublicationProofAppendedIterator, error) {

	logs, sub, err := _ProofChain.contract.FilterLogs(opts, "BlockSpecimenPublicationProofAppended")
	if err != nil {
		return nil, err
	}
	return &ProofChainBlockSpecimenPublicationProofAppendedIterator{contract: _ProofChain.contract, event: "BlockSpecimenPublicationProofAppended", logs: logs, sub: sub}, nil
}

// WatchBlockSpecimenPublicationProofAppended is a free log subscription operation binding the contract event 0x3632b6df680763cbf71da60bfdbd60c965f5d33db6b5ca24acc84378cf86a4c3.
//
// Solidity: event BlockSpecimenPublicationProofAppended(uint64 seq, address extractWorker, uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash)
func (_ProofChain *ProofChainFilterer) WatchBlockSpecimenPublicationProofAppended(opts *bind.WatchOpts, sink chan<- *ProofChainBlockSpecimenPublicationProofAppended) (event.Subscription, error) {

	logs, sub, err := _ProofChain.contract.WatchLogs(opts, "BlockSpecimenPublicationProofAppended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofChainBlockSpecimenPublicationProofAppended)
				if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenPublicationProofAppended", log); err != nil {
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

// ParseBlockSpecimenPublicationProofAppended is a log parse operation binding the contract event 0x3632b6df680763cbf71da60bfdbd60c965f5d33db6b5ca24acc84378cf86a4c3.
//
// Solidity: event BlockSpecimenPublicationProofAppended(uint64 seq, address extractWorker, uint64 chainID, uint64 chainHeightPos, uint64 chainHeightLen, uint64 specimenSize, uint256 specimenHash)
func (_ProofChain *ProofChainFilterer) ParseBlockSpecimenPublicationProofAppended(log types.Log) (*ProofChainBlockSpecimenPublicationProofAppended, error) {
	event := new(ProofChainBlockSpecimenPublicationProofAppended)
	if err := _ProofChain.contract.UnpackLog(event, "BlockSpecimenPublicationProofAppended", log); err != nil {
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
