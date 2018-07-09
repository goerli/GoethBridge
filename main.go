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
)

type Resp struct {
	jsonrpc string
	id int
	result string
}

type Call struct {
	Jsonrpc string `json:"jsonrpc"`
	Method string `json:"method"`
	Params []string `json:"params"`
	Id int `json:"id"`
}

type LogParams struct {
	FromBlock string `json:"fromBlock"`
}

func getLogs(url string, jsonParams string, client *http.Client) (*http.Response, error) {
     	jsonStr := `{"jsonrpc":"2.0","method":"eth_getLogs","params":[` + string(jsonParams) + `],"id":74}`
        jsonBytes := []byte(jsonStr)
        fmt.Println(string(jsonBytes))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
        req.Header.Set("Content-Type", "application/json")
        resp, err := client.Do(req)
	if err != nil { return nil, err }
	return resp, nil
}

func getBlockNumber(url string, client *http.Client) (string, error) {
        // get starting block
        var jsonBytes = []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}`)
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
        req.Header.Set("Content-Type", "application/json")
        blockNumResp, err := client.Do(req)
        if err != nil {
                return "", err
        }
        defer blockNumResp.Body.Close()

        // print out response of eth_blockNumber
        fmt.Println("response Status:", blockNumResp.Status)
        fmt.Println("response Headers:", blockNumResp.Header)
        blockNumBody, _ := ioutil.ReadAll(blockNumResp.Body)
        fmt.Println("response Body:", string(blockNumBody))

        // parse json for result
        jsonBody := []byte(string(blockNumBody))
        res, _, _, _ := jsonparser.Get(jsonBody, "result")
        //resStr := "0x" + hex.EncodeToString(res)
        startBlock := string(res)
        fmt.Println("starting block number: " + startBlock + "\n")
	return startBlock, nil
}

func main() {
	// hard coded to geth running at address:port
	url := "http://127.0.0.1:8545"
        client := &http.Client{}

	startBlock, err := getBlockNumber(url, client)
	if err != nil { fmt.Println(err) }
	//format json call
	var params LogParams
	//params.FromBlock = "0x1" //start at first block
	params.FromBlock = startBlock
	jsonParams, _ := json.Marshal(params)
	fmt.Println(string(jsonParams))

	// poll filter every 500ms for changes
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			resp, _ := getLogs(url, string(jsonParams), client)
			defer resp.Body.Close()

			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("response Body:", string(body))

			params.FromBlock, err = getBlockNumber(url, client)
			jsonParams, _ := json.Marshal(params)
			fmt.Println(string(jsonParams))
		}
	}()

	time.Sleep(60 * time.Second)
	ticker.Stop()
}
