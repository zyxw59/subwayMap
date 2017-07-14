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
)

type Canvas struct {
	Writer io.Writer
}

func (c *Canvas) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.Writer, a...)
}

func (c *Canvas) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.Writer, format, a...)
}

func (c *Canvas) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.Writer, a...)
}

func (c *Canvas) PrintInit(width, height float64, stylesheet string) {
	c.Println(header)
	c.Printf(stylesheetfmt, stylesheet)
	c.Printf(svginitfmt, width, height)
	c.Println(svgns)
}

func (c *Canvas) PrintDefs(rbase, rsep float64, paths ...[]corner.Path) {
	// begin line definitions
	c.Println("<defs>")
	for _, ps := range paths {
		for _, p := range ps {
			c.Println(p.Path(rbase, rsep))
		}
	}
	// end line definitions
	c.Println("</defs>")
}

func (c *Canvas) PrintPaths(paths ...[]corner.Path) {
	for _, ps := range paths {
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

func (c *Canvas) PrintClose() {
	c.Println("</svg>")
}

func (c *Canvas) PrintAll(width, height float64, stylesheet string, rbase, rsep float64, paths ...[]corner.Path) {
	c.PrintInit(width, height, stylesheet)
	c.PrintDefs(rbase, rsep, paths...)
	c.PrintPaths(paths...)
	c.PrintClose()
}
