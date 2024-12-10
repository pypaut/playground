package d03

import (
	"reflect"
	"testing"
)

func TestScanMemory(t *testing.T) {
	inputMemory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	got := ScanMemory(inputMemory)

	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}

func TestExtractMulInstructions(t *testing.T) {
	inputMemory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []string{
		"mul(2,4)",
		"mul(5,5)",
		"mul(11,8)",
		"mul(8,5)",
	}
	got := ExtractMulInstructions(inputMemory)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestExecuteInstruction(t *testing.T) {
	cases := []struct {
		instruction string
		result      int
	}{
		{"mul(2,4)", 8},
		{"mul(5,5)", 25},
		{"mul(11,8)", 88},
		{"mul(8,5)", 40},
	}

	for _, c := range cases {
		got := ExecuteInstruction(c.instruction)
		if got != c.result {
			t.Fatalf("expected %d, got %d", c.result, got)
		}
	}

}

func TestScanMemoryWithDos(t *testing.T) {
	inputMemory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48
	got := ScanMemoryWithDos(inputMemory)

	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
