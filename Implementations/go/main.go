package main

import (
	"fmt"
)

func parseError(inp string) {
	fmt.Printf("\033[31merror\033[0m: %s\n", inp)
}

func (tok *token) prettyPrint() {
	types := []string{"NULL TOKEN", "KEYWORD", "OPENER", "CLOSER", "STRING", "INT"}

	fmt.Printf("location: (%d, %d)\n", tok.lineno, tok.charno)
	fmt.Printf("type:     %s\n", types[tok.t])
	fmt.Printf("value:    %s\n", tok.value)
}

func exprPrint(e expr, off string) {
	switch e.(type) {
	case *atom:
		fmt.Printf("%s%v\n", off, e.(*atom))
	case *expression:
		fmt.Println(off + "[")
		for i:=0; i < len(e.(*expression).children); i++ {
			exprPrint(e.(*expression).children[i], off + "  ")
		}
		fmt.Println(off + "]")
	}
}

func main() {
	l := lexer{
		"[[JAM somefunc [arg1, arg2] [] NOBOT][JAM main [] [[!print \"hello world\"]] NOBOT]]",
		0,
	}

	toks := l.lex()

	for i:=0; i < len(toks); i++ {
		toks[i].prettyPrint()
		fmt.Printf("%d: -------------------------\n", i+1)
	}

	p := parser{
		toks,
		1,
	}

	res := p.parse()
	exprPrint(&res, "")
	funcs := getfuncs(res.children)
	if len(funcs) > 0 {
		fmt.Println(funcs)
		exprPrint(funcs["main"].body, "")
	}
}
