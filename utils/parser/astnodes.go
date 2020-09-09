package parser

import (
	"strconv"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 3, 64)
}

/* Base Node */

type Node interface {
	Walk(v Visitor)
}

//////////////////

type ScopeVar struct {
	Name string
}

func NewScopeVar(name string) *ScopeVar {
	return &ScopeVar{Name: name}
}

func (as *ScopeVar) Walk(v Visitor) {
	v.EnterNode(as)

	v.LeaveNode(as)
}

////////////////

/* Binary Node */

type BinaryNode struct {
	operation string
	op1       Node
	op2       Node
}

func (bn *BinaryNode) New(operation string, exp1, exp2 Node) *BinaryNode {
	return &BinaryNode{operation: operation, op1: exp1, op2: exp2}
}

func (bn *BinaryNode) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}

	bn.op1.Walk(v)
	bn.op2.Walk(v)

	v.LeaveNode(bn)
}

////////////////////////////

/* Unary Node */

type UnaryNode struct {
	operation string
	op1       Node
}

func (un *UnaryNode) New(operation string, op1 Node) *UnaryNode {
	return &UnaryNode{operation: operation, op1: op1}
}

func (un *UnaryNode) Walk(v Visitor) {
	if !v.EnterNode(un) {
		return
	}

	un.op1.Walk(v)

	v.LeaveNode(un)
}

////////////////////////////

/* Name Node */

type UsingVariableNode struct {
	name string
}

func (uvn *UsingVariableNode) Walk(v Visitor) {
	if !v.EnterNode(uvn) {
		return
	}

	v.LeaveNode(uvn)
}

///////////////

/* Value Node */

type ValueNode struct {
	value Value
}

func (vn *ValueNode) Walk(v Visitor) {
	if !v.EnterNode(vn) {
		return
	}

	v.LeaveNode(vn)
}

////////////////

/* Binary Node */

type ConditionalNode struct {
	operation string
	op1       Node
	op2       Node
}

func (bn *ConditionalNode) New(operation string, exp1, exp2 Node) *ConditionalNode {
	return &ConditionalNode{operation: operation, op1: exp1, op2: exp2}
}

func (bn *ConditionalNode) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}

	bn.op1.Walk(v)
	bn.op2.Walk(v)

	v.LeaveNode(bn)
}

//////////////////

/* Assignment Node */

type DeclarationNode struct {
	Variable   string
	Expression Node
}

func (as *DeclarationNode) New(variable string, node Node) *DeclarationNode {
	return &DeclarationNode{Variable: variable, Expression: node}
}

func (as *DeclarationNode) Walk(v Visitor) {
	if !v.EnterNode(as) {
		return
	}

	as.Expression.Walk(v)

	v.LeaveNode(as)
}

//////////////////

/* Print Node */

type PrintNode struct {
	node Node
}

func (ps *PrintNode) Walk(v Visitor) {
	if !v.EnterNode(ps) {
		return
	}

	ps.node.Walk(v)

	v.LeaveNode(ps)
}

//////////////////

/* If Node */

type IfNode struct {
	node     Node
	ifStmt   Node
	elseStmt Node
}

func (is *IfNode) Walk(v Visitor) {
	if !v.EnterNode(is) {
		return
	}

	is.node.Walk(v)
	is.ifStmt.Walk(v)
	is.elseStmt.Walk(v)

	v.LeaveNode(is)
}
