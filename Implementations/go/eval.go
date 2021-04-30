package main

type atom struct {
	data string
	t string
}

type expr interface {
	eval() atom
}

type expression struct {
	children []expr
}

func (a *atom) eval() atom {
	return *a
}

func (e *expression) eval() atom {
	for i := 0; i < len(e.children); i++ {
	}
	return atom{}
}
