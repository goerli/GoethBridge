package main

import (
	"fmt"
	"net/http"
	//"math/big"
	"bytes"
	"io/ioutil"
	"time"
	"jsonparser"
	"encoding/json"
	//"encoding/hex"
	"path/filepath"
	"strings"
	"log"
	"flag"

	//"github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
)

// used for json format a response from an RPC call
type Resp struct {
	jsonrpc string
	id int
	result string
}

// used to json format an RPC call
type Call struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []string `json:"params"`
	Id int `json:"id"`
}

// used for getLogs json formatting
type LogParams struct {
	FromBlock string `json:"fromBlock"`
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
	//fmt.Println("response Body:", string(blockNumBody))

	// parse json for result
	startBlock, err := parseJsonForResult(string(blockNumBody))
	if err != nil {
		return "", nil
	}
	return startBlock, nil
}

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

func main() {
	/* flags */

	// -v
	// default = false
	// if verbosity = true, print out waiting for logs
	verbosePtr := flag.Bool("v", false, "a bool representing verbosity of output")
	configPtr := flag.String("config", "./config.json", "a string of the path to the config file") 

	// @todo: read url and port from config file.
	urlPtr := flag.String("url", "127.0.0.1", "a string of the url of the client")
	portPtr := flag.String("port", "8545", "a string of the port of the client")

	flag.Parse()
	configStr := *configPtr
	fmt.Println("config path: ", configStr)

	verbose := *verbosePtr
	if verbose { fmt.Println("verbose: ", verbose) }

	chains := flag.Args()
	if len(chains) == 0 {
		chains = append(chains,"33")
	}
	fmt.Println("chains to connect to: ", chains)

	clientAddr := *urlPtr
	port := *portPtr
	//fmt.Println("url: ", clientAddr, ":", port)

	url := "http://" + clientAddr + ":" + port
	fmt.Println("listening at: " + url)
    client := &http.Client{}
	var params LogParams

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

	// config file reading
	path, _ = filepath.Abs(configStr)
	file, err = ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read file:", err)	
	}

	for i := 0; i < len(chains); i++ {
		chainStr, err := parseJsonForEntry(string(file), chains[i])
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
	}

	// checking for abi methods
	// bridgeMethods := bridgeabi.Methods
	// transferMethod := bridgeMethods["transfer"]
	// transferSig := transferMethod.Sig()
	// s := string(transferSig[:])
	// fmt.Println(s)

	// checking abi for events
	bridgeEvents := bridgeabi.Events
	depositEvent := bridgeEvents["Deposit"]
	depositHash := depositEvent.Id()
	depositId := depositHash.Hex()
	//fmt.Println("deposit event id: ", depositId) // this is the deposit event to watch for

	creationEvent := bridgeEvents["ContractCreation"]
	creationHash := creationEvent.Id()
	creationId := creationHash.Hex()
	//fmt.Println("contract creation event id: ", creationId)
	fmt.Println("listening for events...")

	logsFound := make(map[string]bool)

	// poll filter every 500ms for changes
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			if verbose { fmt.Println(t) }

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
				fmt.Println(err)
			}
			if verbose { fmt.Println("logsResult: " + logsResult + "\n") }
			//fmt.Println(len(logsResult))

			// if there are new logs, parse for event info
			if len(logsResult) != 2 {
				txHash, _ := parseJsonForEntry(logsResult[1:len(logsResult)-1], "transactionHash")
				//fmt.Println(txHash + "\n")
				if logsFound[txHash] != true { 
					logsFound[txHash] = true
					fmt.Println("new logs found!")

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
					topics,err := parseJsonForEntry(logsResult[1:len(logsResult)-1], "topics")
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("topics: ", topics[2:68])
					//fmt.Println("length of topics: ", len(topics)-4) len = 66: 0x + 64 hex chars = 32 bytes

					if strings.Compare(topics[2:68],depositId) == 0 { 
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
				 	} else if strings.Compare(topics[2:68],creationId) == 0 {
						fmt.Println("*** bridge contract creation\n")
					}
				}
			}

		}
	}()

	// bridge timeout. eventually, change so it never times out
	time.Sleep(6000 * time.Second)
	ticker.Stop()
}