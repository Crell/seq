package seq

import "testing"
import "fmt"

var _ = fmt.Printf

func TestStatementGetParticipants(t *testing.T) {
	s := &statement{from: participant("A"), to: participant("B"), label: "Label"}

	participants := s.getParticipants()

	if participants[0] != "A" {
		t.Error("First participant of the statement was wrong.")
	}
	if participants[1] != "B" {
		t.Error("Second participant of the statement was wrong.")
	}
}

func TestExtractParticipants(t *testing.T) {

	s1 := &statement{from: "A", to: "B", label: "Label"}
	s2 := &statement{from: "C", to: "D", label: "Other"}

	d := &diagram{}

	d.addStatement(s1)
	d.addStatement(s2)

	if d.Participants[0] != "A" {
		t.Error("First partipant is wrong.")
	}
	if d.Participants[1] != "B" {
		t.Error("Second partipant is wrong.")
	}
	if d.Participants[2] != "C" {
		t.Error("Third partipant is wrong.")
	}
	if d.Participants[3] != "D" {
		t.Error("Fourth partipant is wrong.")
	}
}
