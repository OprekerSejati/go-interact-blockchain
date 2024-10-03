package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "your-Infura-Rpc-linkhere"
var haradhatUrl = "http://localhost:8545"

func main() {

	// get the last blocknumber here
	// both mainnet dan local hardhat

	//client, err := ethclient.DialContext(context.Background(), infuraURL)
	client, err := ethclient.DialContext(context.Background(), haradhatUrl)
	if err != nil {
		log.Fatalf("Error to create a ether client:&v", err)
	}
	defer client.Close()
	//var block *types.Block

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:&v", err)
	}

	fmt.Println("block number : ", block.Number())

	// check for balance
	addr := "0xcd3B766CCDd6AE721141F452C550Ca635964ce71"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:&v", err)
	}
	fmt.Println("The balance:", balance)
}
