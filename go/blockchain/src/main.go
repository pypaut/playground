package main

import (
	"fmt"

	"blockchain/src/blockchain"
	"blockchain/src/transaction"
)

func main() {
	myBlockchain := blockchain.NewBlockchain(0)
	myBlockchain.AddNewBlock([]*transaction.Transaction{})
	myBlockchain.Dump()
	isValid, errMessage := myBlockchain.IsValid()
	fmt.Printf("Blockchain is correct: %t (%s)\n", isValid, errMessage)
}
