# Multi Directional Bridge

this is a go implementation of a generic bridge between blockchains. the bridge will be able to connect any networks using native fuel or any token.

### todo
* implement support for non EVM based blockchains

### requirements
go 1.9.1

go-ethereum
`go get github.com/ethereum/go-ethereum`

jsonparser

in $GOPATH/src
`git clone https://github.com/buger/jsonparser`

# to get the bridge
`go get github.com/ChainSafeSystems/gobridge`

# to run
`go run main.go 3 42`
  
  the arguments after `go run main.go` are the IDs of the networks you want to listen on
  
  the IDs and chain info are in the config.json file

* 1: mainnet

* 3: ropsten

* 4: rinkeby

* 42: kovan

* 31: rootstock testnet
  
  additional flags:
 `go run main.go -a 3 42`
 
 `go run main.go --config ./config.json 3 42`
 
 `-a` read logs from every contract on the network (not really useful, mostly for testing)
 
 `-v` verbose output
 
 `--config` specify path to config file
 
 `--keystore` specify path to keystore file

# interacting with the contract

for all the following, you should have another terminal open running the bridge listener with `go run main.go CHAINID1 CHAINID2...`

`go run main.go fund CHAINID` this will open up a prompt for you to make a deposit on the specified chain

`go run main.go deposit CHAINID` this will open up a prompt for you to make a deposit on the specified chain id

`go run main.go pay CHAINID` pay the bridge contract for a later withdraw on the specified chain

`go run main.go withdraw CHAINID` this will withdraw ether that was paid to the bridge contract previously 
 
 `--keystore` specify path to keystore directory
 
 `--password` specify password to account; this assumes that there's the same account for every chain

# issues

you may encounter a "Failed to read file" error referencing Bridge.json. If this happens, run the following:
```
cd truffle/
npm install -g truffle
npm install truffle-hdwallet-provider
cp secrets-example.json secrets.json
truffle compile
```

this sets up needed dependencies for truffle and compiles the contracts.
