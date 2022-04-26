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
	previousHash []byte
}

func NewGenesisBlock(data string) *Block {
	newBlock := &Block{
		index:        0,
		timestamp:    time.Now().Unix(),
		data:         data,
		previousHash: []byte{byte('0')},
	}

	newBlock.hash = newBlock.ComputeHash()
	return newBlock
}

func (b *Block) GetIndex() int {
	return b.index
}

func (b *Block) ComputeHash() [32]byte {
	stringToConvert := fmt.Sprintf("%d%x%d%s", b.index, b.previousHash, b.timestamp, b.data)
	return sha256.Sum256([]byte(stringToConvert))
}
