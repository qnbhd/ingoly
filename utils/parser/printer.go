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
		color.Green("!--> Block")
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

	case *AssignNode:
		color.Green("!--> Assign Variable Parse (Parse) var '" + s.Variable + "' " + "Line " + strconv.Itoa(s.Line))
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

	case *IntNumber:
		color.Blue("!--> Integer (Number) Value: %d, Line: %d", s.value, s.Line)
		return false

	case *FloatNumber:
		color.Blue("!--> Float (Number) Value: %3.3f, Line: %d", s.value, s.Line)
		return false

	case *String:
		color.Blue("!--> String (String) Value: %s, Line: %d", s.value, s.Line)
		return false

	case *Boolean:
		color.Blue("!--> Boolean (Boolean) Value: %t, Line: %d", s.value, s.Line)
		return false

	case *ConditionalNode:
		color.Green("!--> Logical Operation (Operation) '" + s.operation + "' " + "Line " + strconv.Itoa(s.Line))
		w.IndentLevel++
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *KeywordOperatorNode:
		color.Magenta("!--> %s Operator (Keyword) "+"Line "+strconv.Itoa(s.Line), s.operator)
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

		w.IndentLevel--
		return false

	case *ForNode:
		color.Green("!--> For Block [IterVar: '%s'] Line: %d", s.iterVar, s.Line)

		w.IndentLevel++

		color.Green("!--> Start Section " + "Line " + strconv.Itoa(s.Line+1))
		s.start.Walk(w)

		color.Green("!--> Stop Section " + "Line " + strconv.Itoa(s.Line+1))
		s.stop.Walk(w)

		color.Green("!--> Step Section " + "Line " + strconv.Itoa(s.Line+1))
		s.step.Walk(w)

		w.IndentLevel--

		color.Green("!--> Iter Code " + "Line " + strconv.Itoa(s.Line+1))
		s.stmt.Walk(w)

		return false

	case *Break:
		color.HiCyan("!--> Break (Statement)" + "Line " + strconv.Itoa(s.Line+1))
		return false

	case *Continue:
		color.HiCyan("!--> Continue (Statement)" + "Line " + strconv.Itoa(s.Line+1))
		return false

	case *While:
		color.Green("!--> While Block (Statement) Line: %d", s.Line)

		color.Green("!--> Cycle condition " + "Line " + strconv.Itoa(s.Line+1))
		s.condition.Walk(w)

		color.Green("!--> Cycle Statement " + "Line " + strconv.Itoa(s.Line+1))
		s.stmt.Walk(w)

		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
