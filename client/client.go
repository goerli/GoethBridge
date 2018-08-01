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
	//"sync"
	//"path/filepath"

	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

/* global variables */
var events *Events // events to listen for
var keys *keystore.KeyStore // keystore; used to sign txs
var flags map[string]bool // command line flags
//var allChains []*Chain //[]*Chain
var logsRead = map[string]bool{}

type Chain struct {
	Url string
	Id *big.Int
	Contract *common.Address
	GasPrice *big.Int
	From *common.Address
	Password string
	Client *ethclient.Client
	Nonce uint64
}

type Withdrawal struct {
	Recipient string
	Value *big.Int
	FromChain string
	Data string
}

// events to listen for
type Events struct {
	DepositId string
  	CreationId string
 	WithdrawId string
	BridgeSetId string
	BridgeFundedId string
}

/****** helpers ********/

// pads zeroes on front of a string until it's 32 bytes or 64 hex characters long
func padTo32Bytes(s string) (string) {
	l := len(s)
	for {
		if l == 64 {
			return s
		} else {
			s = "0" + s
			l += 1
		}
	}
}

func padBigTo32Bytes(n *big.Int) (string) {
	nBytes := n.Bytes()
	nHexStr := hex.EncodeToString(nBytes)
	return padTo32Bytes(nHexStr)
}

func padIntTo32Bytes(n int64) (string) {
	nBig := new(big.Int).SetInt64(n)
	return padBigTo32Bytes(nBig)
}

// set w.Data
func setWithdrawalData(w *Withdrawal) (*Withdrawal) {
		valueBytes := w.Value.Bytes()
		valueString := hex.EncodeToString(valueBytes)
		valueString = padTo32Bytes(valueString)
		if len(valueString) != 64 {
			fmt.Println("value formatted incorrectly")
		}
		w.Data = w.Recipient + valueString + w.FromChain
		return w
}

// find the index in allChains of a chain with a particular Id
// return index i if chain in allChains, otherwise return -1
func findChainIndex(id *big.Int, allChains []*Chain) int {
	for i, chain := range allChains {
		if chain.Id.Cmp(id) == 0 { return i }
	}
	return -1
}

func FindChain(id *big.Int, allChains []*Chain) (*Chain) {
	for _, chain := range allChains {
		if chain.Id.Cmp(id) == 0 { return chain }
	}
	return nil
}

/***** client functions ******/

func Filter(chain *Chain, allChains []*Chain, filter *ethereum.FilterQuery, logsDone chan bool) {
	logs, err := chain.Client.FilterLogs(context.Background(), *filter)
	if err != nil {
		fmt.Println(err)
	}

	if len(logs) != 0 {
		//fmt.Println(len(logs))
		go ReadLogs(chain, allChains, logs, logsDone)
	}

	logsDone <- true
}

func ReadLogs(chain *Chain, allChains []*Chain, logs []types.Log, logsDone chan bool) {
	//logs := <-ch
	//fmt.Println(logs)
	for _, log := range logs {
		txHash := log.TxHash.Hex()
		if(!logsRead[txHash]) {
			fmt.Println("\nlogs found on chain", chain.Id, "at block", log.BlockNumber)
			fmt.Println("contract address: ", log.Address.Hex())
			for _, topics := range log.Topics {
				topic := topics.Hex()
				fmt.Println("topics: ", topic)

				if strings.Compare(topic, events.DepositId) == 0 { 
					fmt.Println("*** deposit event")
					fmt.Println("txHash: ", txHash)
					withdrawDone := make(chan bool)
					go HandleDeposit(chain, allChains, log.TxHash, withdrawDone)
					<-withdrawDone
			 	} else if strings.Compare(topic, events.CreationId) == 0 {
					fmt.Println("*** bridge contract creation")
				} else if strings.Compare(topic, events.WithdrawId) == 0 {
					fmt.Println("*** withdraw event")
					txHash := log.TxHash.Hex()
					fmt.Println("txHash: ", txHash)
					// receiver, value, toChain := readDepositData(data)
					// fmt.Println("receiver: ", receiver) 
					// fmt.Println("value: ", value) // in hexidecimal
					// fmt.Println("to chain: ", toChain) // in hexidecimal
				} else if strings.Compare(topic, events.BridgeSetId) == 0 {
					fmt.Println("*** set bridge event")
					fmt.Println("txHash: ", txHash)
				} else if strings.Compare(topic, events.BridgeFundedId) == 0 {
					fmt.Println("*** funded bridge event")
					fmt.Println("txHash: ", txHash)
				}
			}
			logsRead[txHash] = true
		}
	}
	logsDone <- true
}

func HandleDeposit(chain *Chain, allChains []*Chain, txHash common.Hash, withdrawDone chan bool) {
	tx, isPending, err := chain.Client.TransactionByHash(context.Background(), txHash)
	if isPending {
		// wait
	}
	if err != nil {
		fmt.Println(err)
	}

	withdrawal := new(Withdrawal)

	data := hex.EncodeToString(tx.Data())
	//fmt.Println("data: ", data)
	//fmt.Println(len(data))
	if len(data) > 72 {
		receiver := data[32:72];
		toChain := data[72:136]
		value := tx.Value()
		// receiver, value, toChain := readDepositData(data)
		fmt.Println("receiver: ", receiver) 
		fmt.Println("value: ", value) // in hexidecimal
		fmt.Println("to chain: ", toChain) // in hexidecimal

		withdrawal.Recipient = data[32:72]
		withdrawal.FromChain = toChain
		withdrawal.Value = value

		fromChain := new(big.Int)
		fromChain.SetString(toChain, 16)
		fmt.Println("chain to withdraw to: ", fromChain)
		//fmt.Println(fromChain)
		//chainIndex := IdsToChainIndex[fromChain]
		//fmt.Println("chain to withdraw to: ", allChains[fromChain])
		// idx := findChainIndex(chain.Id, allChains)
		// fmt.Println("deposit chain id: ", chain.Id, allChains[idx])

		idx := findChainIndex(fromChain, allChains)
		//fmt.Println("withdraw chain id: ", fromChain, allChains[idx])

		if idx == -1 {
			fmt.Println("could not find chain to withdraw to")
		} else {
			Withdraw(allChains[idx], withdrawal)
		}
	}
	withdrawDone <- true
}

func FundPrompt(chain *Chain) {
	var value int64
	var confirm int64
	fmt.Println("\nfunding the bridge contract on chain", chain.Id)
	fmt.Println("note that funding of the bridge cannot be withdrawn")
	fmt.Println("enter value of funding, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	valBig := big.NewInt(value)
	fmt.Println("confirm funding on chain", chain.Id, "with value", value, "wei")
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}
	FundBridge(chain, valBig)
}

func DepositPrompt(chain *Chain) {
	var value int64
	var to int64
	var confirm int64
	fmt.Println("\ndepositing to the bridge contract on chain", chain.Id)
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of deposit, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	fmt.Println("enter chain id to withdraw on")
	fmt.Scanln(&to)
	if to == -1 { 
		return
	}

	valBig := big.NewInt(value)

	toHex := fmt.Sprintf("%x", to)
	//fmt.Println(toHex)
	fmt.Println("confirm deposit on chain", chain.Id, "with value", value, "wei, withdrawing to chain", to)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}
	Deposit(chain, valBig, toHex)
}

func WithdrawToPrompt(chain *Chain) {
	var value int64
	var to int64
	var confirm int64
	fmt.Println("\nwithdrawing to other chains from the bridge contract on chain", chain.Id)
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of withdraw, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	fmt.Println("enter chain id to withdraw on")
	fmt.Scanln(&to)
	if to == -1 { 
		return
	}

	fmt.Println("confirm deposit on chain", chain.Id, "with value", value, "wei, withdrawing to chain", to)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	valBig := big.NewInt(value)
	toHex := fmt.Sprintf("%x", to)
	WithdrawTo(chain, valBig, toHex)
}

func PayBridgePrompt(chain *Chain) {
	var value int64
	var confirm int64
	fmt.Println("\npaying bridge contract on chain", chain.Id)
	fmt.Println("note that bridge payments can later be withdrawn")
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of payment, in wei")
	fmt.Scanln(&value)
	if value == -1 {
		return
	}

	fmt.Println("confirm payment to bridge on chain", chain.Id, "with value", value, "wei")
	fmt.Scanln(&confirm)
	if confirm == -1 {
		return
	}

	valBig := big.NewInt(value)
	PayBridge(chain, valBig)
}

func Prompt(chain *Chain, ks *keystore.KeyStore, fl map[string]bool, donePrompt chan bool) {
	keys = ks
	flags = fl

	/* prompt for fund bridge info */
	if flags["fund"] {
		FundPrompt(chain)
	}

	/* prompt for deposit info */
	if flags["deposit"] {
		DepositPrompt(chain)
	}

	if flags["pay"] {
		PayBridgePrompt(chain)
	}

	if flags["withdraw"] {
		WithdrawToPrompt(chain)
	}

	//donePrompt.Done()
	donePrompt <- true
}
// main goroutine
// starts a client to listen on every chain 
func Listen(chain *Chain, ac []*Chain, e *Events, doneClient chan bool, ks *keystore.KeyStore, fl map[string]bool) {
	// set up global vars
	events = e
	keys = ks
	flags = fl
	allChains := ac

	// dial client
	client, err := ethclient.Dial(chain.Url)
	if err != nil {
		log.Fatal(err)
	}
	chain.Client = client

	fmt.Println("listening at: " + chain.Url)

	fromBlock := big.NewInt(1)
	filter := new(ethereum.FilterQuery)

	// every second, check for new logs and update block number
	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C{
		if flags["v"] { fmt.Println(t) }

		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			//log.Fatal(err)
			//fmt.Println("could not get block with ethclient.. trying http request")
			blockNum, err := getBlockNumber(chain.Url)
			if err != nil {
				log.Fatal(err)
			}
			if flags["v"] { fmt.Println("latest block: ", blockNum) }
			fromBlock, _ = new(big.Int).SetString(blockNum[2:], 16)

			//fmt.Println("fromBlock: ", fromBlock)
		} else if fromBlock != block.Number() {
			if err != nil { log.Fatal(err) }
			if flags["v"] { fmt.Println("latest block: ", block.Number()) }
			fromBlock = block.Number()
		}

		filter.FromBlock = fromBlock
		if !flags["a"] {
			contractArr := make([]common.Address, 1)
			contractArr = append(contractArr, *chain.Contract)
			filter.Addresses = contractArr
		}
		logsDone := make(chan bool)
		go Filter(chain, allChains, filter, logsDone)
		<-logsDone
	}
 
	// bridge timeout. eventually, change so it never times out
	time.Sleep(6000 * time.Second)
	ticker.Stop()
	doneClient <- true
}