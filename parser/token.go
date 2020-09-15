package parser

import "fmt"

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
}

func (tk Token) ToString() string {
	return fmt.Sprintf("[ %s, %s, %d]", tk.Type.String(), tk.Lexeme, tk.Line)
}
