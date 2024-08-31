package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethernum-demo/03.contract"
	"log"
)

// 0xB53D4d29b1BEF33602EE71249af56c61c22879Dd
func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xB53D4d29b1BEF33602EE71249af56c61c22879Dd")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
}
