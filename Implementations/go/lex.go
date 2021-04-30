package main

import (
	"strings"
	"strconv"
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
	tokKeyword = 1 // keywords, function names and variables
	tokOpen = 2 // [
	tokClose = 3 // ]
	tokString = 4 // string
	tokNumber = 5 // numbers. Right now only ints TODO
	tokEnders = " []\n" // characters, that terminate token
)

// returns next byte in token
func (l *lexer) peek() byte {
	l.cursor++
	return l.inp[l.cursor-1]
}

// returns strings starting at cursor
func (l *lexer) peekCustom(ender byte) string {
	tr := ""
	for l.cursor < len(l.inp) && l.inp[l.cursor] != ender {
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
			tr = append(tr, token{lineno, charno, tokString, l.peekCustom('"')})
			charno += l.cursor - old
		case '{':
			old := l.cursor
			l.peekCustom('}')
			charno += l.cursor - old
		case '\n':
			charno = -1
			lineno++
		default:
			old := l.cursor
			val := l.peekTok()
			if _, err := strconv.Atoi(val); err == nil {
				tr = append(tr, token{lineno, charno, tokNumber, val})
			} else {
				tr = append(tr, token{lineno, charno, tokKeyword, val})
			}
			charno += l.cursor - old
		}
		charno++
	}

	return tr
}
