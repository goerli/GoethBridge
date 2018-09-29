package main

import (
)

func main() {
	err := core.Migrate("default", "Bridge")
	if err != nil {
		fmt.Println("could not deploy Bridge.sol to default network")
	}
}