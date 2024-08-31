package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	//连接以太坊节点
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	//配置私钥
	privateKey, err := crypto.HexToECDSA("xxx")
	if err != nil {
		log.Fatal(err)
	}
	//获取私钥对应的账户地址和nonce
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

	//设置转账金额
	value := big.NewInt(100000000000000) //
	gasLimit := uint64(21000)
	//获取gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//定义eth收款账户地址
	toAddress := common.HexToAddress("0xEe014d7DfeB2e46Fef57CA4aDa42e79397edA76e")
	var data []byte

	//创建交易
	txData := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Value:    value,
		Data:     data,
	}

	tx := types.NewTx(txData)

	//获取ChainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//使用交易发送者的私钥签署交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("tx sent : %s", signedTx.Hash().Hex())

}
