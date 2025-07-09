package pokedex

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    " PikacHu ",
			expected: []string{"pikachu"},
		},
		{
			input:    " Charmander  BLAStOiSE	RaIcHu ",
			expected: []string{"charmander", "blastoise", "raichu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v results, got %v", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word :%s, got %s.", expectedWord, word)
			}
		}
	}
}
