# GöethBridge

This is a go implementation of a generic bridge between blockchains. Made specifically to allow testnet eth from rinkeby, kovan, and ropsten to be burned for Göeth (Görli eth). 

### todo
* allow for Göeth to be converted back to the other testnets. 

### requirements
go 1.9.1

go-ethereum
`go get github.com/ethereum/go-ethereum`

leth
`go get github.com/ChainSafeSystems/leth`

# to get the bridge
`go get github.com/goerli/GoethBridge`

# to run
#### generic instructions for bridge, needs to be updated!
```
cd $GOPATH/src/github.com/goerli/GoethBridge
cd leth && leth compile
cd ..
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

# issues

you may encounter a "Failed to read file" error referencing Bridge.abi. If this happens, run the following:
```
cd leth
leth compile
```

if for some reason leth isn't working, you can also use solc.
inside ChainBridge/leth:
```
mkdir build
solc --abi contracts/Bridge.sol -o build
```

