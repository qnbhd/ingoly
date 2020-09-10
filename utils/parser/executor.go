package parser

import (
	"errors"
	"fmt"
)

type Executor struct {
	ctx   *BlockContext
	stack *Stack
}

func NewExecutor() Executor {
	return Executor{NewBlockContext(), NewStack()}
}

func (w *Executor) GetFromContext(name string) (Value, error) {
	res, ok := w.ctx.Vars[name]
	if ok {
		return res, nil
	}
	return NumberValue{0}, errors.New("no needed variable in ctx")
}

func (w Executor) EnterNode(n Node) bool {

	switch s := n.(type) {
	case *BinaryNode:
		s.op1.Walk(w)
		result1, _ := w.stack.Pop()
		s.op2.Walk(w)
		result2, _ := w.stack.Pop()

		switch v := result1.(type) {
		case StringValue:
			string1 := v.AsString()
			if s.operation == "+" {
				w.stack.Push(StringValue{string1 + result2.AsString()})
			}

		case NumberValue:
			number1 := result1.AsNumber()
			number2 := result2.AsNumber()

			switch s.operation {
			case "+":
				w.stack.Push(NumberValue{number1 + number2})
			case "-":
				w.stack.Push(NumberValue{number1 - number2})
			case "*":
				w.stack.Push(NumberValue{number1 * number2})
			case "/":
				if number2 != 0 {
					w.stack.Push(NumberValue{number1 / number2})
				}
			}
		}

		return false

	case *DeclarationNode:
		s.Expression.Walk(w)
		w.ctx.Vars[s.Variable], _ = w.stack.Pop()
		return false

	case *UnaryNode:
		s.op1.Walk(w)
		top, _ := w.stack.Pop()
		switch s.operation {
		case "-":
			inst := NumberValue{top.AsNumber()}
			w.stack.Push(inst)
		case "+":
			w.stack.Push(top)
		}

		return false

	case *UsingVariableNode:
		w.stack.Push(w.ctx.Vars[s.name])
		return false

	case *ValueNode:
		w.stack.Push(s.value)
		return false

	case *ConditionalNode:
		s.op1.Walk(w)
		result1, _ := w.stack.Pop()
		s.op2.Walk(w)
		result2, _ := w.stack.Pop()

		switch v := result1.(type) {
		case StringValue:
			string1 := v.AsString()
			if s.operation == "==" {
				res := 0.
				if string1 == result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			} else if s.operation == ">" {
				res := 0.
				if string1 > result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			} else if s.operation == "<" {
				res := 0.
				if string1 < result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			} else if s.operation == "<=" {
				res := 0.
				if string1 <= result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			} else if s.operation == ">=" {
				res := 0.
				if string1 >= result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			} else if s.operation == "!=" {
				res := 0.
				if string1 != result2.AsString() {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			}
		case NumberValue:
			number1 := result1.AsNumber()
			number2 := result2.AsNumber()

			switch s.operation {
			case "=":
				res := 0.
				if number1 == number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			case "<":
				res := 0.
				if number1 < number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			case ">":
				res := 0.
				if number1 > number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			case ">=":
				res := 0.
				if number1 >= number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			case "<=":
				res := 0.
				if number1 <= number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			case "!=":
				res := 0.
				if number1 != number2 {
					res = 1
				}
				w.stack.Push(NumberValue{res})
			}
		}
		return false

	case *PrintNode:
		s.node.Walk(w)
		res, _ := w.stack.Pop()
		fmt.Println(res.AsString())
		return false

	case *IfNode:
		s.node.Walk(w)
		condition, _ := w.stack.Pop()

		if condition.AsNumber() != 0 {
			s.ifStmt.Walk(w)
		} else if s.elseStmt != nil {
			s.elseStmt.Walk(w)
		}

		return false
	}

	return true
}

func (w Executor) LeaveNode(n Node) {

}
