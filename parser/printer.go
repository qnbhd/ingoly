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

func defaultInfoPrint(colorFunctor func(format string, a ...interface{}), line int, mainInformation ...string) {
	formattedString := "!--> "
	for _, item := range mainInformation {
		formattedString += fmt.Sprintf("%s", item)
	}
	colorFunctor("%s |> Line %d", formattedString, line)
}

func (w Printer) EnterNode(n Node) bool {

	w.PrintIndent(0)

	switch s := n.(type) {
	case *BlockNode:
		defaultInfoPrint(color.Green, s.Line, "Block")
		w.IndentLevel++
		for _, node := range s.Nodes {
			node.Walk(w)
		}
		return false
	case *BinaryNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Binary Operation (Operation) '%s'", s.operation))
		w.IndentLevel++
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *DeclarationNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Declaration Variable (Statement) ['%s']", s.Variable))
		w.IndentLevel++
		s.Expression.Walk(w)
		return false

	case *AssignNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Assign Variable (Statement) ['%s']", s.Variable))
		w.IndentLevel++
		s.Expression.Walk(w)
		return false

	case *UnaryNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Unary Operation (Operation) '%s'", s.operation))
		w.IndentLevel++
		s.op1.Walk(w)
		return false

	case *ScopeVar:
		defaultInfoPrint(color.Blue, s.Line, fmt.Sprintf("Using Variable (Value) ['%s']", s.Name))
		return false

	case *IntNumber:
		defaultInfoPrint(color.Blue, s.Line, fmt.Sprintf("Integer Number (Number) Value: '%d'", s.Value))
		return false

	case *FloatNumber:
		defaultInfoPrint(color.Blue, s.Line, fmt.Sprintf("Float Number (Number) Value: '%3.3f'", s.Value))
		return false

	case *String:
		defaultInfoPrint(color.Blue, s.Line, fmt.Sprintf("String (String) Value: '%s'", string(s.value)))
		return false

	case *Boolean:
		defaultInfoPrint(color.Blue, s.Line, fmt.Sprintf("Boolean (Boolean) Value: '%t'", s.value))
		return false

	case *ConditionalNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Logical Operation (Operation) '%s''", s.operation))
		w.IndentLevel++
		s.op1.Walk(w)
		s.op2.Walk(w)
		return false

	case *FunctionalNode:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Operator '%s' (Operator)", s.operator))
		w.IndentLevel++
		for _, arg := range s.arguments {
			arg.Walk(w)
		}
		return false

	case *IfNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("If-Else Block (Block)"))
		w.IndentLevel += 2
		w.PrintIndent(-1)
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("If Condition (Condition)"))
		s.node.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		defaultInfoPrint(color.Green, s.Line+1, fmt.Sprintf("If Case (Case)"))
		s.ifStmt.Walk(w)
		w.IndentLevel -= 2

		if s.elseStmt != nil {
			w.IndentLevel += 2
			w.PrintIndent(-1)
			defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Else Case (Case)"))
			s.elseStmt.Walk(w)
			w.IndentLevel -= 2
		}

		w.IndentLevel--
		return false

	case *ForNode:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("For Block [IterVar: '%s'] (Block)", s.iterVar))

		w.IndentLevel += 2
		w.PrintIndent(-1)
		defaultInfoPrint(color.Green, s.Line+1, fmt.Sprintf("Start Section (Section)"))
		s.start.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		defaultInfoPrint(color.Green, s.Line+1, fmt.Sprintf("Stop Section (Section)"))
		s.stop.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		defaultInfoPrint(color.Green, s.Line+1, fmt.Sprintf("Step Section (Section)"))
		s.step.Walk(w)
		w.IndentLevel -= 2

		w.IndentLevel += 2
		w.PrintIndent(-1)
		color.Green("!--> Iter Code " + "Line " + strconv.Itoa(s.Line+1))
		s.stmt.Walk(w)
		w.IndentLevel -= 2

		return false

	case *Break:
		defaultInfoPrint(color.HiCyan, s.Line, fmt.Sprintf("Break (Statement)"))
		return false

	case *Continue:
		defaultInfoPrint(color.HiCyan, s.Line, fmt.Sprintf("Continue (Statement)"))
		return false

	case *FunctionDeclareNode:
		defaultInfoPrint(color.Green, s.Line,
			fmt.Sprintf("Declaration Function ['%s'] (Statement) [annotation %s]", s.name, s.returnAnnotation))

		w.IndentLevel++
		if len(s.args) != 0 {
			w.PrintIndent(0)
			color.HiGreen("!--> Arg Names: ")
			for _, item := range s.args {
				w.PrintIndent(1)
				color.Blue("+- %s [annotation: %s]", item.Name, item.Annotation)
			}
		}
		s.body.Walk(w)
		return false

	case *While:
		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("While-Block (Block)"))

		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Cycle Condition (Condition)"))
		s.condition.Walk(w)

		defaultInfoPrint(color.Green, s.Line, fmt.Sprintf("Cycle Body (Block)"))
		s.stmt.Walk(w)

		return false

	case *Return:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Return (Statement)"))

		w.IndentLevel++
		s.value.Walk(w)
		w.IndentLevel--
		return false

	case *Nil:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Nil Value"))
		return false

	case *Array:
		defaultInfoPrint(color.Magenta, s.Line,
			fmt.Sprintf("Array [el annotation: %s]", s.elementsTypeAnnotation))
		w.IndentLevel++
		for _, item := range s.Elements {
			item.Walk(w)
		}
		w.IndentLevel--
		return false

	case *CollectionAccess:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Array Access to '%s' array", s.variableName))

		w.IndentLevel++
		w.PrintIndent(0)
		color.HiGreen("+- Index")
		w.IndentLevel++
		s.index.Walk(w)
		w.IndentLevel -= 2

		return false

	case *Class:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Class Declaring Declare %s", s.className))

		w.IndentLevel++

		for _, item := range s.fields {
			w.PrintIndent(0)
			color.HiGreen("+- Field: %s [annotation: %s]", item.Name, item.Annotation)
		}

		for _, item := range s.methods {
			item.Walk(w)
		}

		return false

	case *ClassAccessRHS:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Class RHS Access to '%s' [field: %s]", s.structName, s.structField))
		return false

	case *ClassAccessLHS:
		defaultInfoPrint(color.Magenta, s.Line, fmt.Sprintf("Class LHS Access to '%s' [field: %s]", s.structName, s.structField))

		w.IndentLevel++
		s.stmt.Walk(w)
		return false

	case *ClassScopeMethodAccess:
		defaultInfoPrint(color.Magenta, s.Line,
			fmt.Sprintf("Var '%s' Executing Method '%s'", s.objName, s.methodToExecute))

		w.IndentLevel++
		for _, item := range s.arguments {
			item.Walk(w)
		}

		return false
	}

	return true
}

func (w Printer) LeaveNode(n Node) {
	w.IndentLevel--
}
