package parser

func (ps *Parser) Node() Node {

	if ps.match(IF) {
		return ps.IfElseBlock()
	} else if ps.match(FOR) {
		return ps.ForBlock()
	} else if ps.match(WHILE) {
		return ps.While()
	} else if ps.match(BREAK) {
		return &Break{Line: ps.get(0).Line}
	} else if ps.match(CONTINUE) {
		return &Continue{Line: ps.get(0).Line}
	} else if ps.match(DECLARE) {
		return ps.FuncDeclaration()
	} else if ps.match(RETURN) {
		return ps.ReturnStmt()
	} else if ps.match(NIL) {
		return &Nil{ps.get(0).Line}
	}
	return ps.AssignNode()
}

func (ps *Parser) AssignNode() Node {

	if ps.match(VAR) && ps.get(0).Type == NAME &&
		ps.get(1).Type == COLONEQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.match(NAME)

		_, ok := ps.consume(COLONEQUAL)
		if ok == nil {
			return &DeclarationNode{Variable: variable, Expression: ps.Expression(), Line: line}
		}
	} else if ps.get(0).Type == NAME && ps.get(1).Type == EQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.consume(NAME)
		_, ok := ps.consume(EQUAL)
		if ok == nil {
			return &AssignNode{Variable: variable, Expression: ps.Expression(), Line: line}
		}
	}

	return ps.Expression()
}

func (ps *Parser) Function() Node {
	targetFuncName := ps.get(0).Lexeme
	ps.consume(NAME)
	ps.consume(LPAR)

	var args []Node
	res := FunctionalNode{args, targetFuncName, ps.get(0).Line}

	for !ps.match(RPAR) {
		res.arguments = append(res.arguments, ps.Expression())
		ps.match(COMMA)
	}

	return &res
}

func (ps *Parser) IfElseBlock() Node {
	line := ps.get(0).Line

	condition := ps.Expression()
	ifStmt := ps.StatementOrBlock()
	var elseStmt Node
	if ps.match(ELSE) {
		elseStmt = ps.StatementOrBlock()
	} else {
		elseStmt = nil
	}
	return &IfNode{condition, ifStmt, elseStmt, line}
}

func (ps *Parser) ForBlock() Node {
	line := ps.get(0).Line

	iterVar, _ := ps.consume(NAME)
	ps.consume(IN)
	ps.consume(LSQB)

	start := ps.Expression()
	ps.consume(SEMI)

	stop := ps.Expression()

	var step Node
	if ps.get(0).Type == SEMI {
		ps.consume(SEMI)
		step = ps.Expression()
	} else {
		step = &IntNumber{1, line}
	}

	strict := false
	switch ps.get(0).Type {
	case RSQB:
		ps.consume(RSQB)
		strict = true
	case RPAR:
		ps.consume(RPAR)
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
	ps.consume(NAME)
	ps.consume(LPAR)

	var argNames []VarWithAnnotation

	for !ps.match(RPAR) {
		varName, _ := ps.consume(NAME)
		varAnnotation, _ := ps.consume(NAME)
		resultVar := VarWithAnnotation{varName.Lexeme, varAnnotation.Lexeme}
		ps.match(COMMA)
		argNames = append(argNames, resultVar)
	}

	returnAnnotation := "nil"

	if ps.get(0).Type == ARROW {
		ps.consume(ARROW)
		res, _ := ps.consume(NAME)
		returnAnnotation = res.Lexeme
	}

	body := ps.StatementOrBlock()

	return &FunctionDeclareNode{name,
		argNames,
		returnAnnotation,
		body,
		line}
}

func (ps *Parser) ReturnStmt() Node {

	returnValue := ps.Expression()

	return &Return{returnValue, ps.get(0).Line}
}
