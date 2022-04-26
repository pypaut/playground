package blockchain

import (
	"blockchain/src/block"
)

type Blockchain struct {
	chain []*block.Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		chain: []*block.Block{block.NewGenesisBlock("Genesis")},
	}
}

func (bc *Blockchain) AddNewBlock(data string) {
	newBlock := block.NewBlock(data, bc.GetLatestBlock())
	bc.chain = append(bc.chain, newBlock)
}

func (bc *Blockchain) GetLatestBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Dump() {
	for _, b := range bc.chain {
		b.Dump()
	}
}
