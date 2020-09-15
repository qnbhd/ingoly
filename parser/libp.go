package parser

import (
	"errors"
	"ingoly/errpull"
)

func contains(s []string, searchTerm string) bool {
	for _, item := range s {
		if item == searchTerm {
			return true
		}
	}
	return false
}

func (ps *Parser) match(tokenType TokenType) bool {
	current := ps.get(0)
	if tokenType != current.Type {
		return false
	}
	ps.pos++
	return true
}

func (ps *Parser) consume(tokenType TokenType) (Token, error) {
	current := ps.get(0)
	err := errors.New("parsing: " + tokenType.String() + " was expected")

	if tokenType != current.Type {
		line := ps.get(0).Line
		inn := errpull.NewInnerError(err, line)
		ps.ErrorsPull.Errors = append(ps.ErrorsPull.Errors, inn)
		return Token{Type: NIL}, err
	}
	ps.pos++
	return current, nil
}

func (ps *Parser) get(relativePosition int) Token {
	position := ps.pos + relativePosition
	if position >= ps.size {
		return Token{Type: EOF, Lexeme: ""}
	}
	return ps.Tokens[position]
}
