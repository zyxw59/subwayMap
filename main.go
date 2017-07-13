package main

import (
	"fmt"
	"github.com/zyxw59/subwayMap/corner"
	"io"
	"os"
)

type Printer struct {
	Writer io.Writer
}

func (p *Printer) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(p.Writer, a...)
}

func (p *Printer) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(p.Writer, format, a...)
}

func (p *Printer) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(p.Writer, a...)
}

func main() {
	var (
		paths               []corner.Path
		width, height, rsep float64
		canvas              Printer
	)
	width = 2000
	height = 2000
	rsep = 10.0
	canvas = Printer{os.Stdout}
	A := *corner.NewCorner(100, 100, corner.South, corner.South)
	B := *corner.NewCorner(100, 400, corner.South, corner.East)
	C := *corner.NewCorner(400, 400, corner.East, corner.East)
	corners := []corner.Corner{A, B, C}
	offsets := []int{0, 0}
	paths = append(paths, *corner.NewPath("abc", corners, offsets))
	printPaths(canvas, width, height, rsep, paths)
}

func printPaths(canvas Printer, width, height, rsep float64, paths []corner.Path) {
	canvas.Println(header)
	canvas.Printf(stylesheetfmt, "lines")
	canvas.Printf(svginitfmt, width, height)
	canvas.Println(svgns)
	// begin line definitions
	canvas.Println("<defs>")
	for _, p := range paths {
		canvas.Println(p.Path(rsep))
	}
	// end line definitions
	canvas.Println("</defs>")
	canvas.Println("</svg>")
}

const (
	header        = "<?xml version='1.0' encoding='utf-8' ?>"
	stylesheetfmt = "<?xml-stylesheet type='text/css' href='%s.css'?>\n"
	svginitfmt    = "<svg width=\"%v\" height=\"%v\"\n"
	svgns         = "xmlns='http://www.w3.org/2000/svg'\nxmlns:xlink='http://www.w3.org/1999/xlink'>"
)
