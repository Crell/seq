package seq

import "github.com/ajstarks/svgo"
import "io"
import "fmt"

var _ = fmt.Printf

func (d*diagram) makeSvg(out io.Writer) {

	width := 500
	height := 500
	canvas := svg.New(out)
	canvas.Start(width, height)

	canvas.Def()
	for _, participant := range d.participants {
		nameLength := len(participant)

		canvas.Rect(1, 1, 10, nameLength + 10)
		canvas.Text(1, 1, string(participant))
	}
	canvas.DefEnd()

	canvas.End()
//	canvas.Circle(width/2, height/2, 100)
//	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
}
