package parser

import "ingoly/utils/tokenizer"

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
		return &UnaryNode{"-", ps.atomic(), line}
	}
	return ps.atomic()
}
