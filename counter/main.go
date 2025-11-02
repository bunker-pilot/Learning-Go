package main

import (
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
	Printfile(data)
}

func Printfile(file *os.File)  {
	const bufferSize =1024
	buffer := make([]byte , bufferSize)
	for{
		size , err := file.Read(buffer)
		if err !=nil{
			break
		}
		fmt.Print(string(buffer[:size]))
		
	}
}

func CountWords(data []byte) int{
	words := bytes.Fields(data)
	return len(words)
}