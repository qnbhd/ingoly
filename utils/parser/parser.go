package parser

import (
	"errors"
	"ingoly/utils/errpull"
	"ingoly/utils/tokenizer"
	"strconv"
)

type Parser struct {
	Tokens     []tokenizer.Token
	variables  *BlockContext
	size       int
	pos        int
	ErrorsPull *errpull.ErrorsPull
}

func (ps *Parser) New(tokens []tokenizer.Token) *Parser {
	return &Parser{Tokens: tokens, variables: NewBlockContext(), size: len(tokens), pos: 0,
		ErrorsPull: errpull.NewErrorsPull()}
}

func (ps *Parser) Parse() (Ast, *errpull.ErrorsPull) {
	ast := Ast{[]Node{}}

	for !ps.match(tokenizer.EOF) {
		ast.Tree = append(ast.Tree, ps.Node())
	}

	return ast, ps.ErrorsPull
}

func (ps *Parser) Block() Node {
	block := BlockNode{}
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

func (ps *Parser) Node() Node {
	line := ps.get(0).Line

	if ps.match(tokenizer.PRINT) {
		res := &PrintNode{node: ps.Expression(), Line: line}
		return res
	} else if ps.match(tokenizer.IF) {
		return ps.IfElseBlock()
	} else if ps.match(tokenizer.FOR) {
		return ps.ForBlock()
	}
	return ps.AssignNode()
}

func (ps *Parser) AssignNode() Node {

	if ps.match(tokenizer.VAR) && ps.get(0).Type == tokenizer.NAME &&
		ps.get(1).Type == tokenizer.COLONEQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.match(tokenizer.NAME)

		_, ok := ps.consume(tokenizer.COLONEQUAL)
		if ok == nil {
			return &DeclarationNode{Variable: variable, Expression: ps.Expression(), Line: line}
		}
	}

	panic("Eq err")
}

func (ps *Parser) IfElseBlock() Node {
	line := ps.get(0).Line

	condition := ps.Expression()
	ifStmt := ps.StatementOrBlock()
	var elseStmt Node
	if ps.match(tokenizer.ELSE) {
		elseStmt = ps.StatementOrBlock()
	} else {
		elseStmt = nil
	}
	return &IfNode{condition, ifStmt, elseStmt, line}
}

func (ps *Parser) ForBlock() Node {
	line := ps.get(0).Line

	iterVar, _ := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.IN)
	ps.consume(tokenizer.LSQB)

	startT, _ := ps.consume(tokenizer.NUMBER)
	start, _ := strconv.ParseFloat(startT.Lexeme, 64)

	ps.consume(tokenizer.SEMI)
	stopT, _ := ps.consume(tokenizer.NUMBER)
	stop, _ := strconv.ParseFloat(stopT.Lexeme, 64)

	step := 1.
	if ps.get(0).Type == tokenizer.SEMI {
		ps.consume(tokenizer.SEMI)
		stepT, _ := ps.consume(tokenizer.NUMBER)
		step, _ = strconv.ParseFloat(stepT.Lexeme, 64)
	}

	strict := false
	switch ps.get(0).Type {
	case tokenizer.RSQB:
		ps.consume(tokenizer.RSQB)
		strict = true
	case tokenizer.RPAR:
		ps.consume(tokenizer.RPAR)
		strict = false
	}

	stmt := ps.StatementOrBlock()

	return &ForNode{
		iterVar: iterVar.Lexeme,
		start:   NumberValue{start},
		stop:    NumberValue{stop},
		step:    NumberValue{step},
		stmt:    stmt,
		Line:    line,
		strict:  strict,
	}

}

func (ps *Parser) Expression() Node {
	return ps.LogicalOr()
}

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

		if ps.match(tokenizer.LESS) {
			result = &ConditionalNode{"<", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.GREATER) {
			result = &ConditionalNode{">", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.EQEQUAL) {
			result = &ConditionalNode{"==", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.LESSEQUAL) {
			result = &ConditionalNode{"<=", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.GREATEREQUAL) {
			result = &ConditionalNode{">=", result, ps.Additive(), line}
			continue
		}
		if ps.match(tokenizer.NOTEQUAL) {
			result = &ConditionalNode{"!=", result, ps.Additive(), line}
			continue
		}
		break
	}

	return result
}

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
		return &UnaryNode{"-", ps.PRIMARY(), line}
	}
	return ps.PRIMARY()
}

func (ps *Parser) PRIMARY() Node {
	current := ps.get(0)
	line := current.Line

	if ps.match(tokenizer.NUMBER) {
		lex, _ := strconv.ParseFloat(current.Lexeme, 64)
		return &ValueNode{value: NumberValue{lex}, Line: line}
	}
	if ps.match(tokenizer.STRING) {
		return &ValueNode{value: StringValue{value: current.Lexeme}, Line: line}
	}
	if ps.match(tokenizer.LPAR) {
		result := ps.Expression()
		ps.consume(tokenizer.RPAR)
		return result
	}
	if ps.match(tokenizer.NAME) {
		return &UsingVariableNode{name: current.Lexeme, Line: line}
	}

	panic("WTF")
}

func (ps *Parser) match(tokenType tokenizer.TokenType) bool {
	current := ps.get(0)
	if tokenType != current.Type {
		return false
	}
	ps.pos++
	return true
}

func (ps *Parser) consume(tokenType tokenizer.TokenType) (tokenizer.Token, error) {
	current := ps.get(0)
	err := errors.New("parsing: " + tokenType.String() + " was expected")

	if tokenType != current.Type {
		line := ps.get(0).Line
		inn := errpull.NewInnerError(err, line)
		ps.ErrorsPull.Errors = append(ps.ErrorsPull.Errors, inn)
		return tokenizer.Token{Type: tokenizer.NIL}, err
	}
	ps.pos++
	return current, nil
}

func (ps *Parser) get(relativePosition int) tokenizer.Token {
	position := ps.pos + relativePosition
	if position >= ps.size {
		return tokenizer.Token{Type: tokenizer.EOF, Lexeme: ""}
	}
	return ps.Tokens[position]
}
