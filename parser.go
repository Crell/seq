package seq

import (
	"bufio"
	"fmt"
	"regexp"
)

var _ = fmt.Printf

type participant string

type NotParsableLine struct {
	text string
}

func (e *NotParsableLine) Error() string {
	return fmt.Sprintf("Could not parse text: %s", e.text)
}

type statement struct {
	from participant
	to   participant
}

func Parse(r *bufio.Reader, out chan *statement) string {

	scanner := bufio.NewScanner(r)

	go scan(scanner, out)

	return "end"
}

func scan(scanner *bufio.Scanner, out chan *statement) {
	for scanner.Scan() {
		text := scanner.Text()
		stmt, err := parseLine(text)
		if err == nil {
			out <- stmt
		}

	}
}

func parseLine(text string) (*statement, error) {

	forwardArrow, _ := regexp.Compile(`(\w+)\s*->\s*(\w+):\s*(\w+)`)
	backArrow, _ := regexp.Compile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`)

	switch {
	case forwardArrow.MatchString(text):
		matches := forwardArrow.FindStringSubmatch(text)
		return &statement{from: participant(matches[1]), to: participant(matches[2])}, nil
	case backArrow.MatchString(text):
		matches := backArrow.FindStringSubmatch(text)
		return &statement{from: participant(matches[1]), to: participant(matches[2])}, nil
	}

	return nil, &NotParsableLine{text: text}
}
