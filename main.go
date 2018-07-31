package main

import (
	"fmt"
	"net/http"
	"math/big"
	"bytes"
	"io/ioutil"
	//"time"
	"jsonparser"
	"encoding/json"
	"encoding/hex"
	//"encoding/binary"
	"path/filepath"
	"strings"
	"log"
	"flag"
	//"strconv"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/rlp"

    //"github.com/noot/multi_directional_bridge/rlp"
    "github.com/noot/multi_directional_bridge/transaction"
    "github.com/noot/multi_directional_bridge/client"
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

/* global vars */
// flags
var verbose bool
var readAll bool
//var clients []*client.Chain

// events to listen for
var DepositId string
var CreationId string
var WithdrawId string
var BridgeSetId string

// keystore
var ks *keystore.KeyStore

// channels 
var setBridgeDone chan bool

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

	withdrawEvent := bridgeEvents["Withdraw"]
	withdrawHash := withdrawEvent.Id()
	WithdrawId = withdrawHash.Hex()
	fmt.Println("withdraw event id: ", WithdrawId)

	bridgeSetEvent := bridgeEvents["BridgeSet"]
	bridgeSetHash := bridgeSetEvent.Id()
	BridgeSetId = bridgeSetHash.Hex()
	fmt.Println("set bridge event id: ", BridgeSetId)
}

// starts a goroutine to listen on every chain 
// func listen(urls []string, chains []string) {
// 	var params LogParams

// 	logsFound := make(map[string]bool)

// 	// poll filter every 500ms for changes
// 	ticker := time.NewTicker(100 * time.Millisecond)
// 	for i, _ := range urls {
// 		go func(chain string, url string, contractAddr string) {
// 			client := &http.Client{}
// 			fmt.Println("listening at: " + url)

// 			go txClient(url, chain, contractAddr, gasPrices[i])

// 			for t := range ticker.C{
// 				if verbose { fmt.Println(t) }

// 				if !readAll { 
// 					params.Address = contracts[i]
// 				} 
// 				params.FromBlock, _ = getBlockNumber(url, client)
// 				if verbose { fmt.Println("getting logs from block number: " + params.FromBlock + "\n") }
// 				jsonParams, _ := json.Marshal(params)
// 	            //fmt.Println("jsonParams: " + string(jsonParams))

// 				//get logs from params.FromBlock
// 				resp, _ := getLogs(url, string(jsonParams), client)
// 				defer resp.Body.Close()

// 				//fmt.Println("response Status:", resp.Status)
// 				//fmt.Println("response Headers:", resp.Header)
// 				body, _ := ioutil.ReadAll(resp.Body)
// 				//fmt.Println("response Body:", string(body))
	 
// 				// parse for getLogs result
// 				//logsResult := parseJsonForResult(string(body))
// 				logsResult, err := parseJsonForEntry(string(body), "result")
// 				if err != nil {
// 					fmt.Println("could not parse logs")
// 					fmt.Println(err)
// 				}
// 				if verbose { fmt.Println("logsResult: " + logsResult + "\n") }
// 				//fmt.Println(len(logsResult))

// 				// if there are new logs, parse for event info
// 				if len(logsResult) > 2 {
// 					txHash, _ := parseJsonForEntry(logsResult[1:len(logsResult)-1], "transactionHash")
// 					//fmt.Println(txHash + "\n")
// 					if logsFound[txHash] != true { 
// 						logsFound[txHash] = true
// 						fmt.Println("\nnew logs found for chain", chain)

// 						//logs <- logsResult
// 						//readLogs(logs)
// 						//go readLogs(logs)
// 						//<-exit

// 						// get logs contract address
// 						address, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "address")
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						// this is not actually a good way to listen for events from a  contract
// 						// this could be used to confirm a log, but for listening to events from
// 						// one contract, we would specify the address in our call to eth_getLogs
// 						fmt.Println("contract addr: ", address)
// 						//fmt.Println("length of address: ", len(address))
// 						for i := 0; i < len(chains); i++ {
// 							if strings.Compare(address[1:41], chains[i]) == 0 {
// 								fmt.Println("bridge contract event heard on chain ", chains[i])
// 							}
// 						}

// 						// read topics of log
// 						topics, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "topics")
// 						if err != nil {
// 							fmt.Println(err)
// 						}
// 						fmt.Println("topics: ", topics[2:68])
// 						//fmt.Println("length of topics: ", len(topics)-4) len = 66: 0x + 64 hex chars = 32 bytes

// 						if strings.Compare(topics[2:68],DepositId) == 0 { 
// 							fmt.Println("*** deposit event ", topics[2:68])
// 							data, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "data")
// 							if err != nil {
// 								fmt.Println(nil)
// 							}

// 							receiver, value, toChain := readDepositData(data)
// 							fmt.Println("receiver: ", receiver) 
// 							fmt.Println("value: ", value) // in hexidecimal
// 							fmt.Println("to chain: ", toChain) // in hexidecimal
// 					 	} else if strings.Compare(topics[2:68],CreationId) == 0 {
// 							fmt.Println("*** bridge contract creation")
// 						} else if strings.Compare(topics[2:68],WithdrawId) == 0 {
// 							fmt.Println("*** withdraw event")
// 							data, err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "data")
// 							if err != nil {
// 								fmt.Println(nil)
// 							}
// 							receiver, value, toChain := readDepositData(data)
// 							fmt.Println("receiver: ", receiver) 
// 							fmt.Println("value: ", value) // in hexidecimal
// 							fmt.Println("to chain: ", toChain) // in hexidecimal
// 						} else if strings.Compare(topics[2:68],BridgeSetId) == 0 {
// 							fmt.Println("*** set bridge event")
// 							setBridgeDone <- true
// 						}
// 					}
// 				}
// 			}
// 		}(chains[i], urls[i], contracts[i])
// 	} 

// 	// bridge timeout. eventually, change so it never times out
// 	time.Sleep(6000 * time.Second)
// 	ticker.Stop()
// }

func setBridge(url string, chain string, contractAddr string, gasPrice *big.Int) (string) {
	client := &http.Client{}
	accounts := ks.Accounts()

	//data := make([]byte, 32)
	//tx := types.NewTransaction(uint64(100), accounts[1].Address, big.NewInt(int64(0)), uint64(4600000), gasPrices[0], data)
	tx := new(transaction.Tx)
	tx.From = "0x" + hex.EncodeToString(accounts[0].Address[:])
	tx.To = contractAddr
	tx.GasPrice = "0x" + hex.EncodeToString(gasPrice.Bytes())
	//tx.Nonce = "0x131"
	//tx.Data = "0xb6b55f250000000000000000000000000000000000000000000000000000000000000021" // call Deposit(uint _toChain)
	tx.Data = "0x8dd148020000000000000000000000008f9b540b19520f8259115a90e4b4ffaeac642a30" //call setBridge(address _addr)
	//txSigned, err := ks.SignTx(accounts[0], tx, big.NewInt(int64(33))) // chainId
	txJson, err := json.Marshal(tx)
	//fmt.Println(string(txJson))

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
	return txHash
}

func setBridgeRaw(url string, chain string, contractAddr string, gasPrice *big.Int, nonce uint64) (string, error) {
	client := &http.Client{}
	accounts := ks.Accounts()

	contract := new(common.Address)
	contractBytes, err := hex.DecodeString(contractAddr[2:])
	if err != nil {
		return "", err
	}
	contract.SetBytes(contractBytes)

	// data, err := hex.DecodeString("8dd148020000000000000000000000008f9b540b19520f8259115a90e4b4ffaeac642a30")
	// if err != nil {
	// 	fmt.Println(err)
	// } 
	data := new([]byte)
	// NewTransaction(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte)
	tx := types.NewTransaction(nonce, *contract, big.NewInt(int64(0)), uint64(4600000), gasPrice, *data)
	txSigned, err := ks.SignTx(accounts[0], tx, big.NewInt(int64(33))) // chainId
	txData, err := rlp.EncodeToBytes(txSigned)

	txRlpData := hex.EncodeToString(txData)
	var txRes *types.Transaction
    rawtx,err := hex.DecodeString(txRlpData)
    rlp.DecodeBytes(rawtx, &txRes)
    //fmt.Println(txRes)

	//fmt.Println("length of rlp encoded data: ", len("0x" + txRlpData))
	resp, err := sendRawTx("0x" + txRlpData, url, client)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(hex.EncodeToString(body))
	txHash, err := parseJsonForEntry(string(body), "result")
	if err != nil {
		return "", err
	}
	//fmt.Println("txHash: " + txHash + "\n")

	return txHash, nil
}

func withdraw(url string, chain string, contractAddr string, data string, gasPrice *big.Int) (string) {
	client := &http.Client{}
	accounts := ks.Accounts()

	//data := make([]byte, 32)
	//tx := types.NewTransaction(uint64(100), accounts[1].Address, big.NewInt(int64(0)), uint64(4600000), gasPrices[0], data)
	tx := new(transaction.Tx)
	tx.From = "0x" + hex.EncodeToString(accounts[0].Address[:])
	//tx.To = "0x" + hex.EncodeToString(accounts[1].Address[:])
	tx.To = contractAddr
	tx.GasPrice = "0x" + hex.EncodeToString(gasPrice.Bytes())
	//tx.Value = "0x333"
	//tx.Nonce = "0x131"
	tx.Data = data //call setBridge(address _addr)
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

func rlpDecodeTx(rawTxData string) (*types.Transaction) {
	tx := new(types.Transaction)
    rawtx,err := hex.DecodeString(rawTxData)
    if err != nil {
    	fmt.Println(err)
    }
    rlp.DecodeBytes(rawtx, &tx)
	return tx
}

func txClient(url string, chain string, contractAddr string, gasPrice *big.Int){
	fmt.Println("client started for chain", chain)
	accounts := ks.Accounts()

	nonce := getNonce(accounts[0].Address[:], url)[2:]
	if(len(nonce) % 2 == 1) {
		nonce = "0" + nonce
	}
	//fmt.Println(nonce)
	nonceBytes, err := hex.DecodeString(nonce[:])
	if err != nil {
		fmt.Println(err)
	}
	nonceBig := new(big.Int)
	nonceBig.SetBytes(nonceBytes)
    nonceUint := nonceBig.Uint64()
    fmt.Println("nonce: ", nonceUint)

	txHash := setBridge(url, chain, contractAddr, gasPrice)
	fmt.Println("set bridge tx hash: ", txHash)

	<-setBridgeDone

	client := &http.Client{}
	txReceipt, err := getTxReceipt(txHash, url, client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txReceipt)

	rawTxData, err := getRawTx(txHash, url, client)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rawTxData)
	fmt.Println("length of received raw data: ", len(rawTxData))

	// decode
	tx := rlpDecodeTx(rawTxData[2:])
	fmt.Println("decoded data: ", tx)
	//fmt.Println(tx.To().Hex())

	txHash, err = setBridgeRaw(url, chain, contractAddr, gasPrice, nonceUint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("set bridge raw tx hash: ", txHash)

	// data := "0xb5c5f6720000000000000000000000008f9b540b19520f8259115a90e4b4ffaeac642a30000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000021"
	// txHash = withdraw(url, chain, contractAddr, data, gasPrice)
	// fmt.Println("withdraw tx hash: ", txHash)

	//client := &http.Client{}
	//data, err := hex.DecodeString("b6b55f250000000000000000000000000000000000000000000000000000000000000021")
	//data, err := hex.DecodeString("8dd148020000000000000000000000005fea67eb73c9e3edac55f22c8833bcc683b70d5d") //call setBridge(address _addr)
	//txGeth := types.NewTransaction(nonceUint, accounts[1].Address, big.NewInt(int64(33)), GAS_LIMIT, gasPrice, data)
	//txSigned, err := ks.SignTx(accounts[0], txGeth, big.NewInt(int64(33))) // chainId

	//txSignedJson, err := json.Marshal(txSigned)
	//fmt.Println(string(txSignedJson))
	//txNonce, err := parseJsonForEntry(string(txSignedJson), "nonce")
	//to, err := parseJsonForEntry(string(txSignedJson), "to")
	//value, err := parseJsonForEntry(string(txSignedJson), "value")
	//gas, err := parseJsonForEntry(string(txSignedJson), "gas")

	//txGasPrice, err := parseJsonForEntry(string(txSignedJson), "gasPrice") // should be ok
	// v, err := parseJsonForEntry(string(txSignedJson), "v")
	// r, err := parseJsonForEntry(string(txSignedJson), "r")
	// s, err := parseJsonForEntry(string(txSignedJson), "s")

	// tx := new(Tx)
	// tx.From = "0x" + hex.EncodeToString(accounts[0].Address[:])
	// //tx.To = to
	// tx.To = contractAddr
	// tx.GasPrice = "0x" + hex.EncodeToString(gasPrice.Bytes())
	// tx.Gas = gas
	// //tx.Value = value
	// //tx.Nonce = "0x" + strconv.FormatInt(int64(nonceUint), 16)
	// //tx.Nonce = "0x131"
	// // withdraw data
	// //tx.Data = "0xb5c5f6720000000000000000000000008f9b540b19520f8259115a90e4b4ffaeac642a30000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000021"
	// tx.Data = "0x8dd148020000000000000000000000008f9b540b19520f8259115a90e4b4ffaeac642a30" // call setBridge(address)
	// //tx.Data = "0xb6b55f250000000000000000000000000000000000000000000000000000000000000021" // call Deposit(uint _toChain)
	// // tx.V = v
	// // tx.R = r
	// // tx.S = s
	// txJson, err := json.Marshal(tx)
	// //fmt.Println(string(txJson))

	// //fmt.Println(txBytes)big
	// txJsonStr := string(txJson)
	// resp, err := sendTx(txJsonStr, url, client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// //fmt.Println(hex.EncodeToString(body))
	// logsResult, err := parseJsonForEntry(string(body), "result")
	// if err != nil {
	// 	fmt.Println("could not parse logs")
	// 	fmt.Println(err)
	// }
	// //fmt.Println("logsResult: " + logsResult + "\n")
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
	passwordPtr := flag.String("password", "password", "a string of the password to the first keystore account") 

	flag.Parse()
	configStr := *configPtr
	fmt.Println("config path: ", configStr)

	verbose = *verbosePtr
	if verbose { fmt.Println("verbose: ", verbose) }

	readAll = *readAllPtr
	if readAll { fmt.Println("read from all contracts? ", readAll)}

	chains := flag.Args()
	if len(chains) == 0 {
		chains = append(chains,"33")
	}
	fmt.Println("chains to connect to: ", chains)

	keystorePath := *keysPtr
	fmt.Println("keystore path: ", keystorePath)

	password := *passwordPtr

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
		clients[i].From = fromAccount
	}

	/* keys */
	ks = newKeyStore(keystorePath)
	ksaccounts := ks.Accounts()
	for i, account := range ksaccounts {
		fmt.Println("account", i, ":", account.Address.Hex())
	}
	//fromAddress := common.HexToAddress(fromAccounts[0])

	err = ks.Unlock(ksaccounts[0], password)
	if err != nil {
		fmt.Println("could not unlock account 0")
		fmt.Println(err)
	}

	// if(ks.HasAddress(fromAddress)) {
	// 	account := new(accounts.Account)
	// 	account.Address = fromAddress
	// 	err = ks.Unlock(*account, password)
	// 	if err != nil {
	// 		fmt.Println("could not unlock account 0")
	// 		fmt.Println(err)
	// 	}
	// } else {
	// 	log.Fatal("account not found in keystore")
	// }

	/* channels */
	setBridgeDone = make(chan bool)
	doneClient := make(chan bool)

	/* listener */
	fmt.Println("\nlistening for events...")
	for _, chain := range clients {
		//fmt.Println(chain)
		go client.Listen(chain, doneClient)
	}

	<-doneClient
	/* client */
	//fmt.Println("\nclient started...")
	//client(chainUrls, chains, keyStore)
}