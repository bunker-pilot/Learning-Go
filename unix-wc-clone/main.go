package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)


func main() {
	opts := DisplayOptions{}
	log.SetFlags(0)
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
		count.Print(os.Stdout,opts,name)
		totals = totals.Add(count)
	}
	if len(filenames) == 0{
		GetCounts(os.Stdin).Print( os.Stdout, opts, "")
	}
	if len(filenames) >1 {
		totals.Print(os.Stdout,opts,"totals")
		}
	if errorhappend{
		os.Exit(1)
	}

}
