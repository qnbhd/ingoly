package parser

import "ingoly/utils/tokenizer"

func (ps *Parser) Node() Node {

	if ps.match(tokenizer.IF) {
		return ps.IfElseBlock()
	} else if ps.match(tokenizer.FOR) {
		return ps.ForBlock()
	} else if ps.match(tokenizer.WHILE) {
		return ps.While()
	} else if ps.match(tokenizer.BREAK) {
		return &Break{Line: ps.get(0).Line}
	} else if ps.match(tokenizer.CONTINUE) {
		return &Continue{Line: ps.get(0).Line}
	} else if ps.match(tokenizer.DECLARE) {
		return ps.FuncDeclaration()
	} else if ps.match(tokenizer.RETURN) {
		return ps.ReturnStmt()
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
	} else if ps.get(0).Type == tokenizer.NAME && ps.get(1).Type == tokenizer.EQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.consume(tokenizer.NAME)
		_, ok := ps.consume(tokenizer.EQUAL)
		if ok == nil {
			return &AssignNode{Variable: variable, Expression: ps.Expression(), Line: line}
		}
	}

	return ps.Expression()
}

func (ps *Parser) Function() Node {
	targetFuncName := ps.get(0).Lexeme
	ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.LPAR)

	var args []Node
	var res Node

	for !ps.match(tokenizer.RPAR) {
		args = append(args, ps.Expression())
		res = &FunctionalNode{args, targetFuncName, ps.get(0).Line}
		ps.match(tokenizer.COMMA)
	}

	return res
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

	start := ps.Expression()
	ps.consume(tokenizer.SEMI)

	stop := ps.Expression()

	var step Node
	if ps.get(0).Type == tokenizer.SEMI {
		ps.consume(tokenizer.SEMI)
		step = ps.Expression()
	} else {
		step = &IntNumber{1, line}
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
		start:   start,
		stop:    stop,
		step:    step,
		stmt:    stmt,
		Line:    line,
		strict:  strict,
	}

}

func (ps *Parser) While() Node {
	condition := ps.Expression()
	stmt := ps.StatementOrBlock()
	return &While{condition, stmt, ps.get(0).Line}
}

func (ps *Parser) FuncDeclaration() Node {
	name := ps.get(0).Lexeme
	line := ps.get(0).Line
	ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.LPAR)

	var argNames []string
	for !ps.match(tokenizer.RPAR) {
		res, _ := ps.consume(tokenizer.NAME)
		ps.match(tokenizer.COMMA)
		argNames = append(argNames, res.Lexeme)
	}

	body := ps.StatementOrBlock()

	return &FunctionDeclareNode{name, argNames, body, line}
}

func (ps *Parser) ReturnStmt() Node {

	returnValue := ps.Expression()

	return &Return{returnValue, ps.get(0).Line}
}
