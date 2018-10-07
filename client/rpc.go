package client

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/types"
)

/* rpc methods 
/* used when making http requests
*/

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
	receipt, err := ParseJsonForResult(string(body))
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
	rawTx, err := ParseJsonForResult(string(body))
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
	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)

	// parse json for result
	startBlock, err := ParseJsonForResult(string(blockNumBody))
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
	blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)

	// parse json for result
	res, err := ParseJsonForResult(string(blockNumBody))
	if err != nil {
		return "", nil
	}
	return res, nil
}

func getNonce(address []byte, url string) (string) {
	// get nonce
	client := &http.Client{}
	txCountData := fmt.Sprintf("\"0x%x\", \"latest\"", address)
	resp, err := getTxCount(txCountData, url, client)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	nonce, err := ParseJsonForResult(string(body))
	if err != nil {
		fmt.Println("could not parse logs")
		fmt.Println(err)
	}

	return nonce
}

/*****  helpers *****/
type JsonRpcResponse struct {
	Id int					`json:"id,omitempty"`
	Jsonrpc string 			`json:"jsonrpc,omitempty"`
	Result string 			`json:"result"`
}

// this function parses jsonStr for the entry "result" and returns its value as a string
func ParseJsonForResult(jsonStr string) (string, error) {
	jsonBody := []byte(jsonStr)
	resp := new(JsonRpcResponse)
	err := json.Unmarshal(jsonBody, resp)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
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