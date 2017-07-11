package corner

import "math"

// A Corner is a point at the intersection of two line segments. It maintains a
// list of offsets of paths using it, as well as an incoming and an outgoing
// Direction.
type Corner struct {
	Point   Point
	in      Direction
	out     Direction
	offsets []int
}

// AddOffsets appends offsets to a Corner
func (c *Corner) AddOffsets(offsets ...int) {
	c.offsets = append(c.offsets, offsets...)
}

// Rounded generates a rounded corner with a given parallel offset, and a given
// radius increment value. It returns the radius, the start and end points of
// the arc, and the sweep flag.
func (c *Corner) Rounded(offset int, rsep float64) (float64, Point, Point, int) {
	var start, end Point
	var r, o float64
	var sweep int
	o = float64(offset) * rsep
	if c.in == c.out {
		start = c.in.Basis(0, o, c.Point)
		return 0, start, start, 0
	}
	theta := c.out.Minus(c.in).Angle()
	if theta > math.Pi {
		// the line is turning ccw, count lowest offset
		r = rsep * float64(offset-minIntSlice(c.offsets)+1)
	} else {
		r = rsep * float64(maxIntSlice(c.offsets)-offset+1)
	}
	l := r / math.Tan(theta/2)
	if l < 0 {
		sweep = 1
	}
	start = c.in.Basis(-math.Abs(l), o, c.Point)
	end = c.out.Basis(math.Abs(l), o, c.Point)
	return r, start, end, sweep
}

// A Point is a point in 2-d space, with an x coordinate and a y coordinate
type Point struct {
	x float64
	y float64
}

// DirectionTo returns the direction to another point
func (p Point) DirectionTo(o Point) Direction {
	theta := math.Atan2(p.y-o.y, p.x-o.x)
	return thetaDirection{theta}
}
