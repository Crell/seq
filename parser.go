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
	from  participant
	to    participant
	label string
}

type Parser struct {
	out     chan *statement
	rules   map[string]*regexp.Regexp
}

func NewParser(out chan *statement) *Parser {
	p := &Parser{out: out}

	p.rules = map[string]*regexp.Regexp{
		"forwardArrow": regexp.MustCompile(`(\w+)\s*->\s*(\w+):\s*(\w+)`),
		"backArrow":    regexp.MustCompile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`),
	}

	//p.forwardArrow := regexp.MustCompile(`(\w+)\s*->\s*(\w+):\s*(\w+)`)
	//p.backArrow := regexp.MustCompile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`)

	return p
}

func (p *Parser) Parse(inStream *bufio.Reader) {
	scanner := bufio.NewScanner(inStream)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			stmt, err := p.parseLine(text)
			if err == nil {
				p.out <- stmt
			}
		}
	}()
}

func (p *Parser) parseLine(text string) (*statement, error) {

	forwardArrow := regexp.MustCompile(`(\w+)\s*->\s*(\w+):\s*(\w+)`)
	backArrow := regexp.MustCompile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`)

	switch {
	case p.rules["forwardArrow"].MatchString(text):
		matches := forwardArrow.FindStringSubmatch(text)
		return &statement{from: participant(matches[1]), to: participant(matches[2]), label: matches[3]}, nil
	case p.rules["backArrow"].MatchString(text):
		matches := backArrow.FindStringSubmatch(text)
		return &statement{from: participant(matches[1]), to: participant(matches[2]), label: matches[3]}, nil
	}

	return nil, &NotParsableLine{text: text}
}
