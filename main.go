package main

import (
	"fmt"
	"os"
	"github.com/zyxw59/subwayMap/corner"
	"github.com/ajstarks/svgo"
)

func main() {
	var paths []corner.Path
	width := 2000
	height := 2000
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	// begin line definitions
	canvas.Def()
	for p := range paths {
		canvas.Path(fmt.Sprint(p))
	}
	// end line definitions
	canvas.DefEnd()
	canvas.End()
}

const header = `<?xml version="1.0" encoding="utf-8" ?>
<?xml-stylesheet type="text/css" href="lines.css"?>
<svg height="2000" width="2000"
viewBox="-1000,-1000 2000,2000"
xmlns="http://www.w3.org/2000/svg"
xmlns:ev="http://www.w3.org/2001/xml-events"
xmlns:xlink="http://www.w3.org/1999/xlink">`
