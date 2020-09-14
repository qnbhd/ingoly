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

func (w Printer) PrintIndent(relAdd int) {
	for i := 0; i < w.IndentLevel+relAdd; i++ {
		fmt.Print("   ")
	}
}

func (w Printer) EnterNode(n Node) bool {

	w.PrintIndent(0)

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

	case *FunctionalNode:
		color.Magenta("!--> %s Operator (Keyword) "+"Line "+strconv.Itoa(s.Line), s.operator)
		w.IndentLevel++
		for _, arg := range s.arguments {
			arg.Walk(w)
		}
		return false

	case *IfNode:
		color.Green("!--> If Else Block " + "Line " + strconv.Itoa(s.Line))

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> If Condition " + "Line " + strconv.Itoa(s.Line))
		s.node.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> If Case " + "Line " + strconv.Itoa(s.Line+1))
		s.ifStmt.Walk(w)
		w.IndentLevel -= 2

		if s.elseStmt != nil {
			w.IndentLevel += 2
			w.PrintIndent(-1)
			color.Green("!--> Else Case ")
			s.elseStmt.Walk(w)
			w.IndentLevel -= 2
		}

		w.IndentLevel--
		return false

	case *ForNode:
		color.Green("!--> For Block [IterVar: '%s'] Line: %d", s.iterVar, s.Line)

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> Start Section " + "Line " + strconv.Itoa(s.Line+1))
		s.start.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> Stop Section " + "Line " + strconv.Itoa(s.Line+1))
		s.stop.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> Step Section " + "Line " + strconv.Itoa(s.Line+1))
		s.step.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> Iter Code " + "Line " + strconv.Itoa(s.Line+1))
		s.stmt.Walk(w)
		w.IndentLevel -= 2

		return false

	case *Break:
		color.HiCyan("!--> Break (Statement)" + " Line " + strconv.Itoa(s.Line+1))
		return false

	case *Continue:
		color.HiCyan("!--> Continue (Statement)" + " Line " + strconv.Itoa(s.Line+1))
		return false

	case *FunctionDeclareNode:
		color.Green("!--> Declaration Function ['%s'] (Statement) Line: %d ", s.name, s.Line)
		color.HiGreen("   !--> Arg Names: ")
		for _, item := range s.argNames {
			color.Blue("      +- %s", item)
		}
		w.IndentLevel++
		s.body.Walk(w)
		w.IndentLevel--
		return false

	case *While:
		color.Green("!--> While Block (Statement) Line: %d", s.Line)

		color.Green("!--> Cycle condition " + "Line " + strconv.Itoa(s.Line+1))
		s.condition.Walk(w)

		color.Green("!--> Cycle Statement " + "Line " + strconv.Itoa(s.Line+1))
		s.stmt.Walk(w)

		return false

	case *Return:
		color.Magenta("!--> Return (Statement) Line: %d", s.Line)

		w.IndentLevel++
		s.value.Walk(w)
		w.IndentLevel--
		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
