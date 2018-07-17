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
	event ContractCreation(address _owner);

	// currently, we only specify the receiving address on the other chain
	// and the value of the deposit.
	// @todo: add `address _toChain` so that the user can specify which
	// chain they wish to withdraw on. this makes the bridge "multi-directional.
	event Deposit(address _receiver, uint _value); 

	// @todo: similarly to the Deposit event, we eventually wish to add
	// a `address _fromChain` argument
	event Withdraw(address _receiver, uint _value); 

	constructor() public {
		emit ContractCreation(msg.sender);
	}

	/* bridge functions */
	function () public payable {
		emit Deposit(msg.sender, msg.value);
	}

	function withdraw(address _receiver, uint _value) public {
		emit Withdraw(_receiver, _value);
		_receiver.transfer(_value);
	}

	/* erc20 functions */
	event Transfer(address _to, uint _value);

	function transfer(address _to, uint _value) public returns (bool) {
		emit Transfer(_to, _value);
		return true;
	}
}
