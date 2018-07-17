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
	address owner;
	// this is the public key address of the bridge's own keypair, not
	// the address of the BridgeSafe or any contract.
	address public bridge;

	event ContractCreation(address _owner);

	// currently, we only specify the receiving address on the other chain
	// and the value of the deposit.
	// @todo: add `address _toChain` so that the user can specify which
	// chain they wish to withdraw on. this makes the bridge "multi-directional.
	event Deposit(address _receiver, uint _value); 
	event DepositErc20(address _receiver, uint _value); 

	// @todo: similarly to the Deposit event, we eventually wish to add
	// a `address _fromChain` argument
	event Withdraw(address _receiver, uint _value); 
	event WithdrawErc20(address _receiver, uint _value); 

	constructor() public {
		owner = msg.sender;
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
		emit Deposit(msg.sender, msg.value);
	}

	function tokenFallback(address _sender, address _origin, uint _value, bytes _data) public returns (bool ok) {
		emit DepositErc20(_origin, _value);
		return true;
	}

	function setBridge(address _addr) public onlyOwner {
		bridge = _addr;
		//owner = 0x0; // remove owner after bridge is set?
		// possibly risky 

	}

	function withdraw(address _receiver, uint _value) public onlyBridge {
		emit Withdraw(_receiver, _value);
		_receiver.transfer(_value);
	}

	function withdrawErc20(address _receiver, uint _value) public onlyBridge {
		emit WithdrawErc20(_receiver, _value);
		transfer(_receiver, _value);
	}

	/* erc20 functions */
	mapping(address => uint) balance;
	uint totalSupply;

	event Transfer(address _to, uint _value);

	function transfer(address _to, uint _value) public returns (bool) {
		require(balance[msg.sender] >= _value);
		balance[msg.sender] -= _value;
		balance[_to] += _value;
		emit Transfer(_to, _value);
		return true;
	}
}
