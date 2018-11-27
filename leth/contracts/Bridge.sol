pragma solidity ^0.5.0;

/* Bridge Smart Contract 
* @noot
* this contract is a generic bridge contract that will emit a Deposit() 
* event upon receiving ether. this Deposit() event will then be picked up by the bridge,
* which acts as a client. the bridge will then submit a transaction to withdraw ether
* on the other chain. when this withdraw is completed, a Withdraw() event will be emitted.
* 
* this contract will be deployed on both sides of the bridge.
*/

contract Bridge {
	address public owner;
	address public bridge;

	uint256 threshold = 1; // the number of signatures that must be reached for a withdraw to take place

	mapping(address => bool) authorities;
	mapping(address => uint256) balance;
	mapping(bytes32 => uint256) withdrawal; // withdrawal tx hash to number of signatures
	mapping(bytes32 => mapping(address => bool)) signedWithdrawal; // tx hash to authority address to whether they have already signed

	event ContractCreation(address _owner);
	event BridgeSet(address _addr);
	event BridgeFunded(address _addr);
	event Paid(address _addr, uint _value);

	event AuthorityAdded(address _addr);
	event AuthorityRemoved(address _addr);
	event ThresholdUpdated(uint256 _threshold);

	event Deposit(address _recipient, uint _value, uint _toChain); 
	event Withdraw(address _recipient, uint _value, uint _fromChain); 
	event SignedForWithdraw(bytes32 _txHash, address _authority);

	constructor() public {
		owner = msg.sender;
		bridge = msg.sender;
		emit ContractCreation(msg.sender);
	}

	/* admin */
	modifier onlyOwner() {
		require(msg.sender == owner);
		_;
	}

	modifier onlyAuthority() {
		require(isAuthority(msg.sender));
		_;
	}

	function setThreshold(uint256 _threshold) public onlyOwner {
		threshold = _threshold;
		emit ThresholdUpdated(threshold);
	}

	function increaseThreshold() public onlyOwner {
		threshold++;
		emit ThresholdUpdated(threshold);
	}

	function decreaseThreshold() public onlyOwner {
		if (threshold > 0) {
			threshold--;
			emit ThresholdUpdated(threshold);
		}
	}

	function addAuthority(address _addr) public onlyOwner {
		authorities[_addr] = true;
		emit AuthorityAdded(_addr);
	}

	function removeAuthority(address _addr) public onlyOwner {
		authorities[_addr] = false;
		emit AuthorityRemoved(_addr);
	}

	function isAuthority(address _addr) public returns (bool) {
		return authorities[_addr];
	}

	/* bridge functions */
	function () external payable {
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

	function withdraw(address payable _recipient, uint _value, uint _fromChain, bytes32 _txHash) public onlyAuthority {
		// make sure authority has not already signed for this withdraw
		require(!signedWithdrawal[_txHash][msg.sender]);
		withdrawal[_txHash]++;
		signedWithdrawal[_txHash][msg.sender] = true;
		emit SignedForWithdraw(_txHash, msg.sender);

		// if enough authorities have signed, execute the withdraw
		if(withdrawal[_txHash] >= threshold) {
			_recipient.transfer(_value);
			emit Withdraw(_recipient, _value, _fromChain);
		}
	}
}