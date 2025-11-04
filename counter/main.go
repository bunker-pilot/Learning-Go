package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	total := 0
	count := 0
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

func CountWordsInFile(filename string) (int, error){
	data , err := os.Open(filename)
	if err != nil{
		return 0 , err
	}
	defer data.Close()
	return  CountWords(data) , nil
}

func CountWords(file io.Reader) int  {
	word := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		word +=1
	}
	return word 
}