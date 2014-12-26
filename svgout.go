package seq

import "io"
import "fmt"
import "text/template"

var _ = fmt.Printf

func (d *diagram) makeSvg(out io.Writer) {

	funcMap := template.FuncMap{
	}

	tmpl, err := template.New("svgtemplate.tpl.svg").Funcs(funcMap).ParseFiles("./svgtemplate.tpl.svg")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(out, d)

	if err != nil {
		panic(err)
	}
}

func (p participant) SvgXCoord(index int) int {
	return index * 7
}

func (d *diagram) SvgWidth() int {
	var width int

	for _, p := range d.Participants {
		width += p.SvgWidth()
	}

	// Give 2 em of spacing between each participant.
	width += len(d.Participants) * 2

	return width
}

func (p participant) SvgWidth() int {
	return len(p) + 1
}
