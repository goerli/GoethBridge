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
	event DepositErc20(address _recipient, uint _value); 

	event Withdraw(address _recipient, uint _value, uint _fromChain); 
	event WithdrawErc20(address _recipient, uint _value); 

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

	function tokenFallback(address _sender, address _origin, uint _value, bytes _data) public returns (bool ok) {
		emit DepositErc20(_origin, _value);
		return true;
	}

	function setBridge(address _addr) public onlyOwner {
		bridge = _addr;
		emit BridgeSet(bridge);
	}

	function withdraw(address _recipient, uint _value, uint _fromChain) public onlyBridge {
		emit Withdraw(_recipient, _value, _fromChain);
		_recipient.transfer(_value);
	}

	function withdrawErc20(address _recipient, uint _value) public onlyBridge {
		emit WithdrawErc20(_recipient, _value);
		transfer(_recipient, _value);
	}

	/* erc20 functions */
	mapping(address => uint) tokenBalance;
	uint totalSupply;

	event Transfer(address _to, uint _value);

	function transfer(address _to, uint _value) public returns (bool) {
		require(tokenBalance[msg.sender] >= _value);
		tokenBalance[msg.sender] -= _value;
		tokenBalance[_to] += _value;
		emit Transfer(_to, _value);
		return true;
	}
}
