package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("xxx")

	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000)

	gasLimit := uint64(21000) // in units

	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xEe014d7DfeB2e46Fef57CA4aDa42e79397edA76e")

	var data []byte

	txData := types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Value:    value,
		Data:     data,
	}

	tx := types.NewTx(&txData)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println(rawTxHex)
}
