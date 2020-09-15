package tokenizer

import (
	"errors"
	"ingoly/errpull"
	"strings"
	"unicode"
)

const PIX = "+-*(){}[]:;/=<>!&|"

type Reserved struct {
	Operators map[string]TokenType
}

func GetReservedOperators() *Reserved {
	var ops Reserved
	ops.Operators = map[string]TokenType{}

	ops.Operators["+"] = PLUS
	ops.Operators["-"] = MINUS
	ops.Operators["*"] = STAR
	ops.Operators["/"] = SLASH
	ops.Operators["%"] = PERCENT
	ops.Operators["^"] = CIRCUMFLEX
	ops.Operators["("] = LPAR
	ops.Operators[")"] = RPAR
	ops.Operators["{"] = LBRACE
	ops.Operators["}"] = RBRACE
	ops.Operators["["] = LSQB
	ops.Operators["]"] = RSQB
	ops.Operators["="] = EQUAL
	ops.Operators["<"] = LESS
	ops.Operators[">"] = GREATER
	ops.Operators["."] = DOT
	ops.Operators[","] = COMMA
	ops.Operators[":"] = COLON
	ops.Operators[";"] = SEMI
	ops.Operators["~"] = TILDE

	ops.Operators["!"] = EXCL
	ops.Operators["&"] = AMPER
	ops.Operators["|"] = VBAR

	ops.Operators[":="] = COLONEQUAL
	ops.Operators["=="] = EQEQUAL
	ops.Operators["!="] = NOTEQUAL
	ops.Operators["<="] = LESSEQUAL
	ops.Operators[">="] = GREATEREQUAL
	ops.Operators["**"] = DOUBLESTAR
	ops.Operators["||"] = VBARVBAR

	ops.Operators["+="] = PLUSEQUAL
	ops.Operators["-="] = MINEQUAL
	ops.Operators["*="] = STAREQUAL
	ops.Operators["/="] = SLASHEQUAL
	ops.Operators["%="] = PERCENTEQUAL
	ops.Operators["&="] = AMPEREQUAL
	ops.Operators["|="] = VBAREQUAL
	ops.Operators["^="] = CIRCUMFLEXEQUAL

	ops.Operators["&&"] = AMPERAMPER
	ops.Operators["->"] = ARROW

	return &ops
}

func GetReservedKeywords() *Reserved {
	var ops Reserved
	ops.Operators = map[string]TokenType{}

	ops.Operators["if"] = IF
	ops.Operators["else"] = ELSE
	ops.Operators["for"] = FOR
	ops.Operators["var"] = VAR
	ops.Operators["range"] = RANGE
	ops.Operators["in"] = IN
	ops.Operators["true"] = TRUE
	ops.Operators["false"] = FALSE
	ops.Operators["break"] = BREAK
	ops.Operators["continue"] = CONTINUE
	ops.Operators["while"] = WHILE
	ops.Operators["declare"] = DECLARE
	ops.Operators["return"] = RETURN
	ops.Operators["nil"] = NIL

	return &ops
}

type Tokenizer struct {
	Input            []rune
	tokens           []Token
	reservedOps      *Reserved
	reservedKeywords *Reserved
	pos              int
	currentLine      int
	ErrorsPull       *errpull.ErrorsPull
}

func New(input string) *Tokenizer {
	return &Tokenizer{Input: []rune(input), tokens: []Token{},
		reservedOps:      GetReservedOperators(),
		reservedKeywords: GetReservedKeywords(),
		pos:              0,
		currentLine:      1,
		ErrorsPull:       errpull.NewErrorsPull()}
}

func (lx *Tokenizer) addToken(tokenType TokenType, lexeme string, line int) {
	lx.tokens = append(lx.tokens, Token{tokenType, lexeme, line})
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
			return errors.New("lexing: invalid number")
		}

	}

	lx.addToken(NUMBER, strings.Trim(buffer.String(), ` \`), lx.currentLine)
	return nil
}

func (lx *Tokenizer) tokenizeOperator() error {
	current := lx.peek(0)
	if current == '/' {
		if lx.peek(1) == '/' {
			lx.next()
			lx.next()
			err := lx.tokenizeComment()
			return err
		} else if lx.peek(1) == '*' {
			lx.next()
			lx.next()
			err := lx.tokenizeMultiLineComment()
			return err
		}
	}

	var buffer strings.Builder

	for {
		text := buffer.String()
		if _, ok := lx.reservedOps.Operators[text+string(current)]; !ok && !(text == "") {
			lx.addToken(lx.reservedOps.Operators[text], "", lx.currentLine)
			return nil
		}
		buffer.WriteRune(current)
		current = lx.next()
	}

}

func (lx *Tokenizer) tokenizeComment() error {
	current := lx.peek(0)

	for current != '\r' && current != '\n' && current != '\x00' {
		current = lx.next()
	}

	return nil

}

func (lx *Tokenizer) tokenizeMultiLineComment() error {
	current := lx.peek(0)

	for {
		if current == '\x00' {
			return errors.New("lexing: missing closing tag in multi-line comment")
		}
		if current == '*' && lx.peek(1) == '/' {
			lx.next()
			lx.next()
			break
		}
		current = lx.next()
	}

	return nil

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

	if result, ok := lx.reservedKeywords.Operators[word]; ok {
		lx.addToken(result, "", lx.currentLine)
		return nil
	}
	lx.addToken(NAME, builder.String(), lx.currentLine)

	return nil
}

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

	lx.addToken(STRING, builder.String(), lx.currentLine)
	return nil
}

func (lx *Tokenizer) Tokenize() ([]Token, *errpull.ErrorsPull) {
	for lx.pos < lx.Length() {
		current := lx.peek(0)

		if current == '\n' {
			lx.currentLine++
		}

		var err error
		if unicode.IsDigit(current) {
			err = lx.tokenizeNumber()
		} else if unicode.IsLetter(current) {
			err = lx.tokenizeWord()
		} else if current == '"' {
			err = lx.tokenizeText()
		} else if strings.Index(PIX, string(current)) != -1 {
			err = lx.tokenizeOperator()
		} else {
			lx.next()
		}

		if err != nil {
			inn := errpull.NewInnerError(err, lx.currentLine)
			lx.ErrorsPull.Errors = append(lx.ErrorsPull.Errors, inn)
		}

	}

	lx.addToken(EOF, "", lx.currentLine)
	return lx.tokens, lx.ErrorsPull
}
