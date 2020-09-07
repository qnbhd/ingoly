package tokenizer

type Token struct {
	Type   TokenType
	Lexeme string
}

func (tk Token) ToString() string {
	return "[" + tk.Type.String() + ", " + tk.Lexeme + "]"
}
