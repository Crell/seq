package seq

import "testing"
import "strings"
import "bufio"

func TestParse(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A -> B: Test"))

	c := make(chan *statement)

	Parse(s, c)
	x := <-c

	if x.from != "A" {
		t.Error("From property incorrect")
	}
	if x.to != "B" {
		t.Error("To property incorrect")
	}

}
