package parser

func (ps *Parser) Block() Node {
	block := BlockNode{}
	block.Line = ps.get(0).Line
	ps.consume(LBRACE)
	for !ps.match(RBRACE) {
		block.Nodes = append(block.Nodes, ps.Node())
	}

	return &block
}

func (ps *Parser) StatementOrBlock() Node {

	if ps.get(0).Type == LBRACE {
		return ps.Block()
	}
	return ps.Node()
}

func (ps *Parser) Expression() Node {
	return ps.LogicalOr()
}
