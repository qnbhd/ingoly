package parser

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
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
	i := 0
	for i < argCount && !w.Stack.IsEmpty() {
		res, _ := w.Stack.Pop()
		args = append(args, res)
		i += 1
	}

	reverseAny(args)

	for idx, arg := range args {
		switch exp := arg.(type) {
		case *IntNumber:
			fmt.Print(exp.Value)
		case *FloatNumber:
			fmt.Print(exp.Value)
		case *Boolean:
			fmt.Print(exp.value)
		case *String:
			fmt.Print(exp.value)
		case *Nil:
			fmt.Print("nil")
		case *Array:
			fmt.Print("[")
			for _, item := range exp.Elements {
				item.Walk(w)
			}
			__InBoxPrint(w, exp, len(exp.Elements), line)
			fmt.Print("]")
		}
		if idx != len(args)-1 {
			fmt.Print(" ")
		}
	}
}

func __InBoxPrintln(w Executor, curNode Node, argCount, line int) {
	__InBoxPrint(w, curNode, argCount, line)
	fmt.Println(" ")
}

func __InBoxType(w Executor, curNode Node, line int) {

	arg, _ := w.Stack.Pop()

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

	w.Stack.Push(&String{[]rune(result), line})
}

func __InBoxMathFunc(w Executor, curNode Node, argCount, line int, functor func(x float64) float64) {

	opNode, _ := w.Stack.Pop()

	var target float64
	switch op := opNode.(type) {
	case *IntNumber:
		target = float64(op.Value)
	case *FloatNumber:
		target = op.Value
	case *Boolean:
		err := errors.New("invalid argument for func")
		w.CreatePullError(err, line)
		return
	case *String:
		err := errors.New("invalid argument for func")
		w.CreatePullError(err, line)
		return
	}

	w.Stack.Push(&FloatNumber{Value: functor(target), Line: line})
}

func __InBoxSin(w Executor, curNode Node, argCount, line int) {
	__InBoxMathFunc(w, curNode, argCount, line, math.Sin)
}
func __InBoxCos(w Executor, curNode Node, argCount, line int) {
	__InBoxMathFunc(w, curNode, argCount, line, math.Cos)
}
func __InBoxSqrt(w Executor, curNode Node, argCount, line int) {
	__InBoxMathFunc(w, curNode, argCount, line, math.Sqrt)
}
func __InBoxAbs(w Executor, curNode Node, argCount, line int) {
	__InBoxMathFunc(w, curNode, argCount, line, math.Abs)
}
func __InBoxExp(w Executor, curNode Node, argCount, line int) {
	__InBoxMathFunc(w, curNode, argCount, line, math.Exp)
}

func __InBoxLen(w Executor, curNode Node, argCount, line int) {

	opCollection, ok := w.Stack.Pop()

	if !ok {
		err := errors.New("argument error for len operator")
		w.CreatePullError(err, line)
		return
	}

	var result int
	switch typedCollection := opCollection.(type) {
	case *Array:
		result = len(typedCollection.Elements)
	case *String:
		result = len(typedCollection.value)
	default:
		err := errors.New("invalid argument (must be collection) for len operator")
		w.CreatePullError(err, line)
		return
	}

	w.Stack.Push(&IntNumber{
		Value: result,
		Line:  line,
	})

}

func __InBoxInput(w Executor, curNode Node, argCount, line int) {

	reader := bufio.NewReader(os.Stdin)
	gettedString, _ := reader.ReadString('\n')
	gettedString = strings.Trim(gettedString, "\n")

	w.Stack.Push(&String{
		value: []rune(gettedString),
		Line:  line,
	})

}
