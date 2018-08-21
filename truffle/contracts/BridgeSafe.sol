pragma solidity ^0.4.23;

/*
* Bridge Safe Smart Contract
* @noot
* the bridge will have access to funds in this contract. this contract will
* need to have balance so that the bridge can pay tx fees and broadcast transactions. 
*
* this is not currently used.
*/

contract BridgeSafe {
	address owner;
	event ContractCreation(address _owner);

	constructor() public {
		owner = msg.sender;
		emit ContractCreation(msg.sender);
	}

	modifier onlyOwner() {
		require(owner == msg.sender);
		_;
	}

	function () public payable {
		// pay me
	}

	function transferOwnership(address _addr) public onlyOwner {
		owner = _addr;
	}
}