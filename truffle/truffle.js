/*
 * NB: since truffle-hdwallet-provider 0.0.5 you must wrap HDWallet providers in a 
 * function when declaring them. Failure to do so will cause commands to hang. ex:
 * ```
 * mainnet: {
 *     provider: function() { 
 *       return new HDWalletProvider(mnemonic, 'https://mainnet.infura.io/<infura-key>') 
 *     },
 *     network_id: '1',
 *     gas: 4500000,
 *     gasPrice: 10000000000,
 *   },
 */
var secrets = require("./secrets.json")
var HDWalletProvider = require("truffle-hdwallet-provider");

module.exports = {
    networks: {
	testnet: {
		host: "localhost",
		port: 8545,
		network_id: "*",
		gasLimit: 4700000
	},
	testnet2: {
		host: "localhost",
		port: 8546,
		network_id: "*",
		gasLimit: 4700000
	},
   	ropsten: {
      		provider: new HDWalletProvider(secrets.mnemonic, "https://ropsten.infura.io/"),
      		network_id: "*",
      		gas: 1000000,
      		gasLimit: 67000000
      		//gasPrice: web3.utils.toWei("20", "gwei") 
   	},
	rinkeby: {
                provider: new HDWalletProvider(secrets.mnemonic, "https://rinkeby.infura.io/gpcq2PXJhM3TALrZmuhX"),
                network_id: "*",
                gas: 1000000,
                gasLimit: 67000000
        },
        kovan: {
                provider: new HDWalletProvider(secrets.mnemonic, "https://kovan.infura.io/gpcq2PXJhM3TALrZmuhX"),
                network_id: "*",
                gas: 1000000,
                gasLimit: 67000000
        },
   	mainnet: {
      		provider: new HDWalletProvider(secrets.mnemonic, "https://mainnet.infura.io/gpcq2PXJhM3TALrZmuhX"),
      		network_id: 1,
      		gas: 1000000,
      		gasLimit: 67000000
   	}  
    }
};
