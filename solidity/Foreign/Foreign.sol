pragma solidity ^0.5.0;

/*
 * To be deployed on the network you are bridging out of (Foreign => Home).
 */

contract Foreign {

	uint256 constant MAX_DEPOSIT = 100 ether;

	mapping(address => uint256) deposited; // how much ether that was deposited today
	mapping(address => uint256) depositTime; // time from their first deposit today

	event ContractCreation(address _owner);
	event Deposit(address _recipient, uint _value, uint _toChain);

	constructor() public {
		emit ContractCreation(msg.sender);
	}

	function deposit(address _recipient, uint _toChain) public payable {
		// if they haven't made a deposit in the last day, reset their amount and time
		if (depositTime[msg.sender] < now + 1 days) {
			depositTime[msg.sender] = now;
			deposited[msg.sender] = 0;
		}

		// cannot deposit more than the maximum in one day
		require(deposited[msg.sender] + msg.value < MAX_DEPOSIT, "Exceeds daily maximum.");

		// burn ether, update deposit balance for today, and emit event for bridge
		address(0).transfer(msg.value);
		deposited[msg.sender] += msg.value;
		emit Deposit(_recipient, msg.value, _toChain);
	}
}