package corner

import (
	"fmt"
	"math"
	"strings"
)

type Segment struct {
	Start   Point
	End     Point
	offsets []float64
	rbase   float64
	rsep    float64
}

func (c *Canvas) NewSegment(x0, y0, x1, y1 float64) *Segment {
	return &Segment{
		Start: Point{x0, y0},
		End:   Point{x1, y1},
		rbase: c.rbase,
		rsep:  c.rsep,
	}
}

func (s *Segment) Direction() Direction {
	return s.Start.DirectionTo(s.End)
}

func (s *Segment) AddOffset(offset float64) {
	s.offsets = append(s.offsets, offset)
}

func (s *Segment) String() string {
	return fmt.Sprintf("Segment((%s), (%s))", s.Start, s.End)
}

// Sequence produces a sequence of Segments from a sequence of Points
func (c *Canvas) Sequence(points ...Point) []*Segment {
	segs := make([]*Segment, len(points)-1)
	for i, p := range points[:len(points)-1] {
		segs[i] = c.NewSegment(p.X, p.Y, points[i+1].X, points[i+1].Y)
	}
	return segs
}

func (s *Segment) endPoint(offset float64) string {
	p := s.Direction().Basis(0, offset*s.rsep, s.End)
	return fmt.Sprintf("L %s", p)
}

func (s *Segment) startPoint(offset float64) string {
	p := s.Direction().Basis(0, offset*s.rsep, s.Start)
	return fmt.Sprintf("M %s\n", p)
}

func (s *Segment) ArcTo(other *Segment, in, out float64) string {
	inr := in * s.rsep
	outr := out * s.rsep
	if s.End != other.Start {
		return s.endPoint(in) + other.startPoint(out)
	}
	inDir := s.Direction()
	outDir := other.Direction()
	if inDir.Equal(outDir) {
		if in == out {
			// nothing to do here
			return ""
		}
		// otherwise, parallel shifts
		delta := math.Abs(out-in) * s.rsep
		p0 := inDir.Basis(-delta, inr, s.End)
		p1 := inDir.Basis(0, inr, s.End)
		p2 := outDir.Basis(0, outr, s.End)
		p3 := outDir.Basis(delta, outr, s.End)
		return fmt.Sprintf("L %s C %s %s %s\n", p0, p1, p2, p3)
	}
	// rounded corner
	var (
		inDelta, outDelta float64
		sweep             int
	)
	theta := inDir.Minus(outDir) / 2
	if theta > math.Pi/2 {
		sweep = 1
		inDelta = maxFloatSlice(s.offsets) - in
		outDelta = maxFloatSlice(other.offsets) - out
	} else {
		sweep = 0
		inDelta = in - minFloatSlice(s.offsets)
		outDelta = out - minFloatSlice(other.offsets)
	}
	r := s.rsep*math.Min(inDelta, outDelta) + s.rbase
	l := math.Abs(r * math.Tan(theta))
	alpha := 1 / math.Sin(theta*2)
	p := outDir.Basis(-alpha*inr, 0, inDir.Basis(alpha*outr, 0, s.End))
	start := inDir.Basis(-l, 0, p)
	end := outDir.Basis(l, 0, p)
	return fmt.Sprintf("L %s A %v,%v 0 0 %v %s\n", start, r, r, sweep, end)
}

func (s *Segment) LabelAt(point Point, posSide bool, text string) *Label {
	var offset float64
	if posSide {
		offset = maxFloatSlice(s.offsets) + labelFudge
	} else {
		offset = minFloatSlice(s.offsets) - labelFudge
	}
	lines := strings.Split(text, "\n")
	anchorPoint := s.Direction().Basis(0, offset*s.rsep, point)
	var anchor Anchor
	switch {
	case anchorPoint.X < point.X:
		anchor += Right
	case anchorPoint.X > point.X:
		anchor += Left
	}
	switch {
	case anchorPoint.Y < point.Y:
		anchor += Bottom
	case anchorPoint.Y > point.Y:
		anchor += Top
	}
	return &Label{
		Lines:  lines,
		Point:  anchorPoint,
		Anchor: anchor,
		id:     fmt.Sprintf("%v-%v-%v", point.X, point.Y, strings.Fields(text)[0]),
		class:  "label",
	}
}

// LabelAtX places a label along a segment at a specified x-value
func (s *Segment) LabelAtX(x float64, posSide bool, text string) *Label {
	// find y such that (x, y) is on the line between s.Start and s.End
	y := ((s.Start.X-x)*s.End.Y - (s.End.X-x)*s.Start.Y) / (s.Start.X - s.End.X)
	return s.LabelAt(Point{x, y}, posSide, text)
}

// LabelAtY places a label along a segment at a specified y-value
func (s *Segment) LabelAtY(y float64, posSide bool, text string) *Label {
	// find x such that (x, y) is on the line between s.Start and s.End
	x := ((s.Start.Y-y)*s.End.X - (s.End.Y-y)*s.Start.X) / (s.Start.Y - s.End.Y)
	return s.LabelAt(Point{x, y}, posSide, text)
}

func SegmentConcat(slices ...[]*Segment) []*Segment {
	sum := 0
	for _, sl := range slices {
		sum += len(sl)
	}
	segs := make([]*Segment, 0, sum)
	for _, sl := range slices {
		segs = append(segs, sl...)
	}
	return segs
}
