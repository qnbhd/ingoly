package parser

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
)

type Printer struct {
	IndentLevel int
}

func NewPrinter() Printer {
	return Printer{IndentLevel: 0}
}

func (w Printer) EnterNode(n Node) bool {

	for i := 0; i < w.IndentLevel; i++ {
		fmt.Print("   ")
	}

	switch s := n.(type) {
	case *BlockNode:
		w.IndentLevel++
		for _, node := range s.Nodes {
			node.Walk(w)
		}
		return false
	case *BinaryNode:
		color.Green("!--> Binary Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *DeclarationNode:
		color.Green("!--> Declaration Variable Parse (Parse) var '" + s.Variable + "' " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.Expression.Walk(w)
		return false

	case *UnaryNode:
		color.Green("!--> Unary Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.op1.Walk(w)
		return false

	case *UsingVariableNode:
		color.Blue("!--> Using Variable (Value) '" + s.name + "' " + "Line " + strconv.Itoa(s.Line))
		return false

	case *ScopeVar:
		color.Green("!--> Scope Variable (Value) '" + s.Name + "' " + "Line " + strconv.Itoa(s.Line))
		return false

	case *ValueNode:
		color.Blue("!--> Value Node (Value) '" + s.value.AsString() + "' " + "Line " + strconv.Itoa(s.Line))
		return false

	case *ConditionalNode:
		color.Green("!--> Logical Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *PrintNode:
		color.Magenta("!--> Print Operator (Keyword) " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.node.Walk(w)
		return false

	case *IfNode:
		color.Green("!--> If Else Block " + "Line " + strconv.Itoa(s.Line))
		s.node.Walk(w)
		color.Green("!--> If Case " + "Line " + strconv.Itoa(s.Line+1))
		s.ifStmt.Walk(w)

		if s.elseStmt != nil {
			color.Green("!--> Else Case ")
			s.elseStmt.Walk(w)
		}

		return false

	case *ForNode:

		color.Green("!--> For Block [iterVar '%s'] [%3.3f; %3.3f; %3.3f] Line: %d",
			s.iterVar, s.start.AsNumber(), s.stop.AsNumber(), s.step.AsNumber(), s.Line)
		s.stmt.Walk(w)
		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
