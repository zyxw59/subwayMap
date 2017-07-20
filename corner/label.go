package corner

import (
	"fmt"
	"strings"
)

const (
	// number of line widths between a line and a label
	labelFudge    = 1
	labelFmt      = "<text id='%s' x='%f' y='%f' dominant-baseline='%s' text-anchor='%s'>%s</text>"
	firstTspanFmt = "<tspan x='%f' dy='%fem'>%s</tspan>"
	tspanFmt      = "<tspan x='%f' dy='1em'>%s</tspan>"
)

type Anchor int8

const (
	Left   Anchor = -3
	Top    Anchor = -1
	Right  Anchor = 3
	Bottom Anchor = 1
)

type Label struct {
	Lines  []string
	Point  Point
	Anchor Anchor
	id     string
	class  string
}

func NewLabel(lines []string, point Point, anchor Anchor, id, class string) *Label {
	return &Label{
		Lines:  lines,
		Point:  point,
		Anchor: anchor,
		id:     id,
		class:  class,
	}
}

func (l *Label) Def() string {
	baseline := "middle"
	align := "middle"
	lines := make([]string, len(l.Lines))
	firstDy := -float64(len(lines)-1) / 2
	switch l.Anchor {
	case Top + Right, Right, Bottom + Right:
		align = "end"
	case Top + Left, Left, Bottom + Left:
		align = "start"
	}
	switch l.Anchor {
	case Bottom + Left, Bottom, Bottom + Right:
		baseline = "alphabetic"
		firstDy = -float64(len(lines) - 1)
	case Top + Left, Top, Top + Right:
		baseline = "hanging"
		firstDy = 0
	}
	for i, t := range l.Lines {
		if i == 0 {
			lines[i] = fmt.Sprintf(firstTspanFmt, l.Point.X, firstDy, t)
		} else {
			lines[i] = fmt.Sprintf(tspanFmt, l.Point.X, t)
		}
	}
	return fmt.Sprintf(labelFmt, l.id, l.Point.X, l.Point.Y, baseline, align, strings.Join(lines, ""))
}

func (p *Label) Use() string {
	return fmt.Sprintf(usefmt, p.id, p.class)
}

func (p *Label) Usebg() string {
	return fmt.Sprintf(usefmt, p.id, "bg "+p.class)
}
