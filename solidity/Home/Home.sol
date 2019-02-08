pragma solidity ^0.5.0;

/*
 * To be deployed on the network you are bridging into (Foreign => Home).
 */

contract Home {
	address public owner;
	address public bridge;

	mapping(bytes32 => bool) withdrawSubmitted;

	event ContractCreation(address _owner);
	event BridgeSet(address _addr);
	event BridgeFunded(address _addr);
	event Withdraw(address _recipient, uint _value, uint _fromChain, bytes32 _txHash);

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
	function () external payable onlyBridge {
		emit BridgeFunded(msg.sender);
	}

	function setBridge(address _addr) public onlyOwner {
		bridge = _addr;
		emit BridgeSet(bridge);
	}

	function withdraw(address payable _recipient, uint _value, uint _fromChain, bytes32 _txHash) public onlyBridge {
		require(!withdrawSubmitted[_txHash]);
		withdrawSubmitted[_txHash] = true;
		_recipient.transfer(_value);
		emit Withdraw(_recipient, _value, _fromChain, _txHash);
	}
}
