package seq

import (
	"fmt"
)

var _ = fmt.Printf

type participant string

type statement struct {
	from  participant
	to    participant
	label string
}

type diagram struct {
	Participants []participant
	statements   []*statement
}

func (d *diagram) addStatement(s *statement) {
	d.Participants = append(d.Participants, s.getParticipants()...)
	d.statements = append(d.statements, s)
}

func (s *statement) getParticipants() []participant {
	participants := make([]participant, 0, 2)

	participants = append(participants, s.from, s.to)

	return participants
}
