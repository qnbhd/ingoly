package parser

func (ps *Parser) LogicalOr() Node {
	result := ps.LogicalAnd()

	for {
		line := ps.get(0).Line
		if ps.match(VBARVBAR) {
			result = &ConditionalNode{"||", result, ps.LogicalAnd(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) LogicalAnd() Node {
	result := ps.Equality()

	for {
		line := ps.get(0).Line
		if ps.match(AMPERAMPER) {
			result = &ConditionalNode{"&&", result, ps.Equality(), line}
			continue
		}
		break
	}

	return result
}

func (ps *Parser) Equality() Node {
	result := ps.Conditional()
	line := ps.get(0).Line

	if ps.match(EQEQUAL) {
		result = &ConditionalNode{"==", result, ps.Conditional(), line}
	}
	if ps.match(NOTEQUAL) {
		result = &ConditionalNode{"!=", result, ps.Conditional(), line}
	}

	return result
}

func (ps *Parser) Conditional() Node {
	result := ps.Additive()

	for {
		line := ps.get(0).Line

		if ps.match(LESS) {
			result = &ConditionalNode{"<", result, ps.Additive(), line}
			continue
		}
		if ps.match(GREATER) {
			result = &ConditionalNode{">", result, ps.Additive(), line}
			continue
		}
		if ps.match(EQEQUAL) {
			result = &ConditionalNode{"==", result, ps.Additive(), line}
			continue
		}
		if ps.match(LESSEQUAL) {
			result = &ConditionalNode{"<=", result, ps.Additive(), line}
			continue
		}
		if ps.match(GREATEREQUAL) {
			result = &ConditionalNode{">=", result, ps.Additive(), line}
			continue
		}
		if ps.match(NOTEQUAL) {
			result = &ConditionalNode{"!=", result, ps.Additive(), line}
			continue
		}
		break
	}

	return result
}
