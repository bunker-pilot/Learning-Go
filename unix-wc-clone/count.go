package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
type Counts struct{
	Bytes int
	Words int
	Lines int
}
type DisplayOptions struct{
	ShowBytes bool
	ShowWords bool
	ShowLines bool
}
func (d DisplayOptions) ShouldShowLines()bool{
	if !d.ShowBytes && !d.ShowLines && !d.ShowWords{
		return true
	}
	return d.ShowLines
}
func (d DisplayOptions) ShouldShowWords()bool{
	if !d.ShowLines && !d.ShowBytes && !d.ShowWords{
		return true
	}
	return d.ShowWords
}
func (d DisplayOptions) ShoulShowBytes()bool{
	if !d.ShowBytes && !d.ShowLines && !d.ShowWords{
		return  true
	}
	return d.ShowBytes
}

func (c Counts) Add(other Counts) Counts{
	c.Lines += other.Lines
	c.Words += other.Words
	c.Bytes += other.Bytes
	return c
}
func (c Counts) Print(w io.Writer ,opts DisplayOptions ,suffixes ...string){
	xs := []string{}
	if opts.ShouldShowLines(){
		xs = append(xs, strconv.Itoa(c.Lines))
	}
	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.Words))
	}
	if opts.ShoulShowBytes(){
		xs =append(xs, strconv.Itoa(c.Bytes))
	}
	xs = append(xs, suffixes...)
	line :=strings.Join(xs , " ")
	fmt.Fprintln(w  , line)
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