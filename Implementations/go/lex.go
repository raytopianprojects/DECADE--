package main

import (
	"strings"
)

// basic lexer type
type lexer struct {
	inp string // input, that is lexed
	cursor int // current position in input
}

type token struct {
	lineno int // line, where token is located
	charno int // collumn, where token is located
	t int // type of token
	value string // value of token
}

const (
	tokNull = 0 // error. currently unused
	tokNormal = 1 // normal token. keywords, strings, numbers etc.
	tokOpen = 2 // [
	tokClose = 3 // ]
	tokEnders = " []\n" // characters, that terminate token
)

// returns next byte in token
func (l *lexer) peek() byte {
	l.cursor++
	return l.inp[l.cursor-1]
}

// returns strings starting at cursor
func (l *lexer) peekStr() string {
	tr := ""
	for l.cursor < len(l.inp) && l.inp[l.cursor] != '"' {
		tr += string(l.inp[l.cursor])
		l.cursor++
	}
	l.cursor++
	return tr
}

// returns token starting at cursor
func (l *lexer) peekTok() string {
	tr := ""
	l.cursor--
	for l.cursor < len(l.inp) && !strings.Contains(tokEnders, string(l.inp[l.cursor])) {
		tr += string(l.inp[l.cursor])
		l.cursor++
	}
	return tr
}

// main lexer loop
func (l *lexer) lex() []token {
	tr := []token{}
	var charno, lineno int

	for l.cursor < len(l.inp) {
		switch l.peek() {
		case ' ':
			continue
		case '[':
			tr = append(tr, token{lineno, charno, tokOpen, "["})
		case ']':
			tr = append(tr, token{lineno, charno, tokClose, "]"})
		case '"':
			old := l.cursor
			tr = append(tr, token{lineno, old, tokNormal, l.peekStr()})
		case '\n':
			charno = -1
			lineno++
		default:
			old := l.cursor
			tr = append(tr, token{lineno, old, tokNormal, l.peekTok()})
		}
		charno++
	}

	return tr
}
