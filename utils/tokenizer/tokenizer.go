package tokenizer

import (
	"errors"
	"strings"
	"unicode"
)

const PIX = "+-*/=<>"

type Tokenizer struct {
	Input  []rune
	tokens []Token
	pos    int
}

func New(input string) *Tokenizer {
	return &Tokenizer{Input: []rune(input), tokens: []Token{}, pos: 0}
}

func (lx *Tokenizer) addToken(tokenType TokenType, lexeme string) {
	lx.tokens = append(lx.tokens, Token{tokenType, lexeme})
}

func (lx Tokenizer) Length() int {
	return len(lx.Input)
}

func (lx *Tokenizer) peek(relativePosition int) rune {
	pos := lx.pos + relativePosition
	if pos >= lx.Length() {
		return 0
	}
	result := lx.Input[pos]
	return result
}

func (lx *Tokenizer) next() rune {
	lx.pos++
	return lx.peek(0)
}

func (lx *Tokenizer) tokenizeNumber() error {
	current := lx.peek(0)
	var buffer strings.Builder

	for unicode.IsDigit(current) {
		buffer.WriteRune(current)
		current = lx.next()

		if current == '.' && (strings.Index(buffer.String(), ".")) == -1 {
			buffer.WriteRune(current)
			current = lx.next()
		} else if current == '.' && (strings.Index(buffer.String(), ".")) != -1 {
			return errors.New("invalid number")
		}

	}

	lx.addToken(NUMBER, strings.Trim(buffer.String(), ` \`))
	return nil
}

func (lx *Tokenizer) tokenizeOperator() error {
	current := lx.peek(0)
	tokenType := tokenOneSym(current)
	if tokenType != NIL {
		lx.addToken(tokenType, "")
		lx.next()
		return nil
	}
	return errors.New("invalid operator")
}

func (lx *Tokenizer) tokenizeWord() error {
	var builder strings.Builder
	current := lx.peek(0)
	for {
		if !(unicode.IsLetter(current) || unicode.IsDigit(current)) && current != '_' {
			break
		}
		builder.WriteRune(current)
		current = lx.next()
	}

	word := builder.String()
	switch word {
	case "print":
		lx.addToken(PRINT, "")
	case "if":
		lx.addToken(IF, "")
	case "else":
		lx.addToken(ELSE, "")
	default:
		lx.addToken(NAME, builder.String())
	}
	return nil
}

func tokenOneSym(sym rune) TokenType {
	switch sym {
	case '+':
		return PLUS
	case '-':
		return MINUS
	case '*':
		return STAR
	case '/':
		return SLASH
	case '(':
		return LPAR
	case ')':
		return RPAR
	case '=':
		return EQUAL
	case '>':
		return GREATER
	case '<':
		return LESS
	case ':':
		return COLON
	}
	return NIL
}

//func tokenTwoSym(twoSym string) TokenType {
//	switch twoSym {
//	case "+=":
//		return PLUSEQUAL
//	case "-=":
//		return MINEQUAL
//	}
//	return NIL
//}

func (lx *Tokenizer) tokenizeText() error {
	lx.next()

	var builder strings.Builder
	current := lx.peek(0)

	for {
		if current == '\\' {
			current = lx.next()
			switch current {
			case '"':
				current = lx.next()
				builder.WriteRune('"')
				continue
			case 'n':
				current = lx.next()
				builder.WriteRune('\n')
				continue
			case 't':
				current = lx.next()
				builder.WriteRune('\t')
				continue
			}
			builder.WriteRune('\\')
			continue
		}
		if current == '"' {
			break
		}
		builder.WriteRune(current)
		current = lx.next()
	}
	lx.next()

	lx.addToken(STRING, builder.String())
	return nil
}

func (lx *Tokenizer) Tokenize() []Token {
	for lx.pos < lx.Length() {
		current := lx.peek(0)

		if unicode.IsDigit(current) {
			_ = lx.tokenizeNumber()
		} else if unicode.IsLetter(current) {
			_ = lx.tokenizeWord()
		} else if current == '"' {
			_ = lx.tokenizeText()
		} else if strings.Index(PIX, string(current)) != -1 {
			_ = lx.tokenizeOperator()
		} else {
			lx.next()
		}

	}

	lx.addToken(EOF, "")
	return lx.tokens
}
