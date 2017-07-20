package corner

import (
	"fmt"
	"io"
)

const (
	header        = "<?xml version='1.0' encoding='utf-8' ?>"
	stylesheetfmt = "<?xml-stylesheet type='text/css' href='%s.css'?>\n"
	svginitfmt    = "<svg width=\"%v\" height=\"%v\"\n"
	svgns         = "xmlns='http://www.w3.org/2000/svg'\nxmlns:xlink='http://www.w3.org/1999/xlink'>"
	usefmt        = "<use xlink:href='#%s' class='%s' />\n"
	whitebg       = "class='whitebg'"
)

type Element interface {
	Def() string
	Use() string
	Usebg() string
}

type Canvas struct {
	writer     io.Writer
	Width      float64
	Height     float64
	Stylesheet string
	rbase      float64
	rsep       float64
	elements   [][]Element
}

// NewCanvas initializes a new Canvas
func NewCanvas(writer io.Writer, width, height float64, stylesheet string, rbase, rsep float64) *Canvas {
	c := &Canvas{
		writer:     writer,
		Width:      width,
		Height:     height,
		Stylesheet: stylesheet,
		rbase:      rbase,
		rsep:       rsep,
	}
	c.printInit()
	return c
}

// AddElements adds the specified Elements as their own layer
func (c *Canvas) AddElements(elements ...Element) {
	c.elements = append(c.elements, elements)
	c.Println("<defs>")
	for _, e := range elements {
		c.Println(e.Def())
	}
	c.Println("</defs>")
}

// Close finishes writing the SVG code
func (c *Canvas) Close() {
	c.printElements()
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

// printElements prints <use> tags to draw the elements defined earlier. Within
// each layer, first the elements are drawn with class="e.Class() bg" to stroke outlines,
// then with class="e.Class()" to draw the elements in foreground
func (c *Canvas) printElements() {
	for _, es := range c.elements {
		c.Println("<g>")
		for _, e := range es {
			c.Println(e.Usebg())
		}
		for _, e := range es {
			c.Println(e.Use())
		}
		c.Println("</g>")
	}
}

// printClose prints the closing </svg> tag
func (c *Canvas) printClose() {
	c.Println("</svg>")
}
