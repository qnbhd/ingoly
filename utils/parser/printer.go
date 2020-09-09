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
	w.IndentLevel++

	switch s := n.(type) {
	case *BinaryNode:
		color.Green("!--> Binary Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *DeclarationNode:
		color.Green("!--> Declaration Variable Statement (Statement) var '" + s.Variable + "' " + "Line " + strconv.Itoa(s.Line))
		s.Expression.Walk(w)
		return false

	case *UnaryNode:
		color.Green("!--> Unary Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
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
		switch s := s.value.(type) {
		case *NumberValue:
			fmt.Print(s.value)
		case *StringValue:
			fmt.Print(s.value)
		}
		return false

	case *ConditionalNode:
		color.Green("!--> Logical Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *PrintNode:
		color.Magenta("!--> Print Operator (Keyword) " + "Line " + strconv.Itoa(s.Line))
		s.node.Walk(w)
		return false

	case *IfNode:
		color.Green("!--> If Else Block" + "Line " + strconv.Itoa(s.Line))
		s.node.Walk(w)
		color.Green("!--> If Case " + "Line " + strconv.Itoa(s.Line+1))

		s.ifStmt.Walk(w)

		if s.elseStmt != nil {
			color.Green("!--> Else Case ")

			s.elseStmt.Walk(w)
		}
		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
