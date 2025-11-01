package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	data , err := os.ReadFile("./words.txt")
	log.SetFlags(0)
	if err != nil{
		log.Fatalln("Failed to read file:", err)
	}
	fmt.Println(CountWords(data))
}

func CountWords(data []byte) int{
	words := bytes.Fields(data)
	return len(words)
}