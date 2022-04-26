package blockchain

import "testing"

func TestNewBlockchainHasGenesisBlock(t *testing.T) {
	bc := NewBlockchain()
	b := bc.GetLatestBlock()

	if b.GetIndex() != 0 || b.GetData() != "Genesis" {
		t.Fatal("New blockchain should contain genesis block\n")
	}
}