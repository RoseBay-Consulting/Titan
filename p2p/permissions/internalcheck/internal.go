package internalcheck
import	(

"fmt"
"github.com/ethereum/go-ethereum/log"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/ethclient"
"github.com/ethereum/go-ethereum/p2p/permissions/internalcheck/permissioncontract"
)



func IsNodePermissioned(nodename string, currentNode string, direction string) bool {
	return(isNodePermissioned(nodename, currentNode, direction))
}

// check if a given node is permissioned to connect to the change
func isNodePermissioned(nodename string, currentNode string, direction string) bool {

	fmt.Println("nodename is ")
	fmt.Println(nodename)
	fmt.Println("currentnode is ")
	fmt.Println(currentNode)
	//fmt.Println("Node , Passed for checking")
	if(checkInLedger(nodename)) {

		//Ready for handshake
		log.Debug("isNodePermissioned")
			return true

	}else{
		log.Debug("isNodePermissioned, ledger empty")

		return false
	}
}

func checkInLedger(nodename string) bool {
	client,	err	:=	ethclient.Dial("http://192.168.88.20:8545")
	if	err	!=	nil	{
	fmt.Println("error in checkledger : ethclient.dial")
	}

	//address is contract
	address:= common.HexToAddress("0xc787a882f5f7037d02c6db8c155f379d48a745c6")
	instance,	err	:= permissioncontract.NewPermissions(address,	client)
	if	err	!=	nil	{
		log.Debug("checkInLedger, error on getting istance")
	}
	key	:=	[32]byte{}

	copy(key[:], []byte(nodename))

	result,	err	:=	instance.Nodeconformations(nil,key)
	if	err	!=	nil	{
		fmt.Println(err)
		log.Debug("checkInLedger, error on getting result from ledger")
	}
	fmt.Println(result)
	return result
}
