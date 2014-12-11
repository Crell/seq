package seq

import svg "github.com/metaleap/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
import "io"
import "fmt"
import "encoding/xml"
import "bytes"
import xsdt "github.com/metaleap/go-xsd/types"

// Wrap the generated SVG type per instructions from xsd-makepkg
type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func (d *diagram) makeSvg(out io.Writer) {

	img := &SvgDoc{}

	d.addDefs(img)

	// Now render the SVG to an XML string and write it back to
	// the provided stream.
	output, err := xml.MarshalIndent(img, "", "  ")
	if err != nil {
		out.Write(bytes.NewBufferString(fmt.Sprintf("error: %v\n", err)).Bytes())
		return
	}

	out.Write(bytes.NewBufferString(xml.Header).Bytes())
	out.Write(output)
}

func (d *diagram) addDefs(img *SvgDoc) {
	defs := &svg.TdefsType{}

	for i, participant := range d.participants {
		defs.Gs = append(defs.Gs, participantToGroupDef(i, participant))
	}

	img.Defses = append(img.Defses, defs)

}

/*
  Render a participant into something roughly like this:
	<g id="participant-2">
	  <rect x="1" y="1" width="5em" height="1.5em" fill="transparent" stroke="black" stroke-width="1px" />
	  <text x="1em" y="1em" style="font-size: 1em; text-anchor: middle;">Abcdef</text>
	</g>
*/
func participantToGroupDef(i int, participant participant) *svg.TgType {
	group := &svg.TgType{}
	group.Id = xsdt.Id(fmt.Sprintf("participant-%d", i))

	// I don't know why we can't declare these in the struct. The compiler
	// says undefined field? Something quirky with the heavy nesting,
	// maybe?
	rect := &svg.TrectType{}
	rect.X = "1"
	rect.Y = "1"
	rect.Height = "1.5em"
	rect.Fill = "transparent"
	rect.Stroke = "black"
	rect.StrokeWidth = "1px"
	rect.Width = svg.TLengthType(fmt.Sprintf("%dem", len(participant)+2))

	group.Rects = append(group.Rects, rect)

	text := &svg.TtextType{}
	text.X = "1em"
	text.Y = "1em"
	text.Style = "font-size: 1em; text-anchor: middle"
	text.XsdGoPkgCDATA = string(participant)
	group.Texts = append(group.Texts, text)

	return group
}
