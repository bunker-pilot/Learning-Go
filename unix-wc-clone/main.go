package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	total := 0
	filenames := os.Args[1:]
	errorhappend := false
	for _ , name := range filenames{
		count , err := CountFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			continue
		}
		fmt.Println(name , ":", count)
		total += count.Words
	}
	if len(filenames) == 0{
		input := os.Stdin
		counts := GetCounts(input)
		fmt.Printf("words: %v, lines: %v, bytes: %v" , counts.Words , counts.Lines , counts.Bytes)
	}
	if len(filenames) >1 {
			fmt.Println("total:", total)
		}
	if errorhappend{
		os.Exit(1)
	}

}
