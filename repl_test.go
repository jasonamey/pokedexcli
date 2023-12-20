package main

import (
	"testing"
)

type Case struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []Case{
		{
			input: "HELlo World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}

}
