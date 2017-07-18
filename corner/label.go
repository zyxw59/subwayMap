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

type Label struct {
	Text    string
	point   Point
	dir     Direction
	offset  int
	posSide bool
	id      string
	class   string
}

func (l *Label) Element(rbase, rsep float64) string {
	baseline := "middle"
	align := "middle"
	lines := strings.Split(l.Text, "\n")
	firstDy := -float64(len(lines)-1) / 2
	var offset float64
	if l.posSide {
		offset = float64(l.offset) + labelFudge
	} else {
		offset = float64(l.offset) - labelFudge
	}
	anchor := l.dir.Basis(offset*rsep, 0, l.point)
	switch {
	case anchor.X < l.point.X:
		align = "end"
	case anchor.X > l.point.X:
		align = "start"
	}
	switch {
	case anchor.Y < l.point.Y:
		baseline = "alphabetic"
		firstDy = -float64(len(lines) - 1)
	case anchor.Y > l.point.Y:
		baseline = "hanging"
		firstDy = 0
	}
	for i, l := range lines {
		if i == 0 {
			lines[i] = fmt.Sprintf(firstTspanFmt, anchor.X, firstDy, l)
		} else {
			lines[i] = fmt.Sprintf(tspanFmt, anchor.X, l)
		}
	}
	return fmt.Sprintf(labelFmt, l.id, anchor.X, anchor.Y, baseline, align, strings.Join(lines, ""))
}

func (l *Label) Id() string {
	return l.id
}

func (l *Label) Class() string {
	return l.class
}
