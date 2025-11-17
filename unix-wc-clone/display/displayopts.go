package display

import "strings"

type Options struct {
	args NewOptions
}
type NewOptions struct{
	ShowHeaders bool
	ShowBytes   bool
	ShowWords   bool
	ShowLines   bool
}
func New (args NewOptions) Options{
	return Options{
		args: args,
	}
}
func (d Options) ShouldShowHeaders () bool{
	return d.args.ShowHeaders
}
func (d Options) ShowHeaders() string {
	headers := []string{}
	if !d.args.ShowLines && !d.args.ShowBytes&& !d.args.ShowWords {
		return "Lines\tWords\tBytes\t"
	}
	if d.args.ShowLines {
		headers = append(headers, "Lines")
	}
	if d.args.ShowWords {
		headers = append(headers, "Words")
	}
	if d.args.ShowBytes{
		headers = append(headers, "Bytes")
	}
	what := strings.Join(headers, "\t") + "\t"
	return what
}
func (d Options) ShouldShowLines() bool {
	if !d.args.ShowBytes&& !d.args.ShowLines && !d.args.ShowWords {
		return true
	}
	return d.args.ShowLines
}
func (d Options) ShouldShowWords() bool {
	if !d.args.ShowLines && !d.args.ShowBytes&& !d.args.ShowWords {
		return true
	}
	return d.args.ShowWords
}
func (d Options) ShoulShowBytes() bool {
	if !d.args.ShowBytes&& !d.args.ShowLines && !d.args.ShowWords {
		return true
	}
	return d.args.ShowBytes
}
