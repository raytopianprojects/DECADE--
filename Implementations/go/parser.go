package main

import (
	"math/rand"
	"fmt"
)

type parser struct {
	inp []token
	cursor int
}

func (p *parser) parse() expression {
	out := expression{}

	for p.cursor=p.cursor; p.cursor < len(p.inp); p.cursor++ {
		switch p.inp[p.cursor].t {
		case tokOpen:
			p.cursor++
			res := p.parse()
			out.children = append(out.children, &res)
		case tokClose:
			return out
		case tokString:
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "str", p.inp[p.cursor].lineno, p.inp[p.cursor].charno})
		case tokNumber:
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "int", p.inp[p.cursor].lineno, p.inp[p.cursor].charno})
		case tokNull:
			panic(fmt.Sprintf("lol. Decade-- doesn't give error messages. Here is a number: %d. Try googling it", rand.Int()))
		default:
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "kw", p.inp[p.cursor].lineno, p.inp[p.cursor].charno})
		}
	}
	return out
}

func getfuncs(inp []expr) map[string]fn {
	tr := map[string]fn{}
	for i:=0; i < len(inp); i++ {
		switch inp[i].(type) {
		case *expression:
			e := inp[i].(*expression)

			if len(e.children) != 5 {
				continue
			}

			if e.children[0].(*atom) == nil || e.children[0].(*atom).data != "JAM" {
				continue
			}

			if e.children[1].(*atom) == nil {
				jamKw := e.children[0].(*atom)
				parseError(fmt.Sprintf("invalid function declaration on (%d, %d)", jamKw.lineno, jamKw.charno+len(jamKw.data)))
				return nil
			}

			if e.children[2].(*expression) == nil {
				l, c := e.children[2].getpos()
				parseError(fmt.Sprintf("invalid function declaration on (%d, %d)", l, c))
			}

			if e.children[3].(*expression) == nil {
				l, c := e.children[3].getpos()
				parseError(fmt.Sprintf("invalid function declaration on (%d, %d)", l, c))
			}

			if e.children[4].(*atom) == nil {
				l, c := e.children[4].getpos()
				parseError(fmt.Sprintf("invalid function declaration on (%d, %d)", l, c))
			}

			tr[e.children[1].(*atom).data] = fn{len(e.children[2].(*expression).children), e.children[3].(*expression).children}
		}
	}
	return tr
}
