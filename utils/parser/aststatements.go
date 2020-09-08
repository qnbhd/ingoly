package parser

import "fmt"

type Statement interface {
	Execute() error
	ToString() string
	getNodesList() []Node
}

//////////////////

/* Assignment Statement */

type AssignmentStatement struct {
	Variable   string
	Expression Node
}

func (as *AssignmentStatement) New(variable string, node Node) *AssignmentStatement {
	return &AssignmentStatement{Variable: variable, Expression: node}
}

func (as *AssignmentStatement) Execute() error {
	result, ok := as.Expression.Execute()
	if ok == nil {
		VarTable[as.Variable] = result
	}
	return nil
}

func (as *AssignmentStatement) ToString() string {
	return "ASSIGNMENT STATEMENT (Statement) Identifier: '" + as.Variable + "'"
}

func (as *AssignmentStatement) getNodesList() []Node {
	return []Node{as.Expression}
}

//////////////////

/* Print Statement */

type PrintStatement struct {
	node Node
}

func (ps *PrintStatement) Execute() error {
	result, _ := ps.node.Execute()
	fmt.Println(result.AsString())
	return nil
}

func (ps *PrintStatement) ToString() string {
	return "PRINT OPERATOR (Keyword)"
}

func (ps *PrintStatement) getNodesList() []Node {
	return []Node{ps.node}
}
