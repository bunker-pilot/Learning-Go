package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data , err := os.Open("./words.txt")
	log.SetFlags(0)
	if err != nil{
		log.Fatalln("Failed to read file:", err)
	}
	defer data.Close()
	fmt.Print(CountWords(data))
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