package main

import (
	"fmt"
	"os"
)

func main() {
	data , _ := os.ReadFile("./words.txt")
	content := string(data)

	i:=0
	start :=0
	words := []string{}
	for i < len(content){
		start = i
		for{
			if content[i] != 10 && content[i] !=32{
				i +=1
			} else{
				words = append(words, content[start:i])
				break
			}
		}
		
	}
	fmt.Println(len(words))
}