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
const PermissionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"previousaccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensuspercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addNodeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isadding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"setConsensus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addingmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"acceptProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addingVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeinfo\",\"outputs\":[{\"name\":\"enode\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\"},{\"name\":\"votecount\",\"type\":\"uint256\"},{\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfNegVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetProcess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendNodeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousenode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"checkNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeconformations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuspention\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspentionmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consensuslimit\",\"type\":\"uint256\"}],\"name\":\"LogOfSetConsensus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousaccount\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousenode\",\"type\":\"bytes32\"}],\"name\":\"LogOfResetProcess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfsuspentionConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddNodeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspandNodeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingAcceptProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionAcceptProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"name\":\"LogOfVoteReject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isadding\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"issuspention\",\"type\":\"bool\"}],\"name\":\"LogOfVoteRejectConsensusMeet\",\"type\":\"event\"}]"

// PermissionsBin is the compiled bytecode used for deploying new contracts.
const PermissionsBin = `0x608060405234801561001057600080fd5b506000600181905560028190556005556112248061002f6000396000f3006080604052600436106101325763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663118e696781146101375780631867f60c146101685780631d3a12d71461018f5780632da0bca8146101b85780633549ef17146101e15780635b5a743b146101f957806360c229f71461022057806360c5cc3a1461023557806362f90d3d1461024d5780636a19f6ad1461027457806374a166fb1461028957806375bccb39146102b057806397b1a65d146102d75780639a7bf5bd1461032e5780639f9a19d314610343578063a29cd37214610358578063b1b49f251461037f578063b50d6f7514610394578063c86d87b9146103ac578063c9663fac146103c4578063d6336a33146103eb578063eae897c714610400578063ff5c11f814610415575b600080fd5b34801561014357600080fd5b5061014c61042a565b60408051600160a060020a039092168252519081900360200190f35b34801561017457600080fd5b5061017d610439565b60408051918252519081900360200190f35b34801561019b57600080fd5b506101b6600435600160a060020a036024351660443561043f565b005b3480156101c457600080fd5b506101cd610559565b604080519115158252519081900360200190f35b3480156101ed57600080fd5b506101b6600435610567565b34801561020557600080fd5b506101b6600435600160a060020a03602435166044356105bf565b34801561022c57600080fd5b506101cd61064f565b34801561024157600080fd5b506101b660043561065e565b34801561025957600080fd5b506101b6600435600160a060020a03602435166044356107bf565b34801561028057600080fd5b5061017d61084c565b34801561029557600080fd5b506101b6600435600160a060020a0360243516604435610852565b3480156102bc57600080fd5b506101b6600435600160a060020a03602435166044356108e7565b3480156102e357600080fd5b506102fb600160a060020a036004351660243561097b565b60408051958652600160a060020a0390941660208601529115158484015260608401526080830152519081900360a00190f35b34801561033a57600080fd5b5061017d6109c2565b34801561034f57600080fd5b506101b66109c8565b34801561036457600080fd5b506101b6600435600160a060020a0360243516604435610a5e565b34801561038b57600080fd5b5061017d610b78565b3480156103a057600080fd5b506101cd600435610b7e565b3480156103b857600080fd5b506101cd600435610bab565b3480156103d057600080fd5b506101b6600435600160a060020a0360243516604435610bc0565b3480156103f757600080fd5b506101cd610d2a565b34801561040c57600080fd5b506101cd610d33565b34801561042157600080fd5b5061017d610d43565b600454600160a060020a031681565b60065481565b6000546301000000900460ff16158015610462575060005462010000900460ff16155b151561046d57600080fd5b600054610100900460ff16158015610488575060005460ff16155b151561049357600080fd5b60008181526007602052604090205460ff16156104af57600080fd5b6000805462ff00001961ff0019909116610100171662010000178155600160a060020a0383168082526008602090815260408084208585528252928390208481556001018054600160a060020a031990811684179091556003859055600480549091168317905582518681529081019190915280820183905290517fbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a18499181900360600190a1505050565b600054610100900460ff1681565b60008110158015610579575060648111155b151561058457600080fd5b60068190556040805182815290517f7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b199181900360200190a150565b60005462010000900460ff161580156105e05750600054610100900460ff16155b1561064a576000805462ff0000191662010000179055610601838383610d49565b60408051848152600160a060020a038416602082015280820183905290517fbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a18499181900360600190a15b505050565b60005462010000900460ff1681565b600054610100900460ff1680610676575060005460ff165b151561067e57fe5b600054610100900460ff1615610721576003546000908152600760205260409020805460ff191660011790556106b26109c8565b6005805460010190556106c3610fa2565b6001556106ce610fc0565b60025560045460035460408051848152600160a060020a03909316602084015282810191909152517f7b52eeed61a1d014375d16e12e533304d2324058cc9675f1f07e8528f91605f29181900360600190a15b60005460ff16156107bc576003546000908152600760205260409020805460ff1916905561074d6109c8565b6005805460001901905561075f610fa2565b60015561076a610fc0565b60025560045460035460408051848152600160a060020a03909316602084015282810191909152517e3363ccdf75bea36f096ec2bf58ea526c70b935d8d18986212e22953d3dad889181900360600190a15b50565b6000546301000000900460ff161580156107dc575060005460ff16155b1561064a576000805463ff000000191663010000001790556107ff838383610fdd565b60408051848152600160a060020a038416602082015280820183905290517f2c2828a1fd0ea058db2d253a84397b677abec7901c6c259ccd296fde43063efb9181900360600190a1505050565b60055481565b6000546301000000900460ff1615156001148015610871575060035481145b801561088a5750600454600160a060020a038381169116145b1561064a5761089a838383610fdd565b60408051848152600160a060020a038416602082015280820183905290517feb2dcd353a40a5049e79791269a24c2fda944048e160e31986d6731cf92c5fff9181900360600190a1505050565b60005462010000900460ff1615156001148015610905575060035481145b801561091e5750600454600160a060020a038381169116145b1561064a5761092e838383610d49565b60408051848152600160a060020a038416602082015280820183905290517f41f00c09df436400d006fe5efa9bfbf90af445c5bd1a283fecdbda2b25b2d69a9181900360600190a1505050565b600860209081526000928352604080842090915290825290208054600182015460028301546003909301549192600160a060020a0382169260a060020a90920460ff169185565b60025481565b6000805463ffffffff1916815560048054600160a060020a0390811683526008602081815260408086206003805488529083528187206002018790558554851687529282528086208354875282528086208301959095559254905484519190921681529182015281517f7350a7ef2a3d35fb20258804f114f645300830cdbb8ee760db828576c6fb81e9929181900390910190a1565b6000546301000000900460ff16158015610a81575060005462010000900460ff16155b1515610a8c57600080fd5b600054610100900460ff16158015610aa7575060005460ff16155b1515610ab257600080fd5b60008181526007602052604090205460ff161515610acf57600080fd5b600080546301000000600160ff19909216821763ff0000001916178255600160a060020a0384168083526008602090815260408085208686528252938490208581559092018054600160a060020a031990811683179091556003859055600480549091168217905582518681529182015280820183905290517fbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a1849916060908290030190a1505050565b60035481565b60008181526007602052604081205460ff16151560011415610ba257506001610ba6565b5060005b919050565b60076020526000908152604090205460ff1681565b6000546301000000900460ff16151560011480610bea575060005462010000900460ff1615156001145b8015610bf7575060035481145b8015610c105750600454600160a060020a038381169116145b1561064a57600160a060020a03821660008181526008602090815260408083208584528252918290206003018054600101908190558251878152918201939093528082018490526060810192909252517f7cb4bca1eaf530a5cfb00996c12f042fc783becb68d363236c238c081d0139029181900360800190a1610c92610fc0565b600160a060020a03831660009081526008602090815260408083208584529091529020600301541061064a5760005460408051858152600160a060020a038516602082015280820184905260ff610100840481161515606083015290921615156080830152517f2df2bdee48a50cd1f9880e4e9b7da78d3e20086869d1b784721308ab708154449181900360a00190a161064a6109c8565b60005460ff1681565b6000546301000000900460ff1681565b60015481565b60008181526007602052604090205460ff1615610d6257fe5b60005460ff1615610d6f57fe5b60005462010000900460ff161515610d8357fe5b6000805461ff001916610100178155600154600160a060020a0384168252600860209081526040808420858552909152909120600201541015610e2b57600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a600160a060020a031990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805490910190555b600154600160a060020a03831660009081526008602090815260408083208584529091529020600201541415610f3957600160a060020a038216600081815260086020908152604080832085845282528083208581556001908101805460a060020a600160a060020a031990911690961774ff000000000000000000000000000000000000000019169590951790945560079091529020805460ff19169091179055610ed56109c8565b600580546001019055610ee6610fa2565b600155610ef1610fc0565b60025560408051600160a060020a03841681526020810183905281517f58294fa6855657667fbee9342000d3676b5fd68ca0adc5fcf1387437b1d9a677929181900390910190a15b600381905560048054600160a060020a038416600160a060020a03199091168117909155604080518581526020810192909252818101839052517fb1a88ea9d6a728bc6b853a9af7a2779a8b3debd9666acd708b4cbbf78c881b509181900360600190a1505050565b600554600654606490820204600101908110610fbd57506005545b90565b60018054600554908103909101908110610fbd5750600554610fbd565b600054610100900460ff1615610fef57fe5b6000546301000000900460ff16151561100457fe5b600160a060020a0382166000908152600860209081526040808320848452909152902060019081015460a060020a900460ff1615151461104057fe5b60008181526007602052604090205460ff16151560011461105d57fe5b6000805460ff19166001908117825554600160a060020a0384168252600860209081526040808420858552909152909120600201541061109957fe5b600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a600160a060020a031990911690941774ff000000000000000000000000000000000000000019169390931790925560020180548201908190559054141561118f576000818152600760205260409020805460ff1916905561112a6109c8565b6005805460001901905561113c610fa2565b600155611147610fc0565b60025560408051600160a060020a03841681526020810183905281517f8524a451edd158bf8c11f5f8be72dc9e9453c46860748ea11ea496776000e598929181900390910190a15b600381905560048054600160a060020a038416600160a060020a03199091168117909155604080518581526020810192909252818101839052517f5d17347382fb804e38ff8b76979b7f52282ca8bb785f61658e4e85f6df851d349181900360600190a15050505600a165627a7a7230582016766ff312307be09a7bb8e82b82bec4a7cb4496f4f9b395785fa2ae204ff3a30029`

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

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(_id uint256) returns()
func (_Permissions *PermissionsTransactor) AcceptProposal(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "acceptProposal", _id)
}

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(_id uint256) returns()
func (_Permissions *PermissionsSession) AcceptProposal(_id *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.AcceptProposal(&_Permissions.TransactOpts, _id)
}

// AcceptProposal is a paid mutator transaction binding the contract method 0x60c5cc3a.
//
// Solidity: function acceptProposal(_id uint256) returns()
func (_Permissions *PermissionsTransactorSession) AcceptProposal(_id *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.AcceptProposal(&_Permissions.TransactOpts, _id)
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

// AddNodeProposal is a paid mutator transaction binding the contract method 0x1d3a12d7.
//
// Solidity: function addNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) AddNodeProposal(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addNodeProposal", _id, _account, _enode)
}

// AddNodeProposal is a paid mutator transaction binding the contract method 0x1d3a12d7.
//
// Solidity: function addNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) AddNodeProposal(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddNodeProposal(&_Permissions.TransactOpts, _id, _account, _enode)
}

// AddNodeProposal is a paid mutator transaction binding the contract method 0x1d3a12d7.
//
// Solidity: function addNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) AddNodeProposal(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.AddNodeProposal(&_Permissions.TransactOpts, _id, _account, _enode)
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

// SuspendNodeProposal is a paid mutator transaction binding the contract method 0xa29cd372.
//
// Solidity: function suspendNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactor) SuspendNodeProposal(opts *bind.TransactOpts, _id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "suspendNodeProposal", _id, _account, _enode)
}

// SuspendNodeProposal is a paid mutator transaction binding the contract method 0xa29cd372.
//
// Solidity: function suspendNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsSession) SuspendNodeProposal(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNodeProposal(&_Permissions.TransactOpts, _id, _account, _enode)
}

// SuspendNodeProposal is a paid mutator transaction binding the contract method 0xa29cd372.
//
// Solidity: function suspendNodeProposal(_id uint256, _account address, _enode bytes32) returns()
func (_Permissions *PermissionsTransactorSession) SuspendNodeProposal(_id *big.Int, _account common.Address, _enode [32]byte) (*types.Transaction, error) {
	return _Permissions.Contract.SuspendNodeProposal(&_Permissions.TransactOpts, _id, _account, _enode)
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
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddNode is a free log retrieval operation binding the contract event 0xb1a88ea9d6a728bc6b853a9af7a2779a8b3debd9666acd708b4cbbf78c881b50.
//
// Solidity: e LogOfAddNode(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfAddNode(opts *bind.FilterOpts) (*PermissionsLogOfAddNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddNodeIterator{contract: _Permissions.contract, event: "LogOfAddNode", logs: logs, sub: sub}, nil
}

// WatchLogOfAddNode is a free log subscription operation binding the contract event 0xb1a88ea9d6a728bc6b853a9af7a2779a8b3debd9666acd708b4cbbf78c881b50.
//
// Solidity: e LogOfAddNode(id uint256, accountaddress address, enodeaddress bytes32)
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

// PermissionsLogOfAddNodeProposalIterator is returned from FilterLogOfAddNodeProposal and is used to iterate over the raw logs and unpacked data for LogOfAddNodeProposal events raised by the Permissions contract.
type PermissionsLogOfAddNodeProposalIterator struct {
	Event *PermissionsLogOfAddNodeProposal // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddNodeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddNodeProposal)
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
		it.Event = new(PermissionsLogOfAddNodeProposal)
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
func (it *PermissionsLogOfAddNodeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddNodeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddNodeProposal represents a LogOfAddNodeProposal event raised by the Permissions contract.
type PermissionsLogOfAddNodeProposal struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddNodeProposal is a free log retrieval operation binding the contract event 0xbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a1849.
//
// Solidity: e LogOfAddNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfAddNodeProposal(opts *bind.FilterOpts) (*PermissionsLogOfAddNodeProposalIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddNodeProposal")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddNodeProposalIterator{contract: _Permissions.contract, event: "LogOfAddNodeProposal", logs: logs, sub: sub}, nil
}

// WatchLogOfAddNodeProposal is a free log subscription operation binding the contract event 0xbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a1849.
//
// Solidity: e LogOfAddNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfAddNodeProposal(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddNodeProposal) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddNodeProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddNodeProposal)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddNodeProposal", log); err != nil {
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

// PermissionsLogOfAddingAcceptProposalIterator is returned from FilterLogOfAddingAcceptProposal and is used to iterate over the raw logs and unpacked data for LogOfAddingAcceptProposal events raised by the Permissions contract.
type PermissionsLogOfAddingAcceptProposalIterator struct {
	Event *PermissionsLogOfAddingAcceptProposal // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddingAcceptProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddingAcceptProposal)
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
		it.Event = new(PermissionsLogOfAddingAcceptProposal)
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
func (it *PermissionsLogOfAddingAcceptProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddingAcceptProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddingAcceptProposal represents a LogOfAddingAcceptProposal event raised by the Permissions contract.
type PermissionsLogOfAddingAcceptProposal struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddingAcceptProposal is a free log retrieval operation binding the contract event 0x7b52eeed61a1d014375d16e12e533304d2324058cc9675f1f07e8528f91605f2.
//
// Solidity: e LogOfAddingAcceptProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfAddingAcceptProposal(opts *bind.FilterOpts) (*PermissionsLogOfAddingAcceptProposalIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddingAcceptProposal")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddingAcceptProposalIterator{contract: _Permissions.contract, event: "LogOfAddingAcceptProposal", logs: logs, sub: sub}, nil
}

// WatchLogOfAddingAcceptProposal is a free log subscription operation binding the contract event 0x7b52eeed61a1d014375d16e12e533304d2324058cc9675f1f07e8528f91605f2.
//
// Solidity: e LogOfAddingAcceptProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfAddingAcceptProposal(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddingAcceptProposal) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddingAcceptProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddingAcceptProposal)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddingAcceptProposal", log); err != nil {
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

// PermissionsLogOfAddingConsensusMeetIterator is returned from FilterLogOfAddingConsensusMeet and is used to iterate over the raw logs and unpacked data for LogOfAddingConsensusMeet events raised by the Permissions contract.
type PermissionsLogOfAddingConsensusMeetIterator struct {
	Event *PermissionsLogOfAddingConsensusMeet // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddingConsensusMeetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddingConsensusMeet)
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
		it.Event = new(PermissionsLogOfAddingConsensusMeet)
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
func (it *PermissionsLogOfAddingConsensusMeetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddingConsensusMeetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddingConsensusMeet represents a LogOfAddingConsensusMeet event raised by the Permissions contract.
type PermissionsLogOfAddingConsensusMeet struct {
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddingConsensusMeet is a free log retrieval operation binding the contract event 0x58294fa6855657667fbee9342000d3676b5fd68ca0adc5fcf1387437b1d9a677.
//
// Solidity: e LogOfAddingConsensusMeet(accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfAddingConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfAddingConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddingConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddingConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfAddingConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfAddingConsensusMeet is a free log subscription operation binding the contract event 0x58294fa6855657667fbee9342000d3676b5fd68ca0adc5fcf1387437b1d9a677.
//
// Solidity: e LogOfAddingConsensusMeet(accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfAddingConsensusMeet(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddingConsensusMeet) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddingConsensusMeet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddingConsensusMeet)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddingConsensusMeet", log); err != nil {
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
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddingVote is a free log retrieval operation binding the contract event 0x41f00c09df436400d006fe5efa9bfbf90af445c5bd1a283fecdbda2b25b2d69a.
//
// Solidity: e LogOfAddingVote(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfAddingVote(opts *bind.FilterOpts) (*PermissionsLogOfAddingVoteIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddingVote")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddingVoteIterator{contract: _Permissions.contract, event: "LogOfAddingVote", logs: logs, sub: sub}, nil
}

// WatchLogOfAddingVote is a free log subscription operation binding the contract event 0x41f00c09df436400d006fe5efa9bfbf90af445c5bd1a283fecdbda2b25b2d69a.
//
// Solidity: e LogOfAddingVote(id uint256, accountaddress address, enodeaddress bytes32)
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

// PermissionsLogOfSuspandNodeProposalIterator is returned from FilterLogOfSuspandNodeProposal and is used to iterate over the raw logs and unpacked data for LogOfSuspandNodeProposal events raised by the Permissions contract.
type PermissionsLogOfSuspandNodeProposalIterator struct {
	Event *PermissionsLogOfSuspandNodeProposal // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSuspandNodeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSuspandNodeProposal)
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
		it.Event = new(PermissionsLogOfSuspandNodeProposal)
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
func (it *PermissionsLogOfSuspandNodeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSuspandNodeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSuspandNodeProposal represents a LogOfSuspandNodeProposal event raised by the Permissions contract.
type PermissionsLogOfSuspandNodeProposal struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspandNodeProposal is a free log retrieval operation binding the contract event 0x2c2828a1fd0ea058db2d253a84397b677abec7901c6c259ccd296fde43063efb.
//
// Solidity: e LogOfSuspandNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspandNodeProposal(opts *bind.FilterOpts) (*PermissionsLogOfSuspandNodeProposalIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspandNodeProposal")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspandNodeProposalIterator{contract: _Permissions.contract, event: "LogOfSuspandNodeProposal", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspandNodeProposal is a free log subscription operation binding the contract event 0x2c2828a1fd0ea058db2d253a84397b677abec7901c6c259ccd296fde43063efb.
//
// Solidity: e LogOfSuspandNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfSuspandNodeProposal(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSuspandNodeProposal) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSuspandNodeProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSuspandNodeProposal)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSuspandNodeProposal", log); err != nil {
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

// PermissionsLogOfSuspentionAcceptProposalIterator is returned from FilterLogOfSuspentionAcceptProposal and is used to iterate over the raw logs and unpacked data for LogOfSuspentionAcceptProposal events raised by the Permissions contract.
type PermissionsLogOfSuspentionAcceptProposalIterator struct {
	Event *PermissionsLogOfSuspentionAcceptProposal // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSuspentionAcceptProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSuspentionAcceptProposal)
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
		it.Event = new(PermissionsLogOfSuspentionAcceptProposal)
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
func (it *PermissionsLogOfSuspentionAcceptProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSuspentionAcceptProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSuspentionAcceptProposal represents a LogOfSuspentionAcceptProposal event raised by the Permissions contract.
type PermissionsLogOfSuspentionAcceptProposal struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionAcceptProposal is a free log retrieval operation binding the contract event 0x003363ccdf75bea36f096ec2bf58ea526c70b935d8d18986212e22953d3dad88.
//
// Solidity: e LogOfSuspentionAcceptProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionAcceptProposal(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionAcceptProposalIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionAcceptProposal")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionAcceptProposalIterator{contract: _Permissions.contract, event: "LogOfSuspentionAcceptProposal", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionAcceptProposal is a free log subscription operation binding the contract event 0x003363ccdf75bea36f096ec2bf58ea526c70b935d8d18986212e22953d3dad88.
//
// Solidity: e LogOfSuspentionAcceptProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfSuspentionAcceptProposal(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSuspentionAcceptProposal) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSuspentionAcceptProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSuspentionAcceptProposal)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSuspentionAcceptProposal", log); err != nil {
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
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionNode is a free log retrieval operation binding the contract event 0x5d17347382fb804e38ff8b76979b7f52282ca8bb785f61658e4e85f6df851d34.
//
// Solidity: e LogOfSuspentionNode(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionNode(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionNodeIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionNode")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionNodeIterator{contract: _Permissions.contract, event: "LogOfSuspentionNode", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionNode is a free log subscription operation binding the contract event 0x5d17347382fb804e38ff8b76979b7f52282ca8bb785f61658e4e85f6df851d34.
//
// Solidity: e LogOfSuspentionNode(id uint256, accountaddress address, enodeaddress bytes32)
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
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspentionVote is a free log retrieval operation binding the contract event 0xeb2dcd353a40a5049e79791269a24c2fda944048e160e31986d6731cf92c5fff.
//
// Solidity: e LogOfSuspentionVote(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspentionVote(opts *bind.FilterOpts) (*PermissionsLogOfSuspentionVoteIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspentionVote")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspentionVoteIterator{contract: _Permissions.contract, event: "LogOfSuspentionVote", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspentionVote is a free log subscription operation binding the contract event 0xeb2dcd353a40a5049e79791269a24c2fda944048e160e31986d6731cf92c5fff.
//
// Solidity: e LogOfSuspentionVote(id uint256, accountaddress address, enodeaddress bytes32)
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
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfsuspentionConsensusMeet is a free log retrieval operation binding the contract event 0x8524a451edd158bf8c11f5f8be72dc9e9453c46860748ea11ea496776000e598.
//
// Solidity: e LogOfsuspentionConsensusMeet(accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfsuspentionConsensusMeet(opts *bind.FilterOpts) (*PermissionsLogOfsuspentionConsensusMeetIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfsuspentionConsensusMeet")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfsuspentionConsensusMeetIterator{contract: _Permissions.contract, event: "LogOfsuspentionConsensusMeet", logs: logs, sub: sub}, nil
}

// WatchLogOfsuspentionConsensusMeet is a free log subscription operation binding the contract event 0x8524a451edd158bf8c11f5f8be72dc9e9453c46860748ea11ea496776000e598.
//
// Solidity: e LogOfsuspentionConsensusMeet(accountaddress address, enodeaddress bytes32)
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
