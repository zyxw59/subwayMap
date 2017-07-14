package main

import (
	"github.com/zyxw59/subwayMap/canvas"
	"github.com/zyxw59/subwayMap/corner"
	"os"
)

func main() {
	var (
		paths               []corner.Path
		width, height, rsep float64
		c                   canvas.Canvas
	)
	width = 2000
	height = 2000
	rsep = 10.0
	c = canvas.Canvas{os.Stdout}
	points := []corner.Point{{100, 400}, {200, 400}, {300, 400}}
	ab := []corner.Point{{100, 300}}
	cd := []corner.Point{{100, 500}}
	ac := []corner.Point{{300, 300}}
	bd := []corner.Point{{300, 500}}
	abCorners := corner.Sequence(append(ab, points...)...)
	cdCorners := corner.Sequence(append(cd, points...)...)
	acCorners := corner.Sequence(append(points, ac...)...)
	bdCorners := corner.Sequence(append(points, bd...)...)
	aCorners := append(abCorners[:2], acCorners[1:]...)
	bCorners := append(abCorners[:2], bdCorners[1:]...)
	cCorners := append(cdCorners[:2], acCorners[1:]...)
	dCorners := append(cdCorners[:2], bdCorners[1:]...)
	p1 := *corner.NewPath("a", aCorners, []int{-1, -1,  0,  0})
	p2 := *corner.NewPath("b", bCorners, []int{ 0,  0,  2,  1})
	p3 := *corner.NewPath("c", cCorners, []int{ 0,  1, -1, -1})
	p4 := *corner.NewPath("d", dCorners, []int{ 1,  2,  1,  0})
	paths1 := append(paths, p1)
	paths2 := append(paths, p2)
	paths3 := append(paths, p3)
	paths4 := append(paths, p4)
	c.PrintAll(width, height, "lines", rsep, paths1, paths2, paths3, paths4)
}
