package compute

import (
	"fmt"
	"testing"
)

func TestConvertToRomans(t *testing.T) {
	cases := []struct{
		name string
		arab Arab
		roman Roman
	}{
		{
			name: "1 is I",
			arab: Arab{1},
			roman: Roman{"I"},
		},
		{
			name: "2 is II",
			arab: Arab{2},
			roman: Roman{"II"},
		},
		{
			name: "3 is III",
			arab: Arab{3},
			roman: Roman{"III"},
		},
		{
			name: "4 is IV",
			arab: Arab{4},
			roman: Roman{"IV"},
		},
		{
			name: "5 is V",
			arab: Arab{5},
			roman: Roman{"V"},
		},
		{
			name: "9 is IX",
			arab: Arab{9},
			roman: Roman{"IX"},
		},
		{
			name: "39 is XXXIX",
			arab: Arab{39},
			roman: Roman{"XXXIX"},
		},
		{
			name: "789 is DCCLXXXIX",
			arab: Arab{789},
			roman: Roman{"DCCLXXXIX"},
		},
		{
			name: "2421 is MMCDXXI",
			arab: Arab{2421},
			roman: Roman{"MMCDXXI"},
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s", c.name), func(t* testing.T) {
			romanResult := c.arab.ConvertToRoman()
			if romanResult != c.roman {
				t.Fatalf("Expected %v, got %v\n", c.roman, romanResult)
			}
		})
	}
}