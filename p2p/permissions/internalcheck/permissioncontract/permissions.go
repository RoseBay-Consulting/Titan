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
const PermissionsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"previousaccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"consensuspercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addNodeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"votesforoperatorcounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressofproposalacceptor\",\"type\":\"address\"}],\"name\":\"proposalaccepterEntry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account_of_operator\",\"type\":\"address\"}],\"name\":\"newOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isadding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"setConsensus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterindexaccountpair\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorindexaccountpair\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addingmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"acceptProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NodeCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"operatorEntry\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"addingVote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeinfo\",\"outputs\":[{\"name\":\"enode\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\"},{\"name\":\"votecount\",\"type\":\"uint256\"},{\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfNegVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operatorindex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"suspendNodeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account_of_operator\",\"type\":\"address\"}],\"name\":\"removeOperatorAccount\",\"outputs\":[{\"name\":\"status\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousenode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"checkNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operatorconsensuspercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"nodeconformations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_enode\",\"type\":\"bytes32\"}],\"name\":\"voteReject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuspention\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_percentage\",\"type\":\"uint256\"}],\"name\":\"setOperatorConsensus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"suspentionmutex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressofproposalacceptor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterindex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LimitOfVote\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consensuslimit\",\"type\":\"uint256\"}],\"name\":\"LogOfSetConsensus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousaccount\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousenode\",\"type\":\"bytes32\"}],\"name\":\"LogOfResetProcess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfsuspentionConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddNodeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspendNodeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfAddingAcceptProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"}],\"name\":\"LogOfSuspentionAcceptProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"rejectcount\",\"type\":\"uint256\"}],\"name\":\"LogOfVoteReject\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"enodeaddress\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"isadding\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"issuspention\",\"type\":\"bool\"}],\"name\":\"LogOfVoteRejectConsensusMeet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfRemoveVoterAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfActiveVoterAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consensuslimit\",\"type\":\"uint256\"}],\"name\":\"LogOfOperatorSetConsensus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfRemoveOperatorAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfActiveOperatorAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfResetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"accountaddress\",\"type\":\"address\"}],\"name\":\"LogOfAddressOfProposalAcceptorEntry\",\"type\":\"event\"}]"

// PermissionsBin is the compiled bytecode used for deploying new contracts.
const PermissionsBin = `0x608060405234801561001057600080fd5b506000600181905560028190556005556119ce8061002f6000396000f3006080604052600436106101b65763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663118e696781146101bb57806313e7c9d8146101ec5780631867f60c146102215780631d3a12d7146102485780632705c65e14610271578063274b7767146102865780632771ca74146102a75780632da0bca8146102c85780633549ef17146102dd57806347aa94ea146102f55780634fe75eeb1461031657806360c229f71461033757806360c5cc3a1461034c5780636a19f6ad14610364578063743d3eb51461037957806374a166fb1461038e57806375bccb39146103b557806397b1a65d146103dc5780639a7bf5bd146104335780639f3cad6014610448578063a29cd3721461045d578063a3ec138d14610484578063ad1999ac146104a5578063b1b49f25146104c6578063b50d6f75146104db578063c0a421ed146104f3578063c86d87b914610508578063c9663fac14610520578063d6336a3314610547578063db7882381461055c578063eae897c714610574578063ebdb432f14610589578063f3a91fec1461059e578063f60cd216146105b3578063ff5c11f8146105c8575b600080fd5b3480156101c757600080fd5b506101d06105dd565b60408051600160a060020a039092168252519081900360200190f35b3480156101f857600080fd5b5061020d600160a060020a03600435166105ec565b604080519115158252519081900360200190f35b34801561022d57600080fd5b50610236610601565b60408051918252519081900360200190f35b34801561025457600080fd5b5061026f600435600160a060020a0360243516604435610607565b005b34801561027d57600080fd5b5061023661073a565b34801561029257600080fd5b5061026f600160a060020a0360043516610740565b3480156102b357600080fd5b5061026f600160a060020a03600435166107b3565b3480156102d457600080fd5b5061020d6107e4565b3480156102e957600080fd5b5061026f6004356107f2565b34801561030157600080fd5b50610236600160a060020a0360043516610864565b34801561032257600080fd5b50610236600160a060020a0360043516610876565b34801561034357600080fd5b5061020d610888565b34801561035857600080fd5b5061026f600435610897565b34801561037057600080fd5b50610236610a7b565b34801561038557600080fd5b5061020d610a81565b34801561039a57600080fd5b5061026f600435600160a060020a0360243516604435610ab1565b3480156103c157600080fd5b5061026f600435600160a060020a0360243516604435610b6a565b3480156103e857600080fd5b50610400600160a060020a0360043516602435610c21565b60408051958652600160a060020a0390941660208601529115158484015260608401526080830152519081900360a00190f35b34801561043f57600080fd5b50610236610c68565b34801561045457600080fd5b50610236610c6e565b34801561046957600080fd5b5061026f600435600160a060020a0360243516604435610c74565b34801561049057600080fd5b5061020d600160a060020a0360043516610da6565b3480156104b157600080fd5b5061020d600160a060020a0360043516610dbb565b3480156104d257600080fd5b50610236610ea0565b3480156104e757600080fd5b5061020d600435610ea6565b3480156104ff57600080fd5b50610236610ed2565b34801561051457600080fd5b5061020d600435610ed8565b34801561052c57600080fd5b5061026f600435600160a060020a0360243516604435610eed565b34801561055357600080fd5b5061020d61107a565b34801561056857600080fd5b5061026f600435611083565b34801561058057600080fd5b5061020d6110f5565b34801561059557600080fd5b5061026f611105565b3480156105aa57600080fd5b506101d06111b5565b3480156105bf57600080fd5b506102366111c4565b3480156105d457600080fd5b506102366111ca565b600454600160a060020a031681565b600f6020526000908152604090205460ff1681565b60065481565b336000908152600f602052604090205460ff16156101b6576000546301000000900460ff16158015610642575060005462010000900460ff16155b151561064d57600080fd5b600054610100900460ff16158015610668575060005460ff16155b151561067357600080fd5b60008181526007602052604090205460ff161561068f57600080fd5b6000805462ff00001961ff0019909116610100171662010000178155600160a060020a0383168082526008602090815260408084208585528252928390208481556001018054600160a060020a031990811684179091556003859055600480549091168317905582518681529081019190915280820183905290517fbacda36ac250ad037a62b6a67ee6021bafae3103cfd00792dad9c4e8f71a18499181900360600190a15b505050565b60115481565b336000908152600f602052604090205460ff16156101b65760138054600160a060020a031916600160a060020a03838116919091179182905560408051929091168252517fdb322fa900d5273c470b8847f6a0a42231e31959eba17a6066d5680425255c89916020908290030190a15b50565b336000908152600f602052604090205460ff16156101b6576107d4816111d0565b15156107df57600080fd5b6107b0565b600054610100900460ff1681565b336000908152600f602052604090205460ff16156101b6576000811015801561081c575060648111155b151561082757600080fd5b60068190556040805182815290517f7eb26b746889baac2fb48d4457332fb465f503b675451a7e294120693da45b199181900360200190a16107b0565b60096020526000908152604090205481565b600d6020526000908152604090205481565b60005462010000900460ff1681565b601354600160a060020a03163314156101b657600054610100900460ff16806108c2575060005460ff165b15156108ca57fe5b600054610100900460ff16156109c75760048054600160a060020a0390811660009081526008602090815260408083206003805485529083528184206001908101805474ff0000000000000000000000000000000000000000191660a060020a179055905484526007909252909120805460ff191690911790559054610950911661128a565b61095861133f565b6005805460010190556109696113d5565b6001556109746113f9565b60025560045460035460408051848152600160a060020a03909316602084015282810191909152517f7b52eeed61a1d014375d16e12e533304d2324058cc9675f1f07e8528f91605f29181900360600190a15b60005460ff16156107df576003546000908152600760205260409020805460ff19169055600454610a0090600160a060020a0316611416565b610a0861133f565b60058054600019019055610a1a6113d5565b600155610a256113f9565b60025560045460035460408051848152600160a060020a03909316602084015282810191909152517e3363ccdf75bea36f096ec2bf58ea526c70b935d8d18986212e22953d3dad889181900360600190a16107b0565b60055481565b6000600e5460001415610aaa57610a97336111d0565b1515610aa257600080fd5b506001610aae565b5060005b90565b336000908152600b602052604090205460ff1680610acf5750600a54155b156101b6576000546301000000900460ff1615156001148015610af3575060035481145b8015610b0c5750600454600160a060020a038381169116145b15610b6557610b1c8383836114dd565b60408051848152600160a060020a038416602082015280820183905290517feb2dcd353a40a5049e79791269a24c2fda944048e160e31986d6731cf92c5fff9181900360600190a15b610735565b336000908152600b602052604090205460ff1680610b885750600a54155b156101b65760005462010000900460ff1615156001148015610bab575060035481145b8015610bc45750600454600160a060020a038381169116145b15610b6557610bd483838361170d565b60408051848152600160a060020a038416602082015280820183905290517f41f00c09df436400d006fe5efa9bfbf90af445c5bd1a283fecdbda2b25b2d69a9181900360600190a1610735565b600860209081526000928352604080842090915290825290208054600182015460028301546003909301549192600160a060020a0382169260a060020a90920460ff169185565b60025481565b600e5481565b336000908152600f602052604090205460ff16156101b6576000546301000000900460ff16158015610caf575060005462010000900460ff16155b1515610cba57600080fd5b600054610100900460ff16158015610cd5575060005460ff16155b1515610ce057600080fd5b60008181526007602052604090205460ff161515610cfd57600080fd5b600080546301000000600160ff19909216821763ff0000001916178255600160a060020a0384168083526008602090815260408085208686528252938490208581559092018054600160a060020a031990811683179091556003859055600480549091168217905582518681529182015280820183905290517f48dfbd110e878cbf6e08b7bb004151d52e571f0c464de79f3a07fee91741afe3916060908290030190a1610735565b600b6020526000908152604090205460ff1681565b336000908152600f602052604081205460ff16156101b657600160a060020a0382166000908152600f602052604090205460ff161515610dfa57600080fd5b600160a060020a038083166000818152600f60209081526040808320805460ff19169055600e80546000199081018552601080855283862054878752600d8652848720548752908552948390208054600160a060020a031916959097169490941790955584549092019093558051918252517fda7af3b29c038afb615f78757e855b5c9131654c06a8a6f0ba33b7b727215465929181900390910190a15060015b919050565b60035481565b60008181526007602052604081205460ff16151560011415610eca57506001610e9b565b506000610e9b565b60125481565b60076020526000908152604090205460ff1681565b336000908152600b602052604090205460ff1680610f0b5750600a54155b156101b6576000546301000000900460ff16151560011480610f3a575060005462010000900460ff1615156001145b8015610f47575060035481145b8015610f605750600454600160a060020a038381169116145b15610b6557600160a060020a03821660008181526008602090815260408083208584528252918290206003018054600101908190558251878152918201939093528082018490526060810192909252517f7cb4bca1eaf530a5cfb00996c12f042fc783becb68d363236c238c081d0139029181900360800190a1610fe26113f9565b600160a060020a038316600090815260086020908152604080832085845290915290206003015410610b655760005460408051858152600160a060020a038516602082015280820184905260ff610100840481161515606083015290921615156080830152517f2df2bdee48a50cd1f9880e4e9b7da78d3e20086869d1b784721308ab708154449181900360a00190a1610b6561133f565b60005460ff1681565b336000908152600f602052604090205460ff16156101b657600081101580156110ad575060648111155b15156110b857600080fd5b60128190556040805182815290517fae848b724d627813f2afaa324ae4610e1503087865f73d5d57d2a90f605c17d59181900360200190a16107b0565b6000546301000000900460ff1681565b336000908152600b602052604090205460ff16806111235750600a54155b156101b657611130611974565b6011541415611176576000600e556040805133815290517f8724988bd627c5b9b96f53247d1031353bf31098365566104e129627052ced989181900360200190a16111b3565b6011805460010190556040805133815290517f8724988bd627c5b9b96f53247d1031353bf31098365566104e129627052ced989181900360200190a15b565b601354600160a060020a031681565b600a5481565b60015481565b600160a060020a0381166000908152600f602052604081205460ff16156111f657600080fd5b600160a060020a0382166000818152600f60209081526040808320805460ff19166001908117909155600e80548552601084528285208054600160a060020a031916871790558054868652600d8552948390208590559301909255815192835290517fd0cd818f9f2837a5e7dc53d100353f4ba5b023a116fb5608841996d4771277c19281900390910190a1506001919050565b600160a060020a0381166000908152600b602052604090205460ff16156112b057600080fd5b600160a060020a0381166000818152600b60209081526040808320805460ff19166001908117909155600a80548552600c84528285208054600160a060020a03191687179055805486865260098552948390208590559301909255815192835290517f45b6638126291e2dab4e36d975252dcb55bc211569e9f1169309199d2bbe74e79281900390910190a150565b6000805463ffffffff1916815560048054600160a060020a0390811683526008602081815260408086206003805488529083528187206002018790558554851687529282528086208354875282528086208301959095559254905484519190921681529182015281517f7350a7ef2a3d35fb20258804f114f645300830cdbb8ee760db828576c6fb81e9929181900390910190a1565b6005546006546064908202046001019081106113f45750600554610aae565b610aae565b600180546005549081039091019081106113f45750600554610aae565b600160a060020a0381166000908152600b602052604090205460ff16151561143d57600080fd5b600160a060020a038082166000818152600b60209081526040808320805460ff19169055600a80546000199081018552600c8085528386205487875260098652848720548752908552948390208054600160a060020a031916959097169490941790955584549092019093558051918252517fceac34d6c239fde6d920bdf4ab89260499033335e865e17d611bf081f33d65fe929181900390910190a150565b600054610100900460ff16156114ef57fe5b6000546301000000900460ff16151561150457fe5b600160a060020a0382166000908152600860209081526040808320848452909152902060019081015460a060020a900460ff1615151461154057fe5b60008181526007602052604090205460ff16151560011461155d57fe5b6000805460ff19166001908117825554600160a060020a0384168252600860209081526040808420858552909152909120600201541061159957fe5b600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a600160a060020a031990911690941774ff00000000000000000000000000000000000000001916939093179092556002018054820190819055905414156116a4576000818152600760205260409020805460ff1916905560045461163790600160a060020a0316611416565b61163f61133f565b600580546000190190556116516113d5565b60015561165c6113f9565b60025560408051600160a060020a03841681526020810183905281517f8524a451edd158bf8c11f5f8be72dc9e9453c46860748ea11ea496776000e598929181900390910190a15b600381905560048054600160a060020a038416600160a060020a03199091168117909155604080518581526020810192909252818101839052517f5d17347382fb804e38ff8b76979b7f52282ca8bb785f61658e4e85f6df851d349181900360600190a1505050565b60008181526007602052604090205460ff161561172657fe5b60005460ff161561173357fe5b60005462010000900460ff16151561174757fe5b6000805461ff001916610100178155600154600160a060020a03841682526008602090815260408084208585529091529091206002015410156117ef57600160a060020a038216600081815260086020908152604080832085845290915290208281556001808201805460a060020a600160a060020a031990911690941774ff0000000000000000000000000000000000000000191693909317909255600201805490910190555b600154600160a060020a0383166000908152600860209081526040808320858452909152902060020154141561190b57600160a060020a03828116600081815260086020908152604080832086845282528083208681556001908101805460a060020a600160a060020a031990911690961774ff000000000000000000000000000000000000000019169590951790945560079091529020805460ff1916909117905560045461189f911661128a565b6118a761133f565b6005805460010190556118b86113d5565b6001556118c36113f9565b60025560408051600160a060020a03841681526020810183905281517f58294fa6855657667fbee9342000d3676b5fd68ca0adc5fcf1387437b1d9a677929181900390910190a15b600381905560048054600160a060020a038416600160a060020a03199091168117909155604080518581526020810192909252818101839052517fb1a88ea9d6a728bc6b853a9af7a2779a8b3debd9666acd708b4cbbf78c881b509181900360600190a1505050565b600060646001600a54036012540281151561198b57fe5b049050600a54811015156113f45750600a54610aae5600a165627a7a723058202ef04aacde93fb3e03993643b9f21c3a268850ed6e2f068818a8fcf28dec89140029`

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

// Addressofproposalacceptor is a free data retrieval call binding the contract method 0xf3a91fec.
//
// Solidity: function addressofproposalacceptor() constant returns(address)
func (_Permissions *PermissionsCaller) Addressofproposalacceptor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "addressofproposalacceptor")
	return *ret0, err
}

// Addressofproposalacceptor is a free data retrieval call binding the contract method 0xf3a91fec.
//
// Solidity: function addressofproposalacceptor() constant returns(address)
func (_Permissions *PermissionsSession) Addressofproposalacceptor() (common.Address, error) {
	return _Permissions.Contract.Addressofproposalacceptor(&_Permissions.CallOpts)
}

// Addressofproposalacceptor is a free data retrieval call binding the contract method 0xf3a91fec.
//
// Solidity: function addressofproposalacceptor() constant returns(address)
func (_Permissions *PermissionsCallerSession) Addressofproposalacceptor() (common.Address, error) {
	return _Permissions.Contract.Addressofproposalacceptor(&_Permissions.CallOpts)
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

// Operatorconsensuspercentage is a free data retrieval call binding the contract method 0xc0a421ed.
//
// Solidity: function operatorconsensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsCaller) Operatorconsensuspercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "operatorconsensuspercentage")
	return *ret0, err
}

// Operatorconsensuspercentage is a free data retrieval call binding the contract method 0xc0a421ed.
//
// Solidity: function operatorconsensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsSession) Operatorconsensuspercentage() (*big.Int, error) {
	return _Permissions.Contract.Operatorconsensuspercentage(&_Permissions.CallOpts)
}

// Operatorconsensuspercentage is a free data retrieval call binding the contract method 0xc0a421ed.
//
// Solidity: function operatorconsensuspercentage() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Operatorconsensuspercentage() (*big.Int, error) {
	return _Permissions.Contract.Operatorconsensuspercentage(&_Permissions.CallOpts)
}

// Operatorindex is a free data retrieval call binding the contract method 0x9f3cad60.
//
// Solidity: function operatorindex() constant returns(uint256)
func (_Permissions *PermissionsCaller) Operatorindex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "operatorindex")
	return *ret0, err
}

// Operatorindex is a free data retrieval call binding the contract method 0x9f3cad60.
//
// Solidity: function operatorindex() constant returns(uint256)
func (_Permissions *PermissionsSession) Operatorindex() (*big.Int, error) {
	return _Permissions.Contract.Operatorindex(&_Permissions.CallOpts)
}

// Operatorindex is a free data retrieval call binding the contract method 0x9f3cad60.
//
// Solidity: function operatorindex() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Operatorindex() (*big.Int, error) {
	return _Permissions.Contract.Operatorindex(&_Permissions.CallOpts)
}

// Operatorindexaccountpair is a free data retrieval call binding the contract method 0x4fe75eeb.
//
// Solidity: function operatorindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsCaller) Operatorindexaccountpair(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "operatorindexaccountpair", arg0)
	return *ret0, err
}

// Operatorindexaccountpair is a free data retrieval call binding the contract method 0x4fe75eeb.
//
// Solidity: function operatorindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsSession) Operatorindexaccountpair(arg0 common.Address) (*big.Int, error) {
	return _Permissions.Contract.Operatorindexaccountpair(&_Permissions.CallOpts, arg0)
}

// Operatorindexaccountpair is a free data retrieval call binding the contract method 0x4fe75eeb.
//
// Solidity: function operatorindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Operatorindexaccountpair(arg0 common.Address) (*big.Int, error) {
	return _Permissions.Contract.Operatorindexaccountpair(&_Permissions.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(bool)
func (_Permissions *PermissionsCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "operators", arg0)
	return *ret0, err
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(bool)
func (_Permissions *PermissionsSession) Operators(arg0 common.Address) (bool, error) {
	return _Permissions.Contract.Operators(&_Permissions.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators( address) constant returns(bool)
func (_Permissions *PermissionsCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Permissions.Contract.Operators(&_Permissions.CallOpts, arg0)
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

// Voterindex is a free data retrieval call binding the contract method 0xf60cd216.
//
// Solidity: function voterindex() constant returns(uint256)
func (_Permissions *PermissionsCaller) Voterindex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "voterindex")
	return *ret0, err
}

// Voterindex is a free data retrieval call binding the contract method 0xf60cd216.
//
// Solidity: function voterindex() constant returns(uint256)
func (_Permissions *PermissionsSession) Voterindex() (*big.Int, error) {
	return _Permissions.Contract.Voterindex(&_Permissions.CallOpts)
}

// Voterindex is a free data retrieval call binding the contract method 0xf60cd216.
//
// Solidity: function voterindex() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Voterindex() (*big.Int, error) {
	return _Permissions.Contract.Voterindex(&_Permissions.CallOpts)
}

// Voterindexaccountpair is a free data retrieval call binding the contract method 0x47aa94ea.
//
// Solidity: function voterindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsCaller) Voterindexaccountpair(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "voterindexaccountpair", arg0)
	return *ret0, err
}

// Voterindexaccountpair is a free data retrieval call binding the contract method 0x47aa94ea.
//
// Solidity: function voterindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsSession) Voterindexaccountpair(arg0 common.Address) (*big.Int, error) {
	return _Permissions.Contract.Voterindexaccountpair(&_Permissions.CallOpts, arg0)
}

// Voterindexaccountpair is a free data retrieval call binding the contract method 0x47aa94ea.
//
// Solidity: function voterindexaccountpair( address) constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Voterindexaccountpair(arg0 common.Address) (*big.Int, error) {
	return _Permissions.Contract.Voterindexaccountpair(&_Permissions.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(bool)
func (_Permissions *PermissionsCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "voters", arg0)
	return *ret0, err
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(bool)
func (_Permissions *PermissionsSession) Voters(arg0 common.Address) (bool, error) {
	return _Permissions.Contract.Voters(&_Permissions.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters( address) constant returns(bool)
func (_Permissions *PermissionsCallerSession) Voters(arg0 common.Address) (bool, error) {
	return _Permissions.Contract.Voters(&_Permissions.CallOpts, arg0)
}

// Votesforoperatorcounter is a free data retrieval call binding the contract method 0x2705c65e.
//
// Solidity: function votesforoperatorcounter() constant returns(uint256)
func (_Permissions *PermissionsCaller) Votesforoperatorcounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "votesforoperatorcounter")
	return *ret0, err
}

// Votesforoperatorcounter is a free data retrieval call binding the contract method 0x2705c65e.
//
// Solidity: function votesforoperatorcounter() constant returns(uint256)
func (_Permissions *PermissionsSession) Votesforoperatorcounter() (*big.Int, error) {
	return _Permissions.Contract.Votesforoperatorcounter(&_Permissions.CallOpts)
}

// Votesforoperatorcounter is a free data retrieval call binding the contract method 0x2705c65e.
//
// Solidity: function votesforoperatorcounter() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) Votesforoperatorcounter() (*big.Int, error) {
	return _Permissions.Contract.Votesforoperatorcounter(&_Permissions.CallOpts)
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

// NewOperator is a paid mutator transaction binding the contract method 0x2771ca74.
//
// Solidity: function newOperator(_account_of_operator address) returns()
func (_Permissions *PermissionsTransactor) NewOperator(opts *bind.TransactOpts, _account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "newOperator", _account_of_operator)
}

// NewOperator is a paid mutator transaction binding the contract method 0x2771ca74.
//
// Solidity: function newOperator(_account_of_operator address) returns()
func (_Permissions *PermissionsSession) NewOperator(_account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.NewOperator(&_Permissions.TransactOpts, _account_of_operator)
}

// NewOperator is a paid mutator transaction binding the contract method 0x2771ca74.
//
// Solidity: function newOperator(_account_of_operator address) returns()
func (_Permissions *PermissionsTransactorSession) NewOperator(_account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.NewOperator(&_Permissions.TransactOpts, _account_of_operator)
}

// OperatorEntry is a paid mutator transaction binding the contract method 0x743d3eb5.
//
// Solidity: function operatorEntry() returns(bool)
func (_Permissions *PermissionsTransactor) OperatorEntry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "operatorEntry")
}

// OperatorEntry is a paid mutator transaction binding the contract method 0x743d3eb5.
//
// Solidity: function operatorEntry() returns(bool)
func (_Permissions *PermissionsSession) OperatorEntry() (*types.Transaction, error) {
	return _Permissions.Contract.OperatorEntry(&_Permissions.TransactOpts)
}

// OperatorEntry is a paid mutator transaction binding the contract method 0x743d3eb5.
//
// Solidity: function operatorEntry() returns(bool)
func (_Permissions *PermissionsTransactorSession) OperatorEntry() (*types.Transaction, error) {
	return _Permissions.Contract.OperatorEntry(&_Permissions.TransactOpts)
}

// ProposalaccepterEntry is a paid mutator transaction binding the contract method 0x274b7767.
//
// Solidity: function proposalaccepterEntry(_addressofproposalacceptor address) returns()
func (_Permissions *PermissionsTransactor) ProposalaccepterEntry(opts *bind.TransactOpts, _addressofproposalacceptor common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "proposalaccepterEntry", _addressofproposalacceptor)
}

// ProposalaccepterEntry is a paid mutator transaction binding the contract method 0x274b7767.
//
// Solidity: function proposalaccepterEntry(_addressofproposalacceptor address) returns()
func (_Permissions *PermissionsSession) ProposalaccepterEntry(_addressofproposalacceptor common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.ProposalaccepterEntry(&_Permissions.TransactOpts, _addressofproposalacceptor)
}

// ProposalaccepterEntry is a paid mutator transaction binding the contract method 0x274b7767.
//
// Solidity: function proposalaccepterEntry(_addressofproposalacceptor address) returns()
func (_Permissions *PermissionsTransactorSession) ProposalaccepterEntry(_addressofproposalacceptor common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.ProposalaccepterEntry(&_Permissions.TransactOpts, _addressofproposalacceptor)
}

// RemoveOperatorAccount is a paid mutator transaction binding the contract method 0xad1999ac.
//
// Solidity: function removeOperatorAccount(account_of_operator address) returns(status bool)
func (_Permissions *PermissionsTransactor) RemoveOperatorAccount(opts *bind.TransactOpts, account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "removeOperatorAccount", account_of_operator)
}

// RemoveOperatorAccount is a paid mutator transaction binding the contract method 0xad1999ac.
//
// Solidity: function removeOperatorAccount(account_of_operator address) returns(status bool)
func (_Permissions *PermissionsSession) RemoveOperatorAccount(account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.RemoveOperatorAccount(&_Permissions.TransactOpts, account_of_operator)
}

// RemoveOperatorAccount is a paid mutator transaction binding the contract method 0xad1999ac.
//
// Solidity: function removeOperatorAccount(account_of_operator address) returns(status bool)
func (_Permissions *PermissionsTransactorSession) RemoveOperatorAccount(account_of_operator common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.RemoveOperatorAccount(&_Permissions.TransactOpts, account_of_operator)
}

// ResetOperator is a paid mutator transaction binding the contract method 0xebdb432f.
//
// Solidity: function resetOperator() returns()
func (_Permissions *PermissionsTransactor) ResetOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "resetOperator")
}

// ResetOperator is a paid mutator transaction binding the contract method 0xebdb432f.
//
// Solidity: function resetOperator() returns()
func (_Permissions *PermissionsSession) ResetOperator() (*types.Transaction, error) {
	return _Permissions.Contract.ResetOperator(&_Permissions.TransactOpts)
}

// ResetOperator is a paid mutator transaction binding the contract method 0xebdb432f.
//
// Solidity: function resetOperator() returns()
func (_Permissions *PermissionsTransactorSession) ResetOperator() (*types.Transaction, error) {
	return _Permissions.Contract.ResetOperator(&_Permissions.TransactOpts)
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

// SetOperatorConsensus is a paid mutator transaction binding the contract method 0xdb788238.
//
// Solidity: function setOperatorConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsTransactor) SetOperatorConsensus(opts *bind.TransactOpts, _percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "setOperatorConsensus", _percentage)
}

// SetOperatorConsensus is a paid mutator transaction binding the contract method 0xdb788238.
//
// Solidity: function setOperatorConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsSession) SetOperatorConsensus(_percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.SetOperatorConsensus(&_Permissions.TransactOpts, _percentage)
}

// SetOperatorConsensus is a paid mutator transaction binding the contract method 0xdb788238.
//
// Solidity: function setOperatorConsensus(_percentage uint256) returns()
func (_Permissions *PermissionsTransactorSession) SetOperatorConsensus(_percentage *big.Int) (*types.Transaction, error) {
	return _Permissions.Contract.SetOperatorConsensus(&_Permissions.TransactOpts, _percentage)
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

// PermissionsLogOfActiveOperatorAccountIterator is returned from FilterLogOfActiveOperatorAccount and is used to iterate over the raw logs and unpacked data for LogOfActiveOperatorAccount events raised by the Permissions contract.
type PermissionsLogOfActiveOperatorAccountIterator struct {
	Event *PermissionsLogOfActiveOperatorAccount // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfActiveOperatorAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfActiveOperatorAccount)
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
		it.Event = new(PermissionsLogOfActiveOperatorAccount)
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
func (it *PermissionsLogOfActiveOperatorAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfActiveOperatorAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfActiveOperatorAccount represents a LogOfActiveOperatorAccount event raised by the Permissions contract.
type PermissionsLogOfActiveOperatorAccount struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfActiveOperatorAccount is a free log retrieval operation binding the contract event 0xd0cd818f9f2837a5e7dc53d100353f4ba5b023a116fb5608841996d4771277c1.
//
// Solidity: e LogOfActiveOperatorAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfActiveOperatorAccount(opts *bind.FilterOpts) (*PermissionsLogOfActiveOperatorAccountIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfActiveOperatorAccount")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfActiveOperatorAccountIterator{contract: _Permissions.contract, event: "LogOfActiveOperatorAccount", logs: logs, sub: sub}, nil
}

// WatchLogOfActiveOperatorAccount is a free log subscription operation binding the contract event 0xd0cd818f9f2837a5e7dc53d100353f4ba5b023a116fb5608841996d4771277c1.
//
// Solidity: e LogOfActiveOperatorAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfActiveOperatorAccount(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfActiveOperatorAccount) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfActiveOperatorAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfActiveOperatorAccount)
				if err := _Permissions.contract.UnpackLog(event, "LogOfActiveOperatorAccount", log); err != nil {
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

// PermissionsLogOfActiveVoterAccountIterator is returned from FilterLogOfActiveVoterAccount and is used to iterate over the raw logs and unpacked data for LogOfActiveVoterAccount events raised by the Permissions contract.
type PermissionsLogOfActiveVoterAccountIterator struct {
	Event *PermissionsLogOfActiveVoterAccount // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfActiveVoterAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfActiveVoterAccount)
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
		it.Event = new(PermissionsLogOfActiveVoterAccount)
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
func (it *PermissionsLogOfActiveVoterAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfActiveVoterAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfActiveVoterAccount represents a LogOfActiveVoterAccount event raised by the Permissions contract.
type PermissionsLogOfActiveVoterAccount struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfActiveVoterAccount is a free log retrieval operation binding the contract event 0x45b6638126291e2dab4e36d975252dcb55bc211569e9f1169309199d2bbe74e7.
//
// Solidity: e LogOfActiveVoterAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfActiveVoterAccount(opts *bind.FilterOpts) (*PermissionsLogOfActiveVoterAccountIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfActiveVoterAccount")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfActiveVoterAccountIterator{contract: _Permissions.contract, event: "LogOfActiveVoterAccount", logs: logs, sub: sub}, nil
}

// WatchLogOfActiveVoterAccount is a free log subscription operation binding the contract event 0x45b6638126291e2dab4e36d975252dcb55bc211569e9f1169309199d2bbe74e7.
//
// Solidity: e LogOfActiveVoterAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfActiveVoterAccount(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfActiveVoterAccount) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfActiveVoterAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfActiveVoterAccount)
				if err := _Permissions.contract.UnpackLog(event, "LogOfActiveVoterAccount", log); err != nil {
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

// PermissionsLogOfAddressOfProposalAcceptorEntryIterator is returned from FilterLogOfAddressOfProposalAcceptorEntry and is used to iterate over the raw logs and unpacked data for LogOfAddressOfProposalAcceptorEntry events raised by the Permissions contract.
type PermissionsLogOfAddressOfProposalAcceptorEntryIterator struct {
	Event *PermissionsLogOfAddressOfProposalAcceptorEntry // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfAddressOfProposalAcceptorEntryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfAddressOfProposalAcceptorEntry)
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
		it.Event = new(PermissionsLogOfAddressOfProposalAcceptorEntry)
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
func (it *PermissionsLogOfAddressOfProposalAcceptorEntryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfAddressOfProposalAcceptorEntryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfAddressOfProposalAcceptorEntry represents a LogOfAddressOfProposalAcceptorEntry event raised by the Permissions contract.
type PermissionsLogOfAddressOfProposalAcceptorEntry struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfAddressOfProposalAcceptorEntry is a free log retrieval operation binding the contract event 0xdb322fa900d5273c470b8847f6a0a42231e31959eba17a6066d5680425255c89.
//
// Solidity: e LogOfAddressOfProposalAcceptorEntry(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfAddressOfProposalAcceptorEntry(opts *bind.FilterOpts) (*PermissionsLogOfAddressOfProposalAcceptorEntryIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfAddressOfProposalAcceptorEntry")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfAddressOfProposalAcceptorEntryIterator{contract: _Permissions.contract, event: "LogOfAddressOfProposalAcceptorEntry", logs: logs, sub: sub}, nil
}

// WatchLogOfAddressOfProposalAcceptorEntry is a free log subscription operation binding the contract event 0xdb322fa900d5273c470b8847f6a0a42231e31959eba17a6066d5680425255c89.
//
// Solidity: e LogOfAddressOfProposalAcceptorEntry(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfAddressOfProposalAcceptorEntry(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfAddressOfProposalAcceptorEntry) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfAddressOfProposalAcceptorEntry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfAddressOfProposalAcceptorEntry)
				if err := _Permissions.contract.UnpackLog(event, "LogOfAddressOfProposalAcceptorEntry", log); err != nil {
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

// PermissionsLogOfOperatorSetConsensusIterator is returned from FilterLogOfOperatorSetConsensus and is used to iterate over the raw logs and unpacked data for LogOfOperatorSetConsensus events raised by the Permissions contract.
type PermissionsLogOfOperatorSetConsensusIterator struct {
	Event *PermissionsLogOfOperatorSetConsensus // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfOperatorSetConsensusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfOperatorSetConsensus)
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
		it.Event = new(PermissionsLogOfOperatorSetConsensus)
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
func (it *PermissionsLogOfOperatorSetConsensusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfOperatorSetConsensusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfOperatorSetConsensus represents a LogOfOperatorSetConsensus event raised by the Permissions contract.
type PermissionsLogOfOperatorSetConsensus struct {
	Consensuslimit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfOperatorSetConsensus is a free log retrieval operation binding the contract event 0xae848b724d627813f2afaa324ae4610e1503087865f73d5d57d2a90f605c17d5.
//
// Solidity: e LogOfOperatorSetConsensus(consensuslimit uint256)
func (_Permissions *PermissionsFilterer) FilterLogOfOperatorSetConsensus(opts *bind.FilterOpts) (*PermissionsLogOfOperatorSetConsensusIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfOperatorSetConsensus")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfOperatorSetConsensusIterator{contract: _Permissions.contract, event: "LogOfOperatorSetConsensus", logs: logs, sub: sub}, nil
}

// WatchLogOfOperatorSetConsensus is a free log subscription operation binding the contract event 0xae848b724d627813f2afaa324ae4610e1503087865f73d5d57d2a90f605c17d5.
//
// Solidity: e LogOfOperatorSetConsensus(consensuslimit uint256)
func (_Permissions *PermissionsFilterer) WatchLogOfOperatorSetConsensus(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfOperatorSetConsensus) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfOperatorSetConsensus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfOperatorSetConsensus)
				if err := _Permissions.contract.UnpackLog(event, "LogOfOperatorSetConsensus", log); err != nil {
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

// PermissionsLogOfRemoveOperatorAccountIterator is returned from FilterLogOfRemoveOperatorAccount and is used to iterate over the raw logs and unpacked data for LogOfRemoveOperatorAccount events raised by the Permissions contract.
type PermissionsLogOfRemoveOperatorAccountIterator struct {
	Event *PermissionsLogOfRemoveOperatorAccount // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfRemoveOperatorAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfRemoveOperatorAccount)
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
		it.Event = new(PermissionsLogOfRemoveOperatorAccount)
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
func (it *PermissionsLogOfRemoveOperatorAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfRemoveOperatorAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfRemoveOperatorAccount represents a LogOfRemoveOperatorAccount event raised by the Permissions contract.
type PermissionsLogOfRemoveOperatorAccount struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfRemoveOperatorAccount is a free log retrieval operation binding the contract event 0xda7af3b29c038afb615f78757e855b5c9131654c06a8a6f0ba33b7b727215465.
//
// Solidity: e LogOfRemoveOperatorAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfRemoveOperatorAccount(opts *bind.FilterOpts) (*PermissionsLogOfRemoveOperatorAccountIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfRemoveOperatorAccount")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfRemoveOperatorAccountIterator{contract: _Permissions.contract, event: "LogOfRemoveOperatorAccount", logs: logs, sub: sub}, nil
}

// WatchLogOfRemoveOperatorAccount is a free log subscription operation binding the contract event 0xda7af3b29c038afb615f78757e855b5c9131654c06a8a6f0ba33b7b727215465.
//
// Solidity: e LogOfRemoveOperatorAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfRemoveOperatorAccount(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfRemoveOperatorAccount) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfRemoveOperatorAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfRemoveOperatorAccount)
				if err := _Permissions.contract.UnpackLog(event, "LogOfRemoveOperatorAccount", log); err != nil {
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

// PermissionsLogOfRemoveVoterAccountIterator is returned from FilterLogOfRemoveVoterAccount and is used to iterate over the raw logs and unpacked data for LogOfRemoveVoterAccount events raised by the Permissions contract.
type PermissionsLogOfRemoveVoterAccountIterator struct {
	Event *PermissionsLogOfRemoveVoterAccount // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfRemoveVoterAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfRemoveVoterAccount)
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
		it.Event = new(PermissionsLogOfRemoveVoterAccount)
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
func (it *PermissionsLogOfRemoveVoterAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfRemoveVoterAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfRemoveVoterAccount represents a LogOfRemoveVoterAccount event raised by the Permissions contract.
type PermissionsLogOfRemoveVoterAccount struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfRemoveVoterAccount is a free log retrieval operation binding the contract event 0xceac34d6c239fde6d920bdf4ab89260499033335e865e17d611bf081f33d65fe.
//
// Solidity: e LogOfRemoveVoterAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfRemoveVoterAccount(opts *bind.FilterOpts) (*PermissionsLogOfRemoveVoterAccountIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfRemoveVoterAccount")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfRemoveVoterAccountIterator{contract: _Permissions.contract, event: "LogOfRemoveVoterAccount", logs: logs, sub: sub}, nil
}

// WatchLogOfRemoveVoterAccount is a free log subscription operation binding the contract event 0xceac34d6c239fde6d920bdf4ab89260499033335e865e17d611bf081f33d65fe.
//
// Solidity: e LogOfRemoveVoterAccount(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfRemoveVoterAccount(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfRemoveVoterAccount) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfRemoveVoterAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfRemoveVoterAccount)
				if err := _Permissions.contract.UnpackLog(event, "LogOfRemoveVoterAccount", log); err != nil {
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

// PermissionsLogOfResetOperatorIterator is returned from FilterLogOfResetOperator and is used to iterate over the raw logs and unpacked data for LogOfResetOperator events raised by the Permissions contract.
type PermissionsLogOfResetOperatorIterator struct {
	Event *PermissionsLogOfResetOperator // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfResetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfResetOperator)
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
		it.Event = new(PermissionsLogOfResetOperator)
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
func (it *PermissionsLogOfResetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfResetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfResetOperator represents a LogOfResetOperator event raised by the Permissions contract.
type PermissionsLogOfResetOperator struct {
	Accountaddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfResetOperator is a free log retrieval operation binding the contract event 0x8724988bd627c5b9b96f53247d1031353bf31098365566104e129627052ced98.
//
// Solidity: e LogOfResetOperator(accountaddress address)
func (_Permissions *PermissionsFilterer) FilterLogOfResetOperator(opts *bind.FilterOpts) (*PermissionsLogOfResetOperatorIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfResetOperator")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfResetOperatorIterator{contract: _Permissions.contract, event: "LogOfResetOperator", logs: logs, sub: sub}, nil
}

// WatchLogOfResetOperator is a free log subscription operation binding the contract event 0x8724988bd627c5b9b96f53247d1031353bf31098365566104e129627052ced98.
//
// Solidity: e LogOfResetOperator(accountaddress address)
func (_Permissions *PermissionsFilterer) WatchLogOfResetOperator(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfResetOperator) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfResetOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfResetOperator)
				if err := _Permissions.contract.UnpackLog(event, "LogOfResetOperator", log); err != nil {
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

// PermissionsLogOfSuspendNodeProposalIterator is returned from FilterLogOfSuspendNodeProposal and is used to iterate over the raw logs and unpacked data for LogOfSuspendNodeProposal events raised by the Permissions contract.
type PermissionsLogOfSuspendNodeProposalIterator struct {
	Event *PermissionsLogOfSuspendNodeProposal // Event containing the contract specifics and raw log

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
func (it *PermissionsLogOfSuspendNodeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsLogOfSuspendNodeProposal)
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
		it.Event = new(PermissionsLogOfSuspendNodeProposal)
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
func (it *PermissionsLogOfSuspendNodeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsLogOfSuspendNodeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsLogOfSuspendNodeProposal represents a LogOfSuspendNodeProposal event raised by the Permissions contract.
type PermissionsLogOfSuspendNodeProposal struct {
	Id             *big.Int
	Accountaddress common.Address
	Enodeaddress   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogOfSuspendNodeProposal is a free log retrieval operation binding the contract event 0x48dfbd110e878cbf6e08b7bb004151d52e571f0c464de79f3a07fee91741afe3.
//
// Solidity: e LogOfSuspendNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) FilterLogOfSuspendNodeProposal(opts *bind.FilterOpts) (*PermissionsLogOfSuspendNodeProposalIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "LogOfSuspendNodeProposal")
	if err != nil {
		return nil, err
	}
	return &PermissionsLogOfSuspendNodeProposalIterator{contract: _Permissions.contract, event: "LogOfSuspendNodeProposal", logs: logs, sub: sub}, nil
}

// WatchLogOfSuspendNodeProposal is a free log subscription operation binding the contract event 0x48dfbd110e878cbf6e08b7bb004151d52e571f0c464de79f3a07fee91741afe3.
//
// Solidity: e LogOfSuspendNodeProposal(id uint256, accountaddress address, enodeaddress bytes32)
func (_Permissions *PermissionsFilterer) WatchLogOfSuspendNodeProposal(opts *bind.WatchOpts, sink chan<- *PermissionsLogOfSuspendNodeProposal) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "LogOfSuspendNodeProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsLogOfSuspendNodeProposal)
				if err := _Permissions.contract.UnpackLog(event, "LogOfSuspendNodeProposal", log); err != nil {
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
