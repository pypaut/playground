package main

import (
	"fmt"

	"blockchain/src/blockchain"
)

func main() {
	myBlockchain := blockchain.NewBlockchain()
	myBlockchain.AddNewBlock("Second block")
	myBlockchain.Dump()
	isValid, errMessage := myBlockchain.IsValid()
	fmt.Printf("Blockchain is correct: %t (%s)\n", isValid, errMessage)
}
