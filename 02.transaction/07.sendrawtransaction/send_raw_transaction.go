package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86f2c851010889c7382520894ee014d7dfeb2e46fef57ca4ada42e79397eda76e87038d7ea4c68000808401546d71a0c73424f9e815d2c63eea5880e3403d53e7f4a97feb00a29df0444d8682b312dfa01bc8da3c3f453517a99cf33a400a7f60e18cd4d75fa2ee8b4218166c1a63de8e"
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatal(err)
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
