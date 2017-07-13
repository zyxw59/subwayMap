package corner

import (
	"fmt"
	"math"
)

// A Corner is a point at the intersection of two line segments. It maintains a
// list of offsets of paths using it, as well as an incoming and an outgoing
// Direction.
type Corner struct {
	Point      Point
	in         Direction
	out        Direction
	inOffsets  []int
	outOffsets []int
}

func NewCorner(x, y float64, in, out Direction) (c *Corner) {
	return &Corner{Point: Point{x, y}, in: in, out: out}
}

// AddOffsets appends incoming and outgoing offsets to a Corner
func (c *Corner) AddOffsets(in, out int) {
	c.inOffsets = append(c.inOffsets, in)
	c.outOffsets = append(c.outOffsets, out)
}

// Rounded generates a rounded corner with given parallel offsets before and
// after the corner, and a given radius increment value. It returns the radius,
// the start and end points of the arc, and the sweep flag.
func (c *Corner) Rounded(in, out int, rsep float64) (r float64, start, end Point, sweep int) {
	theta := c.in.Minus(c.out).Angle() / 2
	var inD, outD int
	// I might have this backwards
	if theta > math.Pi/2 {
		sweep = 1
		inD = in - minIntSlice(c.inOffsets) + 1
		outD = in - minIntSlice(c.outOffsets) + 1
	} else {
		sweep = 0
		inD = maxIntSlice(c.inOffsets) - in + 1
		outD = maxIntSlice(c.outOffsets) - out + 1
	}
	r = rsep * math.Min(float64(inD), float64(outD))
	l := math.Abs(r * math.Tan(theta))
	p := c.offset(in, out)
	start = c.in.Basis(-l, 0, p)
	end = c.out.Basis(l, 0, p)
	return r, start, end, sweep
}

func (c *Corner) offset(in, out int) Point {
	alpha := 1 / math.Cos(c.in.Minus(c.out).Angle())
	inVec := alpha * float64(in)
	outVec := alpha * float64(out)
	return c.out.Basis(inVec, 0, c.in.Basis(outVec, 0, c.Point))
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

func (p Point) String() string {
	return fmt.Sprintf("%v %v", p.x, p.y)
}
