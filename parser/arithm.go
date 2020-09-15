package parser

func (ps *Parser) Additive() Node {
	result := ps.Mult()

	for {
		line := ps.get(0).Line
		if ps.match(PLUS) {
			result = &BinaryNode{"+", result, ps.Additive(), line}
			continue
		}
		if ps.match(MINUS) {
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
		if ps.match(STAR) {
			result = &BinaryNode{"*", result, ps.Unary(), line}
			continue
		}
		if ps.match(SLASH) {
			result = &BinaryNode{"/", result, ps.Unary(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) Unary() Node {
	line := ps.get(0).Line
	if ps.match(MINUS) {
		return &UnaryNode{"-", ps.atomic(), line}
	}
	return ps.atomic()
}
