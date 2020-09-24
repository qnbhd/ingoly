package parser

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Indexer struct {
	Ctx          *Context
	currentClass string
}

var (
	bijectionAnnotationTypes = map[string]string{
		"IntNumber":   "int",
		"FloatNumber": "float",
		"Boolean":     "boolean",
		"String":      "string",
		"Nil":         "nil",
	}
)

func NewIndexer() *Indexer {
	newCtx := NewContext()

	newCtx.Functions["print"] = __InBoxPrint
	newCtx.Functions["input"] = __InBoxInput
	newCtx.Functions["println"] = __InBoxPrintln
	newCtx.Functions["int"] = __TypeCastingInt
	newCtx.Functions["float"] = __TypeCastingFloat
	newCtx.Functions["boolean"] = __TypeCastingBoolean
	newCtx.Functions["string"] = __TypeCastingString
	newCtx.Functions["len"] = __InBoxLen

	return &Indexer{Ctx: newCtx, currentClass: ""}
}

func (w Indexer) EnterNode(n Node) bool {

	switch curNode := n.(type) {
	case *FunctionDeclareNode:
		functor := func(w Executor, opNode Node, argCount, line int) {

			//ctxVariables := map[string]Node{}

			reverseAny(curNode.args)

			functionLabel := fmt.Sprintf("__function%d", rand.Int())
			w.lastStructLabel = functionLabel

			functionCtx := NewContext()
			// TODO general access
			functionCtx.Functions = w.mainContext.Functions

			flagCorrectAnnotations := true
			for _, arg := range curNode.args {
				_receivedArg, _ := w.Stack.Pop()
				_receivedArgType := typeof(_receivedArg)[len("*parser."):]

				needed := arg.Annotation
				getted := bijectionAnnotationTypes[_receivedArgType]

				if getted == "" {
					switch __rectype := _receivedArg.(type) {
					case *ClassScope:
						getted = __rectype.Name
					}
				}

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

				returnedValueAnnotation := "*parser.Nil"
				_retArgType := ""
				getted := ""

				if ok {
					returnValue.Walk(w)
					result, _ := w.Stack.Pop()
					returnedValueAnnotation = typeof(result)

					switch __rectype := result.(type) {
					case *ClassScope:
						getted = __rectype.Name
					default:
						_retArgType = returnedValueAnnotation[len("*parser."):]
						getted = bijectionAnnotationTypes[_retArgType]
					}

					w.Stack.Push(result)
				} else {
					w.Stack.Push(&Nil{Line: line})
				}

				needed := curNode.returnAnnotation

				if needed != getted {
					err := errors.New(
						fmt.Sprintf("invalid type of return value in function '%s' (expected %s, getted %s)",
							curNode.name, curNode.returnAnnotation, bijectionAnnotationTypes[_retArgType]))
					w.CreatePullError(err, curNode.Line)
				}

				w.switchMainContext()
			}

		}
		if w.currentClass != "" {
			w.Ctx.ClassesMethods[w.currentClass][curNode.name] = functor
			return false
		}
		w.Ctx.Functions[curNode.name] = functor

	case *Class:

		w.currentClass = curNode.className
		w.Ctx.Classes[curNode.className] = curNode.fields
		w.Ctx.Functions[curNode.className] = func(w Executor, opNode Node, argCount, line int) {

			var Fields []VarWithAnnotation
			for _, item := range curNode.fields {
				Fields = append(Fields, item)
			}

			reverseAny(Fields)

			newStruct := map[string]Node{}

			for _, arg := range Fields {
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
						//flagCorrectAnnotations = false
						return
					}
				}
				_receivedArg.Walk(w)
				simplified, _ := w.Stack.Pop()
				newStruct[arg.Name] = simplified
			}

			//w.mainContext.Vars[curNode.className] = &ClassScope{newStruct, line}

			w.Stack.Push(&ClassScope{curNode.className, newStruct, line})

		}

		w.Ctx.ClassesMethods[curNode.className] = make(map[string]func(w Executor, curNode Node, argCount, line int))

		for _, item := range curNode.methods {
			item.Walk(w)
		}

		w.currentClass = ""

		fmt.Println(w.Ctx.ClassesMethods["People"])
		return false
	}

	return true
}

func (w Indexer) LeaveNode(n Node) {

}
