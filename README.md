# Multi Directional Bridge

this is a go implementation of a generic bridge between blockchains. the bridge will be able to connect any networks using native fuel or any token.

### todo
* implement log filtering

* implement posting of transactions

* implement support for non-ethereum based blockchains

### requirements
go 1.9.1

go-ethereum
`go get github.com/ethereum/go-ethereum`

# to run
`go run main.go 1 3 42`
  
  the arguments after `gr main.go` are the IDs of the networks you want to listen on
  
  the IDs and chain info are in the config.json file
  
  additional flags:
 `go run main.go -a 1 3 42`
 
 `go run main.go --config ./config.json 1 3 42`
 
 `go run main.go --password pswd 1 33`
 
 `-a` read logs from every contract on the network (not really useful, mostly for testing)
 
 `-v` verbose output
 
 `--password` password string to account 0
 
 `--config` specify path to config file
