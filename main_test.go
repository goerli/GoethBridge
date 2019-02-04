package main

import (
	"bytes"
	"fmt"
	"github.com/ChainSafeSystems/ChainBridge/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)


func TestReadAbi(t *testing.T) {
	expected := &client.Events{
		DepositId: common.BytesToHash(crypto.Keccak256([]byte(fmt.Sprint("Deposit(address,uint256,uint256)")))).Hex(),
		CreationId: common.BytesToHash(crypto.Keccak256([]byte(fmt.Sprint("ContractCreation(address)")))).Hex(),
		WithdrawId: common.BytesToHash(crypto.Keccak256([]byte(fmt.Sprint("Withdraw(address,uint256,uint256)")))).Hex(),
		BridgeFundedId: common.BytesToHash(crypto.Keccak256([]byte(fmt.Sprint("BridgeFunded(address)")))).Hex(),
		PaidId: common.BytesToHash(crypto.Keccak256([]byte(fmt.Sprint("Paid(address,uint256)")))).Hex(),
	}

	actual := readAbi(false)

	var out bytes.Buffer
	if actual.DepositId != expected.DepositId {
		out.WriteString(fmt.Sprintf("%s -- got: %v expected: %v\n", "Deposit", actual.DepositId, expected.DepositId))
	}
	if actual.CreationId != expected.CreationId {
		out.WriteString(fmt.Sprintf("%s -- got: %v expected: %v\n", "Create", actual.CreationId, expected.CreationId))
	}
	if actual.WithdrawId != expected.WithdrawId {
		out.WriteString(fmt.Sprintf("%s -- got: %v expected: %v\n", "Withdraw", actual.WithdrawId, expected.DepositId))
	}
	if actual.BridgeFundedId != expected.BridgeFundedId {
		out.WriteString(fmt.Sprintf("%s -- got: %v expected: %v\n", "BridgeFunded", actual.BridgeFundedId, expected.DepositId))
	}
	if actual.PaidId != expected.PaidId {
		out.WriteString(fmt.Sprintf("%s -- got: %v expected: %v\n", "PaidId", actual.PaidId, expected.DepositId))
	}

	if out.String() != "" {
		t.Fatalf(out.String())
	}
}

func TestExists(t *testing.T) {
	path := "./main.go"
	ok, err := exists(path)
	if err != nil {
		t.Fatalf("Error checking if file exists: %s", err)
	}
	if !ok {
		t.Fatalf("File %s does not exists", path)
	}
}
