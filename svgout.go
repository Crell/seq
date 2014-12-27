package seq

import "io"
import "fmt"
import "text/template"

var _ = fmt.Printf

func (d *diagram) makeSvg(out io.Writer) {

	funcMap := template.FuncMap{}

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

func (d *diagram) SvgParticipantLowerYCoord() int {
	return len(d.statements) * 3
}

func (d *diagram) SvgParticipantSequenceLineYEnd() int {
	return len(d.statements) * 3
}

func (p participant) SvgWidth() int {
	return len(p) + 1
}

func (p participant) SvgHeight() float32 {
	return 1.5
}

func (p participant) SvgSequenceLineXCoord(index int) int {
	return index * 7
}

func (p participant) SvgSequenceLineYStart(index int) int {
	return 0
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

func (d *diagram) SvgHeight() float32 {
	var height float32

	// Each participant appears twice, so count it twice for overall height.
	for _, p := range d.Participants {
		height += p.SvgHeight() * 2
	}

	// Give 2 em of spacing between each participant.
	//height += len(d.Participants) * 2

	return height
}
