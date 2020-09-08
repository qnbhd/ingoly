package parser

import (
	"fmt"
	"strings"
)

type Ast struct {
	Tree      []Statement
	variables map[string]Value
}

func (ast Ast) _printRecursiveStatement(currentStatement Statement, indentLevel int) {
	if currentStatement == nil {
		return
	}

	var builder strings.Builder

	for i := 0; i < indentLevel; i++ {
		builder.WriteString("   ")
	}

	builder.WriteString("!--> ")
	builder.WriteString(currentStatement.ToString())
	builder.WriteString(" ")

	fmt.Println(builder.String())

	for _, node := range currentStatement.getNodesList() {
		ast._printRecursiveNode(node, indentLevel+1)
	}
}

func (ast Ast) _printRecursiveNode(currentNode Node, indentLevel int) {
	if currentNode == nil {
		return
	}

	var builder strings.Builder

	for i := 0; i < indentLevel; i++ {
		builder.WriteString("   ")
	}

	builder.WriteString("!--> ")
	builder.WriteString(currentNode.ToString())
	builder.WriteString(" ")

	fmt.Println(builder.String())

	for _, node := range currentNode.getNodesList() {
		ast._printRecursiveNode(node, indentLevel+1)
	}
}

func (ast Ast) PrintRecursive() {
	for _, instruction := range ast.Tree {
		ast._printRecursiveStatement(instruction, 0)
	}
}
