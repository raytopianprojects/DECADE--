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
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "str"})
		case tokNumber:
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "int"})
		case tokNull:
			panic(fmt.Sprintf("lol. Decade-- doesn't give error messages. Here is a number: %d. Try googling it", rand.Int()))
		default:
			out.children = append(out.children, &atom{p.inp[p.cursor].value, "kw"})
		}
	}
	return out
}
