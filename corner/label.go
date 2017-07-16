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

func (l *Label) Element(rbase, rsep float) string {
	baseline := "middle"
	align := "middle"
	anchor := l.dir.Basis((float64(l.offset)+labelFudge)*rsep, 0, l.point)
	switch {
	case anchor.x < l.point.x:
		align = "end"
	case anchor.x > l.point.x:
		align = "start"
	}
	switch {
	case anchor.y < l.point.y:
		baseline = "alphabetic"
	case anchor.y > l.point.y:
		baseline = "hanging"
	}
	return fmt.Sprintf(labelfmt, l.id, anchor.x, anchor.y, baseline, align, l.Text)
}

func (l *Label) Id() string {
	return l.id
}

func (l *Label) Class() string {
	return l.class
}
