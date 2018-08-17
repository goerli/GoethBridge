package client

import (
	"fmt"
	"jsonparser"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

/***** rpc methods ******/
// used when making http requests


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
	receipt, err := ParseJsonForEntry(string(body), "result")
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
	rawTx, err := ParseJsonForEntry(string(body), "result")
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
func getBlockNumber(url string) (string, error) {
	client := &http.Client{}
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

// this function gets the current block by calling "eth_getBlockByNumber"
func getBlockByNumber(url string, number string) (string, error) {
	client := &http.Client{}
	var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["` + number + `",false],"id":1}`)
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
	res, err := parseJsonForResult(string(blockNumBody))
	if err != nil {
		return "", nil
	}
	return res, nil
}

func getBlockRoot(url string, number string) (common.Hash, error) {
	jsonRes, err := getBlockByNumber(url, number)
	if err != nil {
		return *new(common.Hash), err
	}
	root, err := ParseJsonForEntry(jsonRes, "hash")
	if err != nil {
		return *new(common.Hash), err
	}
	rootHash := common.HexToHash(root)
	return rootHash, nil
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
	nonce, err := ParseJsonForEntry(string(body), "result")
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
func ParseJsonForEntry(jsonStr string, get string) (string, error) {
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