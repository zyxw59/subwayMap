package corner

import "math"

// North is up the page, i.e. -y
var North = rectDirection{0, -1}

// South is down the page, i.e. +y
var South = rectDirection{0, 1}

// East is to the right, i.e. +x
var East = rectDirection{-1, 0}

// West is to the left, i.e. -x
var West = rectDirection{1, 0}

/*
var North = thetaDirection{-math.Pi/2}
var South = thetaDirection{math.Pi/2}
var East = thetaDirection{0}
var West = thetaDirection{math.Pi}
*/

// A Direction is a representation of a direction in 2-d space
type Direction interface {
	Angle() float64
	Minus(Direction) float64
	Basis(float64, float64, Point) Point
	Cos() float64
	Sin() float64
}

type thetaDirection struct {
	theta float64
}

func modTau(theta float64) float64 {
	thetaMod := math.Mod(theta, 2*math.Pi)
	if thetaMod < 0 {
		return thetaMod + 2*math.Pi
	}
	return thetaMod
}

func newThetaDirection(theta float64) thetaDirection {
	return thetaDirection{modTau(theta)}
}

func (d thetaDirection) Minus(other Direction) float64 {
	return modTau(d.theta - other.Angle())
}

func (d thetaDirection) Angle() float64 {
	return d.theta
}

func (d thetaDirection) Cos() float64 {
	return math.Cos(d.theta)
}

func (d thetaDirection) Sin() float64 {
	return math.Sin(d.theta)
}

func (d thetaDirection) Basis(u, v float64, o Point) Point {
	c := d.Cos()
	s := d.Sin()
	return Point{o.x + u*c - v*s, o.y + u*s + v*c}
}

type rectDirection struct {
	x float64
	y float64
}

func (d rectDirection) Minus(other Direction) float64 {
	switch other := other.(type) {
	case rectDirection:
		dot := d.x*other.x + d.y*other.y
		cross := d.x*other.y - d.y*other.x
		return modTau(math.Atan2(cross, dot))
	default:
		return modTau(d.Angle() - other.Angle())
	}
}

func (d rectDirection) r() float64 {
	return math.Sqrt(d.x*d.x + d.y*d.y)
}

func (d rectDirection) Cos() float64 {
	return -d.x / d.r()
}

func (d rectDirection) Sin() float64 {
	return d.y / d.r()
}

func (d rectDirection) Angle() float64 {
	return math.Atan2(d.y, d.x)
}

func (d rectDirection) Basis(u, v float64, o Point) Point {
	return Point{o.x - (u*d.x+v*d.y)/d.r(), o.y + (u*d.y-v*d.x)/d.r()}
}
