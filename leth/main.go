package main

import (
	"fmt"
	"github.com/ChainSafeSystems/leth/core"
)

func main() {
	err := core.Migrate("rinkeby", "Bridge")
	if err != nil {
		fmt.Println("could not deploy Bridge.sol to default network")
	}

	// err := core.Migrate("default", "Home")
 //    if err != nil {
 //            fmt.Println("could not deploy Home.sol to default network")
 //    }

	// err = core.Migrate("testnet", "Foreign")
 //    if err != nil {
 //            fmt.Println("could not deploy Foreign.sol to ropsten")
 //    }

	// err := core.Migrate("kovan", "Foreign")
 //    if err != nil {
 //            fmt.Println("could not deploy Foreign.sol to kovan")
 //    }

	// err = core.Migrate("ropsten", "Foreign")
 //    if err != nil {
 //            fmt.Println("could not deploy Foreign.sol to ropsten")
 //    }

   	// err := core.Migrate("rinkeby", "Foreign")
    // if err != nil {
    //         fmt.Println("could not deploy Foreign.sol to rinkeby")
    // }
}
