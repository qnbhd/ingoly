package parser

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"unicode/utf8"
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
			fmt.Print(exp.value)
		case *FloatNumber:
			fmt.Print(exp.value)
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
			fmt.Print(", ")
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

	w.Stack.Push(&String{result, line})
}

func __InBoxSin(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.Stack.Pop()

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

	w.Stack.Push(&FloatNumber{value: math.Sin(target), Line: line})
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
		result = utf8.RuneCountInString(typedCollection.value)
	default:
		err := errors.New("invalid argument (must be collection) for len operator")
		w.CreatePullError(err, line)
		return
	}

	w.Stack.Push(&IntNumber{
		value: result,
		Line:  line,
	})

}
