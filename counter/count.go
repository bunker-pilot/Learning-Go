package main

import (
	"bufio"
	"io"
	"os"
)

func CountWordsInFile(filename string) (int, error) {
	data, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer data.Close()
	return CountWords(data), nil
}

func CountWords(file io.Reader) int {
	word := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word += 1
	}
	return word
}