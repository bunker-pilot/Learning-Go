package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
type Counts struct{
	Bytes int
	Words int
	Lines int
}

func (c Counts) Print(w io.Writer , filename string){
	fmt.Fprintf(w,"%v %v %v %v \n", c.Lines, c.Words, c.Bytes , filename)
}
func GetCounts(file io.ReadSeeker) Counts{
	const OffsetSeek = 0
	words := CountWords(file)
	file.Seek(OffsetSeek , io.SeekStart)
	lines := CountLines(file)
	file.Seek(OffsetSeek, io.SeekStart)
	bytes := Countbytes(file)
	return Counts{
		Bytes: bytes,
		Words: words,
		Lines: lines,
	}
}
func CountFile(filename string) (Counts, error) {
	data, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer data.Close()
	counts := GetCounts(data)
	return counts , nil
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