package parser

import (
	"fmt"
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
		fmt.Println("!--> Binary Operation (Operation) '" + string(s.operation) + "'")
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *AssignmentNode:
		fmt.Println("!--> Assignment Statement (Statement) var '" + s.Variable + "'")
		s.Expression.Walk(w)
		return false

	case *UnaryNode:
		fmt.Println("!--> Unary Operation (Operation) '" + string(s.operation) + "'")
		s.op1.Walk(w)
		return false

	case *UsingVariableNode:
		fmt.Println("!--> Using Variable (Value) '" + s.name + "'")
		return false

	case *ScopeVar:
		fmt.Println("!--> Scope Variable (Value) '" + s.Name + "'")
		return false

	case *ValueNode:
		fmt.Println("!--> Value Node (Value) '" + s.value.AsString() + "'")
		switch s := s.value.(type) {
		case *NumberValue:
			fmt.Print(s.value)
		case *StringValue:
			fmt.Print(s.value)
		}
		return false

	case *ConditionalNode:
		fmt.Println("!--> Logical Operation (Operation) '" + string(s.operation) + "'")
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *PrintNode:
		fmt.Println("!--> Print Operator (Keyword)")
		s.node.Walk(w)
		return false

	case *IfNode:
		fmt.Println("!--> If Else Block")
		s.node.Walk(w)
		fmt.Println("!--> If Block")
		s.ifStmt.Walk(w)
		fmt.Println("!--> Else Block")
		s.elseStmt.Walk(w)
		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
