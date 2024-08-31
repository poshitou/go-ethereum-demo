package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	//获取一个区块里的交易信息
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(5671744))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Transactions().Len())

	for _, tx := range block.Transactions() {
		//fmt.Println("-------------------------")
		//fmt.Println(tx.Hash().Hex())
		//fmt.Println(tx.Value().String())
		//fmt.Println(tx.Gas())
		//fmt.Println(tx.GasPrice().Uint64())
		//fmt.Println(tx.Nonce())
		//fmt.Println(tx.Data())
		//fmt.Println(tx.To().Hex())
		fmt.Println("-------------------------")

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		from, err := types.Sender(types.NewEIP155Signer(chainID), tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(from.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)
	}

	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	transactionCount, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for i := uint(0); i < transactionCount; i++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, pending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash().Hex(), pending)
}
