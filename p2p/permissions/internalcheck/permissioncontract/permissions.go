// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissioncontract

import (
	"strings"
	"math/big"
	//ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)


// PermissionsABI is the input ABI used to generate the binding from.
const PermissionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"previousaccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isadding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addingmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"suspendNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousenode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeinfo\",\"outputs\":[{\"name\":\"enode\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\"},{\"name\":\"votecount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"checkNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeconformations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuspention\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspentionmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogOfAddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogOfSuspentionNode\",\"type\":\"event\"}]"


// PermissionsBin is the compiled bytecode used for deploying new contracts.
const PermissionsBin = `0x608060405234801561001057600080fd5b50600060018190556004556108c48061002a6000396000f3006080604052600436106100c45763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663118e696781146100c95780632da0bca8146100fa57806360c229f7146101235780636a19f6ad14610138578063768493761461015f5780638ec08b2014610185578063b1b49f25146101a9578063b3b64ce2146101be578063b50d6f7514610210578063c86d87b914610228578063d6336a3314610240578063eae897c714610255578063ff5c11f81461026a575b600080fd5b3480156100d557600080fd5b506100de61027f565b60408051600160a060020a039092168252519081900360200190f35b34801561010657600080fd5b5061010f61028e565b604080519115158252519081900360200190f35b34801561012f57600080fd5b5061010f61029c565b34801561014457600080fd5b5061014d6102ab565b60408051918252519081900360200190f35b34801561016b57600080fd5b50610183600435600160a060020a03602435166102b1565b005b34801561019157600080fd5b50610183600435600160a060020a0360243516610341565b3480156101b557600080fd5b5061014d6103c7565b3480156101ca57600080fd5b506101e2600435600160a060020a03602435166103cd565b60408051948552600160a060020a039093166020850152901515838301526060830152519081900360800190f35b34801561021c57600080fd5b5061010f60043561040f565b34801561023457600080fd5b5061010f60043561043c565b34801561024c57600080fd5b5061010f610451565b34801561026157600080fd5b5061010f61045a565b34801561027657600080fd5b5061014d61046a565b600354600160a060020a031681565b600054610100900460ff1681565b60005462010000900460ff1681565b60045481565b60005462010000900460ff16151560011480156102cf575081600254145b80156102e85750600354600160a060020a038281169116145b156102fc576102f78282610470565b61033d565b60005462010000900460ff1615801561031d5750600054610100900460ff16155b1561033d576000805462ff000019166201000017905561033d8282610470565b5050565b6000546301000000900460ff1615156001148015610360575081600254145b80156103795750600354600160a060020a038281169116145b15610388576102f7828261069e565b6000546301000000900460ff161580156103a5575060005460ff16155b1561033d576000805463ff0000001916630100000017905561033d828261069e565b60025481565b60066020908152600092835260408084209091529082529020805460018201546002909201549091600160a060020a0381169160a060020a90910460ff169084565b60008181526005602052604081205460ff1615156001141561043357506001610437565b5060005b919050565b60056020526000908152604090205460ff1681565b60005460ff1681565b6000546301000000900460ff1681565b60015481565b60008281526005602052604090205460ff161561048957fe5b60005460ff161561049657fe5b60005462010000900460ff1615156104aa57fe5b6000805461ff001916610100178155600154838252600660209081526040808420600160a060020a038616855290915290912060020154101561055f576000828152600660209081526040808320600160a060020a03851680855292529091208381556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805490910190555b6001546000838152600660209081526040808320600160a060020a0386168452909152902060020154141561062f576000828152600660209081526040808320600160a060020a0385168085528184528285208781556001808201805474ff00000000000000000000000000000000000000001973ffffffffffffffffffffffffffffffffffffffff1990911685171660a060020a179055865462ffff001916875588875260058652938620805460ff191685179055600480548501815591865291909352600201929092555490555b60028290556003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831690811790915560408051848152602081019290925280517fc77919a1ee986ff5e4c404dbe8c901c90a8c0240e2dd277e3ae9674cf738473c9281900390910190a15050565b600054610100900460ff16156106b057fe5b6000546301000000900460ff1615156106c557fe5b6000828152600660209081526040808320600160a060020a0385168452909152902060019081015460a060020a900460ff1615151461070057fe5b60008281526005602052604090205460ff16151560011461071d57fe5b6000805460ff19166001908117825554838252600660209081526040808420600160a060020a0386168552909152909120600201541061075957fe5b6000828152600660209081526040808320600160a060020a03851680855292529091208381556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805482019081905590541415610829576000828152600560209081526040808320805460ff19169055825463ff0000ff1916835560068252808320600160a060020a03851684529091528120600201556004805460001901908190556001555b60028290556003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03831690811790915560408051848152602081019290925280517fd7ff2ad5befbbf96c8be6d0e743cfc9cf9c181c502d12e3136f57fa0ada673ee9281900390910190a150505600a165627a7a723058206ee804e4f8f6cf866f233d919fd2075b56ae7f2de42cf95b9537c425f5553aa00029`

// DeployPermissions deploys a new Ethereum contract, binding an instance of Permissions to it.
func DeployPermissions(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Permissions, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermissionsBin), backend )
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Permissions{ PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract} }, nil
}


// Permissions is an auto generated Go binding around an Ethereum contract.
type Permissions struct {
	PermissionsCaller     // Read-only binding to the contract
	PermissionsTransactor // Write-only binding to the contract
	PermissionsFilterer   // Log filterer for contract events
}

// PermissionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionsSession struct {
	Contract     *Permissions        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsCallerSession struct {
	Contract *PermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsTransactorSession struct {
	Contract     *PermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PermissionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionsRaw struct {
	Contract *Permissions // Generic contract binding to access the raw methods on
}

// PermissionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionsCallerRaw struct {
	Contract *PermissionsCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionsTransactorRaw struct {
	Contract *PermissionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissions creates a new instance of Permissions, bound to a specific deployed contract.
func NewPermissions(address common.Address, backend bind.ContractBackend) (*Permissions, error) {
	contract, err := bindPermissions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissions{ PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract} }, nil
}

// NewPermissionsCaller creates a new read-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsCaller(address common.Address, caller bind.ContractCaller) (*PermissionsCaller, error) {
	contract, err := bindPermissions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsCaller{contract: contract}, nil
}

// NewPermissionsTransactor creates a new write-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionsTransactor, error) {
	contract, err := bindPermissions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsTransactor{contract: contract}, nil
}

// NewPermissionsFilterer creates a new log filterer instance of Permissions, bound to a specific deployed contract.
func NewPermissionsFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionsFilterer, error) {
	contract, err := bindPermissions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionsFilterer{contract: contract}, nil
}

// bindPermissions binds a generic wrapper to an already deployed contract.
func bindPermissions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.PermissionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transact(opts, method, params...)
}


// LimitOfVote is a free data retrieval call binding the contract method 0xff5c11f8.
//
// Solidity: function LimitOfVote() constant returns(uint256)
func (_Permissions *PermissionsCaller) LimitOfVote(opts *bind.CallOpts ) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "LimitOfVote" )
	return *ret0, err
}

// LimitOfVote is a free data retrieval call binding the contract method 0xff5c11f8.
//
// Solidity: function LimitOfVote() constant returns(uint256)
func (_Permissions *PermissionsSession) LimitOfVote() ( *big.Int,  error) {
	return _Permissions.Contract.LimitOfVote(&_Permissions.CallOpts )
}

// LimitOfVote is a free data retrieval call binding the contract method 0xff5c11f8.
//
// Solidity: function LimitOfVote() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) LimitOfVote() ( *big.Int,  error) {
	return _Permissions.Contract.LimitOfVote(&_Permissions.CallOpts )
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsCaller) NodeCount(opts *bind.CallOpts ) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "NodeCount" )
	return *ret0, err
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsSession) NodeCount() ( *big.Int,  error) {
	return _Permissions.Contract.NodeCount(&_Permissions.CallOpts )
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) NodeCount() ( *big.Int,  error) {
	return _Permissions.Contract.NodeCount(&_Permissions.CallOpts )
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsCaller) Addingmutex(opts *bind.CallOpts ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "addingmutex" )
	return *ret0, err
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsSession) Addingmutex() ( bool,  error) {
	return _Permissions.Contract.Addingmutex(&_Permissions.CallOpts )
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Addingmutex() ( bool,  error) {
	return _Permissions.Contract.Addingmutex(&_Permissions.CallOpts )
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsCaller) CheckNode(opts *bind.CallOpts , _enode [32]byte ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "checkNode" , _enode)
	return *ret0, err
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsSession) CheckNode( _enode [32]byte ) ( bool,  error) {
	return _Permissions.Contract.CheckNode(&_Permissions.CallOpts , _enode)
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsCallerSession) CheckNode( _enode [32]byte ) ( bool,  error) {
	return _Permissions.Contract.CheckNode(&_Permissions.CallOpts , _enode)
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsCaller) Isadding(opts *bind.CallOpts ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "isadding" )
	return *ret0, err
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsSession) Isadding() ( bool,  error) {
	return _Permissions.Contract.Isadding(&_Permissions.CallOpts )
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Isadding() ( bool,  error) {
	return _Permissions.Contract.Isadding(&_Permissions.CallOpts )
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsCaller) Issuspention(opts *bind.CallOpts ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "issuspention" )
	return *ret0, err
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsSession) Issuspention() ( bool,  error) {
	return _Permissions.Contract.Issuspention(&_Permissions.CallOpts )
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Issuspention() ( bool,  error) {
	return _Permissions.Contract.Issuspention(&_Permissions.CallOpts )
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsCaller) Nodeconformations(opts *bind.CallOpts , arg0 [32]byte ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "nodeconformations" , arg0)
	return *ret0, err
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsSession) Nodeconformations( arg0 [32]byte ) ( bool,  error) {
	return _Permissions.Contract.Nodeconformations(&_Permissions.CallOpts , arg0)
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsCallerSession) Nodeconformations( arg0 [32]byte ) ( bool,  error) {
	return _Permissions.Contract.Nodeconformations(&_Permissions.CallOpts , arg0)
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsCaller) Nodeinfo(opts *bind.CallOpts , arg0 [32]byte , arg1 common.Address ) (struct{ Enode [32]byte;Account common.Address;Flag bool;Votecount *big.Int; }, error) {
	ret := new(struct{
		Enode [32]byte
		Account common.Address
		Flag bool
		Votecount *big.Int

	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "nodeinfo" , arg0, arg1)
	return *ret, err
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsSession) Nodeinfo( arg0 [32]byte , arg1 common.Address ) (struct{ Enode [32]byte;Account common.Address;Flag bool;Votecount *big.Int; },  error) {
	return _Permissions.Contract.Nodeinfo(&_Permissions.CallOpts , arg0, arg1)
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsCallerSession) Nodeinfo( arg0 [32]byte , arg1 common.Address ) (struct{ Enode [32]byte;Account common.Address;Flag bool;Votecount *big.Int; },  error) {
	return _Permissions.Contract.Nodeinfo(&_Permissions.CallOpts , arg0, arg1)
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsCaller) Previousaccount(opts *bind.CallOpts ) (common.Address, error) {
	var (
		ret0 = new(common.Address)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "previousaccount" )
	return *ret0, err
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsSession) Previousaccount() ( common.Address,  error) {
	return _Permissions.Contract.Previousaccount(&_Permissions.CallOpts )
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsCallerSession) Previousaccount() ( common.Address,  error) {
	return _Permissions.Contract.Previousaccount(&_Permissions.CallOpts )
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsCaller) Previousenode(opts *bind.CallOpts ) ([32]byte, error) {
	var (
		ret0 = new([32]byte)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "previousenode" )
	return *ret0, err
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsSession) Previousenode() ( [32]byte,  error) {
	return _Permissions.Contract.Previousenode(&_Permissions.CallOpts )
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsCallerSession) Previousenode() ( [32]byte,  error) {
	return _Permissions.Contract.Previousenode(&_Permissions.CallOpts )
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsCaller) Suspentionmutex(opts *bind.CallOpts ) (bool, error) {
	var (
		ret0 = new(bool)

	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "suspentionmutex" )
	return *ret0, err
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsSession) Suspentionmutex() ( bool,  error) {
	return _Permissions.Contract.Suspentionmutex(&_Permissions.CallOpts )
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Suspentionmutex() ( bool,  error) {
	return _Permissions.Contract.Suspentionmutex(&_Permissions.CallOpts )
}



// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactor) AddNode(opts *bind.TransactOpts , _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addNode" , _enode, _account)
}

// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsSession) AddNode( _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts , _enode, _account)
}

// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactorSession) AddNode( _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts , _enode, _account)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactor) SuspendNode(opts *bind.TransactOpts , _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "suspendNode" , _enode, _account)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsSession) SuspendNode( _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts , _enode, _account)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactorSession) SuspendNode( _enode [32]byte , _account common.Address ) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts , _enode, _account)
}




