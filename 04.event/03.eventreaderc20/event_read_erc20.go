package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "go-ethernum-demo/03.contract/erc20"
	"log"
	"math/big"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/A54-3LMyS2s248AfNl7c90Cf-G0cIbJ_")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xA1d7f71cbBb361A77820279958BAC38fC3667c1a")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2400078),
		ToBlock:   big.NewInt(3046791),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := token.TokenMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	logApprovalSig := []byte("Approval(address,address,uint256)")
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)

	for _, vLog := range logs {
		var transferEvent LogTransfer
		var approvalEvent LogApproval
		event, err := contractABI.Unpack("Transfer", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			fmt.Printf("transfer from : %s \n", transferEvent.From)

			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("transfer to : %s \n", transferEvent.To)

			transferEvent.Tokens = event[0].(*big.Int)
			fmt.Printf("tokens : %d \n", transferEvent.Tokens)

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			fmt.Printf("token owner : %s \n", approvalEvent.TokenOwner)

			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("token spender : %s \n", approvalEvent.Spender)

			approvalEvent.Tokens = event[0].(*big.Int)
			fmt.Printf("tokens : %d \n", approvalEvent.Tokens)
		}

	}

}
