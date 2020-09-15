package parser

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
	Line  int
}

func (bn *BlockNode) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}

	for _, node := range bn.Nodes {
		node.Walk(v)
	}

	v.LeaveNode(bn)
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

type ScopeVar struct {
	name string
	Line int
}

func (uvn *ScopeVar) Walk(v Visitor) {
	if !v.EnterNode(uvn) {
		return
	}

	v.LeaveNode(uvn)
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

type AssignNode struct {
	Variable   string
	Expression Node
	Line       int
}

func (as *AssignNode) Walk(v Visitor) {
	if !v.EnterNode(as) {
		return
	}

	as.Expression.Walk(v)

	v.LeaveNode(as)
}

//////////////////

/* Print Node */

type FunctionalNode struct {
	arguments []Node
	operator  string
	Line      int
}

func (ps *FunctionalNode) Walk(v Visitor) {
	if !v.EnterNode(ps) {
		return
	}

	for _, arg := range ps.arguments {
		arg.Walk(v)
	}

	v.LeaveNode(ps)
}

type FunctionDeclareNode struct {
	name             string
	args             []VarWithAnnotation
	returnAnnotation string
	body             Node
	Line             int
}

func (ps *FunctionDeclareNode) Walk(v Visitor) {
	if !v.EnterNode(ps) {
		return
	}

	ps.body.Walk(v)

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
	start   Node
	stop    Node
	step    Node
	strict  bool
	stmt    Node
	Line    int
}

func (fn *ForNode) Walk(v Visitor) {
	if !v.EnterNode(fn) {
		return
	}

	fn.stmt.Walk(v)
	fn.start.Walk(v)
	fn.stop.Walk(v)
	fn.step.Walk(v)

	v.LeaveNode(fn)
}

type While struct {
	condition Node
	stmt      Node
	Line      int
}

func (w *While) Walk(v Visitor) {
	if !v.EnterNode(w) {
		return
	}

	w.condition.Walk(v)
	w.stmt.Walk(v)

	v.LeaveNode(w)
}

type Break struct {
	Line int
}

func (b *Break) Walk(v Visitor) {
	if !v.EnterNode(b) {
		return
	}

	v.LeaveNode(b)
}

type Continue struct {
	Line int
}

func (b *Continue) Walk(v Visitor) {
	if !v.EnterNode(b) {
		return
	}

	v.LeaveNode(b)
}

type Return struct {
	value Node
	Line  int
}

func (r *Return) Walk(v Visitor) {
	if !v.EnterNode(r) {
		return
	}

	r.value.Walk(v)

	v.LeaveNode(r)
}

type Array struct {
	Elements []Node
	Line     int
}

func (ar *Array) Walk(v Visitor) {
	if !v.EnterNode(ar) {
		return
	}

	for _, item := range ar.Elements {
		item.Walk(v)
	}

	v.LeaveNode(ar)
}

type CollectionAccess struct {
	variableName string
	index        Node
	Line         int
}

func (aa *CollectionAccess) Walk(v Visitor) {
	if !v.EnterNode(aa) {
		return
	}

	aa.index.Walk(v)

	v.LeaveNode(aa)
}
