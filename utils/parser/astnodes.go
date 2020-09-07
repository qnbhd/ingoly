package parser

import (
	"errors"
	"strconv"
)

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

/* Base Node */

type Node interface {
	Execute() (float64, error)
	ToString() string
	getNodesList() []Node
}

////////////////////////////

/* Number Node */

type NumberNode struct {
	value float64
}

func (nn *NumberNode) Execute() (float64, error) {
	return nn.value, nil
}

func (nn *NumberNode) ToString() string {
	return "NUMBER_NODE " + FloatToString(nn.value)
}

func (nn *NumberNode) getNodesList() []Node {
	return []Node{nil}
}

////////////////////////////

/* Binary Node */

type BinaryNode struct {
	operation rune
	op1       Node
	op2       Node
}

func (bn *BinaryNode) New(operation rune, exp1, exp2 Node) *BinaryNode {
	return &BinaryNode{operation: operation, op1: exp1, op2: exp2}
}

func (bn *BinaryNode) Execute() (float64, error) {
	result1, _ := bn.op1.Execute()
	result2, _ := bn.op2.Execute()

	switch bn.operation {
	case '+':
		return result1 + result2, nil
	case '-':
		return result1 - result2, nil
	case '*':
		return result1 * result2, nil
	case '/':
		if result2 == 0 {
			return 0, errors.New("division by zero")
		}
		return result1 / result2, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

func (bn *BinaryNode) ToString() string {
	return "BINARY_OPERATION '" + string(bn.operation) + "'"
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

func (un *UnaryNode) Execute() (float64, error) {
	result1, _ := un.op1.Execute()

	switch un.operation {
	case '+':
		return result1, nil
	case '-':
		return -result1, nil
	default:
		return 0, errors.New("unknown operation")
	}
}

func (un *UnaryNode) ToString() string {
	return "UNARY_OPERATION '" + string(un.operation) + "'"
}

func (un *UnaryNode) getNodesList() []Node {
	return []Node{un.op1}
}
