package seq

import "github.com/ajstarks/svgo"
import "io"
import "fmt"

var _ = fmt.Printf

func (d *diagram) makeSvg(out io.Writer) {

	width := 500
	height := 500
	canvas := svg.New(out)
	canvas.Start(width, height)

	d.makeSvgParticpantDefs(canvas)

//
//	<use xlink:href="#first" x="100" y="100" />
//	<use xlink:href="#first" x="100" y="300" />
//	<use xlink:href="#second" x="300" y="100" />
//	<use xlink:href="#second" x="300" y="300" />


	canvas.End()
	//	canvas.Circle(width/2, height/2, 100)
	//	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
}

func (d *diagram) makeSvgParticpantDefs(canvas *svg.SVG) {
	canvas.Def()
	for i, participant := range d.participants {
		nameLength := len(participant)

		canvas.Group(fmt.Sprintf("id:participant-%s", i))
		canvas.Rect(0, 0, 0, 0, fmt.Sprintf("fill:transparent;stroke:black;stroke-width:1px;height:1.5em;width:%dem", nameLength+2))
		canvas.Text(1, 1, string(participant), "font-size:1em;text-anchor:middle")

		canvas.Gend()

	}

	/*
	<g id="first">
	<rect x="1" y="1" width="5em" height="1.5em" fill="transparent" stroke="black" stroke-width="1px" />
	<text x="1em" y="1em" style="font-size: 1em; text-anchor: middle;">Abcdef</text>
	</g>
	*/

	// canvas.Use(x, y, "participant-1"

	canvas.DefEnd()

}
