pragma solidity ^0.5.0;

/*
 * To be deployed on the network you are bridging out of (Foreign => Home).
 */

contract Foreign {
	// Per account daily limit
	uint256 constant MAX_DEPOSIT = 100 ether;
	// Bridge daily limit
	uint256 constant MAX_PER_DAY = 1000000 ether;

	address payable owner;
	uint256 timestamp;
	uint256 dailyTotal;

	mapping(address => uint256) deposited; // how much ether that was deposited today
	mapping(address => uint256) depositTime; // time from their first deposit today

	event ContractCreation(address _owner);
	event Deposit(address _recipient, uint _value, uint _toChain);
	event Withdraw(uint256 _amount);

	constructor() public {
		owner = msg.sender;
		timestamp = now;
		emit ContractCreation(msg.sender);
	}

	function deposit(address _recipient, uint _toChain) public payable {
		if(now < timestamp + 1 days) {
			// Ensure we are not exceeding the limit
			assert(dailyTotal + msg.value <= MAX_PER_DAY);
			dailyTotal += msg.value;
		} else {
			// Reset total and timestamp
			timestamp = now;
			dailyTotal = 0;
		}

		// if they haven't made a deposit in the last day, reset their amount and time
		if (depositTime[msg.sender] < now + 1 days) {
			depositTime[msg.sender] = now;
			deposited[msg.sender] = 0;
		}

		// cannot deposit more than the maximum in one day
		require(deposited[msg.sender] + msg.value < MAX_DEPOSIT, "Exceeds daily maximum.");

		// update deposit balance for account for today, and emit event for bridge
		deposited[msg.sender] += msg.value;
		emit Deposit(_recipient, msg.value, _toChain);
	}

	function withdraw(uint _amount) public returns (bool){
		require(msg.sender == owner);
		require(_amount <= address(this).balance);
		owner.transfer(_amount);
		emit Withdraw(_amount);
		return true;
	}
}