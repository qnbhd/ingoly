package parser

import (
	"errors"
	"ingoly/utils/tokenizer"
	"strconv"
)

type Parser struct {
	Tokens    []tokenizer.Token
	variables *BlockContext
	size      int
	pos       int
}

func (ps *Parser) New(tokens []tokenizer.Token) *Parser {
	return &Parser{Tokens: tokens, variables: NewBlockContext(), size: len(tokens), pos: 0}
}

func (ps *Parser) Parse() Ast {
	ast := Ast{[]Node{}, ps.variables}

	for !ps.match(tokenizer.EOF) {
		ast.Tree = append(ast.Tree, ps.Node())
	}

	return ast
}

func (ps *Parser) Node() Node {
	line := ps.get(0).Line

	if ps.match(tokenizer.PRINT) {
		return &PrintNode{node: ps.Expression(), Line: line}
	} else if ps.match(tokenizer.IF) {
		return ps.IfElseBlock()
	}
	return ps.AssignNode()
}

func (ps *Parser) AssignNode() Node {

	if ps.match(tokenizer.VAR) && ps.get(0).Type == tokenizer.NAME &&
		ps.get(1).Type == tokenizer.COLONEQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.match(tokenizer.NAME)

		_, ok := ps.consume(tokenizer.COLONEQUAL)
		if ok == nil {
			return &DeclarationNode{Variable: variable, Expression: ps.Expression(), Line: line}
		}
	}

	panic("Eq err")
}

func (ps *Parser) IfElseBlock() Node {
	line := ps.get(0).Line

	condition := ps.Expression()
	ps.consume(tokenizer.COLON)
	ifStmt := ps.Node()
	var elseStmt Node
	if ps.match(tokenizer.ELSE) {
		ps.consume(tokenizer.COLON)
		elseStmt = ps.Node()
	} else {
		elseStmt = nil
	}
	return &IfNode{condition, ifStmt, elseStmt, line}
}

func (ps *Parser) Expression() Node {
	return ps.Conditional()
}

func (ps *Parser) Conditional() Node {
	result := ps.Additive()

	for {
		line := ps.get(0).Line

		if ps.match(tokenizer.LESS) {
			result = &ConditionalNode{"<", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.GREATER) {
			result = &ConditionalNode{">", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.EQEQUAL) {
			result = &ConditionalNode{"==", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.LESSEQUAL) {
			result = &ConditionalNode{"<=", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.GREATEREQUAL) {
			result = &ConditionalNode{">=", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.NOTEQUAL) {
			result = &ConditionalNode{"!=", result, ps.Additive(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) Additive() Node {
	result := ps.Mult()

	for {
		line := ps.get(0).Line
		if ps.match(tokenizer.PLUS) {
			result = &BinaryNode{"+", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.MINUS) {
			result = &BinaryNode{"-", result, ps.Additive(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) Mult() Node {
	result := ps.Unary()

	for {
		line := ps.get(0).Line
		if ps.match(tokenizer.STAR) {
			result = &BinaryNode{"*", result, ps.Unary(), line}
			continue
		}
		if ps.match(tokenizer.SLASH) {
			result = &BinaryNode{"/", result, ps.Unary(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) Unary() Node {
	line := ps.get(0).Line
	if ps.match(tokenizer.MINUS) {
		return &UnaryNode{"-", ps.PRIMARY(), line}
	}
	return ps.PRIMARY()
}

func (ps *Parser) PRIMARY() Node {
	current := ps.get(0)
	line := current.Line

	if ps.match(tokenizer.NUMBER) {
		lex, _ := strconv.ParseFloat(current.Lexeme, 64)
		return &ValueNode{value: NumberValue{lex}, Line: line}
	}
	if ps.match(tokenizer.STRING) {
		return &ValueNode{value: StringValue{value: current.Lexeme}, Line: line}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.Expression()
		ps.match(tokenizer.RPAR)
		return result
	}
	if ps.match(tokenizer.NAME) {
		return &UsingVariableNode{name: current.Lexeme, Line: line}
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
