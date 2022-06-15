package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
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

// singleton pattern => 변수를 직접 드러내지 않고 함수를 통해 드러내는 것
func (b block) GetData() string {
	return b.data
}
func (b block) GetHash() string {
	return b.hash
}
func (b block) GetPrevHash() string {
	return b.prevHash
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

//CreateBlock에서 만드는 Blcock의 해쉬를 계산하고 선언해줌
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

//가장 최신의 블록의 해쉬를 가져옴
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

//블록을 만듦
func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

//블록체인을 호출할 때, 처음 호출한다면 b를 만들어줌.
func GetBlockchain() *blockchain {
	if b == nil {
		//once.DO(func()) 는 func()가 몇번, 언제 호출되든 한번만 실행되게 함.
		once.Do(func() {
			// 이 코드는 단 한번만 실행되어야함 => sync Once 사용
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
