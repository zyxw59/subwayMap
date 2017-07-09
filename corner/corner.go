package corner

import "math"

// A Corner is a point at the intersection of two line segments. It maintains a
// list of offsets of paths using it, as well as an incoming and an outgoing
// Direction.
type Corner struct {
    point Point
    in Direction
    out Direction
    offsets []int
}

// AddOffsets appends offsets to a Corner
func (c *Corner) AddOffsets(offsets ...int) {
    c.offsets = append(c.offsets, offsets...)
}

// Rounded generates a rounded corner with a given parallel offset, and a given
// radius increment value. It returns the radius, the start and end points of
// the arc, and the sweep flag.
func (c *Corner) Rounded(offset int, rsep float64) (float64, Point, Point, bool) {
    var start, end Point
    var r float64
    theta := c.out.minus(c.in).angle()
    if theta > math.Pi {
        // the line is turning ccw, count lowest offset
        r = rsep*float64(offset - minIntSlice(c.offsets) + 1)
    } else {
        r = rsep*float64(maxIntSlice(c.offsets) - offset + 1)
    }
    l := r / math.Tan(theta)
    sweep := l < 0
    start = c.in.vector(-math.Abs(l), c.point)
    end = c.out.vector(math.Abs(l), c.point)
    return r, start, end, sweep
}

// A Point is a point in 2-d space, with an x coordinate and a y coordinate
type Point struct {
    x float64
    y float64
}

// DirectionTo returns the direction to another point
func (p Point) DirectionTo (o Point) Direction {
    theta = math.Atan2(p.y - o.y, p.x - o.x)
    return thetaDirection{theta}
}

// A Direction is a representation of a direction in 2-d space
type Direction interface {
    angle() float64
    minus(Direction) Direction
    vector(float64, Point) Point
}

type thetaDirection struct {
    theta float64
}

func (d thetaDirection) minus(other Direction) Direction {
    return thetaDirection{math.Mod(d.theta - other.angle(), 2*math.Pi)}
}

func (d thetaDirection) angle() float64 {
    return d.theta
}

func (d thetaDirection) vector(l float64, o Point) Point {
    dx := l * math.Cos(d.theta)
    dy := l * math.Sin(d.theta)
    return Point{o.x + dx, o.y + dy}
}


func minIntSlice(slice []int) int {
    m := slice[0]
    for _, x := range slice {
        if x < m {
            m = x
        }
    }
    return m
}

func maxIntSlice(slice []int) int {
    m := slice[0]
    for _, x := range slice {
        if x > m {
            m = x
        }
    }
    return m
}
