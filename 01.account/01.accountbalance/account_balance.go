package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {

	ethClient, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("we have a connection")
	//_ = ethClient
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	balance, err := ethClient.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	blockNumber := big.NewInt(5532993)
	balanceAt, err := ethClient.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)

	fBalance := new(big.Float)
	fBalance.SetString(balanceAt.String())

	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	log.Printf("account %s has %f ETH\n", account.Hex(), ethValue)
	fmt.Printf("account %s has %g ETH\n", account.Hex(), ethValue)

}
