package client 

import (
	"fmt"
	"time"
	"encoding/hex"
	//"encoding/json"
	//"io/ioutil"
	"math/big"
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// events to listen for
var events *Events
var keys *keystore.KeyStore

type Chain struct {
	Url string
	Id string
	Contract *common.Address
	GasPrice *big.Int
	From *common.Address
	Client *ethclient.Client
	Nonce uint64
}

// events to listen for
type Events struct {
	DepositId string
  	CreationId string
 	WithdrawId string
	BridgeSetId string
}

/***** client functions ******/
func Filter(chain *Chain, filter *ethereum.FilterQuery, logsDone chan bool) {
	logs, err := chain.Client.FilterLogs(context.Background(), *filter)
	if err != nil {
		fmt.Println(err)
	}
	if len(logs) != 0 {
		//fmt.Println(len(logs))
		go ReadLogs(chain, logs)
	}

	logsDone <- true
}

func ReadLogs(chain *Chain, logs []types.Log) {
	//logs := <-ch
	//fmt.Println(logs)
	for _, log := range logs {
		fmt.Println("\nlogs found on chain", chain.Id)
		fmt.Println("contract address: ", log.Address.Hex())
		for _, topics := range log.Topics {
			topic := topics.Hex()
			fmt.Println("topics: ", topic)

			if strings.Compare(topic,events.DepositId) == 0 { 
				fmt.Println("*** deposit event")
				txHash := log.TxHash.Hex()
				fmt.Println("txHash: ", txHash)
				go ActOnDeposit(chain, log.TxHash)

		 	} else if strings.Compare(topic,events.CreationId) == 0 {
				fmt.Println("*** bridge contract creation")
			} else if strings.Compare(topic,events.WithdrawId) == 0 {
				fmt.Println("*** withdraw event")
				txHash := log.TxHash.Hex()
				fmt.Println("txHash: ", txHash)
				// receiver, value, toChain := readDepositData(data)
				// fmt.Println("receiver: ", receiver) 
				// fmt.Println("value: ", value) // in hexidecimal
				// fmt.Println("to chain: ", toChain) // in hexidecimal
			} else if strings.Compare(topic,events.BridgeSetId) == 0 {
				fmt.Println("*** set bridge event")
			}
		}
	}
}

func ActOnDeposit(chain *Chain, txHash common.Hash) {
	tx, isPending, err := chain.Client.TransactionByHash(context.Background(), txHash)
	if isPending {
		// wait
	}
	if err != nil {
		fmt.Println(err)
	}

	data := hex.EncodeToString(tx.Data())
	//fmt.Println("data: ", data)
	//fmt.Println(len(data))
	if len(data) > 72 {
		receiver := data[32:72]
		toChain := data[72:136]
		value := tx.Value()
		// receiver, value, toChain := readDepositData(data)
		fmt.Println("receiver: ", receiver) 
		fmt.Println("value: ", value) // in hexidecimal
		fmt.Println("to chain: ", toChain) // in hexidecimal
	}
}

func SetBridge(chain *Chain) () {
	client := chain.Client
	accounts := keys.Accounts()

	data, err := hex.DecodeString("8dd14802000000000000000000000000288a9fb92921472d29ab0b3c3e420a8e4bd4f452")
	if err != nil {
		fmt.Println(err)
	} 
	//data := new([]byte)
	// NewTransaction(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte)
	tx := types.NewTransaction(chain.Nonce, *chain.Contract, big.NewInt(int64(0)), uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTx(accounts[0], tx, big.NewInt(int64(33))) // chainId
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}
}

// starts a goroutine to listen on every chain 
func Listen(chain *Chain, e *Events, doneClient chan bool, ks *keystore.KeyStore, flags map[string]bool) {
	//logsFound := make(map[common.Hash]bool)
	//fmt.Println(chain)
	events = e
	keys = ks
	fmt.Println("listening at: " + chain.Url)

	client, err := ethclient.Dial(chain.Url)
	if err != nil {
		log.Fatal(err)
	}
	chain.Client = client

	nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	chain.Nonce = nonce

	//SetBridge(chain)

	fromBlock := big.NewInt(1)

	filter := new(ethereum.FilterQuery)

	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C{
		if flags["v"] { fmt.Println(t) }

		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil { fmt.Println(err) }
		if flags["v"] { fmt.Println("latest block: ", block.Number()) }

		filter.FromBlock = fromBlock
		if !flags["a"] {
			contractArr := make([]common.Address, 1)
			contractArr = append(contractArr, *chain.Contract)
			filter.Addresses = contractArr
		}
		//filter.Addresses = append(filter.Addresses, *chain.Contract)

		logsDone := make(chan bool)
		go Filter(chain, filter, logsDone)
		<-logsDone


		fromBlock = block.Number()
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
	doneClient <- true
}