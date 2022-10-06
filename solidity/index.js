const ethers = require("ethers")
const Wallet = ethers.Wallet
const providers = ethers.providers
const path = require("path")
const fs = require("fs")

// INSTRUCTIONS FOR USE
// 1. Copy the path to your JSON keystore file into the 'kspath' constant.
// 2. Set the 'password' constant to the password to your JSON keystore.
// 3. Set the 'url' constant to the url of your Ethereum node. make sure the account you're using for
// your keystore has some ether on this network.
// 4. Set the 'contract' constant to the name of the contract you wish to deploy; can be set to
// 'Bridge', 'Home', or 'Foreign.'
const kspath = "../keystore/UTC--2018-05-17T21-58-52.188632298Z--8f9b540b19520f8259115a90e4b4ffaeac642a30"
const password = "password"
const url = "http://107.5.111.63:8545"
const contract = "Bridge"

const deploy = async() => {
	let keystore = path.resolve(kspath)
	let walletJson = fs.readFileSync(keystore, 'utf8')
	let wallet = await Wallet.fromEncryptedJson(walletJson, password)

	let provider = new providers.JsonRpcProvider(url, "unspecified")
	wallet = wallet.connect(provider)

	let abifile = path.resolve(`./${contract}/build/${contract}.abi`)
	let abi = fs.readFileSync(abifile, 'utf8')

	let binfile = path.resolve(`./${contract}/build/${contract}.bin`)
	let bin = fs.readFileSync(binfile, 'utf8')

	let bridgeFactory = new ethers.ContractFactory( abi , bin , wallet )
	let bridge = await bridgeFactory.deploy()
	await provider.waitForTransaction(bridge.deployTransaction.hash)
	let receipt = await provider.getTransactionReceipt(bridge.deployTransaction.hash)

	try {
		let results = path.resolve(`./${contract}-deployment.txt`)
		fs.writeFileSync(results, receipt.contractAddress)		
	} catch (e) {
		console.log(`${e.message}: could not save deployed address to file.`)
	}

}

deploy()