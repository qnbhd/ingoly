package tokenizer

type TokenType int

const (
	NIL = iota
	EOF
	NAME
	NUMBER
	STRING
	LPAR
	RPAR
	LSQB
	RSQB
	COLON
	COMMA
	SEMI
	PLUS
	MINUS
	STAR
	SLASH
	VBAR
	VBARVBAR
	AMPER
	AMPERAMPER
	LESS
	GREATER
	EQUAL
	EXCL
	DOT
	PERCENT
	LBRACE
	RBRACE
	EQEQUAL
	NOTEQUAL
	LESSEQUAL
	GREATEREQUAL
	TILDE
	CIRCUMFLEX
	DOUBLESTAR
	PLUSEQUAL
	MINEQUAL
	STAREQUAL
	SLASHEQUAL
	PERCENTEQUAL
	AMPEREQUAL
	VBAREQUAL
	COLONEQUAL
	CIRCUMFLEXEQUAL
	IF
	ELSE
	FOR
	WHILE
	VAR
	RANGE
	IN
	TRUE
	FALSE
	BREAK
	CONTINUE
	DECLARE
	RETURN
	ARROW
	REQUIRE
	REQUIRESTRING
	CLASS
)

func (tt TokenType) String() string {
	return [...]string{"NIL", "EOF", "NAME", "NUMBER", "STRING", "LPAR", "RPAR",
		"LSQB", "RSQB", "COLON", "COMMA", "SEMI", "PLUS", "MINUS", "STAR", "SLASH",
		"VBAR", "VBARVBAR", "AMPER", "AMPERAMPER", "LESS", "GREATER", "EQUAL",
		"EXCL", "DOT", "PERCENT", "LBRACE", "RBRACE", "EQEQUAL", "NOTEQUAL",
		"LESSEQUAL", "GREATEREQUAL", "TILDE", "CIRCUMFLEX", "DOUBLESTAR",
		"PLUSEQUAL", "MINEQUAL", "STAREQUAL", "SLASHEQUAL", "PERCENTEQUAL",
		"AMPEREQUAL", "VBAREQUAL", "COLONEQUAL", "CIRCUMFLEXEQUAL", "IF", "ELSE", "FOR", "WHILE",
		"VAR", "RANGE", "IN", "TRUE", "FALSE", "BREAK", "CONTINUE", "DECLARE", "RETURN", "ARROW",
		"REQUIRE", "REQUIRESTRING", "CLASS"}[tt]
}
