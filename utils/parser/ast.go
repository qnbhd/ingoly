package parser

import (
	"fmt"
	"strings"
)

type Ast struct {
	Tree      []Node
	variables map[string]Value
}

func (ast Ast) _printRecursiveStatement(currentNode Node, indentLevel int) {
	if currentNode == nil {
		return
	}

	var builder strings.Builder

	for i := 0; i < indentLevel; i++ {
		builder.WriteString("   ")
	}

	builder.WriteString("!--> ")
	builder.WriteString(currentNode.ToString())

	fmt.Println(builder.String())

	for _, stmt := range currentNode.getNodesList() {
		ast._printRecursiveStatement(stmt, indentLevel+1)
	}
}

func (ast Ast) PrintRecursive() {
	for _, instruction := range ast.Tree {
		ast._printRecursiveStatement(instruction, 0)
	}
}
