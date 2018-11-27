pragma solidity ^0.5.0;

/* Bridge Smart Contract 
* @noot
* to be deployed on mainnet.
*/

contract Timelock {
	uint256 constant maxDeposit = 100 ether;

	uint256 constant ratio = 1000;

	mapping(address => uint256) deposited; // how much ether address has deposited
	mapping(address => uint256) depositTime; // time from their latest deposit

	event ContractCreation(address _owner);
	event Deposit(address _recipient, uint _value, uint _toChain); 

	constructor() public {
		emit ContractCreation(msg.sender);
	}

	function deposit(address _recipient, uint _toChain) public payable {
		// update deposit balance for today, and emit event for bridge
		deposited[msg.sender] += msg.value;
		depositTime[msg.sender] = now;
		emit Deposit(_recipient, msg.value * ratio, _toChain);
	}

	function release() public {
		require(depositTime[msg.sender] + 4 weeks < now);
		msg.sender.transfer(deposited[msg.sender]);
		deposited[msg.sender] = 0;
	}
}