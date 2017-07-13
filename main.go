package main

import (
	"github.com/zyxw59/subwayMap/canvas"
	"github.com/zyxw59/subwayMap/corner"
	"os"
)

func main() {
	var (
		paths               []corner.Path
		paths2              []corner.Path
		width, height, rsep float64
		c                   canvas.Canvas
	)
	width = 2000
	height = 2000
	rsep = 10.0
	c = canvas.Canvas{os.Stdout}
	points := []corner.Point{{100, 100}, {100, 400}, {400, 400}, {100, 800}}
	corners := corner.Sequence(points...)
	p1 := corner.NewPath("a", corners, []int{1, 0, 0})
	p2 := corner.NewPath("b", corners, []int{-1, 2, 2})
	p3 := corner.NewPath("c", corners, []int{0, 1, 1})
	paths = append(paths, *p1, *p2)
	paths2 = append(paths2, *p3)
	c.PrintAll(width, height, "lines", rsep, paths, paths2)
}
