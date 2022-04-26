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

func (bc *Blockchain) IsValid() (bool, string) {
	for i, b := range bc.chain {
		if i != b.GetIndex() {
			return false, "Wrong index"
		}

		if b.GetHash() != b.ComputeHash() {
			return false, "Wrong hash"
		}

		if i > 0 && b.GetPreviousHash() != bc.chain[i-1].ComputeHash() {
			return false, "Wrong previous hash"
		}
	}

	return true, ""
}
