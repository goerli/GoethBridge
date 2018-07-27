pragma solidity ^0.4.23;

/*
* Bridge Safe Smart Contract
* @noot
* the bridge will have access to funds in this contract. this contract will
* need to have balance so that the bridge can pay tx fees and broadcast transactions. 
*
* @todo: bridge needs its own private key which it will use to become the owner of this contract.
* is it possible to generate a key inside the bridge and use it to deploy this contract,
* or will a user (ie. me) need to deploy this contract, then transfer ownership to the 
* bridge? either way, the bridge needs to have its own key and some fuel to start with.
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