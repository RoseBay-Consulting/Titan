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
const PermissionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"previousaccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensuspercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isadding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"setConsensus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addingmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addingVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeinfo\",\"outputs\":[{\"name\":\"enode\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\"},{\"name\":\"votecount\",\"type\":\"uint256\"},{\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfNegVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetProcess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousenode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"checkNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeconformations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuspention\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspentionmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"adding\",\"type\":\"string\"}],\"name\":\"LogOfAddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"suspend\",\"type\":\"string\"}],\"name\":\"LogOfSuspentionNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consensuslimit\",\"type\":\"uint256\"}],\"name\":\"LogOfSetConsensus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousaccount\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousenode\",\"type\":\"bytes32\"}],\"name\":\"LogOfResetProcess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"LogOfaddingConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"}],\"name\":\"LogOfsuspentionConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"vote\",\"type\":\"bool\"}],\"name\":\"LogOfAddingVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"flag\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"vote\",\"type\":\"bool\"}],\"name\":\"LogOfSuspentionVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"name\":\"LogOfVoteReject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isadding\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"issuspention\",\"type\":\"bool\"}],\"name\":\"LogOfVoteRejectConsensusMeet\",\"type\":\"event\"}]"

// PermissionsBin is the compiled bytecode used for deploying new contracts.
const PermissionsBin = `0x608060405234801561001057600080fd5b5060006001819055600281905560055561101e8061002f6000396000f3006080604052600436106101115763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663118e696781146101165780631867f60c146101475780632da0bca81461016e5780633549ef17146101975780635b5a743b146101b157806360c229f7146101d857806362f90d3d146101ed5780636a19f6ad1461021457806374a166fb1461022957806375bccb391461025057806397b1a65d146102775780639a7bf5bd146102ce5780639f9a19d3146102e3578063b1b49f25146102f8578063b50d6f751461030d578063c86d87b914610325578063c9663fac1461033d578063d6336a3314610364578063eae897c714610379578063ff5c11f81461038e575b600080fd5b34801561012257600080fd5b5061012b6103a3565b60408051600160a060020a039092168252519081900360200190f35b34801561015357600080fd5b5061015c6103b2565b60408051918252519081900360200190f35b34801561017a57600080fd5b506101836103b8565b604080519115158252519081900360200190f35b3480156101a357600080fd5b506101af6004356103c6565b005b3480156101bd57600080fd5b506101af600435600160a060020a036024351660443561041e565b3480156101e457600080fd5b506101836104b1565b3480156101f957600080fd5b506101af600435600160a060020a03602435166044356104c0565b34801561022057600080fd5b5061015c610548565b34801561023557600080fd5b506101af600435600160a060020a036024351660443561054e565b34801561025c57600080fd5b506101af600435600160a060020a0360243516604435610667565b34801561028357600080fd5b5061029b600160a060020a036004351660243561077e565b60408051958652600160a060020a0390941660208601529115158484015260608401526080830152519081900360a00190f35b3480156102da57600080fd5b5061015c6107c5565b3480156102ef57600080fd5b506101af6107cb565b34801561030457600080fd5b5061015c610861565b34801561031957600080fd5b50610183600435610867565b34801561033157600080fd5b50610183600435610894565b34801561034957600080fd5b506101af600435600160a060020a03602435166044356108a9565b34801561037057600080fd5b50610183610a13565b34801561038557600080fd5b50610183610a1c565b34801561039a57600080fd5b5061015c610a2c565b600454600160a060020a031681565b60065481565b600054610100900460ff1681565b600081101580156103d8575060648111155b15156103e357600080fd5b60068190556040805182815290517f7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b199181900360200190a150565b60005462010000900460ff161515600114801561043c575060035481145b80156104555750600454600160a060020a038381169116145b1561046a57610465838383610a32565b6104ac565b60005462010000900460ff1615801561048b5750600054610100900460ff16155b156104ac576000805462ff00001916620100001790556104ac838383610a32565b505050565b60005462010000900460ff1681565b6000546301000000900460ff16151560011480156104df575060035481145b80156104f85750600454600160a060020a038381169116145b1561050857610465838383610d1a565b6000546301000000900460ff16158015610525575060005460ff16155b156104ac576000805463ff000000191663010000001790556104ac838383610d1a565b60055481565b6000546301000000900460ff161515600114801561056d575060035481145b80156105865750600454600160a060020a038381169116145b1561059b57610596838383610d1a565b6105de565b6000546301000000900460ff16151560011480156105ba575060035481145b80156105d35750600454600160a060020a038381169116145b15156105de57600080fd5b60408051848152600160a060020a03841660208201528082018390526001608082015260a060608201819052600a908201527f73757370656e74696f6e0000000000000000000000000000000000000000000060c082015290517fefddc40bbe0322d1ce33bf7ee3b1633be2a6f6548f15be800ff445d8fb692e289181900360e00190a1505050565b60005462010000900460ff1615156001148015610685575060035481145b801561069e5750600454600160a060020a038381169116145b156106b3576106ae838383610a32565b6106f5565b60005462010000900460ff16151560011480156106d1575060035481145b80156106ea5750600454600160a060020a038381169116145b15156106f557600080fd5b60408051848152600160a060020a03841660208201528082018390526001608082015260a0606082018190526006908201527f616464696e67000000000000000000000000000000000000000000000000000060c082015290517f9ad3ed8a72d9c28e71dfb48742bbfeae3336247ee0919e062de2af58b49217449181900360e00190a1505050565b600860209081526000928352604080842090915290825290208054600182015460028301546003909301549192600160a060020a0382169260a060020a90920460ff169185565b60025481565b6000805463ffffffff1916815560048054600160a060020a0390811683526008602081815260408086206003805488529083528187206002018790558554851687529282528086208354875282528086208301959095559254905484519190921681529182015281517f7350a7ef2a3d35fb20258804f114f645300830cdbb8ee760db828576c6fb81e9929181900390910190a1565b60035481565b60008181526007602052604081205460ff1615156001141561088b5750600161088f565b5060005b919050565b60076020526000908152604090205460ff1681565b6000546301000000900460ff161515600114806108d3575060005462010000900460ff1615156001145b80156108e0575060035481145b80156108f95750600454600160a060020a038381169116145b156104ac57600160a060020a03821660008181526008602090815260408083208584528252918290206003018054600101908190558251878152918201939093528082018490526060810192909252517f7cb4bca1eaf530a5cfb00996c12f042fc783becb68d363236c238c081d0139029181900360800190a161097b610fb7565b600160a060020a0383166000908152600860209081526040808320858452909152902060030154106104ac5760005460408051858152600160a060020a038516602082015280820184905260ff610100840481161515606083015290921615156080830152517f2df2bdee48a50cd1f9880e4e9b7da78d3e20086869d1b784721308ab708154449181900360a00190a16104ac6107cb565b60005460ff1681565b6000546301000000900460ff1681565b60015481565b60008181526007602052604090205460ff1615610a4b57fe5b60005460ff1615610a5857fe5b60005462010000900460ff161515610a6c57fe5b6000805461ff001916610100178155600154600160a060020a0384168252600860209081526040808420858552909152909120600201541015610b2157600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805490910190555b600154600160a060020a03831660009081526008602090815260408083208584529091529020600201541415610c6f57600160a060020a038216600081815260086020908152604080832085845282528083208581556001908101805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690961774ff000000000000000000000000000000000000000019169590951790945560079091529020805460ff19169091179055610bd86107cb565b600580546001019055610be9610fd3565b600155610bf4610fb7565b60025560408051600160a060020a03841681526020810183905260608183018190526006908201527f616464696e670000000000000000000000000000000000000000000000000000608082015290517f7486860e51bf4bc8ff65d6a2a93daaa5fafe07f449fec1fee7dd6f4c9a5a9c7d9181900360a00190a15b600381905560048054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff1990911681179091556040805185815260208101929092528181018390526080606083018190526008908301527f6164646974696f6e00000000000000000000000000000000000000000000000060a0830152517f36647f3687846b9a82b6b2a1364f364419cab9eec541c42dc16fa9a7dfacd4b89181900360c00190a1505050565b600054610100900460ff1615610d2c57fe5b6000546301000000900460ff161515610d4157fe5b600160a060020a0382166000908152600860209081526040808320848452909152902060019081015460a060020a900460ff16151514610d7d57fe5b60008181526007602052604090205460ff161515600114610d9a57fe5b6000805460ff19166001908117825554600160a060020a03841682526008602090815260408084208585529091529091206002015410610dd657fe5b600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a73ffffffffffffffffffffffffffffffffffffffff1990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805482019081905590541415610f0c576000818152600760205260409020805460ff19169055610e746107cb565b60058054600019019055610e86610fd3565b600155610e91610fb7565b60025560408051600160a060020a0384168152602081018390526060818301819052600a908201527f73757370656e73696f6e00000000000000000000000000000000000000000000608082015290517f1144ea0430fa79179c20ea22367df2c383b2be2c02d580fa417e764d19083d0d9181900360a00190a15b600381905560048054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff199091168117909155604080518581526020810192909252818101839052608060608301819052600a908301527f73757370656e74696f6e0000000000000000000000000000000000000000000060a0830152517fc6a6d7a20a1eed76135dfd7ae812dc6a4a6fba50230c0263e4b4de2182a4f9709181900360c00190a1505050565b60018054600554908103909101908110610fd057506005545b90565b600554600654606490820204600101908110610fd05750600554610fd05600a165627a7a723058201a6bdfb8df8da0b3b41333edc7ca9a55d6969043a214a0b67d9718832c25b60c0029`

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

// LimitOfNegVote is a free data retrieval call binding the contract method 0x9a7bf5bd.
//
// Solidity: function LimitOfNegVote() constant returns(uint256)
func (_Permissions *PermissionsCaller) LimitOfNegVote(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "LimitOfNegVote")
	return *ret0, err
}

// LimitOfNegVote is a free data retrieval call binding the contract method 0x9a7bf5bd.
//
// Solidity: function LimitOfNegVote() constant returns(uint256)
func (_Permissions *PermissionsSession) LimitOfNegVote() (*big.Int, error) {
	return _Permissions.Contract.LimitOfNegVote(&_Permissions.CallOpts)
}

// LimitOfNegVote is a free data retrieval call binding the contract method 0x9a7bf5bd.
//
// Solidity: function LimitOfNegVote() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) LimitOfNegVote() (*big.Int, error) {
	return _Permissions.Contract.LimitOfNegVote(&_Permissions.CallOpts)
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

// Nodeinfo is a free data retrieval call binding the contract method 0x97b1a65d.
//
// Solidity: function nodeinfo( address,  bytes32) constant returns(enode bytes32, account address, flag bool, votecount uint256, rejectcount uint256)
func (_Permissions *PermissionsCaller) Nodeinfo(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Enode       [32]byte
	Account     common.Address
	Flag        bool
	Votecount   *big.Int
	Rejectcount *big.Int
}, error) {
	ret := new(struct {
		Enode       [32]byte
		Account     common.Address
		Flag        bool
		Votecount   *big.Int
		Rejectcount *big.Int
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "nodeinfo", arg0, arg1)
	return *ret, err
}

// Nodeinfo is a free data retrieval call binding the contract method 0x97b1a65d.
//
// Solidity: function nodeinfo( address,  bytes32) constant returns(enode bytes32, account address, flag bool, votecount uint256, rejectcount uint256)
func (_Permissions *PermissionsSession) Nodeinfo(arg0 common.Address, arg1 [32]byte) (struct {
	Enode       [32]byte
	Account     common.Address
	Flag        bool
	Votecount   *big.Int
	Rejectcount *big.Int
}, error) {
	return _Permissions.Contract.Nodeinfo(&_Permissions.CallOpts, arg0, arg1)
}

// Nodeinfo is a free data retrieval call binding the contract method 0x97b1a65d.
//
// Solidity: function nodeinfo( address,  bytes32) constant returns(enode bytes32, account address, flag bool, votecount uint256, rejectcount uint256)
func (_Permissions *PermissionsCallerSession) Nodeinfo(arg0 common.Address, arg1 [32]byte) (struct {
	Enode       [32]byte
	Account     common.Address
	Flag        bool
	Votecount   *big.Int
	Rejectcount *big.Int
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

// AddNode is a paid mutator transaction binding the contract method 0x5b5a743b.
//
// Solidity: function addNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) AddNode(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addNode", _id, _account, _enode)
}

// AddNode is a paid mutator transaction binding the contract method 0x5b5a743b.
//
// Solidity: function addNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) AddNode(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _id, _account, _enode)
}

// AddNode is a paid mutator transaction binding the contract method 0x5b5a743b.
//
// Solidity: function addNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) AddNode(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddNode(&_Permissions.TransactOpts, _id, _account, _enode)
}

// AddingVote is a paid mutator transaction binding the contract method 0x75bccb39.
//
// Solidity: function addingVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) AddingVote(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addingVote", _id, _account, _enode)
}

// AddingVote is a paid mutator transaction binding the contract method 0x75bccb39.
//
// Solidity: function addingVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) AddingVote(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddingVote(&_Permissions.TransactOpts, _id, _account, _enode)
}

// AddingVote is a paid mutator transaction binding the contract method 0x75bccb39.
//
// Solidity: function addingVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) AddingVote(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddingVote(&_Permissions.TransactOpts, _id, _account, _enode)
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

// SuspendNode is a paid mutator transaction binding the contract method 0x62f90d3d.
//
// Solidity: function suspendNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) SuspendNode(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "suspendNode", _id, _account, _enode)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x62f90d3d.
//
// Solidity: function suspendNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) SuspendNode(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts, _id, _account, _enode)
}

// SuspendNode is a paid mutator transaction binding the contract method 0x62f90d3d.
//
// Solidity: function suspendNode(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) SuspendNode(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNode(&_Permissions.TransactOpts, _id, _account, _enode)
}

// SuspendVote is a paid mutator transaction binding the contract method 0x74a166fb.
//
// Solidity: function suspendVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) SuspendVote(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "suspendVote", _id, _account, _enode)
}

// SuspendVote is a paid mutator transaction binding the contract method 0x74a166fb.
//
// Solidity: function suspendVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) SuspendVote(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendVote(&_Permissions.TransactOpts, _id, _account, _enode)
}

// SuspendVote is a paid mutator transaction binding the contract method 0x74a166fb.
//
// Solidity: function suspendVote(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) SuspendVote(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendVote(&_Permissions.TransactOpts, _id, _account, _enode)
}

// VoteReject is a paid mutator transaction binding the contract method 0xc9663fac.
//
// Solidity: function voteReject(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) VoteReject(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "voteReject", _id, _account, _enode)
}

// VoteReject is a paid mutator transaction binding the contract method 0xc9663fac.
//
// Solidity: function voteReject(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) VoteReject(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.VoteReject(&_Permissions.TransactOpts, _id, _account, _enode)
}

// VoteReject is a paid mutator transaction binding the contract method 0xc9663fac.
//
// Solidity: function voteReject(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) VoteReject(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.VoteReject(&_Permissions.TransactOpts, _id, _account, _enode)
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
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Adding         string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddNode is a free log retrieval operation binding the contract event 0x36647f3687846b9a82b6b2a1364f364419cab9eec541c42dc16fa9a7dfacd4b8.
//
// Solidity: e LogOfAddNode(id uint256, accountaddress address, enodeaddress bytes32, adding string)
func (_Permissions *PermissionsFilterer) FilterLogOfAddNode(opts *bind.FilterOpts) (*PermissionsLogOfAddNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddNodeIterator{contract: _Permissions.contract, event: "LogOfAddNode", logs: logs, sub: sub}, nil
}

// WatchLogOfAddNode is a free log subscription operation binding the contract event 0x36647f3687846b9a82b6b2a1364f364419cab9eec541c42dc16fa9a7dfacd4b8.
//
// Solidity: e LogOfAddNode(id uint256, accountaddress address, enodeaddress bytes32, adding string)
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

// PermissionsLogOfAddingVoteIterator is returned from FilterLogOfAddingVote and is used to iterate over the raw logs and unpacked data for LogOfAddingVote events raised by the Permissions contract.
type PermissionsLogOfAddingVoteIterator struct {
	Event *PermissionsLogOfAddingVote // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddingVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddingVote)
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
		it.Event = new(PermissionsLogOfAddingVote)
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
func (it *PermissionsLogOfAddingVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddingVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddingVote represents a LogOfAddingVote event raised by the Permissions contract.
type PermissionsLogOfAddingVote struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Flag           string
	Vote           bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddingVote is a free log retrieval operation binding the contract event 0x9ad3ed8a72d9c28e71dfb48742bbfeae3336247ee0919e062de2af58b4921744.
//
// Solidity: e LogOfAddingVote(id uint256, accountaddress address, enodeaddress bytes32, flag string, vote bool)
func (_Permissions *PermissionsFilterer) FilterLogOfAddingVote(opts *bind.FilterOpts) (*PermissionsLogOfAddingVoteIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddingVote")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddingVoteIterator{contract: _Permissions.contract, event: "LogOfAddingVote", logs: logs, sub: sub}, nil
}

// WatchLogOfAddingVote is a free log subscription operation binding the contract event 0x9ad3ed8a72d9c28e71dfb48742bbfeae3336247ee0919e062de2af58b4921744.
//
// Solidity: e LogOfAddingVote(id uint256, accountaddress address, enodeaddress bytes32, flag string, vote bool)
func (_Permissions *PermissionsFilterer) WatchLogOfAddingVote(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddingVote) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddingVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddingVote)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddingVote", log); err != nil {
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
	Previousaccount common.Address
	Previousenode   [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogOfResetProcess is a free log retrieval operation binding the contract event 0x7350a7ef2a3d35fb20258804f114f645300830cdbb8ee760db828576c6fb81e9.
//
// Solidity: e LogOfResetProcess(previousaccount address, previousenode bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfResetProcess(opts *bind.FilterOpts) (*PermissionsLogOfResetProcessIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfResetProcess")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfResetProcessIterator{contract: _Permissions.contract, event: "LogOfResetProcess", logs: logs, sub: sub}, nil
}

// WatchLogOfResetProcess is a free log subscription operation binding the contract event 0x7350a7ef2a3d35fb20258804f114f645300830cdbb8ee760db828576c6fb81e9.
//
// Solidity: e LogOfResetProcess(previousaccount address, previousenode bytes32)
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
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Suspend        string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionNode is a free log retrieval operation binding the contract event 0xc6a6d7a20a1eed76135dfd7ae812dc6a4a6fba50230c0263e4b4de2182a4f970.
//
// Solidity: e LogOfSuspentionNode(id uint256, accountaddress address, enodeaddress bytes32, suspend string)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionNode(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionNodeIterator{contract: _Permissions.contract, event: "LogOfSuspentionNode", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionNode is a free log subscription operation binding the contract event 0xc6a6d7a20a1eed76135dfd7ae812dc6a4a6fba50230c0263e4b4de2182a4f970.
//
// Solidity: e LogOfSuspentionNode(id uint256, accountaddress address, enodeaddress bytes32, suspend string)
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

// PermissionsLogOfSuspentionVoteIterator is returned from FilterLogOfSuspentionVote and is used to iterate over the raw logs and unpacked data for LogOfSuspentionVote events raised by the Permissions contract.
type PermissionsLogOfSuspentionVoteIterator struct {
	Event *PermissionsLogOfSuspentionVote // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSuspentionVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSuspentionVote)
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
		it.Event = new(PermissionsLogOfSuspentionVote)
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
func (it *PermissionsLogOfSuspentionVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSuspentionVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSuspentionVote represents a LogOfSuspentionVote event raised by the Permissions contract.
type PermissionsLogOfSuspentionVote struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Flag           string
	Vote           bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionVote is a free log retrieval operation binding the contract event 0xefddc40bbe0322d1ce33bf7ee3b1633be2a6f6548f15be800ff445d8fb692e28.
//
// Solidity: e LogOfSuspentionVote(id uint256, accountaddress address, enodeaddress bytes32, flag string, vote bool)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionVote(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionVoteIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionVote")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionVoteIterator{contract: _Permissions.contract, event: "LogOfSuspentionVote", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionVote is a free log subscription operation binding the contract event 0xefddc40bbe0322d1ce33bf7ee3b1633be2a6f6548f15be800ff445d8fb692e28.
//
// Solidity: e LogOfSuspentionVote(id uint256, accountaddress address, enodeaddress bytes32, flag string, vote bool)
func (_Permissions *PermissionsFilterer) WatchLogOfSuspentionVote(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSuspentionVote) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSuspentionVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSuspentionVote)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSuspentionVote", log); err != nil {
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

// PermissionsLogOfVoteRejectIterator is returned from FilterLogOfVoteReject and is used to iterate over the raw logs and unpacked data for LogOfVoteReject events raised by the Permissions contract.
type PermissionsLogOfVoteRejectIterator struct {
	Event *PermissionsLogOfVoteReject // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfVoteRejectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfVoteReject)
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
		it.Event = new(PermissionsLogOfVoteReject)
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
func (it *PermissionsLogOfVoteRejectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfVoteRejectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfVoteReject represents a LogOfVoteReject event raised by the Permissions contract.
type PermissionsLogOfVoteReject struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Rejectcount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfVoteReject is a free log retrieval operation binding the contract event 0x7cb4bca1eaf530a5cfb00996c12f042fc783becb68d363236c238c081d013902.
//
// Solidity: e LogOfVoteReject(id uint256, accountaddress address, enodeaddress bytes32, rejectcount uint256)
func (_Permissions *PermissionsFilterer) FilterLogOfVoteReject(opts *bind.FilterOpts) (*PermissionsLogOfVoteRejectIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfVoteReject")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfVoteRejectIterator{contract: _Permissions.contract, event: "LogOfVoteReject", logs: logs, sub: sub}, nil
}

// WatchLogOfVoteReject is a free log subscription operation binding the contract event 0x7cb4bca1eaf530a5cfb00996c12f042fc783becb68d363236c238c081d013902.
//
// Solidity: e LogOfVoteReject(id uint256, accountaddress address, enodeaddress bytes32, rejectcount uint256)
func (_Permissions *PermissionsFilterer) WatchLogOfVoteReject(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfVoteReject) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfVoteReject")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfVoteReject)
				if err := _Permissions.contract.UnpackLog(event, "LogOfVoteReject", log); err != nil {
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

// PermissionsLogOfVoteRejectConsensusMeetIterator is returned from FilterLogOfVoteRejectConsensusMeet and is used to iterate over the raw logs and unpacked data for LogOfVoteRejectConsensusMeet events raised by the Permissions contract.
type PermissionsLogOfVoteRejectConsensusMeetIterator struct {
	Event *PermissionsLogOfVoteRejectConsensusMeet // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfVoteRejectConsensusMeetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfVoteRejectConsensusMeet)
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
		it.Event = new(PermissionsLogOfVoteRejectConsensusMeet)
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
func (it *PermissionsLogOfVoteRejectConsensusMeetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfVoteRejectConsensusMeetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfVoteRejectConsensusMeet represents a LogOfVoteRejectConsensusMeet event raised by the Permissions contract.
type PermissionsLogOfVoteRejectConsensusMeet struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Isadding       bool
	Issuspention   bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfVoteRejectConsensusMeet is a free log retrieval operation binding the contract event 0x2df2bdee48a50cd1f9880e4e9b7da78d3e20086869d1b784721308ab70815444.
//
// Solidity: e LogOfVoteRejectConsensusMeet(id uint256, accountaddress address, enodeaddress bytes32, isadding bool, issuspention bool)
func (_Permissions *PermissionsFilterer) FilterLogOfVoteRejectConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfVoteRejectConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfVoteRejectConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfVoteRejectConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfVoteRejectConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfVoteRejectConsensusMeet is a free log subscription operation binding the contract event 0x2df2bdee48a50cd1f9880e4e9b7da78d3e20086869d1b784721308ab70815444.
//
// Solidity: e LogOfVoteRejectConsensusMeet(id uint256, accountaddress address, enodeaddress bytes32, isadding bool, issuspention bool)
func (_Permissions *PermissionsFilterer) WatchLogOfVoteRejectConsensusMeet(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfVoteRejectConsensusMeet) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfVoteRejectConsensusMeet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfVoteRejectConsensusMeet)
				if err := _Permissions.contract.UnpackLog(event, "LogOfVoteRejectConsensusMeet", log); err != nil {
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
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Flag           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfaddingConsensusMeet is a free log retrieval operation binding the contract event 0x7486860e51bf4bc8ff65d6a2a93daaa5fafe07f449fec1fee7dd6f4c9a5a9c7d.
//
// Solidity: e LogOfaddingConsensusMeet(accountaddress address, enodeaddress bytes32, flag string)
func (_Permissions *PermissionsFilterer) FilterLogOfaddingConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfaddingConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfaddingConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfaddingConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfaddingConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfaddingConsensusMeet is a free log subscription operation binding the contract event 0x7486860e51bf4bc8ff65d6a2a93daaa5fafe07f449fec1fee7dd6f4c9a5a9c7d.
//
// Solidity: e LogOfaddingConsensusMeet(accountaddress address, enodeaddress bytes32, flag string)
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
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Flag           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfsuspentionConsensusMeet is a free log retrieval operation binding the contract event 0x1144ea0430fa79179c20ea22367df2c383b2be2c02d580fa417e764d19083d0d.
//
// Solidity: e LogOfsuspentionConsensusMeet(accountaddress address, enodeaddress bytes32, flag string)
func (_Permissions *PermissionsFilterer) FilterLogOfsuspentionConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfsuspentionConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfsuspentionConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfsuspentionConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfsuspentionConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfsuspentionConsensusMeet is a free log subscription operation binding the contract event 0x1144ea0430fa79179c20ea22367df2c383b2be2c02d580fa417e764d19083d0d.
//
// Solidity: e LogOfsuspentionConsensusMeet(accountaddress address, enodeaddress bytes32, flag string)
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
