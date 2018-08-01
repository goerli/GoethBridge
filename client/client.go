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
	//"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
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

// func mapIdsToChain(allChains []*Chain) {
// 	IdsToChainIndex = make(map[*big.Int]int)
// 	for i, chain := range allChains {
// 		IdsToChainIndex[chain.Id] = i
// 	}
// }

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
		fmt.Println("\nlogs found on chain", chain.Id, "at block", log.BlockNumber)
		fmt.Println("contract address: ", log.Address.Hex())
		for _, topics := range log.Topics {
			topic := topics.Hex()
			fmt.Println("topics: ", topic)

			txHash := log.TxHash.Hex()

			if(!logsRead[txHash]) {
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

				logsRead[txHash] = true
			}
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

/****** functions to send transactions ******/

func SetBridge(chain *Chain) () {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	dataStr := "8dd14802000000000000000000000000" + chain.Contract.Hex()[2:] // setbridge function signature + contract addr
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	chain.Nonce = nonce + 1

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, big.NewInt(int64(0)), uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		log.Fatal(err)
	}
	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to set bridge")
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

func Deposit(chain *Chain, value *big.Int, id string) {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	//dataStr := "0x47e7ef24000000000000000000000000ca35b7d915458ef540ade6068dfe2f44e8fa733c0000000000000000000000000000000000000000000000000000000000000003"
	//chainIdBytes := chain.Id.Bytes()
	//chainIdHex := hex.EncodeToString(chainIdBytes)

	chainId := padTo32Bytes(id)	
	dataStr := "47e7ef24000000000000000000000000" + chain.From.Hex()[2:] + chainId // deposit function signature + recipient addr + chain
	//fmt.Println(len(dataStr))
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	chain.Nonce = nonce 

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		log.Fatal(err)
	}
	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to deposit on chain", chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}

	//nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	//chain.Nonce = nonce + 1
}

func Withdraw(chain *Chain, withdrawal *Withdrawal) {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	//dataStr := "b5c5f672000000000000000000000000ca35b7d915458ef540ade6068dfe2f44e8fa733c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	withdrawal = setWithdrawalData(withdrawal)
	//fmt.Println(withdrawal.Data)
	dataStr := "b5c5f672000000000000000000000000" + withdrawal.Data // withdraw function signature + contract addr
	//fmt.Println(len(dataStr))
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, big.NewInt(int64(0)), uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		log.Fatal(err)
	}
	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to withdraw on chain", chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		//fmt.Println("could not send tx")
		//fmt.Println(err)
	}
}

func FundBridge(chain *Chain, value *big.Int) {
	client := chain.Client
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	data, err := hex.DecodeString("c9c0909f") //fund me function sig
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		log.Fatal(err)
	}
	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to fund bridge on chain", chain.Id, "with value", value.String())
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

// main goroutine
// starts a client to listen on every chain 
func Listen(chain *Chain, ac []*Chain, e *Events, doneClient chan bool, ks *keystore.KeyStore, fl map[string]bool) {
	// set up global vars
	events = e
	keys = ks
	flags = fl
	allChains := ac

	fmt.Println("listening at: " + chain.Url)

	// dial client
	client, err := ethclient.Dial(chain.Url)
	if err != nil {
		log.Fatal(err)
	}
	chain.Client = client

	//nonce, err := client.NonceAt(context.Background(), *chain.From, nil)
	//chain.Nonce = nonce 

	//SetBridge(chain)
	value := big.NewInt(7777777700)
	if chain.Id.Cmp(big.NewInt(4)) == 0 {
		Deposit(chain, value, "3")
	}
	// if chain.Id.Cmp(big.NewInt(3)) == 0 {
	// 	Deposit(chain, "4")
	// }
	//value := big.NewInt(1000000000000000)
	//FundBridge(chain, value)

	fromBlock := big.NewInt(1)
	filter := new(ethereum.FilterQuery)

	// every second, check for new logs and update block number
	ticker := time.NewTicker(1000 * time.Millisecond)
	for t := range ticker.C{
		if flags["v"] { fmt.Println(t) }

		block, err := client.BlockByNumber(context.Background(), nil)
		if fromBlock != block.Number() {
			if err != nil { log.Fatal(err) }
			if flags["v"] { fmt.Println("latest block: ", block.Number()) }
			fromBlock = block.Number()

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

		// d1 := block.Number().Bytes()
	 //    err = ioutil.WriteFile("./lastblock", d1, 0644)
	 //    if err != nil {
	 //    	fmt.Println(err)
	 //    }

	 //    path, _ := filepath.Abs("./lastblock")
		// file, err := ioutil.ReadFile(path)
		// if err != nil {
		//     fmt.Println("Failed to read file:", err)
		// }
		// fmt.Println(string(file))
	}
 
	// bridge timeout. eventually, change so it never times out
	time.Sleep(6000 * time.Second)
	ticker.Stop()
	doneClient <- true
}