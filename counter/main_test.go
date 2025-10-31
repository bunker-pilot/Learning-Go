package main

import "testing"

func TestCountwords(t *testing.T) {
	input := "One two three"
	expected := 3
	result := CountWords([]byte(input))
	if result != expected{
		t.Error("Expected:", expected, "got:", result)
	}
	input = ""
	expected = 0
	result = CountWords([]byte(input))
	if result != expected{
		t.Error("expected:", expected, "got:", result)
	}
	input = " "
	expected = 0
	result = CountWords([]byte(input))
	if result != expected{
		t.Error("expected:", expected, "got:", result)
	}
}