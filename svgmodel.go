package seq

import "fmt"

var _ = fmt.Printf

// Values are in svg units (~pixels)
const (
	participantMargin = 5
	statementSpacing  = 3
	characterWidth = 3
	statementHeight = 5
)

type svgParticipant struct {
	participant
	XCoord int
	Width  int
}

func (p *svgParticipant) CenterX() int {
	return p.XCoord + p.SvgWidth()/2
}

func (p *svgParticipant) SvgHeight() float32 {
	return 1.5
}

func (p *svgParticipant) SvgWidth() int {
	return len(p.participant) + 1
}

type svgStatement struct {
	*statement
	YCoord int
}

func (s *svgStatement) spacing() int {
	return 3
}

type svgDiagram struct {
	participants []svgParticipant
	statements   []*svgStatement
}

func (d *svgDiagram) PartcipantHeight() float32 {
	return 1.5
}

func NewSvgDiagram(d *diagram) *svgDiagram {
	svgD := &svgDiagram{}

	var left = 0

	for _, p := range d.Participants {
		svgP := svgParticipant{participant: p}
		svgP.XCoord = left
		left += svgP.Width + participantMargin
		svgD.participants = append(svgD.participants, svgP)
	}

	var top = svgD.PartcipantHeight() + statementSpacing

	for _, s := range d.statements {
		svgS := &svgStatement{statement: s}

		svgS.YCoord = top

		svgD.statements = append(svgD.statements, svgS)
	}

	return svgD
}

/*
participant width = text length + pad
participant height = const
participant X = previous X + const
participant Y = 0
participantBottom Y = last statement + pad
participantLine x = participant center()
participantLine y1 = participant Y + participant height
participantLine Y2 = participantBottom Y

statement height = const (dynamic later?)
statement X1 = statement.from.centerX
statement Y = previous Y1 + previous height + pad
statement X2 = statement.to.centerX
statement label X = ?
statement label Y = ?

diagram width = count(participants) + count(participants - 1)*pad
diagram height = statement total height * count(statements)*pad + participant height
*/

type ParticipantSvg struct {
	participant
	width    int // text length + pad
	height   int // const
	X        int // previous X + const
	lowerY   int // last statement + pad
	previous *ParticipantSvg
}

type ParticipantSvgLine struct {
	X  int // participant Y + participant height
	Y1 int // participant Y + participant height
	Y2 int // participantLowerY Y
}

type StatementSvg struct {
	statement
	height   int // const (dynamic later?)
	x1       int // statement.from.centerX
	x2       int // statement.to.centerX
	y        int // previous Y1 + previous height + pad
	LabelX   int // ?
	LabelY   int // ?
	previous *StatementSvg
}

type DiagramSvg struct {
	participants []ParticipantSvg
	statements   []*StatementSvg
	width        int // count(participants) + count(participants - 1)*pad
	height       int // statement total height * count(statements)*pad + participant height
}

func (s *StatementSvg) Height() int {
	return statementHeight
}

func (s *StatementSvg) X1() int {
	return s.from.centerX()
}

// count(participants) + count(participants - 1)*pad
func (d *DiagramSvg) Width() int {
	if !d.height {
		d.height = 10
	}
	return d.height
}

// statement total height * count(statements)*pad + participant height
func (d *DiagramSvg) Height() int {
	if !d.width {
		d.width = 10
	}
	return d.width
}

func NewDiagramSvg(d *diagram) *DiagramSvg {
	svgD := &DiagramSvg{}

	// First just copy over everything.
	for _, p := range d.Participants {
		svgP := ParticipantSvg{participant: p}
		svgD.participants = append(svgD.participants, svgP)
	}

	for _, s := range d.statements {
		svgS := &StatementSvg{statement: s}
		svgD.statements = append(svgD.statements, svgS)
	}

	// Now iterate and derive information.
	totalRows := len(svgD.statements) + 2
	rowX = 0

	for row := 0; row < totalRows; row++ {

	}

	return svgD
}
