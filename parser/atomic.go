package parser

import (
	"ingoly/parser/tokenizer"
	"strconv"
	"strings"
)

func (ps *Parser) atomic() Node {
	current := ps.get(0)
	line := current.Line

	switch {
	case ps.match(tokenizer.NUMBER):
		if strings.Index(current.Lexeme, ".") != -1 {
			value, _ := strconv.ParseFloat(current.Lexeme, 64)
			return &FloatNumber{Value: value, Line: line}
		}
		value, _ := strconv.Atoi(current.Lexeme)
		return &IntNumber{Value: value, Line: line}
	case ps.match(tokenizer.STRING):
		return &String{value: []rune(current.Lexeme), Line: line}
	case ps.match(tokenizer.LPAR):
		result := ps.Expression()
		ps.consume(tokenizer.RPAR)
		return result
	case ps.match(tokenizer.CLASS):
		return ps.ClassDeclaring()
	case ps.lookahead(0, tokenizer.NAME) && ps.lookahead(1, tokenizer.DOT) &&
		ps.lookahead(2, tokenizer.NAME) && ps.lookahead(3, tokenizer.LPAR):
		return ps.ClassMethod()
	case ps.lookahead(0, tokenizer.NAME) && ps.lookahead(1, tokenizer.DOT) &&
		ps.lookahead(2, tokenizer.NAME):
		return ps.ClassAccess()
	case ps.lookahead(0, tokenizer.NAME) && ps.lookahead(1, tokenizer.LSQB):
		return ps.ArrayElement()
	case ps.lookahead(0, tokenizer.NAME) && ps.lookahead(1, tokenizer.ARROW) &&
		ps.lookahead(2, tokenizer.LSQB):
		return ps.Array()
	case ps.lookahead(0, tokenizer.VBAR):
		ps.consume(tokenizer.VBAR)
		collection := ps.Expression()
		ps.consume(tokenizer.VBAR)
		arg := []Node{collection}
		return &FunctionalNode{arguments: arg, operator: "len", Line: line}
	case ps.match(tokenizer.TRUE):
		return &Boolean{value: true}
	case ps.match(tokenizer.FALSE):
		return &Boolean{value: false}
	case ps.lookahead(0, tokenizer.NAME) && ps.lookahead(1, tokenizer.LPAR):
		return ps.Function()
	case ps.match(tokenizer.NAME):
		return &ScopeVar{Name: current.Lexeme, Line: line}
	default:
		panic("parsing fatal: error unknown terminal")
	}
}
