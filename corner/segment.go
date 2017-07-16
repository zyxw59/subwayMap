package corner

type Segment struct {
	Start   Point
	End     Point
	offsets []int
}

func (s *Segment) Direction() Direction {
	return s.Start.DirectionTo(s.End)
}
