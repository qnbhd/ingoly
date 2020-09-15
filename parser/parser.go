package parser

import (
	"ingoly/errpull"
	"ingoly/parser/tokenizer"
)

type Parser struct {
	Tokens     []tokenizer.Token
	variables  *Context
	size       int
	pos        int
	ErrorsPull *errpull.ErrorsPull
}

func (ps *Parser) New(tokens []tokenizer.Token) *Parser {
	return &Parser{Tokens: tokens, variables: NewContext(), size: len(tokens), pos: 0,
		ErrorsPull: errpull.NewErrorsPull()}
}

func (ps *Parser) Parse() (Ast, *errpull.ErrorsPull) {
	ast := Ast{[]Node{}}

	for !ps.match(tokenizer.EOF) {
		ast.Tree = append(ast.Tree, ps.Node())
	}

	return ast, ps.ErrorsPull
}
