package corner

import (
	"bytes"
	"fmt"
)

// A Path is a series of Corners to be joined with line segments, along with
// parallel offsets along each segment. A Path registers the smaller of the two
// offsets (before and after) with each Corner
type Path struct {
	id      string
	class   string
	corners []*Corner
	offsets []int
}

// NewPath returns a new Path with a given id, list of corners, and list of
// offsets. len(corners) should equal len(offsets) + 1. If one is too long, it
// will be truncated
func NewPath(id, class string, corners []*Corner, offsets []int) *Path {
	lc := len(corners)
	lo := len(offsets)
	switch {
	case lc < lo+1:
		offsets = offsets[:lc-1]
	case lc > lo+1:
		corners = corners[:lo+1]
	}
	path := Path{
		id:      id,
		class:   class,
		corners: corners,
		offsets: offsets,
	}
	// loop over corners and offsets, registering the offset at each corner
	// we skip the first Corner, since the Path doesn't turn here
	for i, o := range offsets[1:] {
		// register the offsets. corners[i+1] is bounded by segments
		// with offsets offsets[i] and offsets[i+1]
		corners[i+1].AddOffsets(offsets[i], o)
	}
	return &path
}

// Element generates the SVG <path> element to draw the Path. rbase determines
// the base radius of corners, and rsep determines the additional radius for
// each concentric Path
func (p *Path) Element(rbase, rsep float64) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("<path id='%s' d='", p.id))
	for i := range p.corners {
		out.WriteString(p.arc(i, rbase, rsep))
	}
	out.WriteString("' />")
	return out.String()
}

func (p *Path) Id() string {
	return p.id
}

func (p *Path) Class() string {
	return p.class
}

func (p *Path) arc(i int, rbase, rsep float64) string {
	if i == 0 {
		return p.corners[i].startPoint(p.offsets[0], rsep)
	}
	if i == len(p.offsets) {
		return p.corners[i].endPoint(p.offsets[i-1], rsep)
	}
	return p.corners[i].Arc(p.offsets[i-1], p.offsets[i], rbase, rsep)
}
