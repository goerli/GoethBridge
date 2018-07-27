package main

import (
	"fmt"
	"net/http"
	"math/big"
	"bytes"
	"io/ioutil"
	"time"
	"jsonparser"
	"encoding/json"
	"encoding/hex"
	//"encoding/binary"
	"path/filepath"
	"strings"
	"log"
	"flag"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	GAS_LIMIT = 4600000
)

/****** keystore methods ******/
func newKeyStore(path string) (*keystore.KeyStore) {
	newKeyStore := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return newKeyStore
}
/***** rpc methods ******/

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

// used for getLogs json formatting
type LogParams struct {
	FromBlock string 	`json:"fromBlock"`
	Address string 		`json:"address,omitempty"`
}

type Tx struct {
	From string 		`json:"from"`
	To string 			`json:"to,omitempty"`
	Gas string			`json:"gas,omitempty"`
	GasPrice string		`json:"gasPrice,omitempty"`
	Value string 		`json:"value,omitempty"`
	Data string 		`json:"data,omitempty"`
	Nonce string 		`json:"nonce,omitempty"`
	V string  			`json:"v,omitempty"`
	R string 			`json:"r,omitempty"`
	S string 			`json:"s,omitempty"`
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

// this function makes the rpc call "eth_getTranscationReceipt" passing in the txHash
func getTxReceipt(txHash string, url string, client *http.Client) (*http.Response, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["` + txHash + `"],"id":74}`
    jsonBytes := []byte(jsonStr)
    //fmt.Println(string(jsonBytes))

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil { return nil, err }
    return resp, nil
}

func getTxCount(txData string, url string, client *http.Client) (*http.Response, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":[` + txData + `],"id":1}`
    jsonBytes := []byte(jsonStr)
    fmt.Println(string(jsonBytes))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    fmt.Println(resp)
    if err != nil { return nil, err }
    return resp, nil
}

func sendTx(txData string, url string, client *http.Client) (*http.Response, error) {
    jsonStr := `{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[` + txData + `],"id":1}`
    jsonBytes := []byte(jsonStr)
    fmt.Println(string(jsonBytes))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    if err != nil { return nil, err }
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    fmt.Println(resp)
    if err != nil { return nil, err }
    return resp, nil
}

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

/*****  helpers *****/
func readDepositData(data string) (string, string, string) {
	length := len(data)
	if length == 194 { // '0x' + 64 + 64 + 64
		recipient := "0x" + data[26:66]
		value := data[66:130]
		toChain :=  "0x" + data[130:194]
		return recipient, value, toChain
	} else {
		return "", "", ""
	}
}

/* global vars */
// flags
var verbose bool
var readAll bool
var chains []string
var contracts []string
var chainUrls []string
var gasPrices []*big.Int
// events to listen for
var DepositId string
var CreationId string
// keystore
var ks *keystore.KeyStore

func readAbi() {
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
	DepositId = depositHash.Hex()
	fmt.Println("deposit event id: ", DepositId) // this is the deposit event to watch for

	creationEvent := bridgeEvents["ContractCreation"]
	creationHash := creationEvent.Id()
	CreationId = creationHash.Hex()
	fmt.Println("contract creation event id: ", CreationId)
}

// starts a goroutine to listen on every chain 
func listen(urls []string, chains []string) {
	var params LogParams

	logsFound := make(map[string]bool)

	// poll filter every 500ms for changes
	ticker := time.NewTicker(100 * time.Millisecond)
	for i, _ := range urls {
		go func(chain string, url string) {
			client := &http.Client{}
			fmt.Println("listening at: " + url)

			go txClient(url, chain, gasPrices[i])

			for t := range ticker.C{
				if verbose { fmt.Println(t) }

				if !readAll { 
					params.Address = contracts[i]
				} 
				params.FromBlock, _ = getBlockNumber(url, client)
				if verbose { fmt.Println("getting logs from block number: " + params.FromBlock + "\n") }
				jsonParams, _ := json.Marshal(params)
	            //fmt.Println("jsonParams: " + string(jsonParams))

				//get logs from params.FromBlock
				resp, _ := getLogs(url, string(jsonParams), client)
				defer resp.Body.Close()

				//fmt.Println("response Status:", resp.Status)
				//fmt.Println("response Headers:", resp.Header)
				body, _ := ioutil.ReadAll(resp.Body)
				//fmt.Println("response Body:", string(body))
	 
				// parse for getLogs result
				//logsResult := parseJsonForResult(string(body))
				logsResult, err := parseJsonForEntry(string(body), "result")
				if err != nil {
					fmt.Println("could not parse logs")
					fmt.Println(err)
				}
				if verbose { fmt.Println("logsResult: " + logsResult + "\n") }
				//fmt.Println(len(logsResult))

				// if there are new logs, parse for event info
				if len(logsResult) > 2 {
					txHash, _ := parseJsonForEntry(logsResult[1:len(logsResult)-1], "transactionHash")
					//fmt.Println(txHash + "\n")
					if logsFound[txHash] != true { 
						logsFound[txHash] = true
						fmt.Println("\nnew logs found for chain", chain)

						//logs <- logsResult
						//readLogs(logs)
						//go readLogs(logs)
						//<-exit

						// get logs contract address
						address, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "address")
						if err != nil {
							fmt.Println(err)
						}
						// this is not actually a good way to listen for events from a  contract
						// this could be used to confirm a log, but for listening to events from
						// one contract, we would specify the address in our call to eth_getLogs
						fmt.Println("contract addr: ", address)
						//fmt.Println("length of address: ", len(address))
						for i := 0; i < len(chains); i++ {
							if strings.Compare(address[1:41], chains[i]) == 0 {
								fmt.Println("bridge contract event heard on chain ", chains[i])
							}
						}

						// read topics of log
						topics, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "topics")
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println("topics: ", topics[2:68])
						//fmt.Println("length of topics: ", len(topics)-4) len = 66: 0x + 64 hex chars = 32 bytes

						if strings.Compare(topics[2:68],DepositId) == 0 { 
							fmt.Println("*** deposit event ", topics[2:68])
							data, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "data")
							if err != nil {
								fmt.Println(nil)
							}
							//fmt.Println("length of data: ", len(data))
							//fmt.Println("data: ", data)
							receiver, value, toChain := readDepositData(data)
							fmt.Println("receiver: ", receiver) 
							fmt.Println("value: ", value) // in hexidecimal
							fmt.Println("to chain: ", toChain) // in hexidecimal
					 	} else if strings.Compare(topics[2:68],CreationId) == 0 {
							fmt.Println("*** bridge contract creation")
						}
					}
				}
			}
		}(chains[i], urls[i])
	} 

	// bridge timeout. eventually, change so it never times out
	time.Sleep(6000 * time.Second)
	ticker.Stop()
}

func getNonce(address []byte, url string) (string) {
	// get nonce
	client := &http.Client{}
	txCountData := fmt.Sprintf("\"0x%x\", \"latest\"", address)
	fmt.Println(txCountData)
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

func setBridge(address string, url string, chain string, gasPrice *big.Int) (string) {
	client := &http.Client{}
	accounts := ks.Accounts()

	//data := make([]byte, 32)
	//tx := types.NewTransaction(uint64(100), accounts[1].Address, big.NewInt(int64(0)), uint64(4600000), gasPrices[0], data)
	tx := new(Tx)
	tx.From = "0x" + hex.EncodeToString(accounts[0].Address[:])
	//tx.To = "0x" + hex.EncodeToString(accounts[1].Address[:])
	tx.To = "0x5fea67eb73c9e3edac55f22c8833bcc683b70d5d"
	tx.GasPrice = "0x" + hex.EncodeToString(gasPrice.Bytes())
	tx.Value = "0x333"
	//tx.Nonce = "0x131"
	tx.Data = "0xb6b55f250000000000000000000000000000000000000000000000000000000000000021" // call Deposit(uint _toChain)
	//tx.Data = "0x8dd148020000000000000000000000005fea67eb73c9e3edac55f22c8833bcc683b70d5d" //call setBridge(address _addr)
	//txSigned, err := ks.SignTx(accounts[0], tx, big.NewInt(int64(33))) // chainId
	txJson, err := json.Marshal(tx)
	//fmt.Println(string(txJson))

	//fmt.Println(txBytes)big
	txJsonStr := string(txJson)
	resp, err := sendTx(txJsonStr, url, client)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(hex.EncodeToString(body))
	txHash, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
	}
	fmt.Println("txHash: " + txHash + "\n")

	return txHash
}

func txClient(url string, chain string, gasPrice *big.Int){
	fmt.Println("client started for chain", chain)
	accounts := ks.Accounts()

	nonce := getNonce(accounts[0].Address[:], url)[2:]
	if(len(nonce) % 2 == 1) {
		nonce = "0" + nonce
	}
	fmt.Println(nonce)
	nonceBytes, err := hex.DecodeString(nonce[:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(nonceBytes)
	nonceBig := new(big.Int)
	nonceBig.SetBytes(nonceBytes)
    nonceUint := nonceBig.Uint64()
	//nonceUint := uint64(14)

	client := &http.Client{}
	data, err := hex.DecodeString("b6b55f250000000000000000000000000000000000000000000000000000000000000021")
	//data, err := hex.DecodeString("8dd148020000000000000000000000005fea67eb73c9e3edac55f22c8833bcc683b70d5d") //call setBridge(address _addr)
	txGeth := types.NewTransaction(nonceUint, accounts[1].Address, big.NewInt(int64(33)), GAS_LIMIT, gasPrice, data)
	txSigned, err := ks.SignTx(accounts[0], txGeth, big.NewInt(int64(33))) // chainId
	txSignedJson, err := json.Marshal(txSigned)
	fmt.Println(string(txSignedJson))
	//txNonce, err := parseJsonForEntry(string(txSignedJson), "nonce")
	//to, err := parseJsonForEntry(string(txSignedJson), "to")
	value, err := parseJsonForEntry(string(txSignedJson), "value")
	gas, err := parseJsonForEntry(string(txSignedJson), "gas")

	//txGasPrice, err := parseJsonForEntry(string(txSignedJson), "gasPrice") // should be ok
	// v, err := parseJsonForEntry(string(txSignedJson), "v")
	// r, err := parseJsonForEntry(string(txSignedJson), "r")
	// s, err := parseJsonForEntry(string(txSignedJson), "s")

	tx := new(Tx)
	tx.From = "0x" + hex.EncodeToString(accounts[0].Address[:])
	//tx.To = to
	tx.To = "0x5fea67eb73c9e3edac55f22c8833bcc683b70d5d"
	tx.GasPrice = "0x" + hex.EncodeToString(gasPrice.Bytes())
	tx.Gas = gas
	tx.Value = value
	tx.Nonce = "0x" + strconv.FormatInt(int64(nonceUint), 16)
	//tx.Nonce = "0x131"
	tx.Data = "0xb6b55f250000000000000000000000000000000000000000000000000000000000000021" // call Deposit(uint _toChain)
	// tx.V = v
	// tx.R = r
	// tx.S = s
	txJson, err := json.Marshal(tx)
	fmt.Println(string(txJson))

	//fmt.Println(txBytes)big
	txJsonStr := string(txJson)
	resp, err := sendTx(txJsonStr, url, client)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(hex.EncodeToString(body))
	logsResult, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
	}
	fmt.Println("logsResult: " + logsResult + "\n")
}

func main() {
	/* read abi of contract in truffle folder */
	readAbi()

	/* flags */
	// -v
	// default = false
	// if verbosity = true, print out waiting for logs
	verbosePtr := flag.Bool("v", false, "a bool representing verbosity of output")
	// would never actually want this, it's just kinda cool
	readAllPtr := flag.Bool("a", false, "a bool representing whether to read logs from every contract or not")
	configPtr := flag.String("config", "./config.json", "a string of the path to the config file") 
	keysPtr := flag.String("keystore", "./keystore", "a string of the path to the keystore directory") 

	flag.Parse()
	configStr := *configPtr
	fmt.Println("config path: ", configStr)

	verbose = *verbosePtr
	if verbose { fmt.Println("verbose: ", verbose) }

	readAll = *readAllPtr
	if readAll { fmt.Println("read from all contracts? ", readAll)}

	chains = flag.Args()
	if len(chains) == 0 {
		chains = append(chains,"33")
	}
	fmt.Println("chains to connect to: ", chains)

	keystorePath := *keysPtr
	fmt.Println("keystore path: ", keystorePath)

	// config file reading
	path, _ := filepath.Abs(configStr)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read file:", err)	
	}

	// read config file for each chain id
	for i, chain := range chains {
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
		fmt.Println("contract address of chain", chains[i], ":", contractAddr)
		contracts = append(contracts, contractAddr)

		url, err := parseJsonForEntry(chainStr, "url")
		if err != nil {
			fmt.Println("could not find url in config file")
			log.Fatal(err)
		}
		fmt.Println("url of chain", chain, ":", url)
		chainUrls = append(chainUrls, url)

		gp, err := parseJsonForEntry(chainStr, "gasPrice")
		if err != nil {
			fmt.Println("could not find gas price in config file")
			log.Fatal(err)
		}
		bigGas := new(big.Int)
		bigGas.SetString(gp, 10)
		gasPrices = append(gasPrices, bigGas)
	}

	/* keys */
	ks = newKeyStore(keystorePath)
	accounts := ks.Accounts()
	for i, account := range accounts {
		fmt.Println("account", i, ":", account.Address.Hex())
	}
	err = ks.Unlock(accounts[0], "password")
	if err != nil {
		fmt.Println("could not unlock account 0")
		fmt.Println(err)
	}

	/* listener */
	fmt.Println("\nlistening for events...")
	listen(chainUrls, chains)

	/* client */
	//fmt.Println("\nclient started...")
	//client(chainUrls, chains, keyStore)
}