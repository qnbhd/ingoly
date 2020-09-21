package parser

import (
	"ingoly/parser/tokenizer"
	"strconv"
	"strings"
)

func (ps *Parser) atomic() Node {
	current := ps.get(0)
	line := current.Line

	if ps.match(tokenizer.NUMBER) {
		if strings.Index(current.Lexeme, ".") != -1 {
			value, _ := strconv.ParseFloat(current.Lexeme, 64)
			return &FloatNumber{Value: value, Line: line}
		}
		value, _ := strconv.Atoi(current.Lexeme)
		return &IntNumber{Value: value, Line: line}
	}
	if ps.match(tokenizer.STRING) {
		return &String{value: []rune(current.Lexeme), Line: line}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.Expression()
		ps.consume(tokenizer.RPAR)
		return result
	}
	if ps.get(0).Type == tokenizer.STRUCT {
		return ps.ClassDeclaring()
	}
	if ps.get(0).Type == tokenizer.NAME && ps.get(1).Type == tokenizer.DOT &&
		ps.get(2).Type == tokenizer.NAME && ps.get(3).Type == tokenizer.LPAR {
		return ps.ClassMethod()
	}
	if ps.get(0).Type == tokenizer.NAME && ps.get(1).Type == tokenizer.DOT &&
		ps.get(2).Type == tokenizer.NAME && ps.get(3).Type != tokenizer.LPAR {
		return ps.ClassAccess()
	}

	if ps.get(0).Type == tokenizer.NAME && ps.get(1).Type == tokenizer.LSQB {
		return ps.ArrayElement()
	}
	if ps.get(0).Type == tokenizer.LSQB {
		return ps.Array()
	}
	if ps.get(0).Type == tokenizer.VBAR {
		line := ps.get(0).Line
		ps.consume(tokenizer.VBAR)
		collection := ps.Expression()
		ps.consume(tokenizer.VBAR)
		arg := []Node{collection}
		return &FunctionalNode{arguments: arg, operator: "len", Line: line}
	}
	if ps.match(tokenizer.TRUE) {
		return &Boolean{value: true}
	}
	if ps.match(tokenizer.FALSE) {
		return &Boolean{value: false}
	}
	if ps.get(0).Type == tokenizer.NAME &&
		ps.get(1).Type == tokenizer.LPAR {
		//contains(__reservedKeywords, current.Lexeme) {
		return ps.Function()
	}
	if ps.match(tokenizer.NAME) {
		return &ScopeVar{Name: current.Lexeme, Line: line}
	}

	panic("WTF")
}
