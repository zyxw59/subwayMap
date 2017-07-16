package corner

import "fmt"

// number of line widths between a line and a label
const (
	labelFudge = 0.5
	labelfmt   = "<text id='%s' x='%f' y='%f' dominant-baseline='%s' text-anchor='%s'>%s</text>"
)

type Label struct {
	Text   string
	point  Point
	dir    Direction
	offset int
	id     string
	class  string
}

func NewLabel(text string, point Point, dir Direction, offset int, id, class string) *Label {
	return &Label{
		Text:  text,
		point: point,
		dir:   dir,
		id:    id,
		class: class,
	}
}

func (l *Label) Element(rbase, rsep float64) string {
	baseline := "middle"
	align := "middle"
	anchor := l.dir.Basis((float64(l.offset)+labelFudge)*rsep, 0, l.point)
	switch {
	case anchor.X < l.point.X:
		align = "end"
	case anchor.X > l.point.X:
		align = "start"
	}
	switch {
	case anchor.Y < l.point.Y:
		baseline = "alphabetic"
	case anchor.Y > l.point.Y:
		baseline = "hanging"
	}
	return fmt.Sprintf(labelfmt, l.id, anchor.X, anchor.Y, baseline, align, l.Text)
}

func (l *Label) Id() string {
	return l.id
}

func (l *Label) Class() string {
	return l.class
}
