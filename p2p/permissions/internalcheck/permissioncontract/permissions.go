// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissioncontract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PermissionsABI is the input ABI used to generate the binding from.
const PermissionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"previousaccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensuspercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isadding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"setConsensus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addingmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"suspendNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensus\",\"outputs\":[{\"name\":\"limit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetProcess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousenode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodeinfo\",\"outputs\":[{\"name\":\"enode\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\"},{\"name\":\"votecount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"checkNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeconformations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuspention\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspentionmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"adding\",\"type\":\"string\"}],\"name\":\"LogOfAddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"suspend\",\"type\":\"string\"}],\"name\":\"LogOfSuspentionNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consensuslimit\",\"type\":\"uint256\"}],\"name\":\"LogOfSetConsensus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousenode\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"previousaccount\",\"type\":\"address\"}],\"name\":\"LogOfResetProcess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enode\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"LogOfaddingConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enode\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"LogOfsuspentionConsensusMeet\",\"type\":\"event\"}]"

// PermissionsBin is the compiled bytecode used for deploying new contracts.
const PermissionsBin = `0x608060405234801561001057600080fd5b5060006001819055600455610b7f8061002a6000396000f3006080604052600436106100f05763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663118e696781146100f55780631867f60c146101265780632da0bca81461014d5780633549ef171461017657806360c229f7146101905780636a19f6ad146101a557806376849376146101ba5780638ec08b20146101de5780638ef3f761146102025780639f9a19d314610217578063b1b49f251461022c578063b3b64ce214610241578063b50d6f7514610293578063c86d87b9146102ab578063d6336a33146102c3578063eae897c7146102d8578063ff5c11f8146102ed575b600080fd5b34801561010157600080fd5b5061010a610302565b60408051600160a060020a039092168252519081900360200190f35b34801561013257600080fd5b5061013b610311565b60408051918252519081900360200190f35b34801561015957600080fd5b50610162610317565b604080519115158252519081900360200190f35b34801561018257600080fd5b5061018e600435610325565b005b34801561019c57600080fd5b50610162610360565b3480156101b157600080fd5b5061013b61036f565b3480156101c657600080fd5b5061018e600435600160a060020a0360243516610375565b3480156101ea57600080fd5b5061018e600435600160a060020a0360243516610405565b34801561020e57600080fd5b5061013b61048b565b34801561022357600080fd5b5061018e6104a9565b34801561023857600080fd5b5061013b610523565b34801561024d57600080fd5b50610265600435600160a060020a0360243516610529565b60408051948552600160a060020a039093166020850152901515838301526060830152519081900360800190f35b34801561029f57600080fd5b5061016260043561056b565b3480156102b757600080fd5b50610162600435610598565b3480156102cf57600080fd5b506101626105ad565b3480156102e457600080fd5b506101626105b6565b3480156102f957600080fd5b5061013b6105c6565b600354600160a060020a031681565b60055481565b600054610100900460ff1681565b60058190556040805182815290517f7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b199181900360200190a150565b60005462010000900460ff1681565b60045481565b60005462010000900460ff1615156001148015610393575060025482145b80156103ac5750600354600160a060020a038281169116145b156103c0576103bb82826105cc565b610401565b60005462010000900460ff161580156103e15750600054610100900460ff16155b15610401576000805462ff000019166201000017905561040182826105cc565b5050565b6000546301000000900460ff1615156001148015610424575060025482145b801561043d5750600354600160a060020a038281169116145b1561044c576103bb82826108a9565b6000546301000000900460ff16158015610469575060005460ff16155b15610401576000805463ff0000001916630100000017905561040182826108a9565b6004546005546064908202046001019081106104a657506004545b90565b6000805463ffffffff1916815560028054825260076020908152604080842060038054600160a060020a0390811687529184528286208501959095559254935481519485529092169083015280517f25f36e43cc5e3db240c0fa2585314a51ed767376db5476ea24254b803a3276309281900390910190a1565b60025481565b60076020908152600092835260408084209091529082529020805460018201546002909201549091600160a060020a0381169160a060020a90910460ff169084565b60008181526006602052604081205460ff1615156001141561058f57506001610593565b5060005b919050565b60066020526000908152604090205460ff1681565b60005460ff1681565b6000546301000000900460ff1681565b60015481565b60008281526006602052604090205460ff16156105e557fe5b60005460ff16156105f257fe5b60005462010000900460ff16151561060657fe5b6000805461ff001916610100178155600154838252600760209081526040808420600160a060020a03861685529091529091206002015410156106bb576000828152600760209081526040808320600160a060020a03851680855292529091208381556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805490910190555b6001546000838152600760209081526040808320600160a060020a03861684529091529020600201541415610806576000828152600760209081526040808320600160a060020a0385168085528184528285208781556001808201805474ff00000000000000000000000000000000000000001973ffffffffffffffffffffffffffffffffffffffff1990911685171660a060020a179055865462ffff001916875588875260068652938620805460ff1916851790556004805490940190935584529091526002015561078c61048b565b60015560408051838152600160a060020a038316602082015260608183018190526006908201527f616464696e670000000000000000000000000000000000000000000000000000608082015290517f103ede1170d1e7714dcce0879255f6f760594d1d78ecf32669fa2790ba6899d99181900360a00190a15b600282905560038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051848152602081019290925260608282018190526008908301527f6164646974696f6e0000000000000000000000000000000000000000000000006080830152517f32a926be1e10de007860c671acf1ad7a1197bae083fa818ba5511288366f74319181900360a00190a15050565b600054610100900460ff16156108bb57fe5b6000546301000000900460ff1615156108d057fe5b6000828152600760209081526040808320600160a060020a0385168452909152902060019081015460a060020a900460ff1615151461090b57fe5b60008281526006602052604090205460ff16151560011461092857fe5b6000805460ff19166001908117825554838252600760209081526040808420600160a060020a0386168552909152909120600201541061096457fe5b6000828152600760209081526040808320600160a060020a03851680855292529091208381556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805482019081905590541415610ab0576000828152600660209081526040808320805460ff19169055825463ff0000ff1916835560078252808320600160a060020a038516845290915281206002015560048054600019019055610a3661048b565b60015560408051838152600160a060020a038316602082015260608183018190526009908201527f737573706574696f6e0000000000000000000000000000000000000000000000608082015290517f020064aba293970f316af376011a01158e4563606fd4d2c8402d59879a896ddd9181900360a00190a15b600282905560038054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051848152602081019290925260608282018190526009908301527f737573706574696f6e00000000000000000000000000000000000000000000006080830152517ff60f2e03adb385c167f2b27bb7b23b038ada3fd8274e0515ab8b69113af3639e9181900360a00190a150505600a165627a7a723058208af7f2ccbbd7e9f56bb6f3069905454af2e44f9b6a85d380f899ba1d77cc56a20029`

// DeployPermissions deploys a new Ethereum contract, binding an instance of Permissions to it.
func DeployPermissions(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Permissions, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermissionsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Permissions{PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract}}, nil
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
	Contract     *Permissions      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsCallerSession struct {
	Contract *PermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsTransactorSession struct {
	Contract     *PermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
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
	return &Permissions{PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract}}, nil
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
func (_Permissions *PermissionsCaller) LimitOfVote(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "LimitOfVote")
	return *ret0, err
}

// LimitOfVote is a free data retrieval call binding the contract method 0xff5c11f8.
//
// Solidity: function LimitOfVote() constant returns(uint256)
func (_Permissions *PermissionsSession) LimitOfVote() (*big.Int, error) {
	return _Permissions.Contract.LimitOfVote(&_Permissions.CallOpts)
}

// LimitOfVote is a free data retrieval call binding the contract method 0xff5c11f8.
//
// Solidity: function LimitOfVote() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) LimitOfVote() (*big.Int, error) {
	return _Permissions.Contract.LimitOfVote(&_Permissions.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsCaller) NodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "NodeCount")
	return *ret0, err
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsSession) NodeCount() (*big.Int, error) {
	return _Permissions.Contract.NodeCount(&_Permissions.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6a19f6ad.
//
// Solidity: function NodeCount() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) NodeCount() (*big.Int, error) {
	return _Permissions.Contract.NodeCount(&_Permissions.CallOpts)
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsCaller) Addingmutex(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "addingmutex")
	return *ret0, err
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsSession) Addingmutex() (bool, error) {
	return _Permissions.Contract.Addingmutex(&_Permissions.CallOpts)
}

// Addingmutex is a free data retrieval call binding the contract method 0x60c229f7.
//
// Solidity: function addingmutex() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Addingmutex() (bool, error) {
	return _Permissions.Contract.Addingmutex(&_Permissions.CallOpts)
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsCaller) CheckNode(opts *bind.CallOpts, _enode [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "checkNode", _enode)
	return *ret0, err
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsSession) CheckNode(_enode [32]byte) (bool, error) {
	return _Permissions.Contract.CheckNode(&_Permissions.CallOpts, _enode)
}

// CheckNode is a free data retrieval call binding the contract method 0xb50d6f75.
//
// Solidity: function checkNode(_enode bytes32) constant returns(bool)
func (_Permissions *PermissionsCallerSession) CheckNode(_enode [32]byte) (bool, error) {
	return _Permissions.Contract.CheckNode(&_Permissions.CallOpts, _enode)
}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() constant returns(limit uint256)
func (_Permissions *PermissionsCaller) Consensus(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "consensus")
	return *ret0, err
}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() constant returns(limit uint256)
func (_Permissions *PermissionsSession) Consensus() (*big.Int, error) {
	return _Permissions.Contract.Consensus(&_Permissions.CallOpts)
}

// Consensus is a free data retrieval call binding the contract method 0x8ef3f761.
//
// Solidity: function consensus() constant returns(limit uint256)
func (_Permissions *PermissionsCallerSession) Consensus() (*big.Int, error) {
	return _Permissions.Contract.Consensus(&_Permissions.CallOpts)
}

// Consensuspercentage is a free data retrieval call binding the contract method 0x1867f60c.
//
// Solidity: function consensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsCaller) Consensuspercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "consensuspercentage")
	return *ret0, err
}

// Consensuspercentage is a free data retrieval call binding the contract method 0x1867f60c.
//
// Solidity: function consensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsSession) Consensuspercentage() (*big.Int, error) {
	return _Permissions.Contract.Consensuspercentage(&_Permissions.CallOpts)
}

// Consensuspercentage is a free data retrieval call binding the contract method 0x1867f60c.
//
// Solidity: function consensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Consensuspercentage() (*big.Int, error) {
	return _Permissions.Contract.Consensuspercentage(&_Permissions.CallOpts)
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsCaller) Isadding(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "isadding")
	return *ret0, err
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsSession) Isadding() (bool, error) {
	return _Permissions.Contract.Isadding(&_Permissions.CallOpts)
}

// Isadding is a free data retrieval call binding the contract method 0x2da0bca8.
//
// Solidity: function isadding() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Isadding() (bool, error) {
	return _Permissions.Contract.Isadding(&_Permissions.CallOpts)
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsCaller) Issuspention(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "issuspention")
	return *ret0, err
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsSession) Issuspention() (bool, error) {
	return _Permissions.Contract.Issuspention(&_Permissions.CallOpts)
}

// Issuspention is a free data retrieval call binding the contract method 0xd6336a33.
//
// Solidity: function issuspention() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Issuspention() (bool, error) {
	return _Permissions.Contract.Issuspention(&_Permissions.CallOpts)
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsCaller) Nodeconformations(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "nodeconformations", arg0)
	return *ret0, err
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsSession) Nodeconformations(arg0 [32]byte) (bool, error) {
	return _Permissions.Contract.Nodeconformations(&_Permissions.CallOpts, arg0)
}

// Nodeconformations is a free data retrieval call binding the contract method 0xc86d87b9.
//
// Solidity: function nodeconformations( bytes32) constant returns(bool)
func (_Permissions *PermissionsCallerSession) Nodeconformations(arg0 [32]byte) (bool, error) {
	return _Permissions.Contract.Nodeconformations(&_Permissions.CallOpts, arg0)
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsCaller) Nodeinfo(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Enode     [32]byte
	Account   common.Address
	Flag      bool
	Votecount *big.Int
}, error) {
	ret := new(struct {
		Enode     [32]byte
		Account   common.Address
		Flag      bool
		Votecount *big.Int
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "nodeinfo", arg0, arg1)
	return *ret, err
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsSession) Nodeinfo(arg0 [32]byte, arg1 common.Address) (struct {
	Enode     [32]byte
	Account   common.Address
	Flag      bool
	Votecount *big.Int
}, error) {
	return _Permissions.Contract.Nodeinfo(&_Permissions.CallOpts, arg0, arg1)
}

// Nodeinfo is a free data retrieval call binding the contract method 0xb3b64ce2.
//
// Solidity: function nodeinfo( bytes32,  address) constant returns(enode bytes32, account address, flag bool, votecount uint256)
func (_Permissions *PermissionsCallerSession) Nodeinfo(arg0 [32]byte, arg1 common.Address) (struct {
	Enode     [32]byte
	Account   common.Address
	Flag      bool
	Votecount *big.Int
}, error) {
	return _Permissions.Contract.Nodeinfo(&_Permissions.CallOpts, arg0, arg1)
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsCaller) Previousaccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "previousaccount")
	return *ret0, err
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsSession) Previousaccount() (common.Address, error) {
	return _Permissions.Contract.Previousaccount(&_Permissions.CallOpts)
}

// Previousaccount is a free data retrieval call binding the contract method 0x118e6967.
//
// Solidity: function previousaccount() constant returns(address)
func (_Permissions *PermissionsCallerSession) Previousaccount() (common.Address, error) {
	return _Permissions.Contract.Previousaccount(&_Permissions.CallOpts)
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsCaller) Previousenode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "previousenode")
	return *ret0, err
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsSession) Previousenode() ([32]byte, error) {
	return _Permissions.Contract.Previousenode(&_Permissions.CallOpts)
}

// Previousenode is a free data retrieval call binding the contract method 0xb1b49f25.
//
// Solidity: function previousenode() constant returns(bytes32)
func (_Permissions *PermissionsCallerSession) Previousenode() ([32]byte, error) {
	return _Permissions.Contract.Previousenode(&_Permissions.CallOpts)
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsCaller) Suspentionmutex(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "suspentionmutex")
	return *ret0, err
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsSession) Suspentionmutex() (bool, error) {
	return _Permissions.Contract.Suspentionmutex(&_Permissions.CallOpts)
}

// Suspentionmutex is a free data retrieval call binding the contract method 0xeae897c7.
//
// Solidity: function suspentionmutex() constant returns(bool)
func (_Permissions *PermissionsCallerSession) Suspentionmutex() (bool, error) {
	return _Permissions.Contract.Suspentionmutex(&_Permissions.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactor) AddNode(opts *bind.TransactOpts, _enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addNode", _enode, _account)
}

// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsSession) AddNode(_enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _enode, _account)
}

// AddNode is a paid mutator transaction binding the contract method 0x76849376.
//
// Solidity: function addNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactorSession) AddNode(_enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _enode, _account)
}

// ResetProcess is a paid mutator transaction binding the contract method 0x9f9a19d3.
//
// Solidity: function resetProcess() returns()
func (_Permissions *PermissionsTransactor) ResetProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "resetProcess")
}

// ResetProcess is a paid mutator transaction binding the contract method 0x9f9a19d3.
//
// Solidity: function resetProcess() returns()
func (_Permissions *PermissionsSession) ResetProcess() (*types.Transaction, error) {
	return _Permissions.Contract.ResetProcess(&_Permissions.TransactOpts)
}

// ResetProcess is a paid mutator transaction binding the contract method 0x9f9a19d3.
//
// Solidity: function resetProcess() returns()
func (_Permissions *PermissionsTransactorSession) ResetProcess() (*types.Transaction, error) {
	return _Permissions.Contract.ResetProcess(&_Permissions.TransactOpts)
}

// SetConsensus is a paid mutator transaction binding the contract method 0x3549ef17.
//
// Solidity: function setConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsTransactor) SetConsensus(opts *bind.TransactOpts, _percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "setConsensus", _percentage)
}

// SetConsensus is a paid mutator transaction binding the contract method 0x3549ef17.
//
// Solidity: function setConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsSession) SetConsensus(_percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.SetConsensus(&_Permissions.TransactOpts, _percentage)
}

// SetConsensus is a paid mutator transaction binding the contract method 0x3549ef17.
//
// Solidity: function setConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsTransactorSession) SetConsensus(_percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.SetConsensus(&_Permissions.TransactOpts, _percentage)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactor) SuspendNode(opts *bind.TransactOpts, _enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "suspendNode", _enode, _account)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsSession) SuspendNode(_enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts, _enode, _account)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x8ec08b20.
//
// Solidity: function suspendNode(_enode bytes32, _account address) returns()
func (_Permissions *PermissionsTransactorSession) SuspendNode(_enode [32]byte, _account common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts, _enode, _account)
}

// PermissionsLogOfAddNodeIterator is returned from FilterLogOfAddNode and is used to iterate over the raw logs and unpacked data for LogOfAddNode events raised by the Permissions contract.
type PermissionsLogOfAddNodeIterator struct {
	Event *PermissionsLogOfAddNode // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddNode)
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
		it.Event = new(PermissionsLogOfAddNode)
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
func (it *PermissionsLogOfAddNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddNode represents a LogOfAddNode event raised by the Permissions contract.
type PermissionsLogOfAddNode struct {
	Enodeaddress   [32]byte
	Accountaddress common.Address
	Adding         string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddNode is a free log retrieval operation binding the contract event 0x32a926be1e10de007860c671acf1ad7a1197bae083fa818ba5511288366f7431.
//
// Solidity: e LogOfAddNode(enodeaddress bytes32, accountaddress address, adding string)
func (_Permissions *PermissionsFilterer) FilterLogOfAddNode(opts *bind.FilterOpts) (*PermissionsLogOfAddNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddNodeIterator{contract: _Permissions.contract, event: "LogOfAddNode", logs: logs, sub: sub}, nil
}

// WatchLogOfAddNode is a free log subscription operation binding the contract event 0x32a926be1e10de007860c671acf1ad7a1197bae083fa818ba5511288366f7431.
//
// Solidity: e LogOfAddNode(enodeaddress bytes32, accountaddress address, adding string)
func (_Permissions *PermissionsFilterer) WatchLogOfAddNode(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddNode) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddNode)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddNode", log); err != nil {
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

// PermissionsLogOfResetProcessIterator is returned from FilterLogOfResetProcess and is used to iterate over the raw logs and unpacked data for LogOfResetProcess events raised by the Permissions contract.
type PermissionsLogOfResetProcessIterator struct {
	Event *PermissionsLogOfResetProcess // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfResetProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfResetProcess)
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
		it.Event = new(PermissionsLogOfResetProcess)
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
func (it *PermissionsLogOfResetProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfResetProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfResetProcess represents a LogOfResetProcess event raised by the Permissions contract.
type PermissionsLogOfResetProcess struct {
	Previousenode   [32]byte
	Previousaccount common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogOfResetProcess is a free log retrieval operation binding the contract event 0x25f36e43cc5e3db240c0fa2585314a51ed767376db5476ea24254b803a327630.
//
// Solidity: e LogOfResetProcess(previousenode bytes32, previousaccount address)
func (_Permissions *PermissionsFilterer) FilterLogOfResetProcess(opts *bind.FilterOpts) (*PermissionsLogOfResetProcessIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfResetProcess")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfResetProcessIterator{contract: _Permissions.contract, event: "LogOfResetProcess", logs: logs, sub: sub}, nil
}

// WatchLogOfResetProcess is a free log subscription operation binding the contract event 0x25f36e43cc5e3db240c0fa2585314a51ed767376db5476ea24254b803a327630.
//
// Solidity: e LogOfResetProcess(previousenode bytes32, previousaccount address)
func (_Permissions *PermissionsFilterer) WatchLogOfResetProcess(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfResetProcess) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfResetProcess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfResetProcess)
				if err := _Permissions.contract.UnpackLog(event, "LogOfResetProcess", log); err != nil {
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

// PermissionsLogOfSetConsensusIterator is returned from FilterLogOfSetConsensus and is used to iterate over the raw logs and unpacked data for LogOfSetConsensus events raised by the Permissions contract.
type PermissionsLogOfSetConsensusIterator struct {
	Event *PermissionsLogOfSetConsensus // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSetConsensusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSetConsensus)
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
		it.Event = new(PermissionsLogOfSetConsensus)
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
func (it *PermissionsLogOfSetConsensusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSetConsensusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSetConsensus represents a LogOfSetConsensus event raised by the Permissions contract.
type PermissionsLogOfSetConsensus struct {
	Consensuslimit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSetConsensus is a free log retrieval operation binding the contract event 0x7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b19.
//
// Solidity: e LogOfSetConsensus(consensuslimit uint256)
func (_Permissions *PermissionsFilterer) FilterLogOfSetConsensus(opts *bind.FilterOpts) (*PermissionsLogOfSetConsensusIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSetConsensus")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSetConsensusIterator{contract: _Permissions.contract, event: "LogOfSetConsensus", logs: logs, sub: sub}, nil
}

// WatchLogOfSetConsensus is a free log subscription operation binding the contract event 0x7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b19.
//
// Solidity: e LogOfSetConsensus(consensuslimit uint256)
func (_Permissions *PermissionsFilterer) WatchLogOfSetConsensus(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSetConsensus) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSetConsensus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSetConsensus)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSetConsensus", log); err != nil {
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

// PermissionsLogOfSuspentionNodeIterator is returned from FilterLogOfSuspentionNode and is used to iterate over the raw logs and unpacked data for LogOfSuspentionNode events raised by the Permissions contract.
type PermissionsLogOfSuspentionNodeIterator struct {
	Event *PermissionsLogOfSuspentionNode // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSuspentionNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSuspentionNode)
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
		it.Event = new(PermissionsLogOfSuspentionNode)
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
func (it *PermissionsLogOfSuspentionNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSuspentionNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSuspentionNode represents a LogOfSuspentionNode event raised by the Permissions contract.
type PermissionsLogOfSuspentionNode struct {
	Enodeaddress   [32]byte
	Accountaddress common.Address
	Suspend        string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionNode is a free log retrieval operation binding the contract event 0xf60f2e03adb385c167f2b27bb7b23b038ada3fd8274e0515ab8b69113af3639e.
//
// Solidity: e LogOfSuspentionNode(enodeaddress bytes32, accountaddress address, suspend string)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionNode(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionNodeIterator{contract: _Permissions.contract, event: "LogOfSuspentionNode", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionNode is a free log subscription operation binding the contract event 0xf60f2e03adb385c167f2b27bb7b23b038ada3fd8274e0515ab8b69113af3639e.
//
// Solidity: e LogOfSuspentionNode(enodeaddress bytes32, accountaddress address, suspend string)
func (_Permissions *PermissionsFilterer) WatchLogOfSuspentionNode(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSuspentionNode) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSuspentionNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSuspentionNode)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSuspentionNode", log); err != nil {
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

// PermissionsLogOfaddingConsensusMeetIterator is returned from FilterLogOfaddingConsensusMeet and is used to iterate over the raw logs and unpacked data for LogOfaddingConsensusMeet events raised by the Permissions contract.
type PermissionsLogOfaddingConsensusMeetIterator struct {
	Event *PermissionsLogOfaddingConsensusMeet // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfaddingConsensusMeetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfaddingConsensusMeet)
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
		it.Event = new(PermissionsLogOfaddingConsensusMeet)
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
func (it *PermissionsLogOfaddingConsensusMeetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfaddingConsensusMeetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfaddingConsensusMeet represents a LogOfaddingConsensusMeet event raised by the Permissions contract.
type PermissionsLogOfaddingConsensusMeet struct {
	Enode   [32]byte
	Account common.Address
	Flag    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogOfaddingConsensusMeet is a free log retrieval operation binding the contract event 0x103ede1170d1e7714dcce0879255f6f760594d1d78ecf32669fa2790ba6899d9.
//
// Solidity: e LogOfaddingConsensusMeet(enode bytes32, account address, flag string)
func (_Permissions *PermissionsFilterer) FilterLogOfaddingConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfaddingConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfaddingConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfaddingConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfaddingConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfaddingConsensusMeet is a free log subscription operation binding the contract event 0x103ede1170d1e7714dcce0879255f6f760594d1d78ecf32669fa2790ba6899d9.
//
// Solidity: e LogOfaddingConsensusMeet(enode bytes32, account address, flag string)
func (_Permissions *PermissionsFilterer) WatchLogOfaddingConsensusMeet(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfaddingConsensusMeet) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfaddingConsensusMeet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfaddingConsensusMeet)
				if err := _Permissions.contract.UnpackLog(event, "LogOfaddingConsensusMeet", log); err != nil {
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

// PermissionsLogOfsuspentionConsensusMeetIterator is returned from FilterLogOfsuspentionConsensusMeet and is used to iterate over the raw logs and unpacked data for LogOfsuspentionConsensusMeet events raised by the Permissions contract.
type PermissionsLogOfsuspentionConsensusMeetIterator struct {
	Event *PermissionsLogOfsuspentionConsensusMeet // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfsuspentionConsensusMeetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfsuspentionConsensusMeet)
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
		it.Event = new(PermissionsLogOfsuspentionConsensusMeet)
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
func (it *PermissionsLogOfsuspentionConsensusMeetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfsuspentionConsensusMeetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfsuspentionConsensusMeet represents a LogOfsuspentionConsensusMeet event raised by the Permissions contract.
type PermissionsLogOfsuspentionConsensusMeet struct {
	Enode   [32]byte
	Account common.Address
	Flag    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogOfsuspentionConsensusMeet is a free log retrieval operation binding the contract event 0x020064aba293970f316af376011a01158e4563606fd4d2c8402d59879a896ddd.
//
// Solidity: e LogOfsuspentionConsensusMeet(enode bytes32, account address, flag string)
func (_Permissions *PermissionsFilterer) FilterLogOfsuspentionConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfsuspentionConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfsuspentionConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfsuspentionConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfsuspentionConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfsuspentionConsensusMeet is a free log subscription operation binding the contract event 0x020064aba293970f316af376011a01158e4563606fd4d2c8402d59879a896ddd.
//
// Solidity: e LogOfsuspentionConsensusMeet(enode bytes32, account address, flag string)
func (_Permissions *PermissionsFilterer) WatchLogOfsuspentionConsensusMeet(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfsuspentionConsensusMeet) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfsuspentionConsensusMeet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfsuspentionConsensusMeet)
				if err := _Permissions.contract.UnpackLog(event, "LogOfsuspentionConsensusMeet", log); err != nil {
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
