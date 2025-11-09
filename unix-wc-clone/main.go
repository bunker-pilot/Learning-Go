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
		count , err := CountWordsInFile(name)
		if err != nil{
			fmt.Fprintln(os.Stderr , "counter:", err)
			errorhappend = true
			continue
		}
		fmt.Println(name , ":", count)
		total += count
	}
	if len(filenames) == 0{
		countedWords := CountWords(os.Stdin)
		fmt.Println(countedWords)
	}
	if len(filenames) >1 {
			fmt.Println("total:", total)
		}
	if errorhappend{
		os.Exit(1)
	}

}
