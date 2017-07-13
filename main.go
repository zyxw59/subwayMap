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

func (canvas *Printer) PrintInit(width, height float64) {
	canvas.Println(header)
	canvas.Printf(stylesheetfmt, "lines")
	canvas.Printf(svginitfmt, width, height)
	canvas.Println(svgns)
}

func (canvas *Printer) PrintDefs(rsep float64, paths ...[]corner.Path) {
	// begin line definitions
	canvas.Println("<defs>")
	for _, ps := range paths {
		for _, p := range ps {
			canvas.Println(p.Path(rsep))
		}
	}
	// end line definitions
	canvas.Println("</defs>")
}

func (canvas *Printer) PrintPaths(paths ...[]corner.Path) {
	for _, ps := range paths {
		canvas.Println("<g>")
		for _, p := range ps {
			canvas.Printf(usefmt, p.Id, "whitebg")
		}
		for _, p := range ps {
			canvas.Printf(usefmt, p.Id, p.Id)
		}
		canvas.Println("</g>")
	}
}

func (canvas *Printer) PrintClose() {
	canvas.Println("</svg>")
}

func main() {
	var (
		paths               []corner.Path
		paths2              []corner.Path
		width, height, rsep float64
		canvas              Printer
	)
	width = 2000
	height = 2000
	rsep = 10.0
	canvas = Printer{os.Stdout}
	points := []corner.Point{{100, 100}, {100, 400}, {400, 400}, {100, 800}}
	corners := corner.Sequence(points...)
	p1 := corner.NewPath("a", corners, []int{1, 0, 0})
	p2 := corner.NewPath("b", corners, []int{-1, 2, 2})
	p3 := corner.NewPath("c", corners, []int{0, 1, 1})
	paths = append(paths, *p1, *p2)
	paths2 = append(paths2, *p3)
	canvas.PrintInit(width, height)
	canvas.PrintDefs(rsep, paths, paths2)
	canvas.PrintPaths(paths, paths2)
	canvas.PrintClose()
}

const (
	header        = "<?xml version='1.0' encoding='utf-8' ?>"
	stylesheetfmt = "<?xml-stylesheet type='text/css' href='%s.css'?>\n"
	svginitfmt    = "<svg width=\"%v\" height=\"%v\"\n"
	svgns         = "xmlns='http://www.w3.org/2000/svg'\nxmlns:xlink='http://www.w3.org/1999/xlink'>"
	usefmt        = "<use xlink:href='#%s' class='%s' />\n"
	whitebg       = "class='whitebg'"
)
