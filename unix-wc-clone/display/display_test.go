package display_test

import (
	"testing"

	"github.com/erfan-flash/Learning-Go/display"
)

func TestShowHeaders(t *testing.T){
	testcases := [] struct{
		name string
		input display.NewOptions
		expect string
	}{{
		name: "default inputs",
		expect: "Lines\tWords\tBytes\t",
	},
	{
		name: "lines only",
		input:   display.NewOptions{ShowLines: true},
		expect: "Lines\t",
	},
	{
		name: "words only",
		input:   display.NewOptions{ShowWords: true},
		expect: "Words\t",
	},
	{
		name: "bytes only",
		input:   display.NewOptions{ShowBytes: true},
		expect: "Bytes\t",
	},
	{
		name: "lines + words",
		input:   display.NewOptions{ShowLines: true, ShowWords: true},
		expect: "Lines\tWords\t",
	},
	{
		name: "lines + bytes",
		input:   display.NewOptions{ShowLines: true, ShowBytes: true},
		expect: "Lines\tBytes\t",
	},
	{
		name: "words + bytes",
		input:   display.NewOptions{ShowWords: true, ShowBytes: true},
		expect: "Words\tBytes\t",
	},
}
	for _ , tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			res := display.New(tc.input).ShowHeaders()
			if res != tc.expect{
				t.Errorf("Expected: %s  Got: %s" , tc.expect , res)
			}
		})
	}
}

func TestShouldShowLines(t *testing.T){
	testcases := [] struct {
		name  string
		input display.NewOptions
		expect bool
	}{{
		name: "Default",
		expect: true,
	},{
		name: "Only Lines",
		input: display.NewOptions{
			ShowLines: true,
		},
		expect: true,
	},{
		name: "Only Words",
		input: display.NewOptions{
			ShowWords: true,
		},
		expect: false,
	},{
		name:  "Only Bytes",
		input: display.NewOptions{
			ShowBytes: true,
		},
		expect: false,
	},
	}
	for _ , tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			res := display.New(tc.input).ShouldShowLines()
			if res != tc.expect {
				t.Errorf("Expected: %v  Got: %v" , tc.expect , res )
			}
		})
	}
}

func TestShouldShowWords (t *testing.T){
	testcases := []struct {
		name string
		input display.NewOptions
		expect bool
	}{{
		name: "Default",
		expect: true,
	},{
		name: "Only Lines",
		input: display.NewOptions{
			ShowLines: true,
		},
		expect: false,
	},{
		name: "Only Words",
		input: display.NewOptions{
			ShowWords: true,
		},
		expect: true,
	},{
		name:  "Only Bytes",
		input: display.NewOptions{
			ShowBytes: true,
		},
		expect: false,
	},
	}
	for _ , tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			res := display.New(tc.input).ShouldShowWords()
			if res != tc.expect {
				t.Errorf("Expected: %v  Got: %v" , tc.expect , res )
			}
		})
    }
}

func TestShouldShowBytes(t *testing.T){
	testcases := []struct {
		name string
		input display.NewOptions
		expect bool
	}{{
		name: "Default",
		expect: true,
	},{
		name: "Only Lines",
		input: display.NewOptions{
			ShowLines: true,
		},
		expect: false,
	},{
		name: "Only Words",
		input: display.NewOptions{
			ShowWords: true,
		},
		expect: false,
	},{
		name:  "Only Bytes",
		input: display.NewOptions{
			ShowBytes: true,
		},
		expect: true,
	},
	}
	for _ , tc := range testcases{
		t.Run(tc.name , func(t *testing.T) {
			res := display.New(tc.input).ShoulShowBytes()
			if res != tc.expect {
				t.Errorf("Expected: %v  Got: %v" , tc.expect , res )
			}
		})
    }
}