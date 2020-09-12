package parser

import "fmt"

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
