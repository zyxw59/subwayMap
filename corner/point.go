package corner

import "fmt"

// A Point is a point in 2-d space, with an x coordinate and a y coordinate
type Point struct {
	X float64
	Y float64
}

// DirectionTo returns the direction to another point
func (p Point) DirectionTo(o Point) Direction {
	return rectDirection{o.X - p.X, o.Y - p.Y}
}

func (p Point) String() string {
	return fmt.Sprintf("%f %f", p.X, p.Y)
}

// Sequence produces a sequence of Corners from a sequence of Points
func Sequence(points ...Point) []*Corner {
	cs := make([]*Corner, len(points))
	dir := points[0].DirectionTo(points[1])
	for i, p := range points[:len(points)-1] {
		cs[i] = &Corner{Point: p, in: dir}
		dir = p.DirectionTo(points[i+1])
		cs[i].out = dir
	}
	cs[len(points)-1] = &Corner{Point: points[len(points)-1], in: dir, out: dir}
	return cs
}
