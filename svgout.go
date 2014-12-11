package seq

import svg "github.com/metaleap/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
import "io"
import "fmt"
import "encoding/xml"
import "bytes"

var _ = fmt.Printf

// Wrap the generated SVG type per instructions from xsd-makepkg
type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func (d *diagram) makeSvg(out io.Writer) {

	img := &SvgDoc{}

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
