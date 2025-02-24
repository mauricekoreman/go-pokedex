package main

import "testing"

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
			input:    "hello word   ",
			expected: []string{"hello", "word"},
		},
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "this is a test   of things ",
			expected: []string{"this", "is", "a", "test", "of", "things"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, but got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected word %q, but got %q", expectedWord, word)
			}
		}
	}
}
