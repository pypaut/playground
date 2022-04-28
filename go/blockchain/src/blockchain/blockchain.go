package blockchain

import (
	"blockchain/src/block"
	"blockchain/src/transaction"
)

type Blockchain struct {
	chain      []*block.Block
	difficulty int64
}

func NewBlockchain(difficulty int64) *Blockchain {
	return &Blockchain{
		chain:      []*block.Block{block.NewGenesisBlock([]*transaction.Transaction{})},
		difficulty: difficulty,
	}
}

func (bc *Blockchain) AddNewBlock(transactions []*transaction.Transaction) {
	newBlock := block.NewBlock(transactions, bc.difficulty, bc.GetLatestBlock())
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
		if b.GetHash() != b.ComputeHash() {
			return false, "Wrong hash"
		}

		if i > 0 && b.GetPreviousHash() != bc.chain[i-1].ComputeHash() {
			return false, "Wrong previous hash"
		}
	}

	return true, ""
}
