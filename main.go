package main

import (
	//"fmt"
	"log"
	//"context"
	//"time"
	"net/rpc"
	//"github.com/ethereum/go-ethereum/rpc"
	//"math/big"
)

func main() {
	// hard coded to geth running at address:port
	client, err := rpc.Dial("tcp", "127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//subch := make(chan *big.Int)
	var reply int
	args := 0
	err = client.Call("eth_gasLimit", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

}
