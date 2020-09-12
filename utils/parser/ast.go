package parser

import (
	"fmt"
	"ingoly/utils/errpull"
)

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

	unique := map[string]string{}
	var filteredPull []errpull.InnerError

	for _, err := range p.ErrorsPull.Errors {
		errTrace := fmt.Sprintf("%q", err.Err)
		errLine := fmt.Sprintf("%d", err.SourceLine)
		hash := errTrace + errLine
		if _, ok := unique[hash]; !ok {
			unique[hash] = ""
			filteredPull = append(filteredPull, err)
		}
	}

	p.ErrorsPull.Errors = filteredPull
	return p.ErrorsPull
}
