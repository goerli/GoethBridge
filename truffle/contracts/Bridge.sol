pragma solidity ^0.4.23;

/* Bridge Smart Contract 
* @noot
* this contract was created to be a generic bridge contract that will emit a Deposit() 
* event upon receiving ether. this Deposit() event will then be picked up by the bridge,
* which acts as a client. the bridge will then submit a transaction to withdraw ether
* on the other chain. when this withdraw is completed, a Withdraw() event will be emitted.
* 
* this contract will be deployed on both sides of the bridge.
*
* @todo: ERC20/223 support
* @todo: testing + security; merkle proofs; valiadators ning off on transactions for PoA 
* chains or private chains/bridges
*/

contract Bridge {
	address public owner;
	// this is the public key address of the bridge's own keypair, not
	// the address of the BridgeSafe or any contract.
	address public bridge;

	mapping(address => uint) balance;

	event ContractCreation(address _owner);
	event BridgeSet(address _addr);
	event BridgeFunded(address _addr);
	event Paid(address _addr, uint _value);

	event Deposit(address _recipient, uint _value, uint _toChain); 
	event Withdraw(address _recipient, uint _value, uint _fromChain); 

	constructor() public {
		owner = msg.sender;
		bridge = msg.sender;
		emit ContractCreation(msg.sender);
	}

	modifier onlyOwner() {
		require(msg.sender == owner);
		_;
	}

	modifier onlyBridge() {
		require(msg.sender == bridge);
		_;
	}

	/* bridge functions */
	function () public payable {
		//revert();
		balance[msg.sender] += msg.value;
		emit Paid(msg.sender, msg.value);
	}

	function fundBridge() public payable {
		// thanks
		emit BridgeFunded(msg.sender);
	}

	function deposit(address _recipient, uint _toChain) public payable {
		emit Deposit(_recipient, msg.value, _toChain);
	}

	function withdrawTo(address _recipient, uint _toChain, uint _value) public {
		require(balance[msg.sender] >= _value);
		balance[msg.sender] -= _value;
		emit Deposit(_recipient, _value, _toChain);
	}

	function setBridge(address _addr) public onlyOwner {
		bridge = _addr;
		emit BridgeSet(bridge);
	}

	function withdraw(address _recipient, uint _value, uint _fromChain) public onlyBridge {
		_recipient.transfer(_value);
		emit Withdraw(_recipient, _value, _fromChain);
	}
}
