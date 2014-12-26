package seq

import "io"
import "fmt"
import "text/template"

var _ = fmt.Printf

func (d *diagram) makeSvg(out io.Writer) {

	funcMap := template.FuncMap{
		"participantWidth": participantWidth,
		"diagramWidth":     diagramWidth,
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

func participantWidth(p participant) int {
	return len(p) + 1
}

func diagramWidth(d diagram) int {
	var width int

	for _, p := range d.Participants {
		width += participantWidth(p)
	}

	// Give 5 points of spacing between each participant.
	width += len(d.Participants) * 5

	return width
}
