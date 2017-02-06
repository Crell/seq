package seq

// Largely copied from https://talks.golang.org/2011/lex.slide

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type token struct {
	typ tokenType
	val string
}

type tokenType int

const (
	tokenError tokenType = iota
	tokenEOF
	tokenArrow
	tokenParticipant
	tokenBeginStatement
	tokenEndStatement
)

const arrow = "->"

func (t token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}
	return fmt.Sprintf("%q", t.val)
}

type lexer struct {
	start  int
	pos    int
	width  int
	input  string
	tokens chan token
}

type stateFunction func(*lexer) stateFunction

type tokenMap map[tokenType]stateFunction

type tokenizer struct {
	tokens tokenMap
}

func (t *tokenizer) addToken(name tokenType, fn stateFunction) {
	t.tokens[name] = fn
}

func (t *tokenizer) run(start token) {
	for state := start; state != nil; {
		state = t.tokens[state](lexer)
	}
}

func lex(input string) (*lexer, chan token) {
	lexer := &lexer{
		input:  input,
		tokens: make(chan token, 2),
	}

	go lexer.run()
	return lexer, lexer.tokens
}

func (lexer *lexer) run() {
	for state := tokenParticipant; state != nil; {
		state = state(lexer)
	}
}

func (lexer *lexer) emit(t tokenType) {
	lexer.tokens <- token{
		typ: t,
		lexer.input[lexer.start:lexer.pos],
	}
	lexer.start = lexer.pos
}

func (lexer *lexer) next() (rune int) {
	if lexer.pos >= len(lexer.input) {
		lexer.width = 0
		return eof
	}
	rune, lexer.width = utf8.DecodeRuneInString(lexer.input[lexer.pos:])
	lexer.pos += lexer.width
	return rune
}

func (lexer *lexer) ignore() {
	lexer.start = lexer.pos
}

func (lexer *lexer) peek() int {
	rune := lexer.next()
	lexer.backup()
	return rune
}

// accept consumes the next rune if it's from the valid set. The "valid" string is
// a list of runes (characters) that are legal. If the next character is one of those
// return true, else false.
func (lexer *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, lexer.next()) >= 0 {
		return true
	}
	lexer.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (lexer *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, lexer.next()) >= 0 {
	}
	lexer.backup()
}

// backup steps back one rune.
// Can be called only once per call of next.
func (lexer *lexer) backup() {
	lexer.pos -= lexer.width
}

func (lexer *lexer) errorf(format string, args ...interface{}) stateFunction {
	lexer.tokens <- token{
		typ: tokenError,
		val: fmt.Sprintf(format, args...),
	}
	// Stop lexer execution entirely.
	return nil
}

func lexParticipant(lexer *lexer) tokenType {
	for {
		if strings.HasPrefix(lexer.input[lexer.pos:], arrow) {
			if lexer.pos > lexer.start {
				lexer.emit(tokenParticipant)
			}
			return tokenArrow
		}
		if lexer.next() == eof {
			break
		}
	}
	lexer.emit(tokenEOF)
	return nil
}

func lexArrow(lexer *lexer) tokenType {
	lexer.pos += len(arrow)
	lexer.emit(tokenArrow)
	return tokenParticipant
}

func beginStatement(lexer *lexer) tokenType {
	// Where is the isAlpha defined? It's in the slides...
	if isAlpha(lexer.peek()) {
		lexer.emit(tokenBeginStatement)
		return tokenParticipant
	}
	lexer.errorf("Participant names must begin with a letter.")
	return nil
}
