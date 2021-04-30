package main

type fn struct {
	argc int
	body *expression
}

type atom struct {
	data string
	t string
	lineno int
	charno int
}

type expr interface {
	eval() atom
	getpos() (int, int)
}

type expression struct {
	children []expr
	lineno int
	charno int
}

func (a *atom) eval() atom {
	return *a
}

func (a *atom) getpos() (int, int) {
	return a.lineno, a.charno
}

func (e *expression) eval() atom {
	for i := 0; i < len(e.children); i++ {
	}
	return atom{}
}

func (e *expression) getpos() (int, int) {
	return e.lineno, e.charno
}
