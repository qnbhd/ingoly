package parser

import "ingoly/parser/tokenizer"

func (ps *Parser) Additive() Node {
	result := ps.Mult()

	for {
		line := ps.get(0).Line
		switch {
		case ps.match(tokenizer.PLUS):
			result = &BinaryNode{"+", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.MINUS):
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
		switch {
		case ps.match(tokenizer.STAR):
			result = &BinaryNode{"*", result, ps.Unary(), line}
			continue
		case ps.match(tokenizer.SLASH):
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
		return &UnaryNode{"-", ps.atomic(), line}
	}
	return ps.atomic()
}
