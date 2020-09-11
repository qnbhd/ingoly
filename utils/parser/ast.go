package parser

import "ingoly/utils/errpull"

type Ast struct {
	Tree []Node
}

func (ast *Ast) Print() {
	p := NewPrinter()
	for _, stmt := range ast.Tree {
		stmt.Walk(p)
	}
}

func (ast *Ast) Execute() *errpull.ErrorsPull {
	p := NewExecutor()
	for _, stmt := range ast.Tree {
		stmt.Walk(p)
	}
	return p.ErrorsPull
}
