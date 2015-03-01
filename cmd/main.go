package main

import (
	"bufio"
	"github.com/Crell/seq"
	"github.com/codegangsta/cli"
	"io"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Seq"
	app.Usage = "A simple sequence diagram generating app"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "in",
			Value: "",
			Usage: "Input file. Leave blank to use standard in.",
		},
		cli.StringFlag{
			Name:  "out",
			Value: "",
			Usage: "Output file. Leave blank to use standard out.",
		},
	}

	app.Action = cliAction

	app.Run(os.Args)
}

func cliAction(c *cli.Context) {
	inFile := c.String("in")
	if inFile != "" {
		// Open the file
	} else {
		// Open stdin
	}

	outFile := c.String("out")
	var outStream io.Writer
	if outFile != "" {
		outStream, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}
		// close outStream on exit and check for its returned error.
		defer func() {
			if err := outStream.Close(); err != nil {
				panic(err)
			}
		}()
	} else {
		// Open stdout
	}

	// Open file, make a bufio.Reader out of it
	inStream := bufio.NewReader(strings.NewReader("A -> B: Test\nC->D:More"))

	diagram := seq.MakeDiagram(inStream)

	diagram.MakeSvg(outStream)
}
