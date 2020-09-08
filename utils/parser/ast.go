package parser

import (
	"fmt"
	"strings"
)

type Ast struct {
	Tree      []Node
	variables map[string]float64
}

func (ast Ast) _printRecursive(currentNode Node, indentLevel int) {
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
		ast._printRecursive(node, indentLevel+1)
	}
}

func (ast Ast) PrintRecursive() {
	for _, instruction := range ast.Tree {
		ast._printRecursive(instruction, 0)
	}
}
