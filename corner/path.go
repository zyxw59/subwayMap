package corner

import (
	"bytes"
	"fmt"
)

// A Path is a series of Corners to be joined with line segments, along with
// parallel offsets along each segment. A Path registers the smaller of the two
// offsets (before and after) with each Corner
type Path struct {
	Id      string
	corners []Corner
	offsets []int
}

func NewPath(id string, corners []Corner, offsets []int) *Path {
	path := Path{Id: id, corners: corners, offsets: offsets}
	// loop over corners and offsets, registering the offset at each corner
	// we skip the first Corner, since the Path doesn't turn here
	for i, o := range offsets[1:] {
		// register the smaller offset. corners[i] is bounded by segments with
		// offsets offsets[i-1] and offsets[i]
		if offsets[i-1] > o {
			corners[i].AddOffsets(o)
		} else {
			corners[i].AddOffsets(offsets[i-1])
		}
	}
	return &path
}

func (p *Path) Path(rsep float64) string {
	var out bytes.Buffer
	r, start, end, sweep := p.rounded(0, rsep)
	out.WriteString(fmt.Sprintf("A %v,%v 0 %v 0 %s %s", r, r, sweep, start, end))
	return out.String()
}

func (p *Path) rounded(i int, rsep float64) (float64, Point, Point, int) {
	var o int
	switch {
	case i == 0:
		fallthrough
	case p.offsets[i-1] > p.offsets[i]:
		o = p.offsets[i]
	default:
		o = p.offsets[i-1]
	}
	return p.corners[i].Rounded(o, rsep)
}
