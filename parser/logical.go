package parser

import "ingoly/parser/tokenizer"

func (ps *Parser) LogicalOr() Node {
	result := ps.LogicalAnd()

	for {
		line := ps.get(0).Line
		if ps.match(tokenizer.VBARVBAR) {
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
		if ps.match(tokenizer.AMPERAMPER) {
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

	if ps.match(tokenizer.EQEQUAL) {
		result = &ConditionalNode{"==", result, ps.Conditional(), line}
	}
	if ps.match(tokenizer.NOTEQUAL) {
		result = &ConditionalNode{"!=", result, ps.Conditional(), line}
	}

	return result
}

func (ps *Parser) Conditional() Node {
	result := ps.Additive()

	for {
		line := ps.get(0).Line

		switch {
		case ps.match(tokenizer.LESS):
			result = &ConditionalNode{"<", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.GREATER):
			result = &ConditionalNode{">", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.EQEQUAL):
			result = &ConditionalNode{"==", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.LESSEQUAL):
			result = &ConditionalNode{"<=", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.GREATEREQUAL):
			result = &ConditionalNode{">=", result, ps.Additive(), line}
			continue
		case ps.match(tokenizer.NOTEQUAL):
			result = &ConditionalNode{"!=", result, ps.Additive(), line}
			continue
		}
		break
	}

	return result
}
