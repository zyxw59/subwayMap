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
	out.WriteString(fmt.Sprintf("M %s\n", end))
	for i := range p.offsets[1:] {
		r, start, end, sweep = p.rounded(i+1, rsep)
		out.WriteString(fmt.Sprintf("L %s A %v,%v 0 %v 0 %s\n", start, r, r, sweep, end))
	}
	r, start, end, sweep = p.rounded(len(p.corners), rsep)
	out.WriteString(fmt.Sprintf("L %s", start))
	return out.String()
}

func (p *Path) rounded(i int, rsep float64) (float64, Point, Point, int) {
	if i == 0{
		point := p.corners[0].offset(0, p.offsets[0])
		return 0, point, point, 0
	}
	if i == len(p.corners){
		point := p.corners[i].offset(0, p.offsets[i-1])
		return 0, point, point, 0
	}
	return p.corners[i].Rounded(p.offsets[i-1], p.offsets[i], rsep)
}
