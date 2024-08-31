package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 查询最近的一个区块头信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 20621283

	// 查询最近的一个区块信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(20621283))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 20621283
	fmt.Println(block.Time())                // 1632187314
	fmt.Println(block.Difficulty().Uint64()) //
	fmt.Println(block.Hash().Hex())          //
	fmt.Println(len(block.Transactions()))   //

	//查询区块的交易数量
	transactionCount, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transactionCount) //
}
