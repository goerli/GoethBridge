package client

import (
	"encoding/hex"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ChainSafeSystems/ChainBridge/logger"
)

// generate the 4-byte identifier from a function signature
func generateSignature(sig string) (string) {
	bytes := []byte(sig)
	hash := crypto.Keccak256(bytes)
	hex := hex.EncodeToString(hash)
	return hex[0:8] // first 4 bytes only
}

// sign a message using chain.From account
func SignMessage(chain *Chain, msg []byte) ([]byte, error) {
	from := new(accounts.Account)
	from.Address = *chain.From
	msg, err := keys.SignHashWithPassphrase(*from, chain.Password, msg)
	if err != nil {
		return nil, err
	} else { return msg, nil }
}

// send a tx to chain with calldata
func SendTx(chain *Chain, value *big.Int, data []byte) (common.Hash, error) {
	client := chain.Client
	from := new(accounts.Account)
	from.Address = *chain.From

	nonce, err := client.PendingNonceAt(context.Background(), *chain.From)
	chain.Nonce = nonce

	tx := types.NewTransaction(chain.Nonce, *chain.Contract, value, uint64(4600000), chain.GasPrice, data)
	txSigned, err := keys.SignTxWithPassphrase(*from, chain.Password, tx, chain.Id)
	if err != nil {
		logger.Error("could not sign tx: %s", err)
		return *new(common.Hash), err
	}

	txHash := txSigned.Hash()
	err = client.SendTransaction(context.Background(), txSigned)
	if err != nil {
		logger.Error("could not send tx: %s", err)
		return *new(common.Hash), err
	}

	return txHash, nil
}

func AddAuthority(chain *Chain, address string) error {
	dataStr := generateSignature("addAuthority(address)") + padTo32Bytes(address[2:]) // setbridge function signature + contract addr
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return err
	} 

	txHash, err := SendTx(chain, big.NewInt(0), data)
	if err != nil {
		return err
	}

	logger.Info("sending tx %s to add authority on %s...", txHash.Hex(), chain.Name)
	return nil
}

// id is the id of the chain to withdraw the deposit on
// ids are in hexidecimal
func Deposit(chain *Chain, value *big.Int, id string) error {
	dataStr := "47e7ef24" + padTo32Bytes(chain.From.Hex()[2:]) + padTo32Bytes(id) // deposit function signature + recipient addr + chain
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return err
	} 

	txHash, err := SendTx(chain, value, data)
	if err != nil {
		return err
	}	

	logger.Info("sending tx %s to deposit on %s...", txHash.Hex(), chain.Name)
	return nil
}

func PayBridge(chain *Chain, value *big.Int) error {
	txHash, err := SendTx(chain, value, []byte{})
	if err != nil {
		return err
	}	

	logger.Info("sending tx %s to pay bridge on %s...", txHash.Hex(), chain.Name)
	return nil
}

// ids are in hexidecimal
func WithdrawTo(chain *Chain, value *big.Int, id string) error {
	dataStr := "5fcbc20e" + padTo32Bytes(chain.From.Hex()[2:]) + padTo32Bytes(id) + padBigTo32Bytes(value)
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return err
	} 

	txHash, err := SendTx(chain, value, data)
	if err != nil {
		return err
	}	

	logger.Info("sending tx %s to deposit on %s...", txHash.Hex(), chain.Name)
	return nil
}

func Withdraw(chain *Chain, withdrawal *Withdrawal) error {
	w := setWithdrawalData(withdrawal)
	dataStr := "4250a6f3" + w.Data 
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		return err
	} 

	txHash, err := SendTx(chain, big.NewInt(0), data)
	if err != nil {
		return err
	}	

	logger.Info("sending tx %s to withdraw on %s...", txHash.Hex(), chain.Name)
	return nil
}

func FundBridge(chain *Chain, value *big.Int) error {
	weiValue := big.NewInt(0)
	weiConversion := big.NewInt(0)
	weiValue.Mul(value, weiConversion.Exp(big.NewInt(10), big.NewInt(18), nil))
	data, err := hex.DecodeString("c9c0909f") //fund me function sig
	if err != nil {
		return err
	} 

	txHash, err := SendTx(chain, weiValue, data)
	if err != nil {
		return err
	}	

	logger.Info("sending tx %s to fund bridge on %s with value %s...", txHash.Hex(), chain.Name, value.String())
	return nil
}