package parser

import (
	"errors"
	"fmt"
	"ingoly/utils/errpull"
	"math"
	"math/rand"
	"strings"
)

func BoolTernary(statement bool, a, b bool) bool {
	if statement {
		return a
	}
	return b
}

func (w *Executor) ClearStackLastNil() {
	__last, ok := w.stack.Pop()

	if ok {
		switch __last.(type) {
		case *Nil:
			return
		default:
			w.stack.Push(__last)
		}
	}
}

type InterruptionsPull struct {
	interruptions []string
}

func NewInterruptionsPull() *InterruptionsPull {
	return &InterruptionsPull{interruptions: []string{}}
}

const EPS = 1e-13

type Executor struct {
	ctx               *BlockContext
	stack             *Stack
	ErrorsPull        *errpull.ErrorsPull
	lastStructLabel   string
	interruptionsPull *InterruptionsPull
}

func NewExecutor() Executor {
	ctx := NewBlockContext()
	ctx.Functions["sin"] = __InBoxSin
	ctx.Functions["print"] = __InBoxPrint
	ctx.Functions["println"] = __InBoxPrintln
	ctx.Functions["int"] = __TypeCastingInt
	ctx.Functions["float"] = __TypeCastingFloat
	ctx.Functions["boolean"] = __TypeCastingBoolean
	ctx.Functions["string"] = __TypeCastingString

	return Executor{ctx,
		NewStack(),
		errpull.NewErrorsPull(),
		"",
		NewInterruptionsPull()}
}

func (w *Executor) CreatePullError(err error, line int) {
	inn := errpull.NewInnerError(err, line)
	w.ErrorsPull.Errors = append(w.ErrorsPull.Errors, inn)
}

func (w Executor) EnterNode(n Node) bool {

	if len(w.interruptionsPull.interruptions) != 0 {
		return false
	}

	switch s := n.(type) {
	case *BinaryNode:
		s.op1.Walk(w)
		op1, ok := w.stack.Pop()

		if op1 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.op2.Walk(w)
		op2, ok := w.stack.Pop()

		if op2 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
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

	case *AssignNode:
		if _, ok := w.ctx.Vars[s.Variable]; !ok {
			err := errors.New("assign to undeclared variable")
			w.CreatePullError(err, s.Line)
		}

		s.Expression.Walk(w)
		result, _ := w.stack.Pop()

		switch w.ctx.Vars[s.Variable].(type) {
		case *IntNumber:
			switch result.(type) {
			case *IntNumber:
				w.ctx.Vars[s.Variable] = result
			case *FloatNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *Boolean:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *String:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			}
		case *FloatNumber:
			switch result.(type) {
			case *IntNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *FloatNumber:
				w.ctx.Vars[s.Variable] = result
			case *Boolean:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *String:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			}
		case *Boolean:
			switch result.(type) {
			case *IntNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *FloatNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *Boolean:
				w.ctx.Vars[s.Variable] = result
			case *String:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			}
		case *String:
			switch result.(type) {
			case *IntNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *FloatNumber:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *Boolean:
				err := errors.New("invalid new type for assign to variable")
				w.CreatePullError(err, s.Line)
			case *String:
				w.ctx.Vars[s.Variable] = result
			}
		}
		return false

	case *UnaryNode:
		s.Walk(w)
		op1, ok := w.stack.Pop()

		if op1 == nil || !ok {
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

		s.op1.Walk(w)
		op1, ok := w.stack.Pop()

		if op1 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.op2.Walk(w)
		op2, ok := w.stack.Pop()

		if op2 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
		}

		var __CmpOp1, __CmpOp2 float64

		switch __T1 := op1.(type) {
		case *IntNumber:
			switch __T2 := op2.(type) {
			case *IntNumber:
				__CmpOp1 = float64(__T1.value)
				__CmpOp2 = float64(__T2.value)
			case *FloatNumber:
				__CmpOp1 = float64(__T1.value)
				__CmpOp2 = __T2.value
			case *Boolean:
				__CmpOp1 = float64(__T1.value)
				switch __T2.value {
				case true:
					__CmpOp2 = 1.0
				case false:
					__CmpOp2 = 0.0
				}
			case *String:
				err := errors.New("invalid condition operation between int and string")
				w.CreatePullError(err, s.Line)
			}
		case *FloatNumber:
			switch __T2 := op2.(type) {
			case *IntNumber:
				__CmpOp1 = __T1.value
				__CmpOp2 = float64(__T2.value)
			case *FloatNumber:
				__CmpOp1 = __T1.value
				__CmpOp2 = __T2.value
			case *Boolean:
				__CmpOp1 = __T1.value
				switch __T2.value {
				case true:
					__CmpOp2 = 1.0
				case false:
					__CmpOp2 = 0.0
				}
			case *String:
				err := errors.New("invalid condition operation between int and string")
				w.CreatePullError(err, s.Line)
			}
		case *Boolean:
			switch __T2 := op2.(type) {
			case *IntNumber:
				switch __T1.value {
				case true:
					__CmpOp1 = 1.0
				case false:
					__CmpOp1 = 0.0
				}
				__CmpOp2 = float64(__T2.value)
			case *FloatNumber:
				switch __T1.value {
				case true:
					__CmpOp1 = 1.0
				case false:
					__CmpOp1 = 0.0
				}
				__CmpOp2 = __T2.value
			case *Boolean:
				switch __T1.value {
				case true:
					__CmpOp1 = 1.0
				case false:
					__CmpOp1 = 0.0
				}
				switch __T2.value {
				case true:
					__CmpOp2 = 1.0
				case false:
					__CmpOp2 = 0.0
				}
			case *String:
				err := errors.New("invalid condition operation between int and string")
				w.CreatePullError(err, s.Line)
			}

		case *String:
			switch __T2 := op2.(type) {
			case *String:
				switch __T1.value == __T2.value {
				case true:
					__CmpOp1 = 1.0
					__CmpOp2 = 1.0
				case false:
					__CmpOp1 = 0.0
					__CmpOp2 = 0.0
				}

			case *IntNumber:
				err := errors.New("invalid condition operation between string and int")
				w.CreatePullError(err, s.Line)
			case *FloatNumber:
				err := errors.New("invalid condition operation between string and float")
				w.CreatePullError(err, s.Line)
			case *Boolean:
				err := errors.New("invalid condition operation between string and boolean")
				w.CreatePullError(err, s.Line)
			}
		}

		var result bool
		switch s.operation {
		case "==":
			result = __CmpOp1 == __CmpOp2
		case "!=":
			result = __CmpOp1 != __CmpOp2
		case "<":
			result = __CmpOp1 < __CmpOp2
		case ">":
			result = __CmpOp1 > __CmpOp2
		case "<=":
			result = __CmpOp1 <= __CmpOp2
		case ">=":
			result = __CmpOp1 >= __CmpOp2
		case "&&":
			result = (__CmpOp1 != 0) && (__CmpOp2 != 0)
		case "||":
			result = (__CmpOp1 != 0) && (__CmpOp2 != 0)
		}

		w.stack.Push(&Boolean{result, s.Line})

		return false

	case *FunctionDeclareNode:

		w.ctx.Functions[s.name] = func(w Executor, curNode Node, argCount, line int) {

			//ctxVariables := map[string]Node{}

			reverseAny(s.argNames)

			for _, arg := range s.argNames {
				_receivedArg, _ := w.stack.Pop()
				w.ctx.Vars[arg] = _receivedArg
			}

			functionLabel := fmt.Sprintf("__function%d", rand.Int())
			w.lastStructLabel = functionLabel

			s.body.Walk(w)
			returnValue, ok := w.stack.Pop()

			for idx, interrupt := range w.interruptionsPull.interruptions {
				if strings.Contains(interrupt, functionLabel) {
					w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
				}
			}

			w.lastStructLabel = ""

			if ok {
				returnValue.Walk(w)
				result, _ := w.stack.Pop()
				w.stack.Push(result)
			} else {
				w.stack.Push(&Nil{Line: line})
			}

		}

		return false

	case *FunctionalNode:

		for _, arg := range s.arguments {
			arg.Walk(w)
		}

		functor, ok := w.ctx.Functions[s.operator]

		if !ok {
			err := errors.New("undeclared function")
			w.CreatePullError(err, s.Line)
			return false
		}

		functor(w, s, 1, s.Line)

		return false

	case *IfNode:

		s.node.Walk(w)
		condition, ok := w.stack.Pop()

		if condition == nil || !ok {
			err := errors.New("condition expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		switch result := condition.(type) {
		case *Boolean:
			conditionResult := result.value
			if conditionResult {
				s.ifStmt.Walk(w)
			} else if s.elseStmt != nil {
				s.elseStmt.Walk(w)
			}
			return false
		}

		err := errors.New("invalid condition")
		w.CreatePullError(err, s.Line)

		return false

	case *ForNode:

		s.start.Walk(w)
		startNode, ok := w.stack.Pop()

		if startNode == nil || !ok {
			err := errors.New("loop start expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		w.ctx.Vars[s.iterVar] = startNode

		s.stop.Walk(w)
		stopNode, ok := w.stack.Pop()

		if stopNode == nil || !ok {
			err := errors.New("loop stop expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.step.Walk(w)
		stepNode, ok := w.stack.Pop()

		if stepNode == nil || !ok {
			err := errors.New("loop step expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		switch st := startNode.(type) {
		case *IntNumber:
			start := st.value
			switch st := stopNode.(type) {
			case *IntNumber:
				stop := st.value
				switch st := stepNode.(type) {
				case *IntNumber:
					step := st.value
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopIII:
					for i := start; BoolTernary(s.strict, i <= stop, i < stop); i += step {
						w.ctx.Vars[s.iterVar] = &IntNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopIII
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopIII
						}

					}
				case *FloatNumber:
					step := st.value
					stepCasted := int(step)
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopIIF:
					for i := start; BoolTernary(s.strict, i <= stop, i < stop); i += stepCasted {
						w.ctx.Vars[s.iterVar] = &IntNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopIIF
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopIIF
						}
					}
				default:
					err := errors.New("invalid loop step")
					w.CreatePullError(err, s.Line)
				}

			case *FloatNumber:
				stop := st.value
				switch st := stepNode.(type) {
				case *IntNumber:
					step := st.value
					startCasted := float64(start)
					stepCasted := float64(step)
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopIFI:
					for i := startCasted; BoolTernary(s.strict, i <= stop, i < stop); i += stepCasted {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)

						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopIFI
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopIFI
						}
					}
				case *FloatNumber:
					step := st.value
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopIFF:
					for i := float64(start); BoolTernary(s.strict, i <= stop, i < stop); i += step {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)

						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopIFF
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopIFF
						}
					}
				default:
					err := errors.New("invalid loop step")
					w.CreatePullError(err, s.Line)
				}

			default:
				err := errors.New("invalid loop stop")
				w.CreatePullError(err, s.Line)
			}
		case *FloatNumber:
			start := st.value
			switch st := stopNode.(type) {
			case *IntNumber:
				stop := st.value
				switch st := stepNode.(type) {
				case *IntNumber:
					step := st.value
					stopCasted := float64(stop)
					stepCasted := float64(step)
					// labeling
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopFII:
					for i := start; BoolTernary(s.strict, i <= stopCasted, i < stopCasted); i += stepCasted {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopFII
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopFII
						}
					}
				case *FloatNumber:
					step := st.value
					stopCasted := float64(stop)
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopFIF:
					for i := start; BoolTernary(s.strict, i <= stopCasted, i < stopCasted); i += step {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopFIF
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopFIF
						}
					}
				default:
					err := errors.New("invalid loop step")
					w.CreatePullError(err, s.Line)
				}

			case *FloatNumber:
				stop := st.value
				switch st := stepNode.(type) {
				case *IntNumber:
					step := st.value
					stepCasted := float64(step)
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopFFI:
					for i := start; BoolTernary(s.strict, i <= stop, i < stop); i += stepCasted {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()

						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopFFI
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopFFI
						}
					}
				case *FloatNumber:
					step := st.value
					loopLabel := fmt.Sprintf("__label%d", rand.Int())
					w.lastStructLabel = loopLabel
				LoopFFF:

					for i := start; BoolTernary(s.strict, i <= stop, i < stop); i += step {
						w.ctx.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.stack.Pop()
						switch res.(type) {
						case *Break:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							break LoopFFF
						case *Continue:
							for idx, interrupt := range w.interruptionsPull.interruptions {
								if strings.Contains(interrupt, loopLabel) {
									w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
									break
								}
							}
							w.lastStructLabel = ""
							continue LoopFFF
						}
					}
				default:
					err := errors.New("invalid loop step")
					w.CreatePullError(err, s.Line)
				}

			default:
				err := errors.New("invalid loop stop")
				w.CreatePullError(err, s.Line)
			}
		default:

			err := errors.New("invalid loop start")
			w.CreatePullError(err, s.Line)
		}

		delete(w.ctx.Vars, s.iterVar)

		return false

	case *Break:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"_break")
		w.stack.Push(s)
		return false

	case *Continue:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"_continue")
		w.stack.Push(s)
		return false

	case *While:
		s.condition.Walk(w)
		condition, ok := w.stack.Pop()

		if condition == nil || !ok {
			err := errors.New("condition expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		switch result := condition.(type) {
		case *Boolean:
			conditionResult := result.value
			loopLabel := fmt.Sprintf("__label%d", rand.Int())
			w.lastStructLabel = loopLabel
		Loop:
			for conditionResult {
				s.stmt.Walk(w)
				res, _ := w.stack.Pop()

				switch res.(type) {
				case *Break:
					for idx, interrupt := range w.interruptionsPull.interruptions {
						if strings.Contains(interrupt, loopLabel) {
							w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
							break
						}
					}
					w.lastStructLabel = ""
					break Loop
				case *Continue:
					for idx, interrupt := range w.interruptionsPull.interruptions {
						if strings.Contains(interrupt, loopLabel) {
							w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
							break
						}
					}
					w.lastStructLabel = ""
					continue Loop
				}

				s.condition.Walk(w)
				condition, _ := w.stack.Pop()

				switch cnd := condition.(type) {
				case *Boolean:
					conditionResult = cnd.value
				}
			}
			return false
		}

		err := errors.New("invalid condition")
		w.CreatePullError(err, s.Line)

		return false

	case *Return:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"__return")
		w.stack.Push(s.value)
		return false

	case *Nil:
		w.stack.Push(&Nil{s.Line})
		return false
	}

	// clearing nil

	switch n.(type) {
	case *DeclarationNode:
	case *AssignNode:
	case *FunctionalNode:
	//case *ForNode:
	//case *While:
	//case *IfNode:
	default:
		return true
	}

	w.ClearStackLastNil()

	return true

}

func (w Executor) LeaveNode(n Node) {

}
