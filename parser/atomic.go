package parser

import (
	"strconv"
	"strings"
)

func (ps *Parser) atomic() Node {
	current := ps.get(0)
	line := current.Line

	if ps.match(NUMBER) {
		if strings.Index(current.Lexeme, ".") != -1 {
			value, _ := strconv.ParseFloat(current.Lexeme, 64)
			return &FloatNumber{value: value, Line: line}
		}
		value, _ := strconv.Atoi(current.Lexeme)
		return &IntNumber{value: value, Line: line}
	}
	if ps.match(STRING) {
		return &String{value: current.Lexeme, Line: line}
	}
	if ps.match(LPAR) {
		result := ps.Expression()
		ps.consume(RPAR)
		return result
	}
	if ps.match(TRUE) {
		return &Boolean{value: true}
	}
	if ps.match(FALSE) {
		return &Boolean{value: false}
	}
	if ps.get(0).Type == NAME &&
		ps.get(1).Type == LPAR {
		//contains(__reservedKeywords, current.Lexeme) {
		return ps.Function()
	}
	if ps.match(NAME) {
		return &ScopeVar{name: current.Lexeme, Line: line}
	}

	panic("WTF")
}
