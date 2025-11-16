package display

import "strings"

type Options struct {
	ShowHeaders bool
	ShowBytes   bool
	ShowWords   bool
	ShowLines   bool
}

func (d Options) ShouldShowHeaders() string {
	headers := []string{}
	if !d.ShowLines && !d.ShowBytes && !d.ShowWords {
		return "Lines\tWords\tBytes\t"
	}
	if d.ShowLines {
		headers = append(headers, "Lines")
	}
	if d.ShowWords {
		headers = append(headers, "Words")
	}
	if d.ShowBytes {
		headers = append(headers, "Bytes")
	}
	what := strings.Join(headers, "\t") + "\t"
	return what
}
func (d Options) ShouldShowLines() bool {
	if !d.ShowBytes && !d.ShowLines && !d.ShowWords {
		return true
	}
	return d.ShowLines
}
func (d Options) ShouldShowWords() bool {
	if !d.ShowLines && !d.ShowBytes && !d.ShowWords {
		return true
	}
	return d.ShowWords
}
func (d Options) ShoulShowBytes() bool {
	if !d.ShowBytes && !d.ShowLines && !d.ShowWords {
		return true
	}
	return d.ShowBytes
}
