package main

import "testing"

func TestCountwords(t *testing.T) {
	input := "One two three "
	expected := 3
	result := CountWords([]byte(input))
	if result != expected{
		t.Fail()
	}
}