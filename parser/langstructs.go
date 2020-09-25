package parser

import (
	"ingoly/parser/tokenizer"
)

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
	} else if ps.match(tokenizer.NIL) {
		return &Nil{ps.get(0).Line}
	}
	return ps.AssignNode()
}

func (ps *Parser) AssignNode() Node {

	if ps.match(tokenizer.VAR) && ps.get(0).Type == tokenizer.NAME &&
		ps.get(1).Type == tokenizer.COLONEQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.match(tokenizer.NAME)

		ps.consume(tokenizer.COLONEQUAL)
		return &DeclarationNode{Variable: variable, Expression: ps.Expression(), Line: line}
	} else if ps.get(0).Type == tokenizer.NAME && ps.get(1).Type == tokenizer.EQUAL {
		line := ps.get(0).Line
		variable := ps.get(0).Lexeme
		ps.consume(tokenizer.NAME)
		ps.consume(tokenizer.EQUAL)
		return &AssignNode{Variable: variable, Expression: ps.Expression(), Line: line}
	}

	return ps.Expression()
}

func (ps *Parser) Function() Node {
	targetFuncName := ps.get(0).Lexeme
	ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.LPAR)

	var args []Node
	res := FunctionalNode{args, targetFuncName, ps.get(0).Line}

	for !ps.match(tokenizer.RPAR) {
		res.arguments = append(res.arguments, ps.Expression())
		ps.match(tokenizer.COMMA)
	}

	return &res
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

	iterVar := ps.consume(tokenizer.NAME)
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

	var argNames []VarWithAnnotation

	for !ps.match(tokenizer.RPAR) {
		varName := ps.consume(tokenizer.NAME)
		varAnnotation := ps.consume(tokenizer.NAME)
		resultVar := VarWithAnnotation{varName.Lexeme, varAnnotation.Lexeme}
		ps.match(tokenizer.COMMA)
		argNames = append(argNames, resultVar)
	}

	returnAnnotation := "nil"

	if ps.get(0).Type == tokenizer.ARROW {
		ps.consume(tokenizer.ARROW)
		res := ps.consume(tokenizer.NAME)
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

func (ps *Parser) Array() Node {
	elementsTypeAnnotation := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.ARROW)

	res := ps.consume(tokenizer.LSQB)
	line := res.Line
	var elements []Node

	for !ps.match(tokenizer.RSQB) {
		elements = append(elements, ps.Expression())
		ps.match(tokenizer.COMMA)
	}

	return &Array{Elements: elements, elementsTypeAnnotation: elementsTypeAnnotation.Lexeme, Line: line}

}

func (ps *Parser) ArrayElement() Node {
	res := ps.consume(tokenizer.NAME)
	varName := res.Lexeme
	ps.consume(tokenizer.LSQB)
	idx := ps.Expression()
	ps.consume(tokenizer.RSQB)

	return &CollectionAccess{variableName: varName, index: idx, Line: res.Line}
}

func (ps *Parser) ClassDeclaring() Node {
	name := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.LBRACE)

	var fields []VarWithAnnotation
	methods := map[string]Node{}

	for !ps.match(tokenizer.RBRACE) {

		switch {
		case ps.lookahead(0, tokenizer.NAME):
			varName := ps.consume(tokenizer.NAME)
			varAnnotation := ps.consume(tokenizer.NAME)
			resultVar := VarWithAnnotation{varName.Lexeme, varAnnotation.Lexeme}
			fields = append(fields, resultVar)
		case ps.lookahead(0, tokenizer.DECLARE):
			ps.consume(tokenizer.DECLARE)
			method := ps.FuncDeclaration()
			switch __T := method.(type) {
			case *FunctionDeclareNode:
				methods[__T.name] = method
			default:
				panic("incorrect class method declaring")
			}
		default:
			panic("incorrect entity at class context")
		}

	}

	return &Class{
		className: name.Lexeme,
		methods:   methods,
		fields:    fields,
		Line:      ps.get(0).Line,
	}

}

func (ps *Parser) ClassAccess() Node {
	line := ps.get(0).Line
	structName := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.DOT)
	structField := ps.consume(tokenizer.NAME)

	if ps.lookahead(0, tokenizer.EQUAL) {
		ps.consume(tokenizer.EQUAL)
		stmt := ps.Expression()
		return &ClassAccessLHS{
			structName:  structName.Lexeme,
			stmt:        stmt,
			structField: structField.Lexeme,
			Line:        line,
		}
	}

	return &ClassAccessRHS{
		structName:  structName.Lexeme,
		structField: structField.Lexeme,
		Line:        line,
	}
}

func (ps *Parser) ClassMethod() Node {
	className := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.DOT)
	methodToExecute := ps.consume(tokenizer.NAME)
	ps.consume(tokenizer.LPAR)

	var args []Node
	res := ClassScopeMethodAccess{
		objName:         className.Lexeme,
		methodToExecute: methodToExecute.Lexeme,
		arguments:       args,
		Line:            className.Line,
	}

	for !ps.match(tokenizer.RPAR) {
		res.arguments = append(res.arguments, ps.Expression())
		ps.match(tokenizer.COMMA)
	}

	return &res
}
