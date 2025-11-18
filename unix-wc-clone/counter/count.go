package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/erfan-flash/Learning-Go/display"
)
type Counts struct{
	bytes int
	words int
	lines int
}
func (c Counts) Add(other Counts) Counts {
	c.lines += other.lines
	c.words += other.words
	c.bytes += other.bytes
	return c
}
func (c Counts) Print(w io.Writer ,opts display.Options ,suffixes ...string){
	xs := []string{}
	var what string
	if opts.ShouldShowLines(){
		xs = append(xs, strconv.Itoa(c.lines))
	}
	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.words))
	}
	if opts.ShoulShowBytes(){
		xs =append(xs, strconv.Itoa(c.bytes))
	}
	if opts.ShouldShowHeaders(){
		what = opts.ShowHeaders()
		fmt.Fprintln(w, what)
	}
	line :=strings.Join(xs , "\t") + "\t"
	suffixesStr := strings.Join(suffixes , " ")
	
	fmt.Fprintln(w  , line , suffixesStr)
}
func GetCounts(file io.Reader) Counts{
	res := Counts{}
	isInsideWord :=false
	reader := bufio.NewReader(file)
	for {
		r , size, err :=reader.ReadRune()
		if err !=nil{
			break
		}
		res.bytes += size
		if r == '\n'{
			res.lines +=1
		}
		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord{
			res.words +=1
		}
		isInsideWord =!isSpace
	}
	return res
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

func Countwords(r io.Reader) int {
	word := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word += 1
	}
	return word
}
func Countlines(r io.Reader) int {
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