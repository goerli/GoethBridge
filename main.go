package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ChainSafeSystems/ChainBridge/client"
	"github.com/ChainSafeSystems/ChainBridge/logger"
)

/* global vars */
var flags map[string]bool
var ks *keystore.KeyStore

type Config struct {
	Chain map[string]*Chain `json:"networks"`
}

type Chain struct {
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	Id         *big.Int `json:"id,omitempty"`
	Contract   string   `json:"contractAddr"`
	GasPrice   *big.Int `json:"gasPrice"`
	From       string   `json:"from"`
	Password   string   `json:"password,omitempty"`
	StartBlock int      `json:"startBlock,omitempty"`
}

/****** keystore methods ******/
func newKeyStore(path string) *keystore.KeyStore {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}

func readAbi(verbose bool) *client.Events {
	e := new(client.Events)

	// read bridge contract abi
	path, _ := filepath.Abs("./leth/build/Bridge.abi")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("Failed to read file:", err)
	}

	bridgeabi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		logger.Error("Invalid abi:", err)
	}

	// checking abi for events
	bridgeEvents := bridgeabi.Events
	depositEvent := bridgeEvents["Deposit"].Id().Hex()
	e.DepositId = depositEvent
	if verbose {
		logger.Info("deposit event id: %s", e.DepositId)
	}

	creationEvent := bridgeEvents["ContractCreation"].Id().Hex()
	e.CreationId = creationEvent
	if verbose {
		logger.Info("contract creation event id: %s", e.CreationId)
	}

	withdrawEvent := bridgeEvents["Withdraw"]
	withdrawHash := withdrawEvent.Id()
	e.WithdrawId = withdrawHash.Hex()
	if verbose {
		logger.Info("withdraw event id: %s", e.WithdrawId)
	}

	bridgeFundedEvent := bridgeEvents["BridgeFunded"]
	bridgeFundedHash := bridgeFundedEvent.Id()
	e.BridgeFundedId = bridgeFundedHash.Hex()
	if verbose {
		logger.Info("bridge funded event id: %s", e.BridgeFundedId)
	}

	paidEvent := bridgeEvents["Paid"]
	paidHash := paidEvent.Id()
	e.PaidId = paidHash.Hex()
	if verbose {
		logger.Info("bridge paid event id: %s", e.PaidId)
	}

	addAuthEvent := bridgeEvents["AuthorityAdded"]
	addAuthHash := addAuthEvent.Id()
	e.AuthorityAddedId = addAuthHash.Hex()
	if verbose {
		logger.Info("added authority id: %s", e.AuthorityAddedId)
	}

	removeAuthEvent := bridgeEvents["AuthorityRemoved"]
	removeAuthHash := removeAuthEvent.Id()
	e.AuthorityRemovedId = removeAuthHash.Hex()
	if verbose {
		logger.Info("removed authority id: %s", e.AuthorityRemovedId)
	}

	thresholdEvent := bridgeEvents["ThresholdUpdated"]
	thresholdHash := thresholdEvent.Id()
	e.ThresholdUpdated = thresholdHash.Hex()
	if verbose {
		logger.Info("threshold updated id: %s", e.ThresholdUpdated)
	}

	signedEvent := bridgeEvents["SignedForWithdraw"].Id().Hex()
	e.SignedForWithdraw = signedEvent
	if verbose {
		logger.Info("signed for withdraw id: %s", e.SignedForWithdraw)
	}

	return e
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

func startup(id *big.Int) *big.Int {
	log_exists, err := exists("log")
	if err != nil {
		logger.Error("%s", err)
	}
	if !log_exists {
		logger.Info("creating log/ directory...")	
		os.Mkdir("./log", os.ModePerm)
	} 

	path, _ := filepath.Abs("./log/" + id.String() + "_lastblock.txt")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Warn("%s", err)
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
	headerPtr := flag.Bool("header", false, "a bool representing whether to print out the header or not")
	verbosePtr := flag.Bool("v", false, "a bool representing verbosity of output")
	readAllPtr := flag.Bool("a", false, "a bool representing whether to read logs from every contract or not")
	configPtr := flag.String("config", "./config.json", "a string of the path to the config file")
	keysPtr := flag.String("keystore", "./keystore", "a string of the path to the keystore directory")
	// password flag assumes you have the same account on every chain
	passwordPtr := flag.String("password", "password", "a string of the password to the account specified in the config file")
	noListenPtr := flag.Bool("no-listen", false, "a bool; if true, do not start the listener")

	/* bridge subcommands */
	depositCommand := flag.NewFlagSet("deposit", flag.ExitOnError)
	fundCommand := flag.NewFlagSet("fund", flag.ExitOnError)
	payCommand := flag.NewFlagSet("pay", flag.ExitOnError)
	withdrawCommand := flag.NewFlagSet("withdrawto", flag.ExitOnError)

	/* admin subcommands */
	addAuthorityCommand := flag.Bool("addauth", false, "add authority")
	removeAuthorityCommand := flag.Bool("removeauth", false, "remove authority")
	increaseThresholdCommand := flag.Bool("incauth", false, "increase threshold")
	decreaseThresholdCommand := flag.Bool("decauth", false, "decrease threshold")

	// subcommands
	if len(os.Args) > 1 {
		switch os.Args[1] {
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
	if header {
		printHeader()
	}

	configStr := *configPtr
	logger.Info("config path: %s", configStr)
	verbose := *verbosePtr
	readAll := *readAllPtr
	if readAll {
		logger.Info("read from all contracts? %s", readAll)
	}

	keystorePath := *keysPtr
	logger.Info("keystore path: %s", keystorePath)

	password := *passwordPtr
	noListen := *noListenPtr

	var isSubCommandParsed [4]bool
	isSubCommandParsed[0] = depositCommand.Parsed()
	isSubCommandParsed[1] = fundCommand.Parsed()
	isSubCommandParsed[2] = payCommand.Parsed()
	isSubCommandParsed[3] = withdrawCommand.Parsed()

	var subCommandArgs [4][]string
	subCommandArgs[0] = depositCommand.Args()
	subCommandArgs[1] = fundCommand.Args()
	subCommandArgs[2] = payCommand.Args()
	subCommandArgs[3] = withdrawCommand.Args()

	var chains []string

	/*
	  Loop through arguments for the subcommand that is parsed, extract the password for either format --password="pass" or --password pass
	  Return the paramester before the index of the password -> [chains]
	*/
	var commandsNotParsed = 0

	for commandIndex, subCommand := range isSubCommandParsed {
		if subCommand {
			for paramIndex, param := range subCommandArgs[commandIndex] {
				if strings.Contains(param, "--password") {
					/*
						Check if the index of the password flag == same length of all subcommand parameters
						If == => --password="keystorePassword"
						else if password flag index == length - 1 =>  --password keystorePassword
					*/
					if paramIndex == len(subCommandArgs[commandIndex])-1 {
						password = subCommandArgs[commandIndex][paramIndex][11:len(subCommandArgs[commandIndex][paramIndex])]
					} else {
						password = subCommandArgs[commandIndex][paramIndex+1]
					}
					chains = subCommandArgs[commandIndex][0:paramIndex]
					break

				} else {
					chains = subCommandArgs[commandIndex]
					password = *passwordPtr
				}
			}
		} else {
			commandsNotParsed++
		}
	}

	if commandsNotParsed == len(isSubCommandParsed) {
		chains = flag.Args()
		if len(chains) == 0 {
			chains = append(chains, "1")
		}
	}

	flags = make(map[string]bool)
	flags["v"] = verbose
	flags["a"] = readAll
	flags["nolisten"] = noListen

	/* keys */
	ks = newKeyStore(keystorePath)
	ksaccounts := ks.Accounts()
	for i, account := range ksaccounts {
		if verbose {
			logger.Info("account %d: %s", i, account.Address.Hex())
		}
	}

	// config file reading
	path, _ := filepath.Abs(configStr)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.FatalError("Failed to read file: %s", err)
	}

	clients := make([]*client.Chain, len(chains))

	// unmarshal config
	config := new(Config)
	err = json.Unmarshal(file, config)
	if err != nil {
		logger.FatalError("could not unmarshal config: %s", err)
	}

	// read config file for each chain id
	for i, name := range chains {
		if _, ok := config.Chain[name]; ok {
			// continue
		} else {
			logger.FatalError("could not find chain %s", name)
		}

		clients[i] = new(client.Chain)
		clients[i].Id = config.Chain[name].Id
		clients[i].Name = name

		// to start at block 0, `rm -rf log/`
		startBlock := startup(clients[i].Id)
		clients[i].StartBlock = startBlock

		contractAddr := config.Chain[name].Contract
		logger.Info("contract address of chain %s: %s", name, contractAddr)
		contract := new(common.Address)
		contractBytes, err := hex.DecodeString(contractAddr[2:])
		if err != nil {
			logger.FatalError("%s", err)
		}
		contract.SetBytes(contractBytes)
		clients[i].Contract = contract

		url := config.Chain[name].Url
		logger.Info("url of chain %s: %s", name, url)
		clients[i].Url = url

		gasPrice := config.Chain[name].GasPrice
		clients[i].GasPrice = gasPrice

		fromAccount := config.Chain[name].From
		logger.Info("account to send txs from on chain %s: %s", name, fromAccount)
		from := new(common.Address)
		fromBytes, err := hex.DecodeString(fromAccount[2:])
		if err != nil {
			logger.FatalError("%s", err)
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

	/* read abi of contract in leth/build */
	events := readAbi(flags["v"])

	if depositCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.DepositPrompt(chain, ks)
		}
		return
	} else if fundCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.FundPrompt(chain, ks)
		}
		return
	} else if payCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.PayBridgePrompt(chain, ks)
		}
		return
	} else if withdrawCommand.Parsed() {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.WithdrawToPrompt(chain, ks)
		}
		return
	} else if *addAuthorityCommand {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.AddAuthorityPrompt(chain, ks)
		}
		return		
	} else if *removeAuthorityCommand {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.RemoveAuthorityPrompt(chain, ks)
		}
		return		
	} else if *increaseThresholdCommand {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.IncreaseThresholdPrompt(chain, ks)
		}
		return		
	} else if *decreaseThresholdCommand {
		for _, name := range chains {
			chain := client.FindChainByName(name, clients)
			if chain == nil {
				logger.FatalError("chain not found in config")
			}
			client.DecreaseThresholdPrompt(chain, ks)
		}
		return				
	}

	/* channels */
	doneClient := make(chan bool)

	/* wait group for interrupt handling */
	wg := new(sync.WaitGroup)
	wg.Add(len(clients))

	if !noListen {
		/* listener */
		logger.Info("listening for events...")
		for _, chain := range clients {
			go client.Listen(chain, clients, events, doneClient, ks, flags, wg)
		}

		<-doneClient
	}
}
