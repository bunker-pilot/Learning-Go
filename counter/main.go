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
	if len(data) ==0{
		return counter
	}
	for _ , i := range data{
		if i == ' '{
			counter +=1
		}
	}
	counter++
	return counter
}