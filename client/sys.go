package client

import (
	"fmt"
	"math/big"
	"os"
	"log"
	"sync"
)

func Cleanup(chain *Chain, lastBlock *big.Int, wg *sync.WaitGroup) {
	f, err := os.Create("log/" + chain.Id.String() + "_lastblock.txt") // creating...
   	if err != nil {
        fmt.Printf("error creating file: %v", err)
        return
    }
    defer f.Close()

	if lastBlock == nil {
		log.Fatal("could not find last block for chain", chain.Id)
	}
	fmt.Println("last block at chain", chain.Id, "is", lastBlock)
	_, err = f.WriteString(fmt.Sprintf("%d\n", lastBlock)) // writing...
	if err != nil {
	    fmt.Printf("error writing string: %v", err)
	}

	wg.Done()
}