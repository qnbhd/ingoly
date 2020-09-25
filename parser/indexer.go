package parser

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type Indexer struct {
	Ctx          *Context
	CurrentClass string
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

	return &Indexer{Ctx: newCtx, CurrentClass: ""}
}

func (v Indexer) EnterNode(n Node) bool {

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

			for _, arg := range curNode.args {
				_receivedArg, _ := w.Stack.Pop()
				_receivedArgType := typeof(_receivedArg)

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
						return
					}
				}
				functionCtx.Vars[arg.Name] = _receivedArg
			}

			w.currentContext = functionCtx

			curNode.body.Walk(w)
			returnValue, ok := w.Stack.Pop()

			for idx, interrupt := range w.interruptionsPull.interruptions {
				if strings.Contains(interrupt, functionLabel) {
					w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
				}
			}

			w.lastStructLabel = ""

			returnedValueAnnotation := "Nil"
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
					_retArgType = returnedValueAnnotation
					getted = bijectionAnnotationTypes[_retArgType]
				}

				w.Stack.Push(result)
			} else {
				getted = "nil"
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

		if v.CurrentClass != "" {
			v.Ctx.ClassesMethods[v.CurrentClass][curNode.name] = functor
			return false
		}
		v.Ctx.Functions[curNode.name] = functor
		return false

	case *Class:

		v.CurrentClass = curNode.className
		v.Ctx.Classes[curNode.className] = curNode.fields
		v.Ctx.Functions[curNode.className] = func(w Executor, opNode Node, argCount, line int) {

			var Fields []VarWithAnnotation
			for _, item := range curNode.fields {
				Fields = append(Fields, item)
			}

			methodLabel := fmt.Sprintf("__method%d", rand.Int())
			w.lastStructLabel = methodLabel

			reverseAny(Fields)

			newStruct := map[string]Node{}

			for _, arg := range Fields {
				_receivedArg, _ := w.Stack.Pop()
				_receivedArgType := typeof(_receivedArg)

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

			for idx, interrupt := range w.interruptionsPull.interruptions {
				if strings.Contains(interrupt, methodLabel) {
					w.interruptionsPull.interruptions = append(w.interruptionsPull.interruptions[:idx], w.interruptionsPull.interruptions[idx+1:]...)
				}
			}

			//w.mainContext.Vars[curNode.objName] = &ClassScope{newStruct, line}

			scope := ClassScope{curNode.className, newStruct, line}
			scope.fields["this"] = &Ref{&scope}

			w.Stack.Push(&scope)

		}

		v.Ctx.ClassesMethods[curNode.className] = make(map[string]func(w Executor, curNode Node, argCount, line int))

		for _, item := range curNode.methods {
			item.Walk(v)
		}

		v.CurrentClass = ""

		return false
	}

	return true
}

func (w Indexer) LeaveNode(n Node) {

}
