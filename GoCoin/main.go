package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
	/*
		B1
			B1Hash = (data +  'X')
		B2
			B2hash = (data + B1Hash)
		B3
			B3Hash = (data + B2Hash)
		==> 이전 블록을 변경하면 이 블록이 무의미해짐 ==> 변경 불가능
	*/
}

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(hexHash)

}
