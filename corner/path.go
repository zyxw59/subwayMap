package corner

import (
	"bytes"
	"fmt"
)

// A Path is a series of Segments to be joined with rounded corners, along with
// parallel offsets along each segment.
type Path struct {
	id       string
	class    string
	segments []*Segment
	offsets  []float64
}

// NewPath returns a new Path with a given id, class, list of segments, and
// list of offsets. len(segments) should equal len(offsets). If one is too
// long, it will be truncated
func NewPath(id, class string, segments []*Segment, offsets []float64) *Path {
	ls := len(segments)
	lo := len(offsets)
	switch {
	case ls < lo:
		offsets = offsets[:ls]
	case ls > lo:
		segments = segments[:lo]
	}
	path := Path{
		id:       id,
		class:    class,
		segments: segments,
		offsets:  offsets,
	}
	// loop over segments and offsets, registering the offset at each
	// segment
	for i, o := range offsets {
		segments[i].AddOffset(o)
	}
	return &path
}

// Def generates the SVG <path> element to draw the Path. rbase determines
// the base radius of corners, and rsep determines the additional radius for
// each concentric Path
func (p *Path) Def() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("<path id='%s' d='", p.id))
	out.WriteString(p.segments[0].startPoint(p.offsets[0]))
	for i, s := range p.segments[1:] {
		out.WriteString(p.segments[i].ArcTo(s, p.offsets[i], p.offsets[i+1]))
	}
	l := len(p.segments) - 1
	out.WriteString(p.segments[l].endPoint(p.offsets[l]))
	out.WriteString("' />")
	return out.String()
}

func (p *Path) Use() string {
	return fmt.Sprintf(usefmt, p.id, p.class)
}

func (p *Path) Usebg() string {
	return fmt.Sprintf(usefmt, p.id, "bg "+p.class)
}

func (p *Path) arc(i int, rbase, rsep float64) string {
	var first, second string
	if i == 0 {
		first = p.segments[0].startPoint(p.offsets[0])
	} else {
		first = p.segments[i-1].ArcTo(p.segments[i], p.offsets[i-1], p.offsets[i])
	}
	if i == len(p.segments)-1 {
		second = p.segments[i].endPoint(p.offsets[i])
	} else {
		second = p.segments[i].ArcTo(p.segments[i+1], p.offsets[i], p.offsets[i+1])
	}
	return first + second
}
