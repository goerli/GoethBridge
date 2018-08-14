package function

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSignature(sig string) (string) {
	bytes := []byte(sig)
	hash := crypto.Keccak256(bytes)
	hex := hex.EncodeToString(hash)
	return hex[0:8] // first 4 bytes only
}