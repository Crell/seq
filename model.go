package seq

import (
	"fmt"
)

var _ = fmt.Printf

type participant string

type Statement struct {
	from  participant
	to    participant
	label string
}

type StatementFeed chan *Statement

type diagram struct {
	Participants []participant
	statements   []*Statement
}

func (d *diagram) addStatement(s *Statement) {
	d.Participants = append(d.Participants, s.getParticipants()...)
	d.statements = append(d.statements, s)
}

func (s *Statement) getParticipants() []participant {
	participants := make([]participant, 0, 2)

	participants = append(participants, s.from, s.to)

	return participants
}
