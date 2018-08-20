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
var smartContractAddress string ="0x5df78fc6c7e1a110d1be47d70b4cf570cc6f4e26"
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
