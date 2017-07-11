package corner

import "math"

// A Direction is a representation of a direction in 2-d space
type Direction interface {
	Angle() float64
	Minus(Direction) Direction
	Basis(float64, float64, Point) Point
}

type thetaDirection struct {
	theta float64
}

func (d thetaDirection) Minus(other Direction) Direction {
	return thetaDirection{math.Mod(d.theta-other.Angle(), 2*math.Pi)}
}

func (d thetaDirection) Angle() float64 {
	return d.theta
}

func (d thetaDirection) Basis(u, v float64, o Point) Point {
	c := math.Cos(d.theta)
	s := math.Sin(d.theta)
	return Point{o.x + u*c - v*s, o.y + u*s + v*c}
}
