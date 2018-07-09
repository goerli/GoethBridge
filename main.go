package main

import (
	"fmt"
	"net/http"
	//"math/big"
	"bytes"
	"io/ioutil"
	"time"
	"encoding/json"
)

type Resp struct {
	jsonrpc string
	id int
	result string
}

func main() {
	// hard coded to geth running at address:port
	url := "http://127.0.0.1:8545"

	var jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_newBlockFilter","params":[],"id":71}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	//jsonBody := []byte(string(body))

	var jsonBody = []byte(`
		{"jsonrpc":"2.0","id":71,"result":"0x7eec0a227d852c28d4141c212578292b"}
	`)

	var respBody Resp
	err = json.Unmarshal(jsonBody, &respBody)
	if err != nil { fmt.Println(err) }
	fmt.Println(respBody)

        //jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_getFilterChanges","params":[],"id":71}`)
	jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":71}`)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			req, _ = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
        		if err != nil {
                		fmt.Println(err)
        		}
   			defer resp.Body.Close()

    			fmt.Println("response Status:", resp.Status)
    			fmt.Println("response Headers:", resp.Header)
    			body, _ := ioutil.ReadAll(resp.Body)
    			fmt.Println("response Body:", string(body))
		}
	}()

	time.Sleep(60 * time.Second)
	ticker.Stop()
}
