package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
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
	const bufferSize =5
	word := 0
	buffer := make([]byte , bufferSize)
	isWord := false
	for{
		size , err := file.Read(buffer)
		if err ==io.EOF{
			break
		}else if err != nil{
			log.Fatal(err)
		}
		isWord = !unicode.IsSpace(rune(buffer[0])) && isWord
		pancake :=CountWords(buffer[:size])
		if isWord{
			word -=1
		}
		word += pancake
		isWord = !unicode.IsSpace(rune(buffer[size -1]))
	}
	return word
}

func CountWords(data []byte) int{
	words := bytes.Fields(data)
	return len(words)
}