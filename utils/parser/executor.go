package parser

//import (
//	"errors"
//	"fmt"
//	"strings"
//)
//
//func BoolTernary(statement bool, a, b bool) bool {
//	if statement {
//		return a
//	}
//	return b
//}
//
//type Executor struct {
//	ctx   *BlockContext
//	stack *Stack
//}
//
//func NewExecutor() Executor {
//	return Executor{NewBlockContext(), NewStack()}
//}
//
//func (w *Executor) GetFromContext(name string) (Node, error) {
//	res, ok := w.ctx.Vars[name]
//	if ok {
//		return res, nil
//	}
//	return NumberNode{0}, errors.New("no needed variable in ctx")
//}
//
//func (w Executor) EnterNode(n Node) bool {
//
//	switch s := n.(type) {
//	case *BinaryNode:
//		s.op1.Walk(w)
//		result1, _ := w.stack.Pop()
//		s.op2.Walk(w)
//		result2, _ := w.stack.Pop()
//
//		switch v := result1.(type) {
//		case StringNode:
//			string1 := v.AsString()
//			if s.operation == "+" {
//				w.stack.Push(StringNode{string1 + result2.AsString()})
//			}
//		case NumberNode:
//			number1 := result1.AsNumber()
//			number2 := result2.AsNumber()
//			switch s.operation {
//			case "+":
//				w.stack.Push(NumberNode{number1 + number2})
//			case "-":
//				w.stack.Push(NumberNode{number1 - number2})
//			case "*":
//				w.stack.Push(NumberNode{number1 * number2})
//			case "/":
//				if number2 != 0 {
//					w.stack.Push(NumberNode{number1 / number2})
//				}
//			}
//		}
//
//		return false
//
//	case *DeclarationNode:
//		s.Expression.Walk(w)
//		w.ctx.Vars[s.Variable], _ = w.stack.Pop()
//		return false
//
//	case *UnaryNode:
//		s.op1.Walk(w)
//		top, _ := w.stack.Pop()
//		switch s.operation {
//		case "-":
//			inst := NumberNode{-top.AsNumber()}
//			w.stack.Push(inst)
//		case "+":
//			w.stack.Push(top)
//		}
//
//		return false
//
//	case *UsingVariableNode:
//		w.stack.Push(w.ctx.Vars[s.name])
//		return false
//
//	case *NodeNode:
//		w.stack.Push(s.Node)
//		return false
//
//	case *ConditionalNode:
//		s.op1.Walk(w)
//		result1, _ := w.stack.Pop()
//		s.op2.Walk(w)
//		result2, _ := w.stack.Pop()
//
//		var number1, number2 float64
//		switch v := result1.(type) {
//		case StringNode:
//			number1 = float64(strings.Compare(v.AsString(), result2.AsString()))
//			number2 = 0
//		default:
//			number1 = result1.AsNumber()
//			number2 = result2.AsNumber()
//		}
//
//		var result bool
//		switch s.operation {
//		case "<":
//			result = number1 < number2
//		case ">":
//			result = number1 > number2
//		case "==":
//			result = number1 == number2
//		case "<=":
//			result = number1 <= number2
//		case ">=":
//			result = number1 >= number2
//		case "!=":
//			result = number1 != number2
//		case "&&":
//			result = (number1 != 0) && (number2 != 0)
//		case "||":
//			result = (number1 != 0) || (number2 != 0)
//		}
//
//		var LogicalResult float64
//		if result {
//			LogicalResult = 1.
//		} else {
//			LogicalResult = 0.
//		}
//		w.stack.Push(NumberNode{LogicalResult})
//		return false
//
//	case *PrintNode:
//		s.node.Walk(w)
//		res, ok := w.stack.Pop()
//		if (res != nil) && ok {
//			fmt.Println(res.AsString())
//		}
//
//		return false
//
//	case *IfNode:
//		s.node.Walk(w)
//		condition, _ := w.stack.Pop()
//
//		if condition.AsNumber() != 0 {
//			s.ifStmt.Walk(w)
//		} else if s.elseStmt != nil {
//			s.elseStmt.Walk(w)
//		}
//
//		return false
//	case *ForNode:
//
//		w.ctx.Vars[s.iterVar] = NumberNode{s.start.AsNumber()}
//
//		for i := s.start.AsNumber(); BoolTernary(s.strict, i <= s.stop.AsNumber(), i < s.stop.AsNumber()); i += s.step.AsNumber() {
//			s.stmt.Walk(w)
//			w.ctx.Vars[s.iterVar] = NumberNode{w.ctx.Vars[s.iterVar].AsNumber() + s.step.AsNumber()}
//		}
//
//		delete(w.ctx.Vars, s.iterVar)
//
//		return false
//	}
//
//	return true
//}
//
//func (w Executor) LeaveNode(n Node) {
//
//}
