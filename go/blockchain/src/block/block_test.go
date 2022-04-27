package block

import (
	"fmt"

	"testing"
)

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

func TestAmountOfZerosAtBeginning(t *testing.T) {
	cases := []struct {
		name      string
		hash      [32]byte
		nbOfZeros int64
	}{
		{
			name:      fmt.Sprintf("%x", [32]byte{0, 0, 0, 0, 0, 3, 4}),
			hash:      [32]byte{0, 0, 0, 0, 0, 3, 4},
			nbOfZeros: 5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expected := c.nbOfZeros
			actual := amountOfZerosAtBeginning(c.hash)

			if expected != actual {
				t.Fatalf("Expected %d zeros, got %d\n", expected, actual)
			}
		})
	}
}

func TestMineBlock(t *testing.T) {
	cases := []struct {
		name       string
		difficulty int64
	}{
		{
			name:       "Difficulty 1",
			difficulty: 1,
		},
		{
			name:       "Difficulty 2",
			difficulty: 2,
		},
		{
			name:       "Difficulty 3",
			difficulty: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := NewGenesisBlock("my data")
			b.Mine(c.difficulty)

			actual := amountOfZerosAtBeginning(b.hash)
			expected := c.difficulty

			if actual != expected {
				t.Fatalf("Expected %d zeros, got %d\n", expected, actual)
			}
		})
	}

}
