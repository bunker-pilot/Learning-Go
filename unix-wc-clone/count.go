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

func CountWords(r io.Reader) int {
	word := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word += 1
	}
	return word
}
func CountLines(r io.Reader) int {
	lines := 0 
	reader := bufio.NewReader(r)
	for {
		r ,_ ,err:=reader.ReadRune()
		if err != nil{
			break
		}
		if r =='\n'{
			lines+=1
		}
	}
	return lines
}

func Countbytes(r io.Reader) int {
	counter , _ := io.Copy(io.Discard , r)
	return int(counter)
}