package main

import (
	"blockchain/src/blockchain"
)

func main() {
	myBlockchain := blockchain.NewBlockchain()
	myBlockchain.AddNewBlock("Second block")
	myBlockchain.Dump()
}
