package counter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/erfan-flash/Learning-Go/display"
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
			result := GetCounts(strings.NewReader(tc.input)).words
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
			result := GetCounts(r).lines
			if result != tc.expect {
				t.Error("Expected:", tc.expect , "Got:",result)
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testcases:= [] struct{
		name string
		input string
		expect int
	}{{
		name: "five words",
		input: "one two three four five",
		expect: 23,
	},{
		name: "Empty",
		input: "",
		expect: 0,
	},{
		name: "All spaces",
		input: "     ",
		expect: 5,
	},{
		name: "New lines and words",
		input: "one\ntwo\nthree\nfour\n\t",
		expect: 20,
	},{
		name: "Unicode",
		input: "ãé",
		expect: 4,
	}}
	for _, tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			r :=strings.NewReader(tc.input)
			result := GetCounts(r).bytes
			if result != tc.expect {
				t.Error("Expected:", tc.expect , "Got:",result)
			}
		})
	}
}

func TestGetCounts(t *testing.T){
	testcases := [] struct {
		name string
		input string
		expect Counts
	}{}

	for _, tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := GetCounts(r)
			if result != tc.expect{
				t.Errorf("Expected: %v , Got: %v" , tc.expect, result)
			}
		} )
	}
}

func TestPrintCounts (t *testing.T){
	type inputs struct{
		counts Counts
		filename []string
		opts display.NewOptions
	}
	testcases := []struct {
		name string
		input inputs
		expect string
	}{
		{
			name: "simple five words.txt",
			input: inputs{
				counts:	Counts{
					lines: 1,
					words: 5,
					bytes: 23,
				},
				filename: []string{"words.txt"},
			},
			expect: "1\t5\t23\t words.txt\n",
		},
		{
			name: "Empty filename",
			input: inputs{
				counts: Counts{
					lines: 2,
					words: 20,
					bytes: 25,
				},
				opts: display.NewOptions{
					ShowBytes: true,
					ShowLines: true,
					ShowWords: true,
				},
			},
			expect: "2\t20\t25\t \n",
		},{
			name: "show lines words.txt",
			input: inputs{
				counts:Counts{
					lines: 1,
					words: 5,
					bytes: 23,
				},
				filename: []string{"words.txt"},
				opts: display.NewOptions{
					ShowBytes: false,
					ShowLines: true,
					ShowWords: false,
				},
			},
			expect: "1\t words.txt\n",
		},{
			name: "show bytes words.txt",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 23,
				},
				filename: []string{"words.txt"},
				opts: display.NewOptions{
					ShowBytes: true,
					ShowLines: false,
					ShowWords: false,
				},
			},
			expect: "23\t words.txt\n",
		},{
			name: "show words words.txt",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 23,
				},
				filename: []string{"words.txt"},
				opts: display.NewOptions{
					ShowBytes: false,
					ShowLines: false,
					ShowWords: true,
				},
			},
			expect: "5\t words.txt\n",
		},
	}

	for _, tc :=range testcases{
		t.Run(tc.name , func(t *testing.T) {
			buffer := &bytes.Buffer{}
			tc.input.counts.Print(buffer  ,display.New(tc.input.opts) ,tc.input.filename...)

			if tc.expect != buffer.String(){
				t.Errorf("Expected: %v  Got : %v", tc.expect , buffer.String())
			}
		})
	}
}

func TestAddCounts(t *testing.T){
	type inputs struct{
		counts Counts
		other Counts
	}
	testCases := []struct {
		name string
		input inputs
		expect Counts
	}{
		{
			name: "Add by one",
			input: inputs{
				counts: Counts{
					lines: 1,
					words: 5,
					bytes: 25,
				},
				other: Counts{
					lines: 1,
					words: 1,
					bytes: 1,
				},
			},
			expect: Counts{
				lines: 2,
				words: 6,
				bytes: 26,
			},
		},
	}
	for _ , tc := range testCases{
		t.Run(tc.name , func(t *testing.T) {
			total := tc.input.counts
			total = total.Add(tc.input.other)
			if total != tc.expect{
				t.Errorf("Expected: %v Got: %v" , tc.expect , total)
			}
		})
	}
}