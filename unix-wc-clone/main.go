package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"

	"github.com/erfan-flash/Learning-Go/counter"
	"github.com/erfan-flash/Learning-Go/display"
)
type FileCountResult struct{
	counter counter.Counts
	filename  string
}

func main() {
	newopts := display.NewOptions{}
	wg := sync.WaitGroup{}
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
	wg.Add(len(filenames))
	ch := make(chan  FileCountResult)
	errorhappend := false
	for _ , name := range filenames{
		go func (name string)  {
		defer wg.Done()
		count , err := counter.CountFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			return
		}
		ch <- FileCountResult{
			filename: name,
			counter: count,
		}
		
	}(name)
	}
	go func ()  {
		wg.Wait()
		close(ch)
	}()
	for res := range ch{
		res.counter.Print(wr,opts,res.filename)
		totals = totals.Add(res.counter)
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
