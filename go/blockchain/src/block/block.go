package block

import (
	"fmt"

	"blockchain/src/transaction"
	"crypto/sha256"
	"time"
)

type Block struct {
	timestamp    int64
	transactions []*transaction.Transaction
	hash         [32]byte
	previousHash [32]byte
	nonce        int
}

func NewGenesisBlock(transactions []*transaction.Transaction) *Block {
	newBlock := &Block{
		timestamp:    time.Now().Unix(),
		transactions: transactions,
		previousHash: [32]byte{},
	}

	newBlock.hash = newBlock.ComputeHash()
	return newBlock
}

func NewBlock(transactions []*transaction.Transaction, difficulty int64, previousBlock *Block) *Block {
	newBlock := &Block{
		timestamp:    time.Now().Unix(),
		transactions: transactions,
		previousHash: previousBlock.GetHash(),
	}

	newBlock.Mine(difficulty)
	return newBlock
}

func (b *Block) GetTransactions() []*transaction.Transaction {
	return b.transactions
}

func (b *Block) GetHash() [32]byte {
	return b.hash
}

func (b *Block) GetPreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) ComputeHash() [32]byte {
	stringToConvert := fmt.Sprintf("%x%d%d", b.previousHash, b.timestamp, b.nonce)
	for _, t := range b.transactions {
		stringToConvert += fmt.Sprintf("%s%s%d", t.GetFromAddress(), t.GetToAddress(), t.GetAmount())
	}
	return sha256.Sum256([]byte(stringToConvert))
}

func (b *Block) Dump() {
	fmt.Printf("{\n")
	fmt.Printf("    Timestamp: %d\n", b.timestamp)
	fmt.Printf("    Transactions: %v\n", b.transactions)
	fmt.Printf("    Hash: %x\n", b.hash)
	fmt.Printf("    PreviousHash: %x\n", b.previousHash)
	fmt.Printf("}\n")
}

func (b *Block) Mine(difficulty int64) {
	newHash := b.ComputeHash()
	for amountOfZerosAtBeginning(newHash) < difficulty {
		b.nonce++
		newHash = b.ComputeHash()
	}

	b.hash = newHash
}

func amountOfZerosAtBeginning(hash [32]byte) int64 {
	var res int64
	for hash[res] == 0 {
		res++
	}

	return res
}
