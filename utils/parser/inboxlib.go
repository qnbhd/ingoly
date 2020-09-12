package parser

import (
	"errors"
	"fmt"
	"math"
)

func __DummyF64(arg float64) float64 {
	return arg
}

func __InBoxPrint(w Executor, curNode, opNode Node, line int) {

	switch exp := opNode.(type) {
	case *IntNumber:
		fmt.Print(exp.value)
	case *FloatNumber:
		fmt.Print(exp.value)
	case *Boolean:
		fmt.Print(exp.value)
	case *String:
		fmt.Print(exp.value)
	}
}

func __InBoxPrintln(w Executor, curNode, opNode Node, line int) {

	switch exp := opNode.(type) {
	case *IntNumber:
		fmt.Println(exp.value)
	case *FloatNumber:
		fmt.Println(exp.value)
	case *Boolean:
		fmt.Println(exp.value)
	case *String:
		fmt.Println(exp.value)
	}
}

func __InBoxType(w Executor, curNode, opNode Node, line int) {

	var result string
	switch opNode.(type) {
	case *IntNumber:
		result = "int"
	case *FloatNumber:
		result = "float"
	case *Boolean:
		result = "boolean"
	case *String:
		result = "string"
	}

	w.stack.Push(&String{result, line})
}

func __InBoxMathFunc(w Executor, curNode, opNode Node, line int, funcName string) {

	var target float64
	switch op := opNode.(type) {
	case *IntNumber:
		target = float64(op.value)
	case *FloatNumber:
		target = op.value
	case *Boolean:
		err := errors.New("invalid argument for func")
		w.CreatePullError(err, line)
		return
	case *String:
		err := errors.New("invalid argument for func")
		w.CreatePullError(err, line)
		return
	}

	var functor func(float64) float64
	switch funcName {
	case "sin":
		functor = math.Sin
	case "cos":
		functor = math.Cos
	case "tan":
		functor = math.Tan
	case "sqrt":
		functor = math.Sqrt
	default:
		functor = __DummyF64
	}

	w.stack.Push(&FloatNumber{functor(target), line})
}
