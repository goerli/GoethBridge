# Multi Directional Bridge

this is a go implementation of a generic bridge between blockchains. the bridge will be able to connect any networks using native fuel or any token.

### todo
* implement support for non EVM based blockchains

### requirements
go 1.9.1

go-ethereum
`go get github.com/ethereum/go-ethereum`

solc/solcjs
`npm i -g solc`

# to get the bridge
`go get github.com/ChainSafeSystems/ChainBridge`

# compile abi
```
cd solidity
mkdir build
solcjs --abi contracts/Bridge.sol -o build
```

# to run
```
cd $GOPATH/src/github.com/ChainSafeSystems/ChainBridge
go build && go install
```

`ChainBridge [networks]`
  
the arguments after `ChainBridge` are the names of the networks you want to listen on as specified in config.json

eg. `ChainBridge ropsten kovan`

* 1: mainnet

* 3: ropsten

* 4: rinkeby

* 42: kovan

* 31: rootstock testnet
  
  additional flags:
 `ChainBridge -a [networks]`
 
 `ChainBridge --config ./config.json [networks]`
 
 `-a` read logs from every contract on the network (not really useful, mostly for testing)
 
 `-v` verbose output
 
 `--config` specify path to config file
 
 `--keystore` specify path to keystore file

# interacting with the contract

for all the following, you should have another terminal open running the bridge listener with `ChainBridge [networks]`

`ChainBridge fund network` this will open up a prompt for you to make a deposit on the specified chain

`ChainBridge deposit network` this will open up a prompt for you to make a deposit on the specified chain id

`ChainBridge pay network` pay the bridge contract for a later withdraw on the specified chain

`ChainBridge withdraw network` this will withdraw ether that was paid to the bridge contract previously 
 
 `--keystore` specify path to keystore directory
 
 `--password` specify password to account; this assumes that there's the same account for every chain

eg. `ChainBridge fund kovan`


