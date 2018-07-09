package main

import (
	"fmt"
	//"net/rpc"
	"net/http"
	//"github.com/ethereum/go-ethereum/rpc"
	//"math/big"
	"bytes"
	"io/ioutil"
	"time"
)

func main() {
	// hard coded to geth running at address:port
	url := "http://127.0.0.1:8545"

	var jsonStr = []byte(`{"jsonrpc":"2.0","method":"eth_mining","params":[],"id":71}`)
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

	time.Sleep(20 * time.Second)
	ticker.Stop()
}
