package seq

import (
	"bufio"
	"fmt"
	"regexp"
)

var _ = fmt.Printf

type participant string

// Error definitions

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

type ruleParser interface {
	action(text string) *statement
	matches(text string) bool
}

// Base "class" for rule parsers.  Ish. Not really.
type parseRule struct {
	regex *regexp.Regexp
}

type forwardArrow struct {
	parseRule
}

func NewForwardArrow() *forwardArrow {
	o := &forwardArrow{}
	o.regex = regexp.MustCompile(`(\w+)\s*->\s*(\w+):\s*(\w+)`)
	return o
}

func (p *forwardArrow) action(text string) *statement {
	matches := p.regex.FindStringSubmatch(text)
	return &statement{from: participant(matches[1]), to: participant(matches[2]), label: matches[3]}
}

func (p *forwardArrow) matches(text string) bool {
	return p.regex.MatchString(text)
}

type backArrow struct {
	parseRule
}

func NewBackArrow() *backArrow {
	o := &backArrow{}
	o.regex = regexp.MustCompile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`)
	return o
}

func (p *backArrow) action(text string) *statement {
	matches := p.regex.FindStringSubmatch(text)
	return &statement{from: participant(matches[1]), to: participant(matches[2]), label: matches[3]}
}

func (p *backArrow) matches(text string) bool {
	return p.regex.MatchString(text)
}

type Parser struct {
	out   chan *statement
	rules []ruleParser
}

func NewParser(out chan *statement) *Parser {
	p := &Parser{out: out}

	p.addRule(NewForwardArrow())
	p.addRule(NewBackArrow())

	return p
}

func (p *Parser) addRule(rule ruleParser) {
	p.rules = append(p.rules, rule)
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
	for _, rule := range p.rules {
		if rule.matches(text) {
			return rule.action(text), nil
		}
	}

	return nil, &NotParsableLine{text: text}
}
