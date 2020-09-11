package parser

import (
	"errors"
	"ingoly/utils/errpull"
	"math"
)

func BoolTernary(statement bool, a, b bool) bool {
	if statement {
		return a
	}
	return b
}

const EPS = 1e-13

type Executor struct {
	ctx        *BlockContext
	stack      *Stack
	ErrorsPull *errpull.ErrorsPull
}

func NewExecutor() Executor {
	return Executor{NewBlockContext(), NewStack(), errpull.NewErrorsPull()}
}

func (w *Executor) CreatePullError(err error, line int) {
	inn := errpull.NewInnerError(err, line)
	w.ErrorsPull.Errors = append(w.ErrorsPull.Errors, inn)
}

func (w Executor) EnterNode(n Node) bool {

	switch s := n.(type) {
	case *BinaryNode:
		s.op1.Walk(w)
		op1, ok := w.stack.Pop()

		if !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
		}

		s.op2.Walk(w)
		op2, ok := w.stack.Pop()

		if !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
		}

		switch __T1 := op1.(type) {
		case *IntNumber:
			switch __T2 := op2.(type) {
			case *IntNumber:
				switch s.operation {
				case "+":
					w.stack.Push(&IntNumber{__T1.value + __T2.value, s.Line})
				case "-":
					w.stack.Push(&IntNumber{__T1.value - __T2.value, s.Line})
				case "*":
					w.stack.Push(&IntNumber{__T1.value * __T2.value, s.Line})
				case "/":
					if __T2.value == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T1Casted := float64(__T1.value)
						__T2Casted := float64(__T2.value)
						w.stack.Push(&FloatNumber{__T1Casted / __T2Casted, s.Line})
					}
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					__T1Casted := float64(__T1.value)
					w.stack.Push(&FloatNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					__T1Casted := float64(__T1.value)
					w.stack.Push(&FloatNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					__T1Casted := float64(__T1.value)
					w.stack.Push(&FloatNumber{__T1Casted * __T2.value, s.Line})
				case "/":
					if math.Abs(__T2.value) < EPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T1Casted := float64(__T1.value)
						w.stack.Push(&FloatNumber{__T1Casted / __T2.value, s.Line})
					}
				}
			case *Boolean:
				switch s.operation {
				case "+":
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1.value * __T2Casted, s.Line})
				case "/":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					__T1Casted := float64(__T1.value)
					if __T2Casted < EPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.stack.Push(&FloatNumber{__T1Casted / __T2Casted, s.Line})
					}
				}
			case *String:
				switch s.operation {
				case "*":
					__Result := ""
					for i := 0; i < __T1.value; i++ {
						__Result += __T2.value
					}
					w.stack.Push(&String{__Result, s.Line})
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "/":
					err := errors.New("invalid binary operation between int and string")
					w.CreatePullError(err, s.Line)
				}
			}
		case *FloatNumber:
			switch __T2 := op2.(type) {
			case *IntNumber:
				switch s.operation {
				case "+":
					__T2Casted := float64(__T2.value)
					w.stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					__T2Casted := float64(__T2.value)
					w.stack.Push(&FloatNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					__T2Casted := float64(__T2.value)
					w.stack.Push(&FloatNumber{__T1.value * __T2Casted, s.Line})
				case "/":
					if __T2.value == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T2Casted := float64(__T2.value)
						w.stack.Push(&FloatNumber{__T1.value / __T2Casted, s.Line})
					}
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					w.stack.Push(&FloatNumber{__T1.value + __T2.value, s.Line})
				case "-":
					w.stack.Push(&FloatNumber{__T1.value - __T2.value, s.Line})
				case "*":
					w.stack.Push(&FloatNumber{__T1.value * __T2.value, s.Line})
				case "/":
					if math.Abs(__T2.value) < EPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.stack.Push(&FloatNumber{__T1.value * __T2.value, s.Line})
					}
				}
			case *Boolean:
				switch s.operation {
				case "+":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "/":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					if math.Abs(__T2Casted) < EPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.stack.Push(&FloatNumber{__T1.value / __T2Casted, s.Line})
					}
				}
			case *String:
				switch s.operation {
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "*":
					fallthrough
				case "/":
					err := errors.New("invalid binary operation between float and string")
					w.CreatePullError(err, s.Line)
				}
			}
		case *Boolean:
			switch __T2 := op2.(type) {
			case *IntNumber:
				switch s.operation {
				case "+":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted * __T2.value, s.Line})
				case "/":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					if __T2.value == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.stack.Push(&IntNumber{__T1Casted + __T2.value, s.Line})
					}
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}
					w.stack.Push(&FloatNumber{__T1Casted * __T2.value, s.Line})
				case "/":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}

					if math.Abs(__T2.value) < EPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.stack.Push(&FloatNumber{__T1Casted / __T2.value, s.Line})
					}
				}
			case *Boolean:
				switch s.operation {
				case "+":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted + __T2Casted, s.Line})
				case "-":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted - __T2Casted, s.Line})
				case "*":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.stack.Push(&IntNumber{__T1Casted * __T2Casted, s.Line})
				case "/":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					if __T2Casted == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					}
					w.stack.Push(&IntNumber{__T1Casted / __T2Casted, s.Line})
				}
			case *String:
				switch s.operation {
				case "*":
					fallthrough
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "/":
					err := errors.New("division by zero")
					w.CreatePullError(err, s.Line)
				}
			}
		case *String:
			switch __T2 := op2.(type) {
			case *IntNumber:
				switch s.operation {
				case "*":
					__Result := ""
					for i := 0; i < __T2.value; i++ {
						__Result += __T1.value
					}
					w.stack.Push(&String{__Result, s.Line})
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "/":
					err := errors.New("invalid operation between string and integer")
					w.CreatePullError(err, s.Line)
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "*":
					fallthrough
				case "/":
					err := errors.New("invalid operation between string and float")
					w.CreatePullError(err, s.Line)
				}
			case *Boolean:
				switch s.operation {
				case "+":
					fallthrough
				case "-":
					fallthrough
				case "*":
					fallthrough
				case "/":
					err := errors.New("invalid operation between string and boolean")
					w.CreatePullError(err, s.Line)
				}
			case *String:
				switch s.operation {
				case "+":
					w.stack.Push(&String{__T1.value + __T2.value, s.Line})
				case "-":
					fallthrough
				case "*":
					fallthrough
				case "/":
					err := errors.New("invalid operation between string and string")
					w.CreatePullError(err, s.Line)
				}
			}
		}

		return false

	case *DeclarationNode:
		s.Expression.Walk(w)
		w.ctx.Vars[s.Variable], _ = w.stack.Pop()
		return false

	case *UnaryNode:
		s.Walk(w)
		op1, ok := w.stack.Pop()

		if !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
		}

		switch op := op1.(type) {
		case *IntNumber:
			switch s.operation {
			case "-":
				w.stack.Push(&IntNumber{-op.value, s.Line})
			case "+":
				w.stack.Push(&IntNumber{op.value, s.Line})
			}
		case *FloatNumber:
			switch s.operation {
			case "-":
				w.stack.Push(&FloatNumber{-op.value, s.Line})
			case "+":
				w.stack.Push(&FloatNumber{op.value, s.Line})
			}
		case *Boolean:
			err := errors.New("invalid unary operation for boolean")
			w.CreatePullError(err, s.Line)
		case *String:
			err := errors.New("invalid unary operation for string")
			w.CreatePullError(err, s.Line)
		}

		return false

	case *UsingVariableNode:
		w.stack.Push(w.ctx.Vars[s.name])
		return false

	case *IntNumber:
		w.stack.Push(s)
		return false

	case *FloatNumber:
		w.stack.Push(s)
		return false

	case *Boolean:
		w.stack.Push(s)
		return false

	case *String:
		w.stack.Push(s)
		return false

	case *ConditionalNode:

		return false

	case *PrintNode:

		return false

	case *IfNode:

		return false

	case *ForNode:

		return false
	}

	return true
}

func (w Executor) LeaveNode(n Node) {

}
