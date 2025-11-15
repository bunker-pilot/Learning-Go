package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)


func main() {
	opts := DisplayOptions{}
	log.SetFlags(0)
	wr := tabwriter.NewWriter(os.Stdout , 0, 8, 1 , ' ', tabwriter.AlignRight)
	flag.BoolVar(
		&opts.ShowHeaders , "headers" , false, "Used to toggle whether or not to show the headers")
	flag.BoolVar(
		&opts.ShowWords,"w" , false ,"Used to toggle whether or not to show the word count")
	flag.BoolVar(
		&opts.ShowBytes,"c" , false ,"Used to toggle whether or not to show the byte count")
	flag.BoolVar(
		&opts.ShowLines,"l" , false ,"Used to toggle whether or not to show the line count")
	flag.Parse()
	totals := Counts{}
	filenames := flag.Args()
	errorhappend := false
	for _ , name := range filenames{
		count , err := CountFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			continue
		}
		count.Print(wr,opts,name)
		totals = totals.Add(count)
	}
	if len(filenames) == 0{
		GetCounts(os.Stdin).Print(wr, opts, "")
	}
	if len(filenames) >1 {
		totals.Print(wr,opts,"totals")
		}
	wr.Flush()
	if errorhappend{
		os.Exit(1)
	}

}
