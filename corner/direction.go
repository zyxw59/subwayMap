package corner

import "math"

// Rose generates a set of n evenly-spaced Directions
func Rose(n int, offset float64) []Direction {
	dirs := make([]Direction, n)
	if offset == 0 {
		// special cases
		north := rectDirection{0, -1}
		south := rectDirection{0, 1}
		east := rectDirection{1, 0}
		west := rectDirection{-1, 0}
		if n == 1 {
			dirs[0] = north
			return dirs
		}
		if n == 2 {
			dirs[0] = north
			dirs[1] = south
			return dirs
		}
		if n == 4 {
			dirs[0] = north
			dirs[1] = east
			dirs[2] = south
			dirs[3] = west
			return dirs
		}
		if n == 8 {
			dirs[0] = north
			dirs[1] = rectDirection{1, -1}
			dirs[2] = east
			dirs[3] = rectDirection{1, 1}
			dirs[4] = south
			dirs[5] = rectDirection{-1, 1}
			dirs[6] = west
			dirs[7] = rectDirection{-1, -1}
			return dirs
		}
	}
	// general case
	alpha := 2 * math.Pi / float64(n)
	for i := range dirs {
		dirs[i] = thetaDirection{float64(i)*alpha + offset}
	}
	return dirs
}

// A Direction is a representation of a direction in 2-d space
type Direction interface {
	Angle() float64
	Minus(Direction) float64
	Basis(float64, float64, Point) Point
	Cos() float64
	Sin() float64
	Equal(Direction) bool
	Normal() Direction
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
	return Point{o.X + u*c - v*s, o.Y + u*s + v*c}
}

func (d thetaDirection) Equal(other Direction) bool {
	return d.theta == other.Angle()
}

func (d thetaDirection) Normal() Direction {
	return newThetaDirection(d.theta + math.Pi/2)
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
		return modTau(math.Atan2(-cross, dot))
	default:
		return modTau(d.Angle() - other.Angle())
	}
}

func (d rectDirection) r() float64 {
	return math.Sqrt(d.x*d.x + d.y*d.y)
}

func (d rectDirection) Cos() float64 {
	return d.x / d.r()
}

func (d rectDirection) Sin() float64 {
	return d.y / d.r()
}

func (d rectDirection) Angle() float64 {
	return math.Atan2(d.y, d.x)
}

func (d rectDirection) Basis(u, v float64, o Point) Point {
	return Point{o.X + (u*d.x-v*d.y)/d.r(), o.Y + (u*d.y+v*d.x)/d.r()}
}

func (d rectDirection) Equal(other Direction) bool {
	switch other := other.(type) {
	case rectDirection:
		return d.x*other.y == d.y*other.x
	default:
		return d.Angle() == other.Angle()
	}
}

func (d rectDirection) Normal() Direction {
	return rectDirection{-d.y, d.x}
}
