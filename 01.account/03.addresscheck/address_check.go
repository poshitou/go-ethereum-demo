package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	// check the address if valid
	//通过正则来判断地址是否合法 ：必须是0x开头的40位16进制字符
	reg := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid : %v \n", reg.MatchString("0x71c7656ec7ab88b098defb751b7401b5f6d8976f"))
	fmt.Printf("is valid : %v \n", reg.MatchString("0x71g7656ec7ab88b098defb751b7401b5f6d8976f"))

	/**
	  Check if the address is an account or smart contract
	    如果地址存储了字节码，那么该地址是一个智能合约，否则就是一个账户地址
	*/

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract : %v \n", isContract)

	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract = len(bytecode) > 0
	fmt.Printf("is contract : %v \n", isContract)
}
