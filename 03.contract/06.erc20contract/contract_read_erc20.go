package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethernum-demo/03.contract"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")

	if err != nil {
		fmt.Println(err)
	}

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xA1d7f71cbBb361A77820279958BAC38fC3667c1a")

	instance, err := store.NewToken(tokenAddress, client)

	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress("0xA1d7f71cbBb361A77820279958BAC38fC3667c1a")
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(balance.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	fmt.Println(value)
}
