package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	totals := Counts{}
	filenames := os.Args[1:]
	errorhappend := false
	for _ , name := range filenames{
		count , err := CountFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			continue
		}
		count.Print(os.Stdout,name)
		totals = Counts{
			Bytes: totals.Bytes + count.Bytes,
			Words: totals.Words + count.Words,
			Lines: totals.Lines + count.Lines,
		}
	}
	if len(filenames) == 0{
		GetCounts(os.Stdin).Print( os.Stdout,"")
	}
	if len(filenames) >1 {
		totals.Print(os.Stdout,"totals")
		}
	if errorhappend{
		os.Exit(1)
	}

}
