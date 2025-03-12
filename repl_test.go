package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "something else       ",
			expected: []string{"something", "else"},
		},
		{
			input: " WHY AM I YELLING  ",
			expected: []string{"why", "am", "i", "yelling"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"length and of expected: %v, and of received: %v, do not match", 
				c.expected, 
				actual,
			)
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match %v", word, expectedWord)
			}
		}
	}
}

