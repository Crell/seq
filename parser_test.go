package seq

import "testing"
import "strings"
import "bufio"

func TestForwardArrow(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A -> B: Test\nC->D:More"))

	c := make(chan *statement)

	Parse(s, c)

	var x *statement

	x = <-c
	if x.from != "A" {
		t.Error("From property incorrect")
	}
	if x.to != "B" {
		t.Error("To property incorrect")
	}

	x = <-c
	if x.from != "C" {
		t.Error("From property incorrect")
	}
	if x.to != "D" {
		t.Error("To property incorrect")
	}

}

func TestBackArrow(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A <- B: Test\nC<-D:More"))

	c := make(chan *statement)

	Parse(s, c)

	var x *statement

	x = <-c
	if x.from != "A" {
		t.Error("From property incorrect")
	}
	if x.to != "B" {
		t.Error("To property incorrect")
	}

	x = <-c
	if x.from != "C" {
		t.Error("From property incorrect")
	}
	if x.to != "D" {
		t.Error("To property incorrect")
	}

}
