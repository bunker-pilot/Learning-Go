package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/erfan-flash/Learning-Go/counter"
	"github.com/erfan-flash/Learning-Go/display"
)


func main() {
	newopts := display.NewOptions{}
	log.SetFlags(0)
	wr := tabwriter.NewWriter(os.Stdout , 0, 8, 1 , ' ', tabwriter.AlignRight)
	flag.BoolVar(
		&newopts.ShowHeaders , "headers" , false, "Used to toggle whether or not to show the headers")
	flag.BoolVar(
		&newopts.ShowWords,"w" , false ,"Used to toggle whether or not to show the word count")
	flag.BoolVar(
		&newopts.ShowBytes,"c" , false ,"Used to toggle whether or not to show the byte count")
	flag.BoolVar(
		&newopts.ShowLines,"l" , false ,"Used to toggle whether or not to show the line count")
	opts := display.New(newopts)
	flag.Parse()
	totals := counter.Counts{}
	filenames := flag.Args()
	errorhappend := false
	for _ , name := range filenames{
		count , err := counter.CountFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			continue
		}
		count.Print(wr,opts,name)
		totals = totals.Add(count)
	}
	if len(filenames) == 0{
		counter.GetCounts(os.Stdin).Print(wr, opts, "")
	}
	if len(filenames) >1 {
		totals.Print(wr,opts,"totals")
		}
	wr.Flush()
	if errorhappend{
		os.Exit(1)
	}

}
