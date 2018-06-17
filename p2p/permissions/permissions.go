package permissions
import (
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/log"
	"io/ioutil"
)

const (
	NODE_NAME_LENGTH    = 32
)

//PermissionNode is for running node
type PermissionNode struct {
	Enode string `json:"enode,omitempty"`
}
//Response holds the response from ledger
type Response struct {
	Permissionresponse string `json:"enode,omitempty"`
}

var results string

var responsevalue bool

//seeking node is incommingnodename
var incommingnodename string

//sendig enode to next api to get the response is the ledger contain the ledger
//it is get method
func send(w http.ResponseWriter, req *http.Request) {
	data := incommingnodename
	json.NewEncoder(w).Encode(&PermissionNode{data})
}


func receive(w http.ResponseWriter, r *http.Request) {

	body,err := ioutil.ReadAll(r.Body)

	response := &Response{}

	if err != nil {
		log.Error("Failed to access nodes", "err", err)
	}

	if err := json.Unmarshal(body, response); err != nil {
		log.Error("Failed to load nodes", "err", err)
	}

	b,err := strconv.ParseBool(response.Permissionresponse)

  	responsevalue = b
  	fmt.Println("responsevalue is :- ")
	fmt.Println(responsevalue)
}


func IsNodePermissioned(nodename string, currentNode string, direction string) bool {
	return(isNodePermissioned(nodename, currentNode, direction))
}

// check if a given node is permissioned to connect to the change
func isNodePermissioned(nodename string, currentNode string, direction string) bool {
incommingnodename = nodename
/*  To-Do :-
	//to make permissioned with 3rd proposal (Permissioned with contract) replace datadir with api endpoint.

	//When we connect with web3js we have to check whether the rpc flag enabled node is present

*/
	if(checkInLedger(nodename)) {

	//Ready for handshake
		log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:NODE_NAME_LENGTH])
		return true

	}else{

	return false
	}
}

//check whether enode is present in the ledger
func checkInLedger(enode_of_seeking_node string) bool{
	router := mux.NewRouter()

	router.HandleFunc("/nodedetails", send).Methods("GET")

	router.HandleFunc("/noderesponse", receive).Methods("POST")

	/*
		//ToDo : closing port after  timeout

		srv := &http.Server{
			Handler:      router,
			Addr:         "127.0.0.1:8000",
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Fatal(srv.ListenAndServe())


	*/

	muxWithMiddlewares := http.TimeoutHandler(router, time.Second*10,"Timeout!")

	http.ListenAndServe(":8083", muxWithMiddlewares)

	if(responsevalue == true){
		return true
	}else {
		return false
	}
}




func parsePermissionedNodes() []*discover.Node {

	var nodes []*discover.Node
	var	permissionnode = PermissionNode{}

 	discoverednode := permissionnode.Enode
		if discoverednode == "" {
			log.Error("parsePermissionedNodes: Node URL blank")
		}
		node, err := discover.ParseNode(discoverednode)

		if err != nil {
			log.Error("parsePermissionedNodes: Node URL", "url", discoverednode, "err", err)
		}
		nodes = append(nodes, node)

	return nodes
}

