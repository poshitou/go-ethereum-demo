package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xB53D4d29b1BEF33602EE71249af56c61c22879Dd")

	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(bytecode))
}
