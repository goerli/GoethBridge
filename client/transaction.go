package client

import (
	"fmt"
	"encoding/hex"
	"context"
	"math/big"
	//"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/ethclient"
)

/****** functions to send transactions ******/

func SetBridge(chain *Chain) () {
	client := chain.Client
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	dataStr := "8dd14802000000000000000000000000" + chain.From.Hex()[2:] // setbridge function signature + contract addr
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, big.NewInt(int64(0)), uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to set bridge")

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}
}

// id is the id of the chain to withdraw the deposit on
// ids are in hexidecimal
func Deposit(chain *Chain, value *big.Int, id string) {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	chainId := padTo32Bytes(id)	
	dataStr := "47e7ef24000000000000000000000000" + chain.From.Hex()[2:] + chainId // deposit function signature + recipient addr + chain
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to deposit on chain", chain.Id)

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}
}

// ids are in hexidecimal
func DepositMulti(chain *Chain, value *big.Int, ids []string, percents []*big.Int) {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	dataStr := "29a956bc000000000000000000000000" + chain.From.Hex()[2:] + 
	"0000000000000000000000000000000000000000000000000000000000000006" + 
	"000000000000000000000000000000000000000000000000000000000000000c"
	dataStr = dataStr + padIntTo32Bytes(int64(len(ids)))
	//fmt.Println(len(dataStr))
	for _, id := range ids {
		idPadded := padTo32Bytes(id)
		dataStr = dataStr + idPadded
	}

	dataStr = dataStr + padIntTo32Bytes(int64(len(percents)))
	//fmt.Println(len(dataStr))
	for _, percent:= range percents {
		percentPadded := padBigTo32Bytes(percent)
		dataStr = dataStr + percentPadded
	}

	//fmt.Println("input data: ", dataStr)

	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(6700000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to deposit on chain", chain.Id)

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}
}

func Withdraw(chain *Chain, withdrawal *Withdrawal) {
	client := chain.Client
	//accounts := keys.Accounts()
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	withdrawal = setWithdrawalData(withdrawal)
	dataStr := "b5c5f672000000000000000000000000" + withdrawal.Data // withdraw function signature + contract addr
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, big.NewInt(int64(0)), uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to withdraw on chain", chain.Id)

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		//fmt.Println("could not send tx")
		//fmt.Println(err)
	}
}

func FundBridge(chain *Chain, value *big.Int) {
	client := chain.Client
	from := new(accounts.Account)
	from.Address = *chain.From
	fmt.Println()

	data, err := hex.DecodeString("c9c0909f") //fund me function sig
	if err != nil {
		fmt.Println(err)
	} 

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		fmt.Println("could not sign tx")
		fmt.Println(err)
	}

	txHash := txSigned.Hash()
	fmt.Println("attempting to send tx", txHash.Hex(), "to fund bridge on chain", chain.Id, "with value", value.String())

	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		fmt.Println("could not send tx")
		fmt.Println(err)
	}
}