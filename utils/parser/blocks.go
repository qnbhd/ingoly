package parser

import "ingoly/utils/tokenizer"

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

	if ps.get(0).Type == tokenizer.LBRACE {
		return ps.Block()
	}
	return ps.Node()
}

func (ps *Parser) Expression() Node {
	return ps.LogicalOr()
}
