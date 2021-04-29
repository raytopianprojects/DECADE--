package main

import (
	"fmt"
)

func (tok *token) prettyPrint() {
	types := []string{"NULL TOKEN", "NORMAL TOKEN", "OPENER", "CLOSER"}

	fmt.Printf("location: (%d, %d)\n", tok.lineno, tok.charno)
	fmt.Printf("type:     %s\n", types[tok.t])
	fmt.Printf("value:    %s\n", tok.value)
}

func main() {
	l := lexer{
		"{first line}\n{comment}[JAM main [] [[!print \"hello world\"]] NOBOT]",
		0,
	}

	toks := l.lex()

	for i:=0; i < len(toks); i++ {
		toks[i].prettyPrint()
		fmt.Println("-------------------------")
	}
}
