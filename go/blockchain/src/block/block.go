package block

import (
	"fmt"

	"crypto/sha256"
	"time"
)

type Block struct {
	index        int
	timestamp    int64
	data         string
	hash         [32]byte
	previousHash [32]byte
}

func NewGenesisBlock(data string) *Block {
	newBlock := &Block{
		index:        0,
		timestamp:    time.Now().Unix(),
		data:         data,
		previousHash: [32]byte{},
	}

	newBlock.hash = newBlock.ComputeHash()
	return newBlock
}

func NewBlock(data string, previousBlock *Block) *Block {
	newBlock := &Block{
		index:        previousBlock.GetIndex() + 1,
		timestamp:    time.Now().Unix(),
		data:         data,
		previousHash: previousBlock.GetHash(),
	}

	newBlock.hash = newBlock.ComputeHash()
	return newBlock
}

func (b *Block) GetIndex() int {
	return b.index
}

func (b *Block) GetData() string {
	return b.data
}

func (b *Block) GetHash() [32]byte {
	return b.hash
}

func (b *Block) GetPreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) ComputeHash() [32]byte {
	stringToConvert := fmt.Sprintf("%d%x%d%s", b.index, b.previousHash, b.timestamp, b.data)
	return sha256.Sum256([]byte(stringToConvert))
}

func (b *Block) Dump() {
	fmt.Printf("{\n")
	fmt.Printf("    Index: %d\n", b.index)
	fmt.Printf("    Timestamp: %d\n", b.timestamp)
	fmt.Printf("    Data: %s\n", b.data)
	fmt.Printf("    Hash: %x\n", b.hash)
	fmt.Printf("    PreviousHash: %x\n", b.previousHash)
	fmt.Printf("}\n")
}
