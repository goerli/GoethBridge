package main

import (
	"fmt"
	"net/http"
	"math/big"
	"bytes"
	"io/ioutil"
	//"time"
	"jsonparser"
	//"encoding/json"
	"encoding/hex"
	//"encoding/binary"
	"path/filepath"
	"strings"
	"log"
	"flag"
	//"sync"
	//"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/rlp"

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

/***** rpc methods ******/
// used initially when making http requests; can delete these later
// used for json format a response from an RPC call
type Resp struct {
	jsonrpc string
	id int
	result string
}

// used to json format an RPC call
type Call struct {
	Jsonrpc string 		`json:"jsonrpc"`
	Method string 		`json:"method"`
	Params []string 	`json:"params"`
	Id int 				`json:"id"`
}

// this function makes the rpc call "eth_getLogs" passing in jsonParams as the json formatted
// parameters to the call
// json parameters: [optional] fromBlock, toBlock
func getLogs(url string, jsonParams string, client *http.Client) (*http.Response, error) {
	jsonStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[` + jsonParams + `],"id":74}`
	jsonBytes := []byte(jsonStr)
	//fmt.Println(string(jsonBytes))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil { return nil, err }
	return resp, nil
}

// this function makes the rpc call "eth_getTransactionReceipt" passing in the txHash
func getTxReceipt(txHash string, url string, client *http.Client) (string, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":1}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println(string(jsonBytes))
    fmt.Println("getting tx receipt...")

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil { return "", err }

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	receipt, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
		return "", err
	}
	return receipt, nil
}

// this function calles the JSON RPC method "eth_getRawTransactionByHash" and returns the raw tx data if there is no error
func getRawTx(txData string, url string, client *http.Client) (string, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getRawTransactionByHash","params":["` + txData + `"],"id":1}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println("\n", string(jsonBytes))
    fmt.Println("getting raw tx data...")

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return "", err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil { return "", err }
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	rawTx, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
		return "", err
	}
	return rawTx, nil
}

func getTxCount(txData string, url string, client *http.Client) (*http.Response, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":[` + txData + `],"id":1}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println(string(jsonBytes))
    fmt.Println("geting nonce...")

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    //fmt.Println(resp)
    if err != nil { return nil, err }
    return resp, nil
}

func sendTx(txData string, url string, client *http.Client) (*http.Response, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[` + txData + `],"id":1}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println("\n", string(jsonBytes))
    fmt.Println("sending tx with params: ", txData)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    //fmt.Println(resp)
    if err != nil { return nil, err }
    return resp, nil
}

func sendRawTx(txData string, url string, client *http.Client) (*http.Response, error) {
	fmt.Println("sending raw transaction with data: ", txData)
    jsonStr := `{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["` + txData + `"],"id":1}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println("\n", string(jsonBytes))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    //fmt.Println(resp)
    if err != nil { return nil, err }
    return resp, nil
}

// this function gets the current block number by calling "eth_blockNumber"
func getBlockNumber(url string, client *http.Client) (string, error) {
	var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	blockNumResp, err := client.Do(req)
	if err != nil {
       	return "", err
	}
	defer blockNumResp.Body.Close()

	// print out response of eth_blockNumber
	//fmt.Println("response Status:", blockNumResp.Status)
	//fmt.Println("response Headers:", blockNumResp.Header)
	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)
	//fmt.Println("responnse Body:", string(blockNumBody))

	// parse json for result
	startBlock, err := parseJsonForResult(string(blockNumBody))
	if err != nil {
		return "", nil
	}
	return startBlock, nil
}

func getNonce(address []byte, url string) (string) {
	// get nonce
	client := &http.Client{}
	txCountData := fmt.Sprintf("\"0x%x\", \"latest\"", address)
	//fmt.Println(txCountData)
	resp, err := getTxCount(txCountData, url, client)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	nonce, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
	}
	//fmt.Println("account nonce: ", nonce)
	//nonceBytes, err := hex.DecodeString(nonce[2:len(nonce)])
	return nonce
}

/*****  helpers *****/
// this function parses jsonStr for the result entry and returns its value as a string
func parseJsonForResult(jsonStr string) (string, error) {
	jsonBody := []byte(string(jsonStr))
	res, _, _, err := jsonparser.Get(jsonBody, "result")
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// this function parses jsonStr for the entry "get" and returns its value as a string
func parseJsonForEntry(jsonStr string, get string) (string, error) {
	jsonBody := []byte(string(jsonStr))
	res, _, _, err := jsonparser.Get(jsonBody, get)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func rlpDecodeTx(rawTxData string) (*types.Transaction) {
	tx := new(types.Transaction)
    rawtx,err := hex.DecodeString(rawTxData)
    if err != nil {
    	fmt.Println(err)
    }
    rlp.DecodeBytes(rawtx, &tx)
	return tx
}

func readAbi() (*client.Events) {
	e := new(client.Events)

	// read bridge contract abi
	path, _ := filepath.Abs("./truffle/build/contracts/Bridge.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
	    fmt.Println("Failed to read file:", err)
	}

	fileAbi, err := parseJsonForEntry(string(file), "abi")
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
	fmt.Println("set bridge event id: ", e.BridgeFundedId)
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

		chainStr, err := parseJsonForEntry(string(file), chain)
		if err != nil {
			fmt.Println("could not find chain in config file")
			log.Fatal(err)
		}

		contractAddr, err := parseJsonForEntry(chainStr, "contractAddr")
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

		url, err := parseJsonForEntry(chainStr, "url")
		if err != nil {
			fmt.Println("could not find url in config file")
			log.Fatal(err)
		}
		fmt.Println("url of chain", chain, ":", url)
		clients[i].Url = url

		gp, err := parseJsonForEntry(chainStr, "gasPrice")
		if err != nil {
			fmt.Println("could not find gas price in config file")
			log.Fatal(err)
		}
		bigGas := new(big.Int)
		bigGas.SetString(gp, 10)
		clients[i].GasPrice = bigGas

		fromAccount, err := parseJsonForEntry(chainStr, "from")
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