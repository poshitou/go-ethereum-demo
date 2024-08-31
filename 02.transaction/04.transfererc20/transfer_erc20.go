package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		panic(err)
	}

	//配置私钥
	privateKey, err := crypto.HexToECDSA("xxx")
	if err != nil {
		panic(err)
	}

	//获取私钥对应的公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//获取公钥对应的账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//获取随机数nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	//获取gasprice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	toAddress := common.HexToAddress("0xEe014d7DfeB2e46Fef57CA4aDa42e79397edA76e")
	tokenAddress := common.HexToAddress("0x4E71E941878CE2afEB1039A0FE16f5eb557571C8")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	estimatedGasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})

	gasLimit := uint64(float64(estimatedGasLimit) * 10)

	if err != nil {
		panic(err)
	}
	fmt.Println(gasLimit)

	txData := types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     data,
	}

	tx := types.NewTx(&txData)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
