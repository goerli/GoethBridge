package rlp 

import (
	"github.com/ethereum/go-ethereum/rlp"
	//"github.com/ethereum/go-ethereum/core/types"

    "github.com/noot/multi_directional_bridge/transaction"
)

func EncodeRawTx(tx *transaction.Tx) ([]byte, error) {
	txb, err := rlp.EncodeToBytes(tx)
	return txb, err
}