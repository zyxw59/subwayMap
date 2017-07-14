package main

import (
	"github.com/zyxw59/subwayMap/canvas"
	"github.com/zyxw59/subwayMap/corner"
	"os"
)

func stY(st int) float64 {
	if st >= 65 {
		return float64(250-st) * 5
	}
	if st >= 0 {
		return float64(65-st)*10 + stY(65)
	}
	return float64(-st)*5 + stY(0)
}

func main() {
	var (
		c canvas.Canvas
	)
	const (
		width  = 2000
		height = 3000
		rbase  = 30
		rsep   = 10
		av11x  = 90
		av10x  = 100
		av8x   = 200
		av7x   = 300
		av6x   = 400
		av5x   = 450
		av4x   = 550
		av2x   = 650
	)

	// Broadway
	bdwySt181 := corner.Point{av11x, stY(181)}
	bdwySt107 := corner.Point{av11x, stY(107)}
	bdwySt104 := corner.Point{av10x, stY(104)}
	bdwySt77 := corner.Point{av10x, stY(77)}
	bdwySt59 := corner.Point{av8x, stY(59)}
	timesSq := corner.Point{av7x, stY(45)}
	// 8 Av
	av8st145 := corner.Point{av8x, stY(145)}
	av8st53 := corner.Point{av8x, stY(53)}
	av8st14 := corner.Point{av8x, stY(12)}
	av8st4 := corner.Point{av6x, stY(8)}
	// 6 Av
	av6st63 := corner.Point{av6x, stY(63)}
	av6st53 := corner.Point{av6x, stY(53)}
	av6st4 := corner.Point{av6x, stY(0)}
	// 63 St
	st63av2 := corner.Point{av2x, stY(63)}
	// 53 St
	st53av2 := corner.Point{av2x, stY(53)}
	// Lower Manhattan
	// 2 Av (F)
	houstonAv2 := corner.Point{av2x, stY(0)}
	// Rector St (1)
	greenwichRector := corner.Point{av7x, stY(-80)}
	// South Ferry (1)
	southFerry := corner.Point{av2x, stY(-90)}
	// Chambers St (A C)
	churchChambers := corner.Point{av6x, stY(-50)}


	// Lines
	// IRT Broadway--7 Av Line
	av7 := corner.Sequence(bdwySt181, bdwySt107, bdwySt104, bdwySt77, bdwySt59, timesSq, greenwichRector, southFerry)
	// 8 Av trunk
	av8 := corner.Sequence(av8st145, av8st14, av8st4, churchChambers)
	// E to Queens
	av8e := corner.Sequence(st53av2, av8st53, av8st14)
	// 6 Av trunk
	av6 := corner.Sequence(av8st145, av8st53, av6st53, av6st4, houstonAv2)
	// M to Queens
	av6m := corner.Sequence(st53av2, av6st53, av6st4)
	// F to Queens
	av6f := corner.Sequence(st63av2, av6st63, av6st53)

	// Paths
	a1 := *corner.NewPath("1", "av7", av7, []int{0, 0, 0, 0, 0, 0, 0})
	av7lines := []corner.Path{a1}
	bA := *corner.NewPath("a", "av8", av8, []int{1, 0, 2, 2, 2})
	bC := *corner.NewPath("c", "av8", av8, []int{0, -1, 1, 1, 1})
	bE := *corner.NewPath("e", "av8", append(av8e[:2], av8[1:]...), []int{-1, 0, -1, 1, 1, 1})
	av8lines := []corner.Path{bA, bC, bE}
	bB := *corner.NewPath("b", "av6", av6, []int{-1, 0, 0, 1})
	bD := *corner.NewPath("d", "av6", av6, []int{2, 0, 0, 1})
	bF := *corner.NewPath("f", "av6", append(av6f[:2], av6[3:]...), []int{1, -1, 0, 1})
	bM := *corner.NewPath("m", "av6", append(av6m[:2], av6[3:]...), []int{0, -1, 0, 1})
	av6lines := []corner.Path{bB, bD, bF, bM}

	// Draw it
	c = canvas.Canvas{os.Stdout}
	c.PrintAll(width, height, "nyc", rbase, rsep, av7lines, av8lines, av6lines)
}
