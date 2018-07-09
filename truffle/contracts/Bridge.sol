pragma solidity ^0.4.23;

contract Bridge {
	event ContractCreation(address _owner);
	event Transfer(address _to, uint _value);
	
	constructor() {
		emit ContractCreation(msg.sender);
	}

	function transfer(address _to, uint _value) public returns (bool) {
		emit Transfer(_to, _value);
		return true;
	}
}
