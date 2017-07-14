package corner

import (
	"fmt"
	"log"
	"math"
)

var _ = log.Print

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

func (c *Corner) Arc(in, out int, rsep float64) (text string) {
	inr := float64(in) * rsep
	outr := float64(out) * rsep
	if c.in.Equal(c.out) {
		if in == out {
			// nothing to do here
			return ""
		}
		// otherwise, parallel shifts
		delta := math.Abs(float64(out-in) * rsep)
		p0 := c.in.Basis(-delta, inr, c.Point)
		p1 := c.in.Basis(0, inr, c.Point)
		p2 := c.out.Basis(0, outr, c.Point)
		p3 := c.out.Basis(delta, outr, c.Point)
		return fmt.Sprintf("L %s C %s %s %s\n", p0, p1, p2, p3)
	}
	// rounded corner
	r, start, end, sweep := c.Rounded(in, out, rsep)
	return fmt.Sprintf("L %s A %v,%v 0 0 %v %s\n", start, r, r, sweep, end)
}

func (c *Corner) endPoint(offset int, rsep float64) string {
	p := c.in.Basis(0, float64(offset)*rsep, c.Point)
	return fmt.Sprintf("L %s", p)
}

func (c *Corner) startPoint(offset int, rsep float64) string {
	p := c.out.Basis(0, float64(offset)*rsep, c.Point)
	return fmt.Sprintf("M %s\n", p)
}

// Rounded generates a rounded corner with given parallel offsets before and
// after the corner, and a given radius increment value. It returns the radius,
// the start and end points of the arc, and the sweep flag.
func (c *Corner) Rounded(in, out int, rsep float64) (r float64, start, end Point, sweep int) {
	theta := c.in.Minus(c.out) / 2
	var inD, outD int
	if theta > math.Pi/2 {
		sweep = 1
		inD = maxIntSlice(c.inOffsets) - in + 1
		outD = maxIntSlice(c.outOffsets) - out + 1
	} else {
		sweep = 0
		inD = in - minIntSlice(c.inOffsets) + 1
		outD = out - minIntSlice(c.outOffsets) + 1
	}
	r = rsep * math.Min(float64(inD), float64(outD))
	l := math.Abs(r * math.Tan(theta))
	p := c.offset(float64(in)*rsep, float64(out)*rsep)
	start = c.in.Basis(-l, 0, p)
	end = c.out.Basis(l, 0, p)
	return r, start, end, sweep
}

func (c *Corner) offset(in, out float64) Point {
	if c.in.Equal(c.out) {
		return c.in.Basis(0, (in+out)/2, c.Point)
	}
	alpha := 1 / math.Sin(c.in.Minus(c.out))
	return c.out.Basis(-alpha*in, 0, c.in.Basis(alpha*out, 0, c.Point))
}

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
