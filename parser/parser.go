package parser

import (
	"ingoly/errpull"
)

type Parser struct {
	Tokens     []Token
	variables  *BlockContext
	size       int
	pos        int
	ErrorsPull *errpull.ErrorsPull
}

func (ps *Parser) New(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, variables: NewBlockContext(), size: len(tokens), pos: 0,
		ErrorsPull: errpull.NewErrorsPull()}
}

func (ps *Parser) Parse() (Ast, *errpull.ErrorsPull) {
	ast := Ast{[]Node{}}

	for !ps.match(EOF) {
		ast.Tree = append(ast.Tree, ps.Node())
	}

	return ast, ps.ErrorsPull
}
