package seq

import (
	"bufio"
	"fmt"
	"regexp"
)

var _ = fmt.Printf

// Error definitions

type NotParsableLine struct {
	text string
}

func (e *NotParsableLine) Error() string {
	return fmt.Sprintf("Could not parse text: %s", e.text)
}

type ruleParser interface {
	action(text string) *statement
	matches(text string) bool
}

// Base "class" for rule parsers.  Ish. Not really.
type parseRule struct {
	regex *regexp.Regexp
}

type forwardArrowRule struct {
	parseRule
}

func NewForwardArrowRule() *forwardArrowRule {
	o := &forwardArrowRule{}
	o.regex = regexp.MustCompile(`(\w+)\s*->\s*(\w+):\s*(\w+)`)
	return o
}

func (p *forwardArrowRule) action(text string) *statement {
	matches := p.regex.FindStringSubmatch(text)
	return &statement{from: participant(matches[1]), to: participant(matches[2]), label: matches[3]}
}

func (p *forwardArrowRule) matches(text string) bool {
	return p.regex.MatchString(text)
}

type backArrowRule struct {
	parseRule
}

func NewBackArrowRule() *backArrowRule {
	o := &backArrowRule{}
	o.regex = regexp.MustCompile(`(\w+)\s*<-\s*(\w+):\s*(\w+)`)
	return o
}

func (p *backArrowRule) action(text string) *statement {
	matches := p.regex.FindStringSubmatch(text)
	return &statement{from: participant(matches[2]), to: participant(matches[1]), label: matches[3]}
}

func (p *backArrowRule) matches(text string) bool {
	return p.regex.MatchString(text)
}

type Parser struct {
	out   chan *statement
	rules []ruleParser
}

func NewParser(out chan *statement) *Parser {
	p := &Parser{out: out}

	p.addRule(NewForwardArrowRule())
	p.addRule(NewBackArrowRule())

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
