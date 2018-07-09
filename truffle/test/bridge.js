var Bridge = artifacts.require("Bridge");

contract('Bridge', function(accounts) {
	var addrA = accounts[0];

	it("should deploy", async() => {
		bridge = await Bridge.new();
	})
	
	it("should call transfer", async() => {
		txHash = await bridge.transfer(addrA, 10)
	})
})
