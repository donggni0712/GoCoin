package main

import (
	"fmt"

	"github.com/donggni0712/GoCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Seccond Block")
	chain.AddBlock("Third Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.GetData())
		fmt.Printf("Hash: %s\n", block.GetHash())
		fmt.Printf("PrevHash: %s\n\n", block.GetPrevHash())
	}
}
