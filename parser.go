package seq

import (
	"bufio"
	//	"io"
	//	"bytes"
)

type statement struct {
	from string
	to   string
}

func Parse(r *bufio.Reader, out chan string) string {

	scanner := bufio.NewScanner(r)

	go scan(scanner, out)

	return "end"

	//	lineBytes, _ := r.ReadSlice('\n')
	//
	//	if err != nil {
	//		// Error handling?
	//	}
	//
	//	//	n := bytes.Index(lineBytes, []byte{0})
	//
	//	line := string(lineBytes[:])

	//return line
	//
	//	if err {
	//		// Error handling?
	//	}
	//
	//	fmt.Scanf(line, )
	//

}

func scan(scanner *bufio.Scanner, out chan string) {
	for scanner.Scan() {
		out <- scanner.Text()
	}

}
