package client 

import (
	"math/big"
    "github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	Url string
	Contract string
	GasPrice *big.Int
	From string
}