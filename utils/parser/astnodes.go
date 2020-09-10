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

type BaseNode struct {
	Ast *Ast
}

func (bn *BaseNode) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}

	for _, node := range bn.Ast.Tree {
		node.Walk(v)
	}
}

type BlockNode struct {
	Nodes []Node
}

func (bn *BlockNode) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}

	for _, node := range bn.Nodes {
		node.Walk(v)
	}
}

type ScopeVar struct {
	Name string
	Line int
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
	Line      int
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
	Line      int
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
	Line int
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
	Line  int
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
	Line      int
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
	Line       int
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
	Line int
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
	Line     int
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

type ForNode struct {
	iterVar string
	start   NumberValue
	stop    NumberValue
	step    NumberValue
	strict  bool
	stmt    Node
	Line    int
}

func (fn *ForNode) Walk(v Visitor) {
	if !v.EnterNode(fn) {
		return
	}

	fn.stmt.Walk(v)

	v.LeaveNode(fn)
}
