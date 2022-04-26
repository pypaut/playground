package block

import "testing"

func TestNewGenesisBlockHasIndex0(t *testing.T) {
	b := NewGenesisBlock("my_data")
	if b.GetIndex() != 0 {
		t.Fatalf("New genesis block should have index 0, but has %d\n", b.GetIndex())
	}
}

func TestGenesisBlockShaCorresponds(t *testing.T) {
	b := NewGenesisBlock("my_data")
	if b.hash != b.ComputeHash() {
		t.Fatal("New genesis block has a broken hash!")
	}
}
