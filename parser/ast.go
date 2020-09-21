package parser

import (
	"fmt"
	"ingoly/errpull"
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

func (ast *Ast) Index() *Context {
	p := NewIndexer()
	for _, stmt := range ast.Tree {
		stmt.Walk(p)
	}
	return p.Ctx
}

func (ast *Ast) Execute(indexedContext *Context) *errpull.ErrorsPull {
	p := NewExecutor(indexedContext)

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

	//fmt.Println("**********")
	//fmt.Println(p.mainContext.Vars)
	//fmt.Println("**********")
	//fmt.Println(p.mainContext.Vars["Point"])
	//fmt.Println("**********")

	return p.ErrorsPull
}
