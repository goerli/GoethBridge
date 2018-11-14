package main

import (
	"fmt"
	"math/big"
	"io/ioutil"
	"encoding/json"
	"encoding/hex"
	"path/filepath"
	"strings"
	"log"
	"flag"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/goerli/GoethBridge/client"
)

/* global vars */
var flags map[string]bool
var ks *keystore.KeyStore

type Config struct {
	Chain map[string]*Chain  `json:"networks"`
} 

type Chain struct {
	Name string 						`json:"name"`
	Url string 							`json:"url"`
	Id *big.Int 						`json:"id,omitempty"`
	Contract string 		 			`json:"contractAddr"`
	GasPrice *big.Int					`json:"gasPrice"`
	From string 		 				`json:"from"`
	Password string 					`json:"password,omitempty"`
	StartBlock int 		 				`json:"startBlock,omitempty"`
}

/****** keystore methods ******/
func newKeyStore(path string) (*keystore.KeyStore) {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}

func readAbi(verbose bool) (*client.Events) {
	e := new(client.Events)

	// read bridge contract abi
	path, _ := filepath.Abs("./leth/build/Bridge.abi")
	file, err := ioutil.ReadFile(path)
	if err != nil {
	    fmt.Println("Failed to read file:", err)
	}

	bridgeabi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
	    fmt.Println("Invalid abi:", err)
	}

	// checking abi for events
	bridgeEvents := bridgeabi.Events
	depositEvent := bridgeEvents["Deposit"]
	depositHash := depositEvent.Id()
	e.DepositId = depositHash.Hex()
	if verbose { fmt.Println("deposit event id: ", e.DepositId) }

	creationEvent := bridgeEvents["ContractCreation"]
	creationHash := creationEvent.Id()
	e.CreationId = creationHash.Hex()
	if verbose { fmt.Println("contract creation event id: ", e.CreationId) }

	withdrawEvent := bridgeEvents["Withdraw"]
	withdrawHash := withdrawEvent.Id()
	e.WithdrawId = withdrawHash.Hex()
	if verbose { fmt.Println("withdraw event id: ", e.WithdrawId) }

	bridgeSetEvent := bridgeEvents["BridgeSet"]
	bridgeSetHash := bridgeSetEvent.Id()
	e.BridgeSetId = bridgeSetHash.Hex()
	if verbose { fmt.Println("set bridge event id: ", e.BridgeSetId) }

	bridgeFundedEvent := bridgeEvents["BridgeFunded"]
	bridgeFundedHash := bridgeFundedEvent.Id()
	e.BridgeFundedId = bridgeFundedHash.Hex()
	if verbose { fmt.Println("bridge funded event id: ", e.BridgeFundedId) }

	paidEvent := bridgeEvents["Paid"]
	paidHash := paidEvent.Id()
	e.PaidId = paidHash.Hex()
	if verbose { fmt.Println("bridge paid event id", e.PaidId) }
	return e
}

func startup(id *big.Int) (*big.Int) {
	path, _ := filepath.Abs("./log/" + id.String() + "_lastblock.txt")
	file, err := ioutil.ReadFile(path)
	if err != nil {
	    fmt.Println("Failed to read file:", err)
	}
	startBlock := new(big.Int)
	startBlock.SetString(string(file), 10)
	return startBlock
}

func printHeader() {
	fmt.Println("██████╗ ██████╗ ██╗██████╗  ██████╗ ███████╗")
	fmt.Println("██╔══██╗██╔══██╗██║██╔══██╗██╔════╝ ██╔════╝")
	fmt.Println("██████╔╝██████╔╝██║██║  ██║██║  ███╗█████╗  ")
	fmt.Println("██╔══██╗██╔══██╗██║██║  ██║██║   ██║██╔══╝  ")
	fmt.Println("██████╔╝██║  ██║██║██████╔╝╚██████╔╝███████╗")
	fmt.Println("╚═════╝ ╚═╝  ╚═╝╚═╝╚═════╝  ╚═════╝ ╚══════╝")
}

func main() {
	/* flags */
	headerPtr := flag.Bool("header", true, "a bool representing whether to print out the header or not")
	verbosePtr := flag.Bool("v", false, "a bool representing verbosity of output")
	readAllPtr := flag.Bool("a", false, "a bool representing whether to read logs from every contract or not")
	configPtr := flag.String("config", "./config.json", "a string of the path to the config file") 
	keysPtr := flag.String("keystore", "./keystore", "a string of the path to the keystore directory") 
	// password flag assumes you have the same account on every chain
	passwordPtr := flag.String("password", "password", "a string of the password to the account specified in the config file") 
	noListenPtr := flag.Bool("no-listen", false, "a bool; if true, do not start the listener")

	depositCommand := flag.NewFlagSet("deposit", flag.ExitOnError)
	fundCommand := flag.NewFlagSet("fund", flag.ExitOnError)
	payCommand := flag.NewFlagSet("payCommand", flag.ExitOnError)
	withdrawCommand := flag.NewFlagSet("withrawCommand", flag.ExitOnError)

	// subcommands
	if len(os.Args) > 1 {
		switch os.Args[1]{
			case "deposit":
				depositCommand.Parse(os.Args[2:])
			case "fund":
				fundCommand.Parse(os.Args[2:])
			case "pay":
				payCommand.Parse(os.Args[2:])
			case "withdraw":
				withdrawCommand.Parse(os.Args[2:])
			default:
				// continue
		}
	}

	flag.Parse()
	header := *headerPtr
	if header { printHeader() }

	configStr := *configPtr
	fmt.Println("config path: ", configStr)

	verbose := *verbosePtr
	if verbose { fmt.Println("verbose: ", verbose) }

	readAll := *readAllPtr
	if readAll { fmt.Println("read from all contracts? ", readAll)}

	keystorePath := *keysPtr
	fmt.Println("keystore path: ", keystorePath)

	password := *passwordPtr
	noListen := *noListenPtr

	var chains []string
	if depositCommand.Parsed() {
		chains = depositCommand.Args()
		fmt.Println("deposit to:", chains)
	} else if fundCommand.Parsed() {
		chains = fundCommand.Args()
		fmt.Println("fund bridge on chains", chains)
	} else if payCommand.Parsed() {
		chains = payCommand.Args()
		fmt.Println("pay bridge on chains", chains)
	} else if withdrawCommand.Parsed() {
		chains = withdrawCommand.Args()
		fmt.Println("withdraw from bridge on chains", chains)
	} else {
		chains = flag.Args()
		if len(chains) == 0 {
			chains = append(chains,"1")
		}
		fmt.Println("chains to connect to: ", chains)
	}

	flags = make(map[string]bool)
	flags["v"] = verbose
	flags["a"] = readAll
	flags["nolisten"] = noListen

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

	// unmarshal config
	config := new(Config)
	err = json.Unmarshal(file, config)
	if err != nil {
		log.Fatal("could not unmarshal config: ", err)
	}

	// read config file for each chain id
	for i, name := range chains {
		if _, ok := config.Chain[name]; ok {
			 // continue  
		} else {
			log.Fatal("could not find chain ", name)
			os.Exit(1)
		}

		clients[i] = new(client.Chain)
		clients[i].Id = config.Chain[name].Id
		clients[i].Name = name 

		// to start at block 0, `rm -rf log/`
		startBlock := startup(clients[i].Id)
		clients[i].StartBlock = startBlock

		contractAddr := config.Chain[name].Contract
		fmt.Println("contract address of chain", name, ":", contractAddr)
		contract := new(common.Address)
		contractBytes, err := hex.DecodeString(contractAddr[2:])
		if err != nil {
			log.Fatal(err)
		}
		contract.SetBytes(contractBytes)
		clients[i].Contract = contract

		url := config.Chain[name].Url
		fmt.Println("url of chain", name, ":", url)
		clients[i].Url = url

		gasPrice := config.Chain[name].GasPrice
		clients[i].GasPrice = gasPrice

		fromAccount := config.Chain[name].From
		fmt.Println("account to send txs from on chain", name, ":", fromAccount)
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

	/* read abi of contract in truffle folder */
	events := readAbi(flags["v"])

	if depositCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			client.DepositPrompt(chain, ks)
		}
		return
	} else if fundCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				log.Fatal("chain not found in config")
			}
			client.FundPrompt(chain, ks)
		}
		return
	} else if payCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			client.PayBridgePrompt(chain, ks)
		}
		return
	} else if  withdrawCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			client.WithdrawToPrompt(chain, ks)
		}
		return	
	}

	/* channels */
	doneClient := make(chan bool)

	/* wait group for interrupt handling */
	wg := new(sync.WaitGroup)
	wg.Add(len(clients))

	if(!noListen) {
		/* listener */
		fmt.Println("\nlistening for events...")
		for _, chain := range clients {
			go client.Listen(chain, clients, events, doneClient, ks, flags, wg)
		}

		<-doneClient
	}
}