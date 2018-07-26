var Bridge = artifacts.require("Bridge");

contract('Bridge', function(accounts) {
	var addrA = accounts[0]

	it("should deploy", async() => {
		bridge = await Bridge.deployed()
		console.log("\tbridge contract address: " + bridge.address)
		assert(bridge != undefined)
		owner = await bridge.owner.call()
		assert(owner == addrA)
	})

	it("should make a ether deposit", async() => {
		let _val = 17
		deposit = await bridge.sendTransaction({from: addrA, value: _val})
	})
	
	// it("should call transfer", async() => {
	// 	txHash = await bridge.transfer(addrA, 10)
	// })
})
