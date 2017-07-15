package canvas

import (
	"fmt"
	"github.com/zyxw59/subwayMap/corner"
	"io"
)

const (
	header        = "<?xml version='1.0' encoding='utf-8' ?>"
	stylesheetfmt = "<?xml-stylesheet type='text/css' href='%s.css'?>\n"
	svginitfmt    = "<svg width=\"%v\" height=\"%v\"\n"
	svgns         = "xmlns='http://www.w3.org/2000/svg'\nxmlns:xlink='http://www.w3.org/1999/xlink'>"
	usefmt        = "<use xlink:href='#%s' class='%s' />\n"
	whitebg       = "class='whitebg'"
	labelFudge    = 0.5
)

type Canvas struct {
	writer     io.Writer
	Width      float64
	Height     float64
	Stylesheet string
	Rbase      float64
	Rsep       float64
	paths      [][]corner.Path
}

// New initializes a new Canvas
func New(writer io.Writer, width, height float64, stylesheet string, rbase, rsep float64) *Canvas {
	c := &Canvas{
		writer:     writer,
		Width:      width,
		Height:     height,
		Stylesheet: stylesheet,
		Rbase:      rbase,
		Rsep:       rsep,
	}
	c.printInit()
	return c
}

// AddPaths adds the specified Paths as their own layer
func (c *Canvas) AddPaths(paths ...corner.Path) {
	c.paths = append(c.paths, paths)
	c.Println("<defs>")
	for _, p := range paths {
		c.Println(p.Path(c.Rbase, c.Rsep))
	}
	c.Println("</defs>")
}

// Close finishes writing the SVG code
func (c *Canvas) Close() {
	c.printPaths()
	c.printClose()
}

func (c *Canvas) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.writer, a...)
}

func (c *Canvas) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.writer, format, a...)
}

func (c *Canvas) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.writer, a...)
}

// printInit prints the XML declaration, stylesheet, and opening <svg> tag
func (c *Canvas) printInit() {
	c.Println(header)
	c.Printf(stylesheetfmt, c.Stylesheet)
	c.Printf(svginitfmt, c.Width, c.Height)
	c.Println(svgns)
}

// printPaths prints <use> tags to draw the paths. Within each layer, first the
// paths are drawn with class="whitebg" to stroke behind the lines, then with
// class="p.Class" to draw the lines themselves
func (c *Canvas) printPaths() {
	for _, ps := range c.paths {
		c.Println("<g>")
		for _, p := range ps {
			c.Printf(usefmt, p.Id, "whitebg")
		}
		for _, p := range ps {
			c.Printf(usefmt, p.Id, p.Class)
		}
		c.Println("</g>")
	}
}

// printClose prints the closing </svg> tag
func (c *Canvas) printClose() {
	c.Println("</svg>")
}
