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
		// register the offsets. corners[i+1] is bounded by segments with
		// offsets offsets[i] and offsets[i+1]
		corners[i+1].AddOffsets(offsets[i], o)
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
	if i == 0{
		return 0, p.corners[0].Point, p.corners[0].Point, 0
	}
	if i == len(p.corners){
		return 0, p.corners[i].Point, p.corners[i].Point, 0
	}
	return p.corners[i].Rounded(p.offsets[i-1], p.offsets[i], rsep)
}
