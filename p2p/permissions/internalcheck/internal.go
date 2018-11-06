package internalcheck
import	(

	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p/permissions/internalcheck/permissioncontract"
)

// var smartContractAddress string ="0x6a44e2c99f909a04c4d08a08279a237e2ad6ee64"
//var smartContractAddress string ="0x9a292d944522596880fa6adcfcf97db3d20ac6a2"
//var smartContractAddress string ="0x5df78fc6c7e1a110d1be47d70b4cf570cc6f4e26"
//var smartContractAddress string ="0xf4b06b25e11407549672ab2a52c863e932a55998"
//var smartContractAddress string ="0x2c07cab2330f325bcdfe02edf1add65eb9c0a846"
//var smartContractAddress string ="0xbfa6d19be3f8695798b492cfa36103e44a85d5b6"
//var smartContractAddress string ="0x55e0040c3f29533b811367975d3b34a214d45504"
//var smartContractAddress string ="0x34e199daf1887985bf5767da7e99a34c5781c48e"
var smartContractAddress string ="0xee1d05bf36147425cfb38d777202263ceaff70cb"
var cc string

func IsNodePermissioned(nodename string, currentNode string, direction string) bool {
	cc=currentNode
	return(isNodePermissioned(nodename, currentNode, direction))
}

// check if a given node is permissioned to connect to the change
func isNodePermissioned(nodename string, currentNode string, direction string) bool {

	//fmt.Println("Node , Passed for checking")
	if(checkInLedger(nodename,direction)) {

		//Ready for handshake
		log.Debug("isNodePermissioned")
		return true

	}else{
		log.Debug("isNodePermissioned, ledger empty")

		return false
	}
}

func checkInLedger(nodename string ,direction string) bool {
	//	client,	err	:=	ethclient.Dial("http://18.136.104.41:8501")
	client,	err	:=	ethclient.Dial(direction)
	if	err	!=	nil	{
		fmt.Println("error in checkledger : ethclient.dial")
	}

	//address is contract
	address:= common.HexToAddress(smartContractAddress)
	instance,	err	:= permissioncontract.NewPermissions(address,	client)
	if	err	!=	nil	{
		log.Debug("checkInLedger, error on getting instance")
	}
	key	:=	[32]byte{}

	copy(key[:], []byte(nodename))

	result,	err	:=	instance.Nodeconformations(nil,key)
	/*fmt.Println("******")
	fmt.Println(cc)
	fmt.Println(nodename)
	fmt.Println(result)
	fmt.Println("******")*/
	if	err	!=	nil	{
		fmt.Println(err)
		log.Debug("checkInLedger, error on getting result from ledger")
	}
	return result
}
