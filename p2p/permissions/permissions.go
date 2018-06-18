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
	"context"
	"sync/atomic"
	"syscall"
	"os"
	"os/signal"
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


type myServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

//sendig enode to next api to get the response is the ledger contain the ledger
func send(w http.ResponseWriter, req *http.Request) {
	data := incommingnodename
	json.NewEncoder(w).Encode(&PermissionNode{data})
}


func (s *myServer)receive(w http.ResponseWriter, r *http.Request) {
	body,err := ioutil.ReadAll(r.Body)
	response := &Response{}
	if err != nil {
		log.Error("receive : faile to access nodes", "err", err)
	}
	if err := json.Unmarshal(body, response); err != nil {
		log.Error("receive : failed to load nodes", "err", err)
	}

	b,err := strconv.ParseBool(response.Permissionresponse)
  	responsevalue = b
	fmt.Println(responsevalue)

  	//newlyadded
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		fmt.Println("Shutdown through API call")
		return
	}
	go func() {
		s.shutdownReq <- true
	}()
}

func IsNodePermissioned(nodename string, currentNode string, direction string) bool {
	return(isNodePermissioned(nodename, currentNode, direction))
}

// check if a given node is permissioned to connect to the change
func isNodePermissioned(nodename string, currentNode string, direction string) bool {

	incommingnodename = nodename
	if(checkInLedger(nodename)) {

	//Ready for handshake
		log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:NODE_NAME_LENGTH])
		return true

	}else{

	return false
	}
}

//check whether enode is present in the ledger
func checkInLedger(enode_of_seeking_node string) bool {
	server := NewServer()
	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Error("checkLedger: ListenAndServer", "err", err)
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown()

	<-done

	if(responsevalue == true){
		return true
	}else {
		return false
	}
}

func NewServer() *myServer {
	//create server
	s := &myServer{
		Server: http.Server{
			Addr:         ":8078",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}

	router := mux.NewRouter()

	router.HandleFunc("/nodedetails", send).Methods("GET")

	router.HandleFunc("/noderesponse", s.receive).Methods("POST")

	//set http server handler
	s.Handler = router

	return s
}

func (s *myServer) WaitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		fmt.Println(sig)
		case sig := <-s.shutdownReq:

		fmt.Println(sig)
	}

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Error("WaitShutdiwn:", "err", err)
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

