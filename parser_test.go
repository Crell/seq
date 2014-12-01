package seq

import "testing"
import "strings"
import "bufio"

func TestParse(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A -> B: Test"))

	c := make(chan string)

	Parse(s, c)

	x := <-c

	if x != "A -> B: Test" {
		t.Error("Wrong string")
	}

	//
	//	if x.from != "A" {
	//		t.Error("From property incorrect")
	//	}
	//	if x.to != "B" {
	//		t.Error("To property incorrect")
	//	}
}
