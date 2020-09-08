package parser

import (
	"errors"
	"fmt"
	"strconv"
)

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 3, 64)
}

/* Base Node */

type Node interface {
	Execute() (Value, error)
	ToString() string
	getNodesList() []Node
}

////////////////

/* Binary Node */

type BinaryNode struct {
	operation rune
	op1       Node
	op2       Node
}

func (bn *BinaryNode) New(operation rune, exp1, exp2 Node) *BinaryNode {
	return &BinaryNode{operation: operation, op1: exp1, op2: exp2}
}

func (bn *BinaryNode) Execute() (Value, error) {
	result1, _ := bn.op1.Execute()
	result2, _ := bn.op2.Execute()

	switch v := result1.(type) {

	case StringValue:
		string1 := v.AsString()
		if bn.operation == '+' {
			return &StringValue{value: string1 + result2.AsString()}, nil
		} else {
			return &StringValue{value: string1}, errors.New("undefined operation")
		}

	case NumberValue:
		number1 := result1.AsNumber()
		number2 := result2.AsNumber()

		switch bn.operation {
		case '+':
			return &NumberValue{value: number1 + number2}, nil
		case '-':
			return &NumberValue{value: number1 - number2}, nil
		case '*':
			return &NumberValue{value: number1 * number2}, nil
		case '/':
			if number2 == 0 {
				return &NumberValue{value: 0}, errors.New("division by zero")
			}
			return &NumberValue{value: number1 / number2}, nil
		default:
			return &NumberValue{value: 0}, errors.New("unknown operation")
		}
	}

	return &NumberValue{value: 0}, errors.New("unknown bin expression")
}

func (bn *BinaryNode) ToString() string {
	return "BINARY OPERATION (Operation) '" + string(bn.operation) + "'"
	//return "[" + bn.op1.ToString() + ", " + bn.op2.ToString() +
	//	", OP:" + string(bn.operation) + "]"
}

func (bn *BinaryNode) getNodesList() []Node {
	return []Node{bn.op1, bn.op2}
}

////////////////////////////

/* Unary Node */

type UnaryNode struct {
	operation rune
	op1       Node
}

func (un *UnaryNode) New(operation rune, op1 Node) *UnaryNode {
	return &UnaryNode{operation: operation, op1: op1}
}

func (un *UnaryNode) Execute() (Value, error) {
	result1, _ := un.op1.Execute()

	switch un.operation {
	case '+':
		return NumberValue{value: result1.AsNumber()}, nil
	case '-':
		return NumberValue{value: -result1.AsNumber()}, nil
	default:
		return NumberValue{value: 0}, errors.New("unknown operation")
	}
}

func (un *UnaryNode) ToString() string {
	return "UNARY OPERATION (Operation) '" + string(un.operation) + "'"
}

func (un *UnaryNode) getNodesList() []Node {
	return []Node{un.op1}
}

////////////////////////////

/* Name Node */

type UsingVariableNode struct {
	name string
}

func (uvn *UsingVariableNode) Execute() (Value, error) {
	if value, ok := VarTable[uvn.name]; ok {
		return value, nil
	}
	return NumberValue{0}, errors.New("unknown identifier")
}

func (uvn *UsingVariableNode) ToString() string {
	return "USING VARIABLE (Identifier) '" + uvn.name + "'"
}

func (uvn *UsingVariableNode) getNodesList() []Node {
	return []Node{}
}

///////////////

/* Value Node */

type ValueNode struct {
	value Value
}

func (vn *ValueNode) Execute() (Value, error) {
	return vn.value, nil
}

func (vn *ValueNode) ToString() string {
	return "VALUE NODE (Value) '" + vn.value.AsString() + "'"
}

func (vn *ValueNode) getNodesList() []Node {
	return []Node{}
}

////////////////

/* Binary Node */

type ConditionalNode struct {
	operation rune
	op1       Node
	op2       Node
}

func (bn *ConditionalNode) New(operation rune, exp1, exp2 Node) *ConditionalNode {
	return &ConditionalNode{operation: operation, op1: exp1, op2: exp2}
}

func (bn *ConditionalNode) Execute() (Value, error) {
	result1, _ := bn.op1.Execute()
	result2, _ := bn.op2.Execute()

	switch v := result1.(type) {

	case StringValue:
		string1 := v.AsString()
		if bn.operation == '>' {
			res := 0.
			if string1 > result2.AsString() {
				res = 1
			}
			return &NumberValue{value: res}, nil
		} else if bn.operation == '<' {
			res := 0.
			if string1 > result2.AsString() {
				res = 1
			}
			return &NumberValue{value: res}, nil
		} else {
			return &NumberValue{value: 0}, errors.New("undefined operation")
		}

	case NumberValue:
		number1 := result1.AsNumber()
		number2 := result2.AsNumber()

		switch bn.operation {
		case '=':
			res := 0.
			if number1 == number2 {
				res = 1
			}
			return &NumberValue{value: res}, nil
		case '<':
			res := 0.
			if number1 < number2 {
				res = 1
			}
			return &NumberValue{value: res}, nil
		case '>':
			res := 0.
			if number1 > number2 {
				res = 1
			}
			return &NumberValue{value: res}, nil
		default:
			return &NumberValue{value: 0}, errors.New("unknown operation")
		}
	}

	return &NumberValue{value: 0}, errors.New("unknown bin expression")
}

func (bn *ConditionalNode) ToString() string {
	return "CONDITIONAL (Operation) '" + string(bn.operation) + "'"
	//return "[" + bn.op1.ToString() + ", " + bn.op2.ToString() +
	//	", OP:" + string(bn.operation) + "]"
}

func (bn *ConditionalNode) getNodesList() []Node {
	return []Node{bn.op1, bn.op2}
}

//////////////////

/* Assignment Node */

type AssignmentNode struct {
	Variable   string
	Expression Node
}

func (as *AssignmentNode) New(variable string, node Node) *AssignmentNode {
	return &AssignmentNode{Variable: variable, Expression: node}
}

func (as *AssignmentNode) Execute() (Value, error) {
	result, ok := as.Expression.Execute()
	if ok == nil {
		VarTable[as.Variable] = result
	}
	return NumberValue{0}, nil
}

func (as *AssignmentNode) ToString() string {
	return "ASSIGNMENT Node (Node) Identifier: '" + as.Variable + "'"
}

func (as *AssignmentNode) getNodesList() []Node {
	return []Node{as.Expression}
}

//////////////////

/* Print Node */

type PrintNode struct {
	node Node
}

func (ps *PrintNode) Execute() (Value, error) {
	result, _ := ps.node.Execute()
	fmt.Println(result.AsString())
	return NumberValue{0}, nil
}

func (ps *PrintNode) ToString() string {
	return "PRINT OPERATOR (Keyword)"
}

func (ps *PrintNode) getNodesList() []Node {
	return []Node{ps.node}
}

//////////////////

/* If Node */

type IfNode struct {
	node     Node
	ifStmt   Node
	elseStmt Node
}

func (is *IfNode) Execute() (Value, error) {
	result, _ := is.node.Execute()

	var err error
	if result.AsNumber() != 0 {
		_, err = is.ifStmt.Execute()
	} else if is.elseStmt != nil {
		_, err = is.elseStmt.Execute()
	}

	return NumberValue{0}, err
}

func (is *IfNode) ToString() string {
	//result := "IF-ELSE Node (Statement) \n"
	//result += "   !-->" + is.node.ToString() + "\n"
	//result += "   !-->" + "IF (Statement)\n"
	//result += "      !-->" + is.ifStmt.ToString() + "\n"
	//if is.elseStmt != nil {
	//	result += "   !-->" + "ELSE (Statement)\n"
	//	result += "      !-->" + is.elseStmt.ToString() + "\n"
	//}
	return "IF ELSE NODE (Statement)"
}

func (is *IfNode) getNodesList() []Node {
	return []Node{is.node, is.ifStmt, is.elseStmt}
}
