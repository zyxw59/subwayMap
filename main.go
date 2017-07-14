package main

import (
	"github.com/zyxw59/subwayMap/canvas"
	"github.com/zyxw59/subwayMap/corner"
	"os"
)

func main() {
	var (
		paths                      []corner.Path
		width, height, rbase, rsep float64
		c                          canvas.Canvas
	)
	width = 2000
	height = 2000
	rbase = 30
	rsep = 10
	c = canvas.Canvas{os.Stdout}
	av8st145 := corner.Point{200, 100}
	av8st63 := corner.Point{200, 500}
	av6st63 := corner.Point{400, 500}
	av8st14 := corner.Point{200, 900}
	av6st4 := corner.Point{400, 1000}
	stChurchStChambers := corner.Point{400, 1200}
	stHoustonAv2 := corner.Point{600, 1000}
	av8 := corner.Sequence(av8st145, av8st14, av6st4, stChurchStChambers)
	av6 := corner.Sequence(av8st145, av8st63, av6st63, av6st4, stHoustonAv2)
	a := *corner.NewPath("a", "av8", av8, []int{0, 0, 1})
	b := *corner.NewPath("b", "av6", av6, []int{-1, 0, 0, 0})
	c.PrintAll(width, height, "nyc", rbase, rsep, append(paths, a), append(paths, b))
}
