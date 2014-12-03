package seq

import "testing"
import "strings"
import "bufio"

func TestForwardArrow(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A -> B: Test\nC->D:More"))
	c := make(chan *statement)

	p := NewParser(c)
	p.Parse(s)

	var x *statement

	x = <-c
	if x.from != "A" {
		t.Error("From property incorrect")
	}
	if x.to != "B" {
		t.Error("To property incorrect")
	}
	if x.label != "Test" {
		t.Error("Label incorrect")
	}

	x = <-c
	if x.from != "C" {
		t.Error("From property incorrect")
	}
	if x.to != "D" {
		t.Error("To property incorrect")
	}
	if x.label != "More" {
		t.Error("Label incorrect")
	}

}

func TestBackArrow(t *testing.T) {

	s := bufio.NewReader(strings.NewReader("A <- B: Test\nC<-D:More"))
	c := make(chan *statement)

	p := NewParser(c)
	p.Parse(s)

	var x *statement

	x = <-c
	if x.from != "B" {
		t.Error("From property incorrect")
	}
	if x.to != "A" {
		t.Error("To property incorrect")
	}
	if x.label != "Test" {
		t.Error("Label incorrect")
	}

	x = <-c
	if x.from != "D" {
		t.Error("From property incorrect")
	}
	if x.to != "C" {
		t.Error("To property incorrect")
	}
	if x.label != "More" {
		t.Error("Label incorrect")
	}

}
