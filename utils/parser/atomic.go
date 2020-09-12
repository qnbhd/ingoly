package parser

import (
	"ingoly/utils/tokenizer"
	"strconv"
	"strings"
)

func (ps *Parser) atomic() Node {
	current := ps.get(0)
	line := current.Line

	if ps.match(tokenizer.NUMBER) {
		if strings.Index(current.Lexeme, ".") != -1 {
			value, _ := strconv.ParseFloat(current.Lexeme, 64)
			return &FloatNumber{value: value, Line: line}
		}
		value, _ := strconv.Atoi(current.Lexeme)
		return &IntNumber{value: value, Line: line}
	}
	if ps.match(tokenizer.STRING) {
		return &String{value: current.Lexeme, Line: line}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.Expression()
		ps.consume(tokenizer.RPAR)
		return result
	}
	if ps.match(tokenizer.TRUE) {
		return &Boolean{value: true}
	}
	if ps.match(tokenizer.FALSE) {
		return &Boolean{value: false}
	}
	if ps.get(0).Type == tokenizer.NAME &&
		ps.get(1).Type == tokenizer.LPAR &&
		contains(__reservedKeywords, current.Lexeme) {
		return ps.Function()
	}
	if ps.match(tokenizer.NAME) {
		return &UsingVariableNode{name: current.Lexeme, Line: line}
	}

	panic("WTF")
}