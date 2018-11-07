pragma solidity ^0.4.22;

contract Permissions{
    /**
    * Node contains details of nodes and its enode address
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
   
   //nodeconformations conforms the node to elisible to  connect to the network
   
   
    mapping(
        bytes32 => bool
            )public nodeconformations;
            
    //nodeinfo is fundamental ledger that holds all the permissions core details

    mapping( 
    address => mapping(
    bytes32 => Node
        )
    )public nodeinfo ;
    
    /**
    * Following ten events are used for tracking from the web3 library and interact with offchain.
    */
        
    event LogOfSetConsensus(uint consensuslimit);
    
    event LogOfResetProcess(address previousaccount, bytes32 previousenode);
    
    event LogOfAddNode(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfAddingConsensusMeet(address accountaddress, bytes32 enodeaddress);
    
    event LogOfAddingVote(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfsuspentionConsensusMeet(address accountaddress, bytes32 enodeaddress);
    
    event LogOfSuspentionNode(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfSuspentionVote(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfAddNodeProposal(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfSuspendNodeProposal(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfAddingAcceptProposal(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfSuspentionAcceptProposal(uint id, address accountaddress, bytes32 enodeaddress);
    
    event LogOfVoteReject(uint id, address accountaddress, bytes32 enodeaddress, uint rejectcount);
    
    event LogOfVoteRejectConsensusMeet(uint id, address accountaddress, bytes32 enodeaddress, bool isadding, bool issuspention);
    
    
    //newly added logs for storing and deleting the accounts in system 
    event LogOfRemoveVoterAccount(address accountaddress);
    
    event LogOfActiveVoterAccount(address accountaddress);
    
    event LogOfOperatorSetConsensus(uint consensuslimit);
    
    //operator operation on the system 

    event LogOfRemoveOperatorAccount(address accountaddress);
    
    event LogOfActiveOperatorAccount(address accountaddress);
    
    event LogOfResetOperator(address accountaddress);
    
    event LogOfSetOperatorCount(uint256 maxnumberofoperator);
    
    event LogOfAddressOfProposalAcceptorEntry(address accountaddress);

    
    
    /**
     * There are three actor actively interacting with the system.
     * 1. cobuna platform operator account : operator account used set the consensus and to propose the node.
     * 2. cobuna stakeholder account : cobuna statakeholder accounts interact with the system for vote the proposal. 
     * 3. accept or reject proposasl at the time of timout
    */
    
    modifier isoperator(){
          if(operators[msg.sender]){
            _;
        }else{
            revert();
        }
       
    }
    modifier isvoter(){
        if(voters[msg.sender]){
            _;
        }else{
            revert();
        }
       
    }
    
    
    modifier isproposalaccepter(){
        if(addressofproposalacceptor == msg.sender){
            _;
        }else {
            revert();
        }
        
        
    }
    
    
    constructor()
        public{
            LimitOfVote = 0;
            LimitOfNegVote = 0;
            NodeCount = 0;
        }
 


    /** 
    * @param _id is id that represent the current transaction and used for the offchain tracking.
    * @param _account is account address of the user
    * @param _enode is enode address of the  node 
    *    addNodeProposal function enters the enode and account of the proposed node.
    *    the node will be eligible to peer with other node when it meets the consensus
    *    untill reach to consensus node will be proposed node. If meets the consensus then
    *    it will be approved node. it will be signified by the nodeconformations[enode_of_proposed_node]
    */
 
    function addNodeProposal(uint _id, address _account, bytes32 _enode)
        public
       isoperator{
            require((suspentionmutex == false) && (addingmutex == false));
            require((isadding == false) && (issuspention == false));
            require(!nodeconformations[_enode]);
             isadding = true;
             addingmutex = true;
             
            nodeinfo[_account][_enode].enode = _enode;
            nodeinfo[_account][_enode].account = _account;
             previousenode = _enode;
            previousaccount = _account;
        emit LogOfAddNodeProposal(_id, _account, _enode);
    }


    /**
    * @param _id is id that represent the current transaction and used for the offchain tracking.
    * @param _account is account address of the user
    * @param _enode is enode address of the  node 
    *    suspendNodeProposal will disable the nodeconformations flag (nodeconformations = false)
    *    while checking in the phase of handshake it will check the nodeconformations status
    */
    
    function suspendNodeProposal(uint _id, address _account, bytes32 _enode)
        public
        isoperator{
            require((suspentionmutex == false) && (addingmutex == false));
            require((isadding == false) && (issuspention == false));
            require(nodeconformations[_enode]);
             issuspention = true;
             suspentionmutex = true;
             
            nodeinfo[_account][_enode].enode = _enode;
            nodeinfo[_account][_enode].account = _account;
             previousenode = _enode;
            previousaccount = _account;
        emit LogOfSuspendNodeProposal(_id, _account, _enode);
    }


    /**
    * @param _id is id that represent the current transaction and used for the offchain tracking.
    * @param _account is account address of the user
    * @param _enode is enode address of the  node 
    *
    *    addnode operation for voteing purpose 
    *    else statement is designed in such a way that it cannot be execute.
    */




    function addingVote(uint _id, address _account, bytes32 _enode)
        public
        isvoter{
         if((addingmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _addNode(_id, _account, _enode);
         
        //LogOfAddingVote indicates that, vote is done
         emit LogOfAddingVote(_id, _account, _enode);

        }
    }




    /**
    * @param _id is id that represent the current transaction and used for the offchain tracking.
    * @param _account is account address of the user
    * @param _enode is enode address of the  node 
    *    suspendVote operation for voteing purpose 
    *    else statement is designed in such a way that it cannot be execute.        
    */

    function suspendVote(uint _id, address _account, bytes32 _enode)
        public
        isvoter{
            if((suspentionmutex == true) && (previousenode == _enode) && (previousaccount == _account)){
                _suspendNode(_id, _account, _enode);
        emit LogOfSuspentionVote(_id, _account, _enode);
        }
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

            
            //nodeconformations is used for check the node from permissions layers in core chain
                nodeconformations[_enode] = true;
                
            //storing account after conformation
            activeVoterAccount(previousaccount);

            //resetProcess() resets the flags used in operation 
                resetProcess();

            //NodeCount counts the number of node that are verified by the network.
            //this will be incremented only if the consensus is reached
                NodeCount++;
            //Again calculating the limit of vote for accept and reject, this is due to when ever there is change in NodeCount then 
            //limit will change 
                LimitOfVote = acceptConsensus();
                LimitOfNegVote = rejectConsensus();
                
                emit LogOfAddingConsensusMeet(_account, _enode);
            }
            previousenode = _enode;
            previousaccount = _account;
         emit LogOfAddNode(_id, _account, _enode);

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
                
                //removing account from the system
                removeVoterAccount(previousaccount);
              
                //resetProcess() resets the flags used in operation 
                resetProcess();
                NodeCount--;
                //LimitOfVote = NodeCount;
                //Dynamic consensus set by the user 
                LimitOfVote = acceptConsensus();
                LimitOfNegVote = rejectConsensus();
            emit LogOfsuspentionConsensusMeet(_account, _enode);
            }
        previousenode = _enode;
        previousaccount = _account;
        emit LogOfSuspentionNode(_id, _account, _enode);
    }
    
    
    
    /**
    * @param _id is id that represent the current transaction and used for the offchain tracking.
    * @param _account is account address of the user
    * @param _enode is enode address of the  node 
    *    Negative vote calculation is required for resetting the process of due to comming deadlock and ensure that the consensus will not reach.
    *    voteReject funciton is called when reject the vote
    * This is equally valid for both adding and suspention.
    */
  
    function voteReject(uint _id, address _account, bytes32 _enode)
        public
        isvoter{
            
            
        if(((suspentionmutex == true) || (addingmutex == true)) && (previousenode == _enode) && (previousaccount == _account)){
        
        nodeinfo[_account][_enode].rejectcount += 1;
        emit LogOfVoteReject(_id, _account, _enode, nodeinfo[_account][_enode].rejectcount);
         
         if( nodeinfo[_account][_enode].rejectcount >= rejectConsensus() ) {
             emit LogOfVoteRejectConsensusMeet(_id, _account, _enode, isadding, issuspention);
              resetProcess();
                }
        }
        
    }
    
   
   /** 
   * acceptConsensus funciton  calculate the total number if vote required to accept the node 
   * @return limit for accept the node 
   */
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
    
   /** 
   * Total consensus required to reject the node conformation. 
   * @return limit for reject the node 
   */
    
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
    
    /**
    * setConsensus sets the consensus percentage
    * @param _percentage , vote required to accept the node.  
    */
     
    function setConsensus(uint _percentage)
        public 
        isoperator{
            //checking range of percentage , 0 <= _percentage => 100 
            require((_percentage >= 0) && (_percentage <= 100 ));
            consensuspercentage = _percentage;    
            emit LogOfSetConsensus(consensuspercentage);
    }

    /**
    * resetProcess is used to avoid from deadlock on current process 
    * This will be executed when we complete either rejection vote or accept vote for the node.
    */
    function resetProcess()
        private{
            issuspention = false;
            isadding = false;
            addingmutex = false;
            suspentionmutex = false;
            //vote count to zero
            nodeinfo[previousaccount][previousenode].votecount = 0;
            nodeinfo[previousaccount][previousenode].rejectcount = 0;
            
            emit LogOfResetProcess(previousaccount, previousenode);
    }


/**
 * acceptProposal() accepts the node if the time is out. this function is trigirred by EOA and check the time period offchain
 * @param _id is id that represent the current transaction and used for the offchain tracking.
*/

    function acceptProposal(uint _id)
        public
        isproposalaccepter{
        //checks if the process is running or not 
        assert(isadding || issuspention);
        if(isadding){
            
            //flag represent the particulat node alreay exist in the the network.
            nodeinfo[previousaccount][previousenode].flag = true;
            
            //nodeconformations is used for check the node from permissions layers in core chain
            nodeconformations[previousenode] = true;
            
            //storing account after conformation
            activeVoterAccount(previousaccount);

            //resetProcess() resets the flags used in operation 
            resetProcess();

            //NodeCount counts the number of node that are verified by the network.
            //this will be incremented only if the consensus is reached
            NodeCount++;
            //Again calculating the limit of vote for accept and reject, this is due to when ever there is change in NodeCount then 
            //limit will change 
            LimitOfVote = acceptConsensus();
            LimitOfNegVote = rejectConsensus();
            emit LogOfAddingAcceptProposal(_id, previousaccount, previousenode);
        }
        if(issuspention){
              
            //set nodeconformations to false and it inticates this node is disabled.
            nodeconformations[previousenode] = false;
            
            //removing account from the system
            removeVoterAccount(previousaccount); 
            
            //resetProcess() resets the flags used in operation 
            resetProcess();
            NodeCount--;
               
            //Again calculating the limit of vote for accept and reject, this is due to when ever there is change in NodeCount then 
            //limit will change 
            LimitOfVote = acceptConsensus();
            LimitOfNegVote = rejectConsensus();
        
            emit LogOfSuspentionAcceptProposal(_id, previousaccount, previousenode);
        }
    }
    
    
   /**
    * Below voters management 
   */ 
   
    
    //voterindexaccountpair map the index for particular account address
    mapping (address => uint256) public voterindexaccountpair;
  
     
   //voterindex holds the current number of voter in the system. This is used to sort the address so that we can iterate in minimal cost.
   uint public voterindex;
   
   //voters are voters for the proposal and this used in permissions for milisus entry
   mapping (address => bool) public voters;
   
   //voterlist is list of voters stored in the system.
   mapping (uint256 => address) voterlist;
   
    /**
     * storeAccount stores the account that comes to interact with the system. 
     * @param account_of_voter is account to store in the system
    */
    function activeVoterAccount(address account_of_voter) private {
        require(!voters[account_of_voter]);
        
        //we use voters[account_of_voter] while checking the condition whether satisfy the voters
        voters[account_of_voter] = true;
        
        //storing on voterlist
        voterlist[voterindex] = account_of_voter;
        voterindexaccountpair[account_of_voter] = voterindex;
        voterindex++;
        emit LogOfActiveVoterAccount(account_of_voter);
        
    }
    
    /**
     * removeVoterAccount remove acccount from the system and replace with account in the last element of the array
     * @param account_of_voter is account to remove fron the system.
    */
    
    function removeVoterAccount(address account_of_voter) private {
        require(voters[account_of_voter]);
        voters[account_of_voter] = false;
        
        //sorting account in such a way that just removed account is now replace with last of array
       
        //copping last element of array of voterlist to vaccent index
        //index-1 is last element of array as we increased index in activeVoterAccount()
        
        voterlist[voterindexaccountpair[account_of_voter]] = voterlist[voterindex-1];
        
        
        //decreasing list as we placed last element in vaccent place
        voterindex--;
        emit LogOfRemoveVoterAccount(account_of_voter);
      
    }
    
    /**
     * 
     * Specially for operator set, reset and change also storing and removing the address of operator 
    */
  
    //operatorindexaccountpair map the index for particular account address
    mapping (address => uint256) public operatorindexaccountpair;
  
     
    //operatorindex holds the current number of operator in the system. This is used to sort the address so that we can iterate in minimal cost.
    uint public operatorindex;
   
   //operators used for avoiding duplication
    mapping (address => bool) public operators;
   
    //operatorlist is list of voters stored in the system.
    mapping (uint256 => address) operatorlist;
   
    /**
     * activeOperatorAccount stores the account that comes to interact with the system. 
     * @param account_of_operator account which need to store in the system
    */
    
    function activeOperatorAccount(address account_of_operator) 
    private
    returns(bool){
        require(!operators[account_of_operator]);
        
        //we use operators[account_of_voter] while checking the condition whether satisfy the operators
        operators[account_of_operator] = true;
        
        //storing on operatorlist
        operatorlist[operatorindex] = account_of_operator;
        operatorindexaccountpair[account_of_operator] = operatorindex;
        operatorindex++;
        emit LogOfActiveOperatorAccount(account_of_operator);
        return true;
    }
    
    /**
     * removeOperatorAccount remove acccount from the system and replace with account in the last element of the array
     * @param account_of_operator is account to remove from the system. 
     * This contract function will only access by the operators.
     * @return status - true on success and false on failure
    */
    
    function removeOperatorAccount(address account_of_operator) 
    public 
    isoperator
    returns(bool status){
        require(operators[account_of_operator]);
        operators[account_of_operator] = false;
        
        //sorting account in such a way that just removed account is now replace with last of array
       
        //copping last element of array of operatorlist to vaccent index
        //index-1 is last element of array as we increased index in activeOperatorAccount()
        
        operatorlist[operatorindexaccountpair[account_of_operator]] = operatorlist[operatorindex-1];
        
        //decreasing list as we placed last element in vaccent place
        operatorindex--;
        emit LogOfRemoveOperatorAccount(account_of_operator);
      return true;
    }
    
    /**
     * operatorEntry() for activities to enter into the system and store the account.
     * there are two conditions
     * 1. for initial entry
     * 2. for manula change for number of operator
     * @return true if successfully stored an account in the system.
    
    */
    function operatorEntry() 
    public
    returns(bool){
        
        //for the first time storing account as operator account 
        //operator account cannot be stored twice. 
        //In case same operator hits the contract twice ,
        //then it is handled on storing time.
        
        if((operatorindex) == 0){
            require(activeOperatorAccount(msg.sender));
            return true;
        }
        
        
        //for the first time operator is stored directly. 
        //if node owner wants to add multiple account as operator
        //Or if this first account private key is lost then we need 
        //another account and system have to work properly.
        
      
        //if max number of operator is set other then default by the operator initially stored in the system.
        //then there will be operator greater then default value. Default value is zero.
      /**
       * operatorindex, maxnumberofoperator,testing(true/false)
       * in case of not changing operator count from default 
       * 1 , 0 , t
       * 
       * in case of changing operator other then zero
       * 
       * 1,1, t
       * 1,2, t
       * 2,1, f
       * 2,2, t
       * 2,3, t
      
      */
        //operatorindex already become 1
        //this if condition will satisfy only at maxnumberofoperator is default value
       
        else if(operatorindex <= maxnumberofoperator){
            require(activeOperatorAccount(msg.sender));
            return true;
        }
        else revert();
     }
    
    //maximum number of operator allowed in the system
    uint256 public maxnumberofoperator ;
    
    //voters vote for resetOperator() to reset the operator entry aggain
    uint256 public votesforoperatorcounter;
   
    //setOperatorCount() will call by the operator that already stored in the system.
    function setOperatorCount(uint256 number_of_operators)
    public 
    isoperator{
        maxnumberofoperator = number_of_operators;
        
        emit LogOfSetOperatorCount(maxnumberofoperator);
        
    }
    
    /**
     * if operator is out of control and need to change by the node owner
     * vote by the node owner to set the new operator. 
     * this function is only execute by the node owners
     * this is only for critical case of operator
     * max number of operator also set to 0 inorder to start work again from begaining
     * 
    */
    
    function resetOperator() 
    public 
    isvoter{
        
       if(votesforoperatorcounter == acceptOperatorConsensus()){
          
       
        //maxnumberofoperator and operatorindex is set to zero so we can start new process for the operator.
        maxnumberofoperator = 0;
        
        operatorindex = 0;
        
        //emit the event so we can track which account involved in resetting the process.
        
        emit LogOfResetOperator(msg.sender);
       }else{
           votesforoperatorcounter++;
           emit LogOfResetOperator(msg.sender);
           
       }
    }
    
    
   //operatorconsensuspercentage is used for number of voters required to set the new consensus from the default value. 
    uint256 public operatorconsensuspercentage;
    
    
    /**
    * setOperatorConsensus sets the consensus percentage
    * @param _percentage , vote required to accept the node.  
    */
    
    function setOperatorConsensus(uint _percentage)
        public 
        isoperator{
            //checking range of percentage , 0 <= _percentage => 100 
            require((_percentage >= 0) && (_percentage <= 100 ));
            operatorconsensuspercentage = _percentage;    
            emit LogOfOperatorSetConsensus(operatorconsensuspercentage);
    }

    
    
   /** 
    * acceptOperatorConsensus funciton  calculate the total number if vote required to accept the Operator 
    * @return limit of vote for accept an Operator 
   */
    function acceptOperatorConsensus()
        private 
        view
        returns(uint limit){
            //just setting  ~50%. 
            limit = ((operatorconsensuspercentage*(voterindex-1)/100));
           //checking overflow
            if (limit >= voterindex){
                return voterindex;
            }else {
                return limit;
            }
    }

    //Below for proposal acceptor.
    
    address public addressofproposalacceptor;
    
    /**
     * proposalaccepterEntry() function replace previous account of proposal acceptor with new address.
     * @param _addressofproposalacceptor is address of new proposal acceptor account
    */
    
    function proposalaccepterEntry(address _addressofproposalacceptor)
    public
    isoperator{ 
        addressofproposalacceptor = _addressofproposalacceptor;
        emit LogOfAddressOfProposalAcceptorEntry(addressofproposalacceptor);
    }

 
}