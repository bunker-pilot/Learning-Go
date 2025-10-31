package main

import "testing"

func TestCountwords(t *testing.T) {

	testCases := []struct{
		name string
		input string
		expected int
	}{
		{
			name: "three words",
			input : "One two three",
			expected : 3,
		},
		{
			name: "Nothing",
			input: "",
			expected: 0,
		},
		{
			name: "single space",
			input: " ",
			expected: 0,
		},
	}
	for _ , tc := range testCases{
		t.Run(tc.name , func(t *testing.T) {
			result := CountWords([]byte(tc.input))
		if result != tc.expected{
		t.Error("Expected:", tc.expected, "got:", result)
		}})
	}
}