package parser

type Ast struct {
	Tree []Node
}

func (ast *Ast) Print() {
	p := NewPrinter()
	for _, stmt := range ast.Tree {
		stmt.Walk(p)
	}
}

//func (ast *Ast) Execute() {
//	p := NewExecutor()
//	for _, stmt := range ast.Tree {
//		stmt.Walk(p)
//	}
//}
