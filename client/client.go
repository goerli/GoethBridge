package client 

import (
	"fmt"
	"time"
	//"encoding/json"
	//"io/ioutil"
	"math/big"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
)

type Chain struct {
	Url string
	Contract *common.Address
	GasPrice *big.Int
	From string
	Client *ethclient.Client
}

func Filter(chain *Chain, filter *ethereum.FilterQuery, ch chan []types.Log) {
	logs, err := chain.Client.FilterLogs(context.Background(), *filter)
	if err != nil {
		fmt.Println(err)
	}
	ch <- logs
}

// starts a goroutine to listen on every chain 
func Listen(chain *Chain, doneClient chan bool) {
	//logsFound := make(map[common.Hash]bool)
	//fmt.Println(chain)
	fmt.Println("listening at: " + chain.Url)

	client, err := ethclient.Dial(chain.Url)
	if err != nil {
		log.Fatal(err)
	}
	chain.Client = client

	fromBlock := new(big.Int)
	fromBlock.SetInt64(int64(0))

	ticker := time.NewTicker(100 * time.Millisecond)
	for t := range ticker.C{
		if false { fmt.Println(t) }

		filter := new(ethereum.FilterQuery)
		filter.FromBlock = fromBlock
		//filter.Addresses = append(filter.Addresses, *chain.Contract)

		ch := make(chan []types.Log)
		go Filter(chain, filter, ch)

		// if(len(logs) > 0) {
		// 	logsFound[logs[]]
		// 	fmt.Println(logs)
		// }

		// go func(err <-chan error){
		// 	e := <-err
		// 	log.Fatal(e)
		// }(sub.Err())

		// log := <-logs
		// fmt.Println(log)
	}
 
	// bridge timeout. eventually, change so it never times out
	time.Sleep(6000 * time.Second)
	ticker.Stop()

	// for t := range ticker.C{

	// 	//fmt.Println(len(logsResult))

	// 	// if there are new logs, parse for event info
	// 	if len(logsResult) > 2 {
	// 		txHash, _ := parseJsonForEntry(logsResult[1:len(logsResult)-1], "transactionHash")
	// 		//fmt.Println(txHash + "\n")
	// 		if logsFound[txHash] != true { 
	// 			logsFound[txHash] = true
	// 			fmt.Println("\nnew logs found for chain", chain)

	// 			//logs <- logsResult
	// 			//readLogs(logs)
	// 			//go readLogs(logs)
	// 			//<-exit

	// 			// get logs contract address
	// 			address, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "address")
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			// this is not actually a good way to listen for events from a  contract
	// 			// this could be used to confirm a log, but for listening to events from
	// 			// one contract, we would specify the address in our call to eth_getLogs
	// 			fmt.Println("contract addr: ", address)
	// 			//fmt.Println("length of address: ", len(address))
	// 			for i := 0; i < len(chains); i++ {
	// 				if strings.Compare(address[1:41], chains[i]) == 0 {
	// 					fmt.Println("bridge contract event heard on chain ", chains[i])
	// 				}
	// 			}

	// 			// read topics of log
	// 			topics, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "topics")
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			fmt.Println("topics: ", topics[2:68])
	// 			//fmt.Println("length of topics: ", len(topics)-4) len = 66: 0x + 64 hex chars = 32 bytes

	// 			if strings.Compare(topics[2:68],DepositId) == 0 { 
	// 				fmt.Println("*** deposit event ", topics[2:68])
	// 				data, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "data")
	// 				if err != nil {
	// 					fmt.Println(nil)
	// 				}

	// 				receiver, value, toChain := readDepositData(data)
	// 				fmt.Println("receiver: ", receiver) 
	// 				fmt.Println("value: ", value) // in hexidecimal
	// 				fmt.Println("to chain: ", toChain) // in hexidecimal
	// 		 	} else if strings.Compare(topics[2:68],CreationId) == 0 {
	// 				fmt.Println("*** bridge contract creation")
	// 			} else if strings.Compare(topics[2:68],WithdrawId) == 0 {
	// 				fmt.Println("*** withdraw event")
	// 				data, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "data")
	// 				if err != nil {
	// 					fmt.Println(nil)
	// 				}
	// 				receiver, value, toChain := readDepositData(data)
	// 				fmt.Println("receiver: ", receiver) 
	// 				fmt.Println("value: ", value) // in hexidecimal
	// 				fmt.Println("to chain: ", toChain) // in hexidecimal
	// 			} else if strings.Compare(topics[2:68],BridgeSetId) == 0 {
	// 				fmt.Println("*** set bridge event")
	// 				setBridgeDone <- true
	// 			}
	// 		}
	// 	}
	// }
}