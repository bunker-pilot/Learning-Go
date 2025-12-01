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
	err error
}

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
	ch := CountFiles(filenames)
	for res := range ch{
		if res.err !=nil{
			fmt.Fprintf(os.Stderr, "counter:", res.err)
			errorhappend = true
			continue
		}
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

func CountFiles (filenames []string) <-chan FileCountResult{
	ch := make(chan FileCountResult)
	wg := sync.WaitGroup{}
	wg.Add(len(filenames))
	for _ , filename := range filenames{
		go func (filename string)  {
			defer wg.Done()
			count , err:= counter.CountFile(filename)
		
		ch <- FileCountResult{
			filename: filename,
			counter: count,
			err: err,
		}
		
		}(filename)
	}
	go func ()  {
		wg.Wait()
		close(ch)
	}()
	return ch
}
