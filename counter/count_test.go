package main_test

import (
	"strings"
	"testing"

	counter "github.com/erfan-flash/Learning-Go"
)

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
		{
			name : "suffix",
			input: "hello   ",
			expected: 1,
		},
		{
			name: "Preix",
			input: "   Hekllo",
			expected: 1,
		},
	}
	for _ , tc := range testCases{
		t.Run(tc.name , func(t *testing.T) {
			result := counter.CountWords(strings.NewReader(tc.input))
		if result != tc.expected{
		t.Error("Expected:", tc.expected, "got:", result)
		}})
	}
}

func TestCountLines(t *testing.T) {
	testcases := []struct{
		name string
		input string
		expect int
	}{
		{
			name: "Simple words, 1 line",
			input: "One Two Three, Simple? \n",
			expect: 1,
		},
		{
			name: "Empty",
			input: "",
			expect: 0,
		},
		{
			name: "No new lines",
			input: "No new lines baby",
			expect: 0,
		},
	}
	for _, tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			r :=strings.NewReader(tc.input)
			result := counter.CountLines(r)
			if result != tc.expect {
				t.Error("Expected:", tc.expect , "Got:",result)
			}
		})
	}
}