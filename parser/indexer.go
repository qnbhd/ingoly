package parser

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Indexer struct {
	Ctx *Context
}

func NewIndexer() *Indexer {
	newCtx := NewContext()
	newCtx.Functions["__builtin__sin"] = __InBoxSin
	newCtx.Functions["__builtin__cos"] = __InBoxCos
	newCtx.Functions["__builtin__sqrt"] = __InBoxSqrt
	newCtx.Functions["__builtin__abs"] = __InBoxAbs
	newCtx.Functions["__builtin__exp"] = __InBoxExp

	newCtx.Functions["print"] = __InBoxPrint
	newCtx.Functions["input"] = __InBoxInput
	newCtx.Functions["println"] = __InBoxPrintln
	newCtx.Functions["int"] = __TypeCastingInt
	newCtx.Functions["float"] = __TypeCastingFloat
	newCtx.Functions["boolean"] = __TypeCastingBoolean
	newCtx.Functions["string"] = __TypeCastingString
	newCtx.Functions["len"] = __InBoxLen

	return &Indexer{Ctx: newCtx}
}

func (w Indexer) EnterNode(n Node) bool {

	switch curNode := n.(type) {
	case *FunctionDeclareNode:
		w.Ctx.Functions[curNode.name] = func(w Executor, opNode Node, argCount, line int) {

			//ctxVariables := map[string]Node{}

			reverseAny(curNode.args)

			functionLabel := fmt.Sprintf("__function%d", rand.Int())
			w.lastStructLabel = functionLabel

			functionCtx := NewContext()
			// TODO general access
			functionCtx.Functions = w.mainContext.Functions

			bijectionAnnotationTypes := map[string]string{}

			bijectionAnnotationTypes["IntNumber"] = "int"
			bijectionAnnotationTypes["FloatNumber"] = "float"
			bijectionAnnotationTypes["Boolean"] = "boolean"
			bijectionAnnotationTypes["String"] = "string"
			bijectionAnnotationTypes["Nil"] = "nil"

			flagCorrectAnnotations := true
			for _, arg := range curNode.args {
				_receivedArg, _ := w.Stack.Pop()
				_receivedArgType := typeof(_receivedArg)[len("*parser."):]

				needed := arg.Annotation
				getted := bijectionAnnotationTypes[_receivedArgType]

				if needed != getted {
					if !(needed == "int" && getted == "float" || needed == "float" && getted == "int") {
						err := errors.New(
							fmt.Sprintf("invalid type of argument '%s' (expected %s, getted %s)",
								arg.Name, needed, getted))
						w.CreatePullError(err, curNode.Line)
						flagCorrectAnnotations = false
					}
				}
				functionCtx.Vars[arg.Name] = _receivedArg
			}

			if flagCorrectAnnotations {

				w.currentContext = functionCtx

				curNode.body.Walk(w)
				returnValue, ok := w.Stack.Pop()

				for idx, interrupt := range w.interruptionsPull.interruptions {
					if strings.Contains(interrupt, functionLabel) {
						w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
					}
				}

				w.lastStructLabel = ""

				returnAnnotation := "*parser.Nil"
				if ok {
					returnValue.Walk(w)
					result, _ := w.Stack.Pop()
					returnAnnotation = typeof(result)
					w.Stack.Push(result)
				} else {
					w.Stack.Push(&Nil{Line: line})
				}

				_retArgType := returnAnnotation[len("*parser."):]
				if curNode.returnAnnotation != bijectionAnnotationTypes[_retArgType] {
					err := errors.New(
						fmt.Sprintf("invalid type of return value in function '%s' (expected %s, getted %s)",
							curNode.name, curNode.returnAnnotation, bijectionAnnotationTypes[_retArgType]))
					w.CreatePullError(err, curNode.Line)
				}

				w.switchMainContext()
			}

		}
	}

	return true
}

func (w Indexer) LeaveNode(n Node) {

}
