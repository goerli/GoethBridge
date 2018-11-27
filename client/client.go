package client 

import (
	"time"
	"encoding/hex"
	"math/big"
	"context"
	"log"
	"strings"
	"os"
	"os/signal"
	"syscall"
	"sync"
	//"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/ChainSafeSystems/ChainBridge/logger"
)

/* global variables */
var events *Events // events to listen for
var keys *keystore.KeyStore // keystore; used to sign txs
var flags map[string]bool // command line flags
var logsRead = map[string]bool{}

type Chain struct {
	Name string 						`json:"name"`
	Url string 							`json:"url"`
	Id *big.Int 						`json:"id,omitempty"`
	Contract *common.Address 			`json:"contractAddr"`
	GasPrice *big.Int 					`json:"gasPrice"`
	From *common.Address 				`json:"from"`
	Password string 					`json:"password,omitempty"`
	Client *ethclient.Client 			`json:"client,omitempty"`
	Nonce uint64 						`json:"nonce,omitempty"`
	StartBlock *big.Int 				`json:"startBlock,omitempty"`
}

type Withdrawal struct {
	Recipient string
	Value *big.Int
	FromChain string
	TxHash string
	Data string
}

// events to listen for
type Events struct {
	DepositId string
  	CreationId string
 	WithdrawId string
	BridgeFundedId string
	PaidId string
	AuthorityAddedId string
	AuthorityRemovedId string
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
		logger.Warn("value formatted incorrectly")
	}
	w.Data = w.Recipient + valueString + w.FromChain + w.TxHash
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

func FindChainByName(name string, allChains []*Chain) (*Chain) {
	for _, chain := range allChains {
		if chain.Name == name {
			return chain
		}
	}
	return nil
}

/***** client functions ******/

func Filter(chain *Chain, allChains []*Chain, filter *ethereum.FilterQuery, logsDone chan bool) {
	logs, err := chain.Client.FilterLogs(context.Background(), *filter)
	if err != nil {
		logger.Error("%s", err)
	}

	if len(logs) != 0 {
		go ReadLogs(chain, allChains, logs, logsDone)
	}

	logsDone <- true
}

func ReadLogs(chain *Chain, allChains []*Chain, logs []types.Log, logsDone chan bool) {
	for _, log := range logs {
		txHash := log.TxHash.Hex()
		if(!logsRead[txHash]) {
			logsRead[txHash] = true

			logger.Event("logs found on %s at block %d", chain.Name, log.BlockNumber)
			logger.Event("contract address: %s", log.Address.Hex())

			for _, topics := range log.Topics {
				topic := topics.Hex()
				if strings.Compare(topic, events.DepositId) == 0 { 
					logger.Event("deposit event: tx hash: %s", txHash)
					withdrawDone := make(chan bool)
					go HandleDeposit(chain, allChains, log.TxHash, withdrawDone)
					<-withdrawDone
			 	} else if strings.Compare(topic, events.CreationId) == 0 {
					logger.Event("bridge contract creation")
				} else if strings.Compare(topic, events.WithdrawId) == 0 {
					logger.Event("withdraw event: tx hash: %s", txHash)
					printWithdraw(chain, log.TxHash)
				} else if strings.Compare(topic, events.AuthorityAddedId) == 0 {
					logger.Event("set bridge event: tx hash: %s", txHash)
				} else if strings.Compare(topic, events.BridgeFundedId) == 0 {
					logger.Event("funded bridge event: tx hash: %s", txHash)
				} else if strings.Compare(topic, events.PaidId) == 0 {
					logger.Event("bridge paid event: tx hash: %s", txHash)
				}
			}
		}
	}
	logsDone <- true
}

func waitOnPending(chain *Chain, txHash common.Hash) (*types.Transaction) {
	for {
		tx, isPending, _ := chain.Client.TransactionByHash(context.Background(), txHash)
		if !isPending { return tx }
	}
}

func printWithdraw(chain *Chain, txHash common.Hash) {
	tx := waitOnPending(chain, txHash)

	data := hex.EncodeToString(tx.Data())
	if len(data) > 136 {
		receiver := data[32:72];
		value := data[72:136]
		fromChain := data[136:200]

		logger.Event("receiver: 0x%s", receiver) 
		logger.Event("value: %s", value) // in hexidecimal
		logger.Event("from chain: %s", fromChain) // in hexidecimal
	}
}

func HandleDeposit(chain *Chain, allChains []*Chain, txHash common.Hash, withdrawDone chan bool) {
	tx := waitOnPending(chain, txHash)

	withdrawal := new(Withdrawal)
	data := hex.EncodeToString(tx.Data())

	if len(data) > 72 {
		receiver := data[32:72]
		toChain := data[72:136]
		value := tx.Value()

		//to, _ := strconv.Atoi(toChain)
		logger.Event("receiver: 0x%s", receiver) 
		logger.Event("value: %d", value) // in hexidecimal
		//logger.Event("to chain: %d", to) // in hexidecimal

		withdrawal.Recipient = data[32:72]
		withdrawal.FromChain = toChain
		withdrawal.Value = value
		withdrawal.TxHash = txHash.Hex()[2:]

		fromChain := new(big.Int)
		fromChain.SetString(toChain, 16)
		logger.Event("chain to withdraw to: %s", fromChain)

		idx := findChainIndex(fromChain, allChains)

		if idx == -1 {
			logger.Error("could not find chain to withdraw to")
		} else {
			err := Withdraw(allChains[idx], withdrawal)
			if err != nil {
				logger.Error("%s", err)
			}
		}
	}

	withdrawDone <- true
}

// main goroutine
// starts a client to listen on every chain 
func Listen(chain *Chain, ac []*Chain, e *Events, doneClient chan bool, ks *keystore.KeyStore, fl map[string]bool, wg *sync.WaitGroup) {
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

	logger.Info("listening at: %s", chain.Url)

	fromBlock := chain.StartBlock

	//lastBlocks[chain.Id] <- fromBlock
	logger.Info("starting block on %s: %s", chain.Name, fromBlock)
	filter := new(ethereum.FilterQuery)

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        Cleanup(chain, fromBlock, wg)
        wg.Wait()
        os.Exit(1)
    }()

	// every second, check for new logs and update block number
	for {
		filter.FromBlock = fromBlock

		// if not reading from all contracts, add the bridge contract address to the filter
		if !flags["a"] {
			contractArr := make([]common.Address, 1)
			contractArr = append(contractArr, *chain.Contract)
			filter.Addresses = contractArr
		}
		logsDone := make(chan bool)
		go Filter(chain, allChains, filter, logsDone)

		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil {
			// could not get block with ethclient.. trying http request
			blockNum, err := getBlockNumber(chain.Url)
			if err != nil {
				logger.Error("getBlockNumber error: %s", err)
			}
			if flags["v"] { 
				logger.Info("latest block on %s: %s", chain.Name, blockNum) 
			}
			fromBlock, _ = new(big.Int).SetString(blockNum[2:], 16)
		} else if fromBlock != block.Number() {
			if err != nil { 
				logger.Error("%s", err) 
			}
			if flags["v"] { 
				logger.Info("latest block on %s: %s", chain.Name, block.Number()) 
			}
			fromBlock = block.Number()
		}

		<-logsDone

		time.Sleep(1 * time.Second)
	}
 	doneClient <- true
}