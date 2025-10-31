package main

import (
	"fmt"
	"os"
)

func main() {
	data , _ := os.ReadFile("./words.txt")
	fmt.Println(CountWords(data))
}

func CountWords(data []byte) int{
	counter :=0
	for _ , i := range data{
		if i ==' '{
			counter +=1
		}
	}
	return counter
}