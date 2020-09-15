package parser

import (
	"errors"
	"fmt"
)

func __TypeCastingInt(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.Stack.Pop()

	switch s := opNode.(type) {
	case *IntNumber:
		w.Stack.Push(&IntNumber{s.value, s.Line})
	case *FloatNumber:
		w.Stack.Push(&IntNumber{int(s.value), s.Line})
	case *Boolean:
		var result int
		switch s.value {
		case true:
			result = 1
		case false:
			result = 0
		}
		w.Stack.Push(&IntNumber{result, s.Line})
	case *String:
		err := errors.New("invalid type casting from string to int")
		w.CreatePullError(err, line)
	}

}

func __TypeCastingFloat(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.Stack.Pop()

	switch s := opNode.(type) {
	case *IntNumber:
		w.Stack.Push(&FloatNumber{float64(s.value), s.Line})
	case *FloatNumber:
		w.Stack.Push(&FloatNumber{s.value, s.Line})
	case *Boolean:
		var result float64
		switch s.value {
		case true:
			result = 1.
		case false:
			result = 0.
		}
		w.Stack.Push(&FloatNumber{result, s.Line})
	case *String:
		err := errors.New("invalid type casting from string to float")
		w.CreatePullError(err, line)
	}
}
func __TypeCastingBoolean(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.Stack.Pop()

	switch s := opNode.(type) {
	case *IntNumber:
		result := s.value != 0
		w.Stack.Push(&Boolean{result, s.Line})
	case *FloatNumber:
		result := s.value != 0
		w.Stack.Push(&Boolean{result, s.Line})
	case *Boolean:
		w.Stack.Push(&Boolean{s.value, s.Line})
	case *String:
		err := errors.New("invalid type casting from string to boolean")
		w.CreatePullError(err, line)
	}
}

func __TypeCastingString(w Executor, curNode Node, argCount, line int) {

	opNode, _ := w.Stack.Pop()

	switch s := opNode.(type) {
	case *IntNumber:
		result := fmt.Sprintf("%d", s.value)
		w.Stack.Push(&String{result, s.Line})
	case *FloatNumber:
		result := fmt.Sprintf("%f", s.value)
		w.Stack.Push(&String{result, s.Line})
	case *Boolean:
		result := fmt.Sprintf("%t", s.value)
		w.Stack.Push(&String{result, s.Line})
	case *String:
		w.Stack.Push(&String{s.value, s.Line})
	}

}
