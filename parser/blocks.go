package parser

import "ingoly/parser/tokenizer"

func (ps *Parser) Block() Node {
	block := BlockNode{}
	block.Line = ps.get(0).Line
	ps.consume(tokenizer.LBRACE)

	for !ps.match(tokenizer.RBRACE) {
		block.Nodes = append(block.Nodes, ps.Node())
	}

	return &block
}

func (ps *Parser) StatementOrBlock() Node {

	if ps.lookahead(0, tokenizer.LBRACE) {
		return ps.Block()
	}
	return ps.Node()
}

func (ps *Parser) Expression() Node {
	return ps.LogicalOr()
}
