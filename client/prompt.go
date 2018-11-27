package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

/*** admin functions ***/
func AddAuthorityPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var address string
	var confirm int64
	fmt.Println("adding authority on ", chain.Name)
	fmt.Println("enter address of authority to add")
	fmt.Scanln(&address)
	if len(address) != 42 { 
		return
	}

	fmt.Println("confirm adding authority ", address, "on", chain.Name)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	AddAuthority(chain, address[2:])
}

func RemoveAuthorityPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var address string
	var confirm int64
	fmt.Println("removing authority on ", chain.Name)
	fmt.Println("enter address of authority to remove")
	fmt.Scanln(&address)
	if len(address) != 42 { 
		return
	}

	fmt.Println("confirm removing authority ", address, "on", chain.Name)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	RemoveAuthority(chain, address[2:])
}

func IncreaseThresholdPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var address string
	var confirm int64
	fmt.Println("increasing authority threshold on ", chain.Name)
	fmt.Println("confirm authority threshold increase", address, "on", chain.Name)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	IncreaseThreshold(chain)
}

func DecreaseThresholdPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var address string
	var confirm int64
	fmt.Println("decreasing authority threshold on ", chain.Name)
	fmt.Println("confirm authority threshold decrease", address, "on", chain.Name)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	DecreaseThreshold(chain)
}

/*** bridge functions ***/
func FundPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var value int64
	var confirm int64
	fmt.Println("\nfunding the bridge contract on chain", chain.Id)
	fmt.Println("note that funding of the bridge cannot be withdrawn")
	fmt.Println("enter value of funding, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	valBig := big.NewInt(value)
	fmt.Println("confirm funding on chain", chain.Id, "with value", value, "wei")
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}
	FundBridge(chain, valBig)
}

func DepositPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var value int64
	var to int64
	var confirm int64
	fmt.Println("\ndepositing to the bridge contract on chain", chain.Id)
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of deposit, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	fmt.Println("enter chain id to withdraw on")
	fmt.Scanln(&to)
	if to == -1 { 
		return
	}

	valBig := big.NewInt(value)

	toHex := fmt.Sprintf("%x", to)
	fmt.Println("confirm deposit on chain", chain.Id, "with value", value, "wei, withdrawing to chain", to)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}
	Deposit(chain, valBig, toHex)
}

func WithdrawToPrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var value int64
	var to int64
	var confirm int64
	fmt.Println("\nwithdrawing to other chains from the bridge contract on chain", chain.Id)
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of withdraw, in wei")
	fmt.Scanln(&value)
	if value == -1 { 
		return
	}
	fmt.Println("enter chain id to withdraw on")
	fmt.Scanln(&to)
	if to == -1 { 
		return
	}

	fmt.Println("confirm deposit on chain", chain.Id, "with value", value, "wei, withdrawing to chain", to)
	fmt.Scanln(&confirm)
	if confirm == -1 { 
		return
	}

	valBig := big.NewInt(value)
	toHex := fmt.Sprintf("%x", to)
	WithdrawTo(chain, valBig, toHex)
}

func PayBridgePrompt(chain *Chain, ks *keystore.KeyStore) {
	keys = ks

	var value int64
	var confirm int64
	fmt.Println("\npaying bridge contract on chain", chain.Id)
	fmt.Println("note that bridge payments can later be withdrawn")
	fmt.Println("type -1 to escape")
	fmt.Println("enter value of payment, in wei")
	fmt.Scanln(&value)
	if value == -1 {
		return
	}

	fmt.Println("confirm payment to bridge on chain", chain.Id, "with value", value, "wei")
	fmt.Scanln(&confirm)
	if confirm == -1 {
		return
	}

	valBig := big.NewInt(value)
	PayBridge(chain, valBig)
}