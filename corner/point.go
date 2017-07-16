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
