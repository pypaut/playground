package blockchain

import (
	"blockchain/src/block"
	"blockchain/src/transaction"
)

type Blockchain struct {
	chain               []*block.Block
	difficulty          int64
	pendingTransactions []*transaction.Transaction
	miningReward        int64
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

func (bc *Blockchain) CreateTransaction(t *transaction.Transaction) {
	bc.pendingTransactions = append(bc.pendingTransactions, t)
}

func (bc *Blockchain) GetBalance(address string) (balance int64) {
	for _, b := range bc.chain {
		for _, t := range b.GetTransactions() {
			if address == t.GetFromAddress() {
				balance -= t.GetAmount()
			} else if address == t.GetToAddress() {
				balance += t.GetAmount()
			}
		}
	}

	return balance
}

func (bc *Blockchain) MinePendingTransactions(rewardAddress string) {
	bc.AddNewBlock(bc.pendingTransactions)
	bc.pendingTransactions = []*transaction.Transaction{
		transaction.NewTransaction("", rewardAddress, bc.miningReward),
	}
}
