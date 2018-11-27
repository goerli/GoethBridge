pragma solidity ^0.5.0;

/* BridgeableToken Smart Contract
*  @noot
*  this contract implements functionality for bridging of an erc20/erc223 token
*  for a token to be bridgeable, the bridge contract needs to be able to transfer 
*  the token to a recipient on the receiving chain.
*  with a usual erc20 implementation, the bridge would need approval to transfer
*  tokens, otherwise the transfer will fail.
*/

import "./lib/StandardToken.sol";

contract BridgeableToken is StandardToken {

	address owner;
	address bridge;

	string public constant name = "Bridgeable Token";
    string public constant symbol = "BETH";
    uint8 public constant decimals = 18; 

	event Deposit(address _recipient, uint _value, uint _toChain); 
	event Withdraw(address _recipient, uint _value, uint _fromChain); 

	constructor() public {
		owner = msg.sender;
		bridge = msg.sender;
		uint256 initialSupply = 33 ** uint256(decimals);
		_mint(owner, initialSupply);
	}

	modifier onlyOwner() {
		require(msg.sender == owner);
		_;
	}

	modifier onlyBridge() {
		require(msg.sender == bridge);
		_;
	}

	function setBridge(address _addr) public onlyOwner {
		bridge = _addr;
	}

	// this function will burn tokens and emit a deposit event to be picked up by the bridge.
	function deposit(address _recipient, uint256 _value, uint256 _toChain) public {
		_burn(msg.sender, _value);
		emit Deposit(_recipient, _value, _toChain);
	}

	// this function is called by the bridge upon seeing a deposit event
	// this function will mint tokens for the recipient and emit a Withdraw event
	function withdraw(address _recipient, uint256 _value, uint256 _fromChain) public onlyBridge {
		_mint(_recipient, _value);
		emit Withdraw(_recipient, _value, _fromChain);
	}
}