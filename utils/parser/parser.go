package parser

import (
	"errors"
	"ingoly/utils/tokenizer"
	"strconv"
)

type Parser struct {
	Tokens    []tokenizer.Token
	variables map[string]Value
	size      int
	pos       int
}

func (ps *Parser) New(tokens []tokenizer.Token) *Parser {
	VarTable = make(map[string]Value)
	return &Parser{Tokens: tokens, variables: VarTable, size: len(tokens), pos: 0}
}

func (ps *Parser) Parse() Ast {
	ast := Ast{[]Node{}, ps.variables}

	for !ps.match(tokenizer.EOF) {
		ast.Tree = append(ast.Tree, ps.Node())
	}

	return ast
}

func (ps *Parser) Node() Node {
	if ps.match(tokenizer.PRINT) {
		return &PrintNode{node: ps.EXPRESSION()}
	} else if ps.match(tokenizer.IF) {
		return ps.IfElseBlock()
	}
	return ps.ASSIGNNode()
}

func (ps *Parser) ASSIGNNode() Node {
	current := ps.get(0)

	if ps.match(tokenizer.NAME) && ps.get(0).Type == tokenizer.EQUAL {
		variable := current.Lexeme
		_, ok := ps.consume(tokenizer.EQUAL)
		if ok == nil {
			return &AssignmentNode{Variable: variable, Expression: ps.EXPRESSION()}
		}
	}
	panic("Eq err")
}

func (ps *Parser) IfElseBlock() Node {
	condition := ps.EXPRESSION()
	ifStmt := ps.Node()
	ps.consume(tokenizer.COLON)
	var elseStmt Node
	if ps.match(tokenizer.ELSE) {
		ps.consume(tokenizer.COLON)
		elseStmt = ps.Node()
	} else {
		elseStmt = nil
	}
	return &IfNode{condition, ifStmt, elseStmt}
}

func (ps *Parser) EXPRESSION() Node {
	return ps.Conditional()
}

func (ps *Parser) Conditional() Node {
	result := ps.ADDITIVE()

	for {
		if ps.match(tokenizer.LESS) {
			result = &ConditionalNode{'<', result, ps.ADDITIVE()}
			continue
		}
		if ps.match(tokenizer.GREATER) {
			result = &ConditionalNode{'>', result, ps.ADDITIVE()}
			continue
		}
		if ps.match(tokenizer.EQUAL) {
			result = &ConditionalNode{'=', result, ps.ADDITIVE()}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) ADDITIVE() Node {
	result := ps.MULT()

	for {
		if ps.match(tokenizer.PLUS) {
			result = &BinaryNode{'+', result, ps.ADDITIVE()}
			continue
		}
		if ps.match(tokenizer.MINUS) {
			result = &BinaryNode{'-', result, ps.ADDITIVE()}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) MULT() Node {
	result := ps.UNARY()

	for {
		if ps.match(tokenizer.STAR) {
			result = &BinaryNode{'*', result, ps.UNARY()}
			continue
		}
		if ps.match(tokenizer.SLASH) {
			result = &BinaryNode{'/', result, ps.UNARY()}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) UNARY() Node {
	if ps.match(tokenizer.MINUS) {
		return &UnaryNode{'-', ps.PRIMARY()}
	}
	return ps.PRIMARY()
}

func (ps *Parser) PRIMARY() Node {
	current := ps.get(0)

	if ps.match(tokenizer.NUMBER) {
		lex, _ := strconv.ParseFloat(current.Lexeme, 64)
		return &ValueNode{value: NumberValue{lex}}
	}
	if ps.match(tokenizer.STRING) {
		return &ValueNode{value: StringValue{current.Lexeme}}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.EXPRESSION()
		ps.match(tokenizer.RPAR)
		return result
	}
	if ps.match(tokenizer.NAME) {
		return &UsingVariableNode{name: current.Lexeme}
	}

	panic("WTF")
}

func (ps *Parser) match(tokenType tokenizer.TokenType) bool {
	current := ps.get(0)
	if tokenType != current.Type {
		return false
	}
	ps.pos++
	return true
}

func (ps *Parser) consume(tokenType tokenizer.TokenType) (tokenizer.Token, error) {
	current := ps.get(0)
	if tokenType != current.Type {
		return tokenizer.Token{Type: tokenizer.NIL},
			errors.New(tokenType.String() + " was expected")
	}
	ps.pos++
	return current, nil
}

func (ps *Parser) get(relativePosition int) tokenizer.Token {
	position := ps.pos + relativePosition
	if position >= ps.size {
		return tokenizer.Token{Type: tokenizer.EOF, Lexeme: ""}
	}
	return ps.Tokens[position]
}
