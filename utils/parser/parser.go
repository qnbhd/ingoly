package parser

import (
	"ingoly/utils/tokenizer"
	"strconv"
)

type Parser struct {
	Tokens    []tokenizer.Token
	variables VarTable
	size      int
	pos       int
}

func (ps *Parser) New(tokens []tokenizer.Token) *Parser {
	return &Parser{Tokens: tokens, variables: VarTable{}, size: len(tokens), pos: 0}
}

func (ps *Parser) Parse() Ast {
	ast := Ast{[]Node{}, ps.variables}

	for !ps.match(tokenizer.EOF) {
		ast.Tree = append(ast.Tree, ps.EXPRESSION())
	}

	return ast
}

func (ps *Parser) EXPRESSION() Node {
	return ps.ADDITIVE()
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
		return &NumberNode{value: lex}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.EXPRESSION()
		ps.match(tokenizer.RPAR)
		return result
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

func (ps *Parser) get(relativePosition int) tokenizer.Token {
	position := ps.pos + relativePosition
	if position >= ps.size {
		return tokenizer.Token{Type: tokenizer.EOF, Lexeme: ""}
	}
	return ps.Tokens[position]
}
