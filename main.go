package main

import (
	"fmt"
	//"net/http"
	"math/big"
	"io/ioutil"
	//"time"
	//"encoding/json"
	"encoding/hex"
	//"encoding/binary"
	"path/filepath"
	"strings"
	"log"
	"flag"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/accounts/abi"

    //"github.com/noot/multi_directional_bridge/rlp"
    //"github.com/noot/multi_directional_bridge/transaction"
    "github.com/noot/multi_directional_bridge/client"
)

/* global vars */
var flags map[string]bool
var ks *keystore.KeyStore

/****** keystore methods ******/
func newKeyStore(path string) (*keystore.KeyStore) {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}

func readAbi() (*client.Events) {
	e := new(client.Events)

	// read bridge contract abi
	path, _ := filepath.Abs("./truffle/build/contracts/Bridge.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
	    fmt.Println("Failed to read file:", err)
	}

	fileAbi, err := client.ParseJsonForEntry(string(file), "abi")
	if err != nil {
		log.Fatal(err)
	}

	bridgeabi, err := abi.JSON(strings.NewReader(fileAbi))
	if err != nil {
	    fmt.Println("Invalid abi:", err)
	}

	// checking abi for events
	bridgeEvents := bridgeabi.Events
	depositEvent := bridgeEvents["Deposit"]
	depositHash := depositEvent.Id()
	e.DepositId = depositHash.Hex()
	fmt.Println("deposit event id: ", e.DepositId) // this is the deposit event to watch for

	creationEvent := bridgeEvents["ContractCreation"]
	creationHash := creationEvent.Id()
	e.CreationId = creationHash.Hex()
	fmt.Println("contract creation event id: ", e.CreationId)

	withdrawEvent := bridgeEvents["Withdraw"]
	withdrawHash := withdrawEvent.Id()
	e.WithdrawId = withdrawHash.Hex()
	fmt.Println("withdraw event id: ", e.WithdrawId)

	bridgeSetEvent := bridgeEvents["BridgeSet"]
	bridgeSetHash := bridgeSetEvent.Id()
	e.BridgeSetId = bridgeSetHash.Hex()
	fmt.Println("set bridge event id: ", e.BridgeSetId)

	bridgeFundedEvent := bridgeEvents["BridgeFunded"]
	bridgeFundedHash := bridgeFundedEvent.Id()
	e.BridgeFundedId = bridgeFundedHash.Hex()
	fmt.Println("bridge funded event id: ", e.BridgeFundedId)
	return e
}

func main() {
	/* read abi of contract in truffle folder */
	events := readAbi()

	/* flags */
	verbosePtr := flag.Bool("v", false, "a bool representing verbosity of output")
	readAllPtr := flag.Bool("a", false, "a bool representing whether to read logs from every contract or not")
	configPtr := flag.String("config", "./config.json", "a string of the path to the config file") 
	keysPtr := flag.String("keystore", "./keystore", "a string of the path to the keystore directory") 
	// password flag assumes you have the same account on every chain
	passwordPtr := flag.String("password", "password", "a string of the password to the account specified in the config file") 

	fundBridgePtr := flag.Bool("fund", false, "a bool; if true, prompt user to fund bridge contract")
	depositPtr := flag.Bool("deposit", false, "a bool; if true, prompt user to deposit to bridge contract")

	flag.Parse()
	configStr := *configPtr
	fmt.Println("config path: ", configStr)

	verbose := *verbosePtr
	if verbose { fmt.Println("verbose: ", verbose) }

	readAll := *readAllPtr
	if readAll { fmt.Println("read from all contracts? ", readAll)}

	chains := flag.Args()
	if len(chains) == 0 {
		chains = append(chains,"33")
	}
	fmt.Println("chains to connect to: ", chains)

	keystorePath := *keysPtr
	fmt.Println("keystore path: ", keystorePath)

	password := *passwordPtr

	fundBridge := *fundBridgePtr
	deposit := *depositPtr

	flags = make(map[string]bool)
	flags["v"] = verbose
	flags["a"] = readAll
	flags["fund"] = fundBridge
	flags["deposit"] = deposit

	/* keys */
	ks = newKeyStore(keystorePath)
	ksaccounts := ks.Accounts()
	for i, account := range ksaccounts {
		if verbose { fmt.Println("account", i, ":", account.Address.Hex()) }
	}

	// config file reading
	path, _ := filepath.Abs(configStr)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read file:", err)	
	}

	clients := make([]*client.Chain, len(chains))
	// read config file for each chain id
	for i, chain := range chains {
		clients[i] = new(client.Chain)

		clients[i].Id = new(big.Int)
		clients[i].Id.SetString(chain, 10)

		chainStr, err := client.ParseJsonForEntry(string(file), chain)
		if err != nil {
			fmt.Println("could not find chain in config file")
			log.Fatal(err)
		}

		contractAddr, err := client.ParseJsonForEntry(chainStr, "contractAddr")
		if err != nil {
			fmt.Println("could not find contractAddr in config file")
			log.Fatal(err)
		}
		fmt.Println("contract address of chain", chain, ":", contractAddr)
		contract := new(common.Address)
		contractBytes, err := hex.DecodeString(contractAddr[2:])
		if err != nil {
			log.Fatal(err)
		}
		contract.SetBytes(contractBytes)
		clients[i].Contract = contract

		url, err := client.ParseJsonForEntry(chainStr, "url")
		if err != nil {
			fmt.Println("could not find url in config file")
			log.Fatal(err)
		}
		fmt.Println("url of chain", chain, ":", url)
		clients[i].Url = url

		gp, err := client.ParseJsonForEntry(chainStr, "gasPrice")
		if err != nil {
			fmt.Println("could not find gas price in config file")
			log.Fatal(err)
		}
		bigGas := new(big.Int)
		bigGas.SetString(gp, 10)
		clients[i].GasPrice = bigGas

		fromAccount, err := client.ParseJsonForEntry(chainStr, "from")
		if err != nil {
			fmt.Println("could not find from account in config file")
			log.Fatal(err)
		}
		fmt.Println("account to send txs from on chain", chain, ":", fromAccount)
		from := new(common.Address)
		fromBytes, err := hex.DecodeString(fromAccount[2:])
		if err != nil {
			log.Fatal(err)
		}
		from.SetBytes(fromBytes)
		clients[i].From = from

		clients[i].Password = password

		/* unlock account */
		// if(ks.HasAddress(*from)) {
		// 	account := new(accounts.Account)
		// 	account.Address = *from
		// 	err = ks.Unlock(*account, password)
		// 	if err != nil {
		// 		fmt.Println("could not unlock account")
		// 		fmt.Println(err)
		// 	} else {
		// 		log.Fatal("account not found in keystore")
		// 	}
		// }
	}

	for _, chain := range clients {
		/* dial client */
		chainClient, err := ethclient.Dial(chain.Url)
		if err != nil {
			log.Fatal(err)
		}
		chain.Client = chainClient
	}

	/* channels */
	doneClient := make(chan bool)
	donePrompt := make(chan bool)

	/* prompt, if flags set */
	for _, chain := range clients {
		/* prompt if flags set & listen */
		go client.Prompt(chain, ks, flags, donePrompt)
		<-donePrompt
	}

	/* listener */
	fmt.Println("\nlistening for events...")
	for _, chain := range clients {
		go client.Listen(chain, clients, events, doneClient, ks, flags)
	}

	<-doneClient
}