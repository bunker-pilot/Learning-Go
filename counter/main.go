package main

import (
	"bufio"
	"bytes"
	"fmt"
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
	fmt.Print(Countfile(data))
}

func Countfile(file *os.File) int  {
	word := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan(){
		word +=1
	}
	return word
}

func CountWords(data []byte) int{
	words := bytes.Fields(data)
	return len(words)
}