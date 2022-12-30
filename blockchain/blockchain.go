package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)


type Block struct {
	Data string
	Hash string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash()  {
	newHash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", newHash)
}

func getLastHash() string {
	length := len(GetBlockChain().blocks)
	if length == 0 {
		return ""
	}
	return GetBlockChain().blocks[length -1].Hash
}

func createNewBlock(data string) *Block{
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddNewBlock(data string) {
	b.blocks = append(b.blocks, createNewBlock(data))
}

func GetBlockChain() *blockchain{
	if b == nil{
		once.Do(func() {
			b = &blockchain{}
			b.AddNewBlock("First")
		})
		
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block{
	return b.blocks
}

