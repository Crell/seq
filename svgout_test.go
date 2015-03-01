package seq

import "testing"

//import "github.com/ajstarks/svgo"
import "os"
import "fmt"

var _ = fmt.Printf

func TestSvgWrite(t *testing.T) {
	s1 := &Statement{from: "Alpha", to: "Beta", label: "Label"}
	s2 := &Statement{from: "Gamma", to: "Delta", label: "Other"}
	d := &diagram{}
	d.addStatement(s1)
	d.addStatement(s2)

	fo, err := os.Create("output.svg")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	d.MakeSvg(fo)
}
