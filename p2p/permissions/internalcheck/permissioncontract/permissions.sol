pragma solidity ^0.4.22;

contract Permissions{
    /*
    *Node contains details of nodes and its enode address
    */
    struct Node{
        //enode is e-node address of the node
        bytes32     enode;
        //account is account associated with node
        address     account;
        //flag represent the node is started to vote or completed
        //flag also represent that this node is already exist in the network
        //flag do not identify this node is currently in peer.
        bool        flag;
        //votecount counts total number of accepted vote for each node
        uint        votecount;
        //rejectcount counts total number if rejected vote for each node
        uint        rejectcount;
    }

    //issuspention for suspention phase is running
    bool public issuspention;

     //isadding for adding phase is running
    bool public isadding;

    //addingmutex prevent form  running multiple adding request
    //if addingmutex == false then we can start new job
    bool public addingmutex;

    //suspentionmutex prevent form  running multiple suspention request
    //if suspentionmutex == false then we can start new job
    bool public suspentionmutex;

    //LimitOfVote is total number of vote need to approve the node
    uint public LimitOfVote;
    uint public LimitOfNegVote;

    //previous holds the previous node information
    //it tracks that is the previous information matches the current transaction
    //this helps to not doing multiple adding and suspend work before successfull completion of job
    bytes32 public previousenode;
    address public previousaccount;

    //NodeCount is total number of approved node in the network
    uint public NodeCount;

    //percentage is the consensus to meet (like 50%)
    uint public consensuspercentage;
    
   //nodeconformations conforms the node to elisibconsensuspercentagele to connect to the network
   
    mapping(
        bytes32 => bool
            )public nodeconformations;

    mapping( 
    address => mapping(
    bytes32 => Node
        )
    )public nodeinfo ;
        
    event LogOfAddNode(uint id, address accountaddress, bytes32 enodeaddress, string adding );
    event LogOfSuspentionNode(uint id, address accountaddress, bytes32 enodeaddress, string suspend);
    event LogOfSetConsensus(uint consensuslimit);
    event LogOfResetProcess(address previousaccount, bytes32 previousenode);
    event LogOfaddingConsensusMeet(address accountaddress, bytes32 enodeaddress, string flag);
    event LogOfsuspentionConsensusMeet(address accountaddress, bytes32 enodeaddress, string flag);
    event LogOfAddingVote(uint id, address accountaddress, bytes32 enodeaddress, string flag, bool vote);
    event LogOfSuspentionVote(uint id, address accountaddress, bytes32 enodeaddress, string flag, bool vote);
    event LogOfVoteReject(uint id, address accountaddress, bytes32 enodeaddress, uint rejectcount);
    
    
    
    constructor()
        public{
            LimitOfVote = 0;
            LimitOfNegVote = 0;
            NodeCount = 0;
        }
 
     //addNode function enters the enode and account of the proposed node.
    //the node will be eligible to peer with other node when it meets the consensus
    //untill reach to consensus node will be proposed node. If meets the consensus then
    //it will be approved node. it will be signified by the nodeconformations[enode_of_proposed_node]
    function addNode(uint _id, address _account, bytes32 _enode)
        public{

            if((addingmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _addNode(_id, _account, _enode);
            }
            else if((addingmutex == false) && (isadding == false)){
                addingmutex = true;
                _addNode(_id, _account,_enode);
            }
    }

    //suspendNode will disable the nodeconformations flag (nodeconformations = false)
    //while checking in the phase of handshake it will check the nodeconformations status
    function suspendNode(uint _id, address _account, bytes32 _enode)
        public{
             if((suspentionmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _suspendNode(_id, _account, _enode);
            }
            else if((suspentionmutex == false) && (issuspention == false)){
                suspentionmutex = true;
                _suspendNode(_id, _account,_enode);
            }
    }


//addnode operation for voteing purpose 
//else statement is designed in such a way that it cannot be execute.
function addingVote(uint _id, address _account, bytes32 _enode)
    public{
         if((addingmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _addNode(_id, _account, _enode);
            }else {
                 require((addingmutex == true) && (previousenode == _enode) && (previousaccount == _account));
            }   
        //LogOfAddingVote indicates that, vote is done
         emit LogOfAddingVote(_id, _account, _enode, "adding", true);
}







//suspendVote operation for voteing purpose 
//else statement is designed in such a way that it cannot be execute.
function suspendVote(uint _id, address _account, bytes32 _enode)
    public{
        if((suspentionmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _suspendNode(_id, _account, _enode);
        }else{
            require((suspentionmutex == true) && (previousenode == _enode) && (previousaccount == _account));
        }        
        emit LogOfSuspentionVote(_id, _account, _enode, "suspention", true);
    }

  //checkNode checks the seeking node is eligible to peer with existing network
    //it will be called by api1--> api2
    function checkNode(bytes32 _enode)
        public
        view
        returns(bool){
            if(nodeconformations[_enode] == true){
              return true;
            }else {
            return false;
            }
    }


    function _addNode(uint _id, address _account, bytes32 _enode)
        private {

            assert(!nodeconformations[_enode]);

            assert(!issuspention);

            assert(addingmutex);

             isadding = true;

            if(nodeinfo[_account][_enode].votecount < LimitOfVote){
            nodeinfo[_account][_enode].enode = _enode;
            nodeinfo[_account][_enode].account = _account;
            nodeinfo[_account][_enode].flag = true;
            nodeinfo[_account][_enode].votecount += 1;
            }
            if(nodeinfo[_account][_enode].votecount == LimitOfVote){
                nodeinfo[_account][_enode].enode = _enode;
                nodeinfo[_account][_enode].account = _account;
                nodeinfo[_account][_enode].flag = true;

            //  addingmutex = false;
            //  now the adding phase is completed
            //  isadding = false;

            //nodeconformations is used for check the node from permissions layers in core chain
                nodeconformations[_enode] = true;

            //resetProcess() resets the flags used in operation 
                resetProcess();

            //NodeCount counts the number of node that are verified by the network.
            //this will be incremented only if the consensus is reached
                NodeCount++;

            //now the node is accepted and hence count is set to 0.
            //if vote count is 0 then we will easily do the process of suspend node
            //    nodeinfo[_account][_enode].votecount = 0;
              
                LimitOfVote = acceptConsensus();
                LimitOfNegVote = rejectConsensus();
                
                emit LogOfaddingConsensusMeet(_account, _enode, "adding");
            }
            previousenode = _enode;
            previousaccount = _account;
         emit LogOfAddNode(_id, _account, _enode, "addition");

    }

    function _suspendNode(uint _id, address _account, bytes32 _enode)
        private{
            //checks if adding is running
            assert(!isadding);

            //checks suspention is running
            assert(suspentionmutex);

           //still flag is enabled and cannot be set to false.
           //only nodeconformations can decide it  is removed or not.
           assert(nodeinfo[_account][_enode].flag == true);

           //checks if node is currently on workable state
           assert(nodeconformations[_enode] == true);

            //if suspention is running then make suspention flag to true
            issuspention = true;

            assert(nodeinfo[_account][_enode].votecount < LimitOfVote);
            nodeinfo[_account][_enode].enode = _enode;
            nodeinfo[_account][_enode].account = _account;
            nodeinfo[_account][_enode].flag = true;
            nodeinfo[_account][_enode].votecount += 1;

            //waiting for final vote and if final vote is equal to limit of vote then changes the states of node count and LimitOfVote and related flags
            if(nodeinfo[_account][_enode].votecount == LimitOfVote){

                //set nodeconformations to false and it inticates this node is disabled.
                nodeconformations[_enode] = false;
            //  issuspention = false;
            //  suspentionmutex = false;
            //  votecount is set to zero so we can further proceed for addition again if suspended
            //  nodeinfo[_account][_enode].votecount = 0;
              
            //resetProcess() resets the flags used in operation 
                resetProcess();
                NodeCount--;
                //LimitOfVote = NodeCount;
                //Dynamic consensus set by the user 
                LimitOfVote = acceptConsensus();
                LimitOfNegVote = rejectConsensus();
            emit LogOfsuspentionConsensusMeet(_account, _enode, "suspension");
            }
        previousenode = _enode;
        previousaccount = _account;
        emit LogOfSuspentionNode(_id, _account, _enode, "suspention");
    }
    
    
    //Negative vote calculation is required for resetting the process of due to comming deadlock and ensure that the consensus will not reach.
    //voteReject funciton is called when reject the vote
    //This is equally valid for both adding and suspention.
    function voteReject(uint _id, address _account, bytes32 _enode)
        public{
            
            
        if(((suspentionmutex == true) || (addingmutex == true)) && (previousenode == _enode) && (previousaccount == _account)){
        
        nodeinfo[_account][_enode].rejectcount += 1;
        emit LogOfVoteReject(_id, _account, _enode, nodeinfo[_account][_enode].rejectcount);
         
         if( nodeinfo[_account][_enode].rejectcount >= rejectConsensus() ) {
              resetProcess();
             
             }
           
        }
        
    }
   
   //funciton acceptConsensus calculate the total number if vote required to accept the node 
    function acceptConsensus()
        private 
        view
        returns(uint limit){
            limit = ((consensuspercentage*NodeCount/100)+1);
           //checking overflow
            if (limit >= NodeCount){
                return NodeCount;
            }else {
                return limit;
            }
    }
    
    
    //Total consensus required to reject the node conformation.
    function rejectConsensus()
    private
    view
    returns(uint limit){
        limit = NodeCount - LimitOfVote + 1;
            //checking overflow
            if (limit >= NodeCount){
                return NodeCount;
            }else {
                return limit;
            }

    }
    
    //setConsensus sets the consensus percentage 
    function setConsensus(uint _percentage)
        public {
            //checking range of percentage , 0 <= _percentage => 100 
            require((_percentage >= 0) && (_percentage <= 100 ));
            consensuspercentage = _percentage;    
            emit LogOfSetConsensus(consensuspercentage);
    }

    //resetProcess is used to avoid from deadlock on current process 
    function resetProcess()
        public{
            issuspention = false;
            isadding = false;
            addingmutex = false;
            suspentionmutex = false;
            //vote count to zero
            nodeinfo[previousaccount][previousenode].votecount = 0;
            nodeinfo[previousaccount][previousenode].rejectcount = 0;
            
            emit LogOfResetProcess(previousaccount, previousenode);
    }
}

