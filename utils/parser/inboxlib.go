package parser

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func __InBoxPrint(w Executor, curNode Node, argCount, line int) {

	var args []Node
	for !w.stack.IsEmpty() {
		res, _ := w.stack.Pop()
		args = append(args, res)
	}

	reverseAny(args)

	for _, arg := range args {
		switch exp := arg.(type) {
		case *IntNumber:
			fmt.Print(exp.value, " ")
		case *FloatNumber:
			fmt.Print(exp.value, " ")
		case *Boolean:
			fmt.Print(exp.value, " ")
		case *String:
			fmt.Print(exp.value, " ")
		case *Nil:
			fmt.Print("nil", " ")
		}

	}
}

func __InBoxPrintln(w Executor, curNode Node, argCount, line int) {

	var args []Node
	for !w.stack.IsEmpty() {
		res, _ := w.stack.Pop()
		args = append(args, res)
	}

	reverseAny(args)

	for _, arg := range args {
		switch exp := arg.(type) {
		case *IntNumber:
			fmt.Print(exp.value, " ")
		case *FloatNumber:
			fmt.Print(exp.value, " ")
		case *Boolean:
			fmt.Print(exp.value, " ")
		case *String:
			fmt.Print(exp.value, " ")
		case *Nil:
			fmt.Print("nil", " ")
		}
	}

	fmt.Println()
}

func __InBoxType(w Executor, curNode Node, line int) {

	arg, _ := w.stack.Pop()

	var result string
	switch arg.(type) {
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

//
//func __InBoxMathFunc(w Executor, curNode, opNode Node, line int, funcName string) {
//
//	var target float64
//	switch op := opNode.(type) {
//	case *IntNumber:
//		target = float64(op.value)
//	case *FloatNumber:
//		target = op.value
//	case *Boolean:
//		err := errors.New("invalid argument for func")
//		w.CreatePullError(err, line)
//		return
//	case *String:
//		err := errors.New("invalid argument for func")
//		w.CreatePullError(err, line)
//		return
//	}
//
//	var functor func(float64) float64
//	switch funcName {
//	case "sin":
//		functor = math.Sin
//	case "cos":
//		functor = math.Cos
//	case "tan":
//		functor = math.Tan
//	case "sqrt":
//		functor = math.Sqrt
//	default:
//		functor = __DummyF64
//	}
//
//	w.stack.Push(&FloatNumber{functor(target), line})
//}

func __InBoxSin(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.stack.Pop()

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

	w.stack.Push(&FloatNumber{value: math.Sin(target), Line: line})
}
