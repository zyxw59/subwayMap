package corner

import (
	"fmt"
	"math"
)

type Segment struct {
	Start   Point
	End     Point
	offsets []int
}

func NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		Start: Point{x0, y0},
		End:   Point{x1, y1},
	}
}

func (s *Segment) Direction() Direction {
	return s.Start.DirectionTo(s.End)
}

func (s *Segment) AddOffset(offset int) {
	s.offsets = append(s.offsets, offset)
}

func (s *Segment) String() string {
	return fmt.Sprintf("Segment((%s), (%s))", s.Start, s.End)
}

// Sequence produces a sequence of Segments from a sequence of Points
func Sequence(points ...Point) []*Segment {
	segs := make([]*Segment, len(points)-1)
	for i, p := range points[:len(points)-1] {
		segs[i] = &Segment{
			Start: p,
			End:   points[i+1],
		}
	}
	return segs
}

func (s *Segment) endPoint(offset int, rsep float64) string {
	p := s.Direction().Basis(0, float64(offset)*rsep, s.End)
	return fmt.Sprintf("L %s", p)
}

func (s *Segment) startPoint(offset int, rsep float64) string {
	p := s.Direction().Basis(0, float64(offset)*rsep, s.Start)
	return fmt.Sprintf("M %s\n", p)
}

func (s *Segment) ArcTo(other *Segment, in, out int, rbase, rsep float64) string {
	inr := float64(in) * rsep
	outr := float64(out) * rsep
	if s.End != other.Start {
		return s.endPoint(in, rsep) + other.startPoint(out, rsep)
	}
	inDir := s.Direction()
	outDir := other.Direction()
	if inDir.Equal(outDir) {
		if in == out {
			// nothing to do here
			return ""
		}
		// otherwise, parallel shifts
		delta := math.Abs(float64(out-in) * rsep)
		p0 := inDir.Basis(-delta, inr, s.End)
		p1 := inDir.Basis(0, inr, s.End)
		p2 := outDir.Basis(0, outr, s.End)
		p3 := outDir.Basis(delta, outr, s.End)
		return fmt.Sprintf("L %s C %s %s %s\n", p0, p1, p2, p3)
	}
	// rounded corner
	var inDelta, outDelta, sweep int
	theta := inDir.Minus(outDir) / 2
	if theta > math.Pi/2 {
		sweep = 1
		inDelta = maxIntSlice(s.offsets) - in
		outDelta = maxIntSlice(other.offsets) - out
	} else {
		sweep = 0
		inDelta = in - minIntSlice(s.offsets)
		outDelta = out - minIntSlice(other.offsets)
	}
	r := rsep*math.Min(float64(inDelta), float64(outDelta)) + rbase
	l := math.Abs(r * math.Tan(theta))
	alpha := 1 / math.Sin(theta*2)
	p := outDir.Basis(-alpha*inr, 0, inDir.Basis(alpha*outr, 0, s.End))
	start := inDir.Basis(-l, 0, p)
	end := outDir.Basis(l, 0, p)
	return fmt.Sprintf("L %s A %v,%v 0 0 %v %s\n", start, r, r, sweep, end)
}

func (s *Segment) LabelAt(point Point, posSide bool, text string) *Label {
	var offset int
	if posSide {
		offset = maxIntSlice(s.offsets)
	} else {
		offset = minIntSlice(s.offsets)
	}
	return &Label{
		Text:    text,
		point:   point,
		dir:     s.Direction().Normal(),
		offset:  offset,
		posSide: posSide,
		id:      fmt.Sprintf("%v-%v-%v", text, point.X, point.Y),
		class:   "",
	}
}

func (s *Segment) LabelAtX(x float64, posSide bool, text string) *Label {
	// find y such that (x, y) is on the line between s.Start and s.End
	y := ((s.Start.X-x)*s.End.Y - (s.End.X-x)*s.Start.Y) / (s.Start.X - s.End.X)
	return s.LabelAt(Point{x, y}, posSide, text)
}

func (s *Segment) LabelAtY(y float64, posSide bool, text string) *Label {
	// find x such that (x, y) is on the line between s.Start and s.End
	x := ((s.Start.Y-y)*s.End.X - (s.End.Y-y)*s.Start.X) / (s.Start.Y - s.End.Y)
	return s.LabelAt(Point{x, y}, posSide, text)
}
