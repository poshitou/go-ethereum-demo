package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethernum-demo/03.contract"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xB53D4d29b1BEF33602EE71249af56c61c22879Dd")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6603929),
		ToBlock:   big.NewInt(6603929),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := store.StoreMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		event, err := contractABI.Unpack("ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		for _, event := range event {
			uint8s := event.([32]uint8)
			fmt.Println(string(uint8s[:]))
		}

		var topics [4]string
		for i, topic := range vLog.Topics {
			topics[i] = topic.Hex()
		}
		fmt.Println(topics)

		eventSignature := []byte("ItemSet(bytes32,bytes32)")

		hash := crypto.Keccak256Hash(eventSignature)

		fmt.Println(hash.Hex())
	}

}
