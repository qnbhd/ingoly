package parser

import (
	"errors"
	"fmt"
	"ingoly/errpull"
	"math"
	"math/rand"
	"reflect"
	"strings"
)

func BoolTernary(statement bool, a, b bool) bool {
	if statement {
		return a
	}
	return b
}

func (w *Executor) ClearStackLastNil() {
	__last, ok := w.Stack.Pop()

	if ok {
		switch __last.(type) {
		case *Nil:
			return
		default:
			w.Stack.Push(__last)
		}
	}
}

func (w *Executor) forkNewContext() *Context {
	newCtx := NewContext()
	for name, node := range w.mainContext.Vars {
		newCtx.Vars[name] = node
	}
	newCtx.Functions = w.mainContext.Functions
	return newCtx
}

func (w *Executor) copyUpdatedVars(ctx *Context) {
	for name, _ := range w.mainContext.Vars {
		w.mainContext.Vars[name] = ctx.Vars[name]
	}
}

func (w *Executor) switchMainContext() {
	w.currentContext = w.mainContext
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

type InterruptionsPull struct {
	interruptions []string
}

func NewInterruptionsPull() *InterruptionsPull {
	return &InterruptionsPull{interruptions: []string{}}
}

const (
	__ArithmeticsEPS = 1e-13
)

type Executor struct {
	currentContext    *Context
	mainContext       *Context
	Stack             *Stack
	ErrorsPull        *errpull.ErrorsPull
	lastStructLabel   string
	interruptionsPull *InterruptionsPull
}

func NewExecutor(ctx *Context) Executor {

	return Executor{ctx,
		ctx,
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
		op1, ok := w.Stack.Pop()

		if op1 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.op2.Walk(w)
		op2, ok := w.Stack.Pop()

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
					w.Stack.Push(&IntNumber{__T1.value + __T2.value, s.Line})
				case "-":
					w.Stack.Push(&IntNumber{__T1.value - __T2.value, s.Line})
				case "*":
					w.Stack.Push(&IntNumber{__T1.value * __T2.value, s.Line})
				case "/":
					if __T2.value == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T1Casted := float64(__T1.value)
						__T2Casted := float64(__T2.value)
						w.Stack.Push(&FloatNumber{__T1Casted / __T2Casted, s.Line})
					}
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					__T1Casted := float64(__T1.value)
					w.Stack.Push(&FloatNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					__T1Casted := float64(__T1.value)
					w.Stack.Push(&FloatNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					__T1Casted := float64(__T1.value)
					w.Stack.Push(&FloatNumber{__T1Casted * __T2.value, s.Line})
				case "/":
					if math.Abs(__T2.value) < __ArithmeticsEPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T1Casted := float64(__T1.value)
						w.Stack.Push(&FloatNumber{__T1Casted / __T2.value, s.Line})
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
					w.Stack.Push(&IntNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.Stack.Push(&IntNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					var __T2Casted int
					switch __T2.value {
					case true:
						__T2Casted = 1
					case false:
						__T2Casted = 0
					}
					w.Stack.Push(&IntNumber{__T1.value * __T2Casted, s.Line})
				case "/":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					__T1Casted := float64(__T1.value)
					if __T2Casted < __ArithmeticsEPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.Stack.Push(&FloatNumber{__T1Casted / __T2Casted, s.Line})
					}
				}
			case *String:
				switch s.operation {
				case "*":
					__Result := ""
					for i := 0; i < __T1.value; i++ {
						__Result += __T2.value
					}
					w.Stack.Push(&String{__Result, s.Line})
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
					w.Stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					__T2Casted := float64(__T2.value)
					w.Stack.Push(&FloatNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					__T2Casted := float64(__T2.value)
					w.Stack.Push(&FloatNumber{__T1.value * __T2Casted, s.Line})
				case "/":
					if __T2.value == 0 {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						__T2Casted := float64(__T2.value)
						w.Stack.Push(&FloatNumber{__T1.value / __T2Casted, s.Line})
					}
				}
			case *FloatNumber:
				switch s.operation {
				case "+":
					w.Stack.Push(&FloatNumber{__T1.value + __T2.value, s.Line})
				case "-":
					w.Stack.Push(&FloatNumber{__T1.value - __T2.value, s.Line})
				case "*":
					w.Stack.Push(&FloatNumber{__T1.value * __T2.value, s.Line})
				case "/":
					if math.Abs(__T2.value) < __ArithmeticsEPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.Stack.Push(&FloatNumber{__T1.value * __T2.value, s.Line})
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
					w.Stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "-":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					w.Stack.Push(&FloatNumber{__T1.value - __T2Casted, s.Line})
				case "*":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					w.Stack.Push(&FloatNumber{__T1.value + __T2Casted, s.Line})
				case "/":
					var __T2Casted float64
					switch __T2.value {
					case true:
						__T2Casted = 1.0
					case false:
						__T2Casted = 0.0
					}
					if math.Abs(__T2Casted) < __ArithmeticsEPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.Stack.Push(&FloatNumber{__T1.value / __T2Casted, s.Line})
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
					w.Stack.Push(&IntNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					w.Stack.Push(&IntNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					var __T1Casted int
					switch __T1.value {
					case true:
						__T1Casted = 1
					case false:
						__T1Casted = 0
					}
					w.Stack.Push(&IntNumber{__T1Casted * __T2.value, s.Line})
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
						w.Stack.Push(&IntNumber{__T1Casted + __T2.value, s.Line})
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
					w.Stack.Push(&FloatNumber{__T1Casted + __T2.value, s.Line})
				case "-":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}
					w.Stack.Push(&FloatNumber{__T1Casted - __T2.value, s.Line})
				case "*":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}
					w.Stack.Push(&FloatNumber{__T1Casted * __T2.value, s.Line})
				case "/":
					var __T1Casted float64
					switch __T1.value {
					case true:
						__T1Casted = 1.0
					case false:
						__T1Casted = 0.0
					}

					if math.Abs(__T2.value) < __ArithmeticsEPS {
						err := errors.New("division by zero")
						w.CreatePullError(err, s.Line)
					} else {
						w.Stack.Push(&FloatNumber{__T1Casted / __T2.value, s.Line})
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
					w.Stack.Push(&IntNumber{__T1Casted + __T2Casted, s.Line})
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
					w.Stack.Push(&IntNumber{__T1Casted - __T2Casted, s.Line})
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
					w.Stack.Push(&IntNumber{__T1Casted * __T2Casted, s.Line})
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
					w.Stack.Push(&IntNumber{__T1Casted / __T2Casted, s.Line})
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
					w.Stack.Push(&String{__Result, s.Line})
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
					w.Stack.Push(&String{__T1.value + __T2.value, s.Line})
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
		w.currentContext.Vars[s.Variable], _ = w.Stack.Pop()
		return false

	case *AssignNode:
		if _, ok := w.currentContext.Vars[s.Variable]; !ok {
			err := errors.New("assign to undeclared variable")
			w.CreatePullError(err, s.Line)
		}

		s.Expression.Walk(w)
		result, _ := w.Stack.Pop()

		switch w.currentContext.Vars[s.Variable].(type) {
		case *IntNumber:
			switch result.(type) {
			case *IntNumber:
				w.currentContext.Vars[s.Variable] = result
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
				w.currentContext.Vars[s.Variable] = result
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
				w.currentContext.Vars[s.Variable] = result
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
				w.currentContext.Vars[s.Variable] = result
			}
		}
		return false

	case *UnaryNode:
		s.Walk(w)
		op1, ok := w.Stack.Pop()

		if op1 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
		}

		switch op := op1.(type) {
		case *IntNumber:
			switch s.operation {
			case "-":
				w.Stack.Push(&IntNumber{-op.value, s.Line})
			case "+":
				w.Stack.Push(&IntNumber{op.value, s.Line})
			}
		case *FloatNumber:
			switch s.operation {
			case "-":
				w.Stack.Push(&FloatNumber{-op.value, s.Line})
			case "+":
				w.Stack.Push(&FloatNumber{op.value, s.Line})
			}
		case *Boolean:
			err := errors.New("invalid unary operation for boolean")
			w.CreatePullError(err, s.Line)
		case *String:
			err := errors.New("invalid unary operation for string")
			w.CreatePullError(err, s.Line)
		}

		return false

	case *ScopeVar:
		_, ok := w.currentContext.Vars[s.name]
		if !ok {
			err := errors.New(fmt.Sprintf("using undefined variable '%s'", s.name))
			w.CreatePullError(err, s.Line)
			return false
		}
		w.Stack.Push(w.currentContext.Vars[s.name])
		return false

	case *IntNumber:
		w.Stack.Push(s)
		return false

	case *FloatNumber:
		w.Stack.Push(s)
		return false

	case *Boolean:
		w.Stack.Push(s)
		return false

	case *String:
		w.Stack.Push(s)
		return false

	case *Array:
		w.Stack.Push(s)
		return false

	case *CollectionAccess:
		result, ok := w.currentContext.Vars[s.variableName]

		if !ok {
			err := errors.New(fmt.Sprintf("using undefined array to access '%s'", s.variableName))
			w.CreatePullError(err, s.Line)
			return false
		}

		s.index.Walk(w)
		index, ok := w.Stack.Pop()

		if !ok {
			err := errors.New(fmt.Sprintf("invalid index for array access to '%s'", s.variableName))
			w.CreatePullError(err, s.Line)
			return false
		}

		switch collection := result.(type) {
		case *Array:
			switch idx := index.(type) {
			case *IntNumber:
				if idx.value < len(collection.Elements) {
					w.Stack.Push(collection.Elements[idx.value])
				} else {
					err := errors.New(
						fmt.Sprintf("index out of range for array access to '%s' [%d] with length %d",
							s.variableName, idx.value, len(collection.Elements)))
					w.CreatePullError(err, s.Line)
					return false
				}
			}
		case *String:
			switch idx := index.(type) {
			case *IntNumber:
				stringCollection := []rune(collection.value)
				if idx.value < len(stringCollection) {
					w.Stack.Push(&String{string(stringCollection[idx.value]), s.Line})
				} else {
					err := errors.New(
						fmt.Sprintf("index out of range for string access to '%s' [%d] with length %d",
							s.variableName, idx.value, len(stringCollection)))
					w.CreatePullError(err, s.Line)
					return false
				}
			}
		}

		return false

	case *ConditionalNode:

		s.op1.Walk(w)
		op1, ok := w.Stack.Pop()

		if op1 == nil || !ok {
			err := errors.New("using var before initialization")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.op2.Walk(w)
		op2, ok := w.Stack.Pop()

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

		w.Stack.Push(&Boolean{result, s.Line})

		return false

	case *FunctionDeclareNode:

		//w.currentContext.Functions[s.name] = func(w Executor, curNode Node, argCount, line int) {
		//
		//	//ctxVariables := map[string]Node{}
		//
		//	reverseAny(s.args)
		//
		//	functionLabel := fmt.Sprintf("__function%d", rand.Int())
		//	w.lastStructLabel = functionLabel
		//
		//	functionCtx := lib.NewContext()
		//	// TODO general access
		//	functionCtx.Functions = w.mainContext.Functions
		//
		//	bijectionAnnotationTypes := map[string]string{}
		//
		//	bijectionAnnotationTypes["IntNumber"] = "int"
		//	bijectionAnnotationTypes["FloatNumber"] = "float"
		//	bijectionAnnotationTypes["Boolean"] = "boolean"
		//	bijectionAnnotationTypes["String"] = "string"
		//	bijectionAnnotationTypes["Nil"] = "nil"
		//
		//	flagCorrectAnnotations := true
		//	for _, arg := range s.args {
		//		_receivedArg, _ := w.Stack.Pop()
		//		_receivedArgType := typeof(_receivedArg)[len("*parser."):]
		//
		//		if arg.Annotation != bijectionAnnotationTypes[_receivedArgType] {
		//			err := errors.New(
		//				fmt.Sprintf("invalid type of argument '%s' (expected %s, getted %s)",
		//					arg.Name, arg.Annotation, bijectionAnnotationTypes[_receivedArgType]))
		//			w.CreatePullError(err, s.Line)
		//			flagCorrectAnnotations = false
		//		}
		//		functionCtx.Vars[arg.Name] = _receivedArg
		//	}
		//
		//	if flagCorrectAnnotations {
		//
		//		w.currentContext = functionCtx
		//
		//		s.body.Walk(w)
		//		returnValue, ok := w.Stack.Pop()
		//
		//		for idx, interrupt := range w.interruptionsPull.interruptions {
		//			if strings.Contains(interrupt, functionLabel) {
		//				w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
		//			}
		//		}
		//
		//		w.lastStructLabel = ""
		//
		//		returnAnnotation := "*parser.Nil"
		//		if ok {
		//			returnValue.Walk(w)
		//			result, _ := w.Stack.Pop()
		//			returnAnnotation = typeof(result)
		//			w.Stack.Push(result)
		//		} else {
		//			w.Stack.Push(&Nil{Line: line})
		//		}
		//
		//		_retArgType := returnAnnotation[len("*parser."):]
		//		if s.returnAnnotation != bijectionAnnotationTypes[_retArgType] {
		//			err := errors.New(
		//				fmt.Sprintf("invalid type of return value in function '%s' (expected %s, getted %s)",
		//					s.name, s.returnAnnotation, bijectionAnnotationTypes[_retArgType]))
		//			w.CreatePullError(err, s.Line)
		//		}
		//
		//		w.switchMainContext()
		//	}
		//
		//}

		return false

	case *FunctionalNode:

		for _, arg := range s.arguments {
			arg.Walk(w)
		}

		functor, ok := w.currentContext.Functions[s.operator]

		if !ok {
			err := errors.New("undeclared function")
			w.CreatePullError(err, s.Line)
			return false
		}

		functor(w, s, 1, s.Line)

		w.ClearStackLastNil()

		return false

	case *IfNode:

		s.node.Walk(w)
		condition, ok := w.Stack.Pop()

		if condition == nil || !ok {
			err := errors.New("condition expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		ifBlockCtx := w.forkNewContext()
		w.currentContext = ifBlockCtx

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

		w.copyUpdatedVars(ifBlockCtx)
		w.switchMainContext()

		return false

	case *ForNode:

		s.start.Walk(w)
		startNode, ok := w.Stack.Pop()

		if startNode == nil || !ok {
			err := errors.New("loop start expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		w.currentContext.Vars[s.iterVar] = startNode

		s.stop.Walk(w)
		stopNode, ok := w.Stack.Pop()

		if stopNode == nil || !ok {
			err := errors.New("loop stop expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		s.step.Walk(w)
		stepNode, ok := w.Stack.Pop()

		if stepNode == nil || !ok {
			err := errors.New("loop step expected")
			w.CreatePullError(err, s.Line)
			return false
		}

		forBlockCtx := w.forkNewContext()
		w.currentContext = forBlockCtx

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
						w.currentContext.Vars[s.iterVar] = &IntNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &IntNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)

						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)

						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()

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
						w.currentContext.Vars[s.iterVar] = &FloatNumber{i, s.Line}
						s.stmt.Walk(w)
						res, _ := w.Stack.Pop()
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

		//delete(w.currentContext.Vars, s.iterVar)
		w.copyUpdatedVars(forBlockCtx)
		w.switchMainContext()

		return false

	case *Break:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"_break")
		w.Stack.Push(s)
		return false

	case *Continue:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"_continue")
		w.Stack.Push(s)
		return false

	case *While:
		s.condition.Walk(w)
		condition, ok := w.Stack.Pop()

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

			whileBlockCtx := w.forkNewContext()
			w.currentContext = whileBlockCtx

		Loop:
			for conditionResult {
				s.stmt.Walk(w)
				res, _ := w.Stack.Pop()

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
				condition, _ := w.Stack.Pop()

				switch cnd := condition.(type) {
				case *Boolean:
					conditionResult = cnd.value
				}
			}

			w.copyUpdatedVars(whileBlockCtx)
			w.switchMainContext()
			return false
		}

		err := errors.New("invalid condition")
		w.CreatePullError(err, s.Line)

		return false

	case *Return:
		w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions, w.lastStructLabel+"__return")
		w.Stack.Push(s.value)
		return false

	case *Nil:
		w.Stack.Push(&Nil{s.Line})
		return false
	}

	return true

}

func (w Executor) LeaveNode(n Node) {

}
