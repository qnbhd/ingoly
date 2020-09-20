package tokenizer

import (
	"errors"
	"fmt"
	"ingoly/errpull"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	//"os"
)

type Requirer struct {
	ErrorsPull *errpull.ErrorsPull
}

func getPathFromPackageName(requireString string) string {
	words := strings.Split(requireString, ".")
	res, _ := os.Getwd()
	targetPath := filepath.Join(res, filepath.Join(words...))
	return targetPath + ".inly"
}

func NewRequirer() *Requirer {
	requirer := Requirer{
		ErrorsPull: errpull.NewErrorsPull(),
	}
	return &requirer
}

func (rq *Requirer) Require(tokens []Token) ([]Token, *errpull.ErrorsPull) {
	requireFlag := false

	for idxToInsert, item := range tokens {
		if item.Type == REQUIRE {
			requireFlag = true
			continue
		}
		if requireFlag {
			if item.Type != REQUIRESTRING {
				err := errors.New("name of package expected at line")
				inn := errpull.NewInnerError(err, item.Line)
				rq.ErrorsPull.Errors = append(rq.ErrorsPull.Errors, inn)
				requireFlag = false
				continue
			}
			packageToRequire := item.Lexeme
			packagePath := getPathFromPackageName(packageToRequire)
			fmt.Println(packagePath)
			if _, err := os.Stat(packagePath); err == nil {
				data, err := ioutil.ReadFile(packagePath)

				if err != nil {
					fmt.Println(err)
					return tokens, rq.ErrorsPull
				}

				lastLine := item.Line
				fmt.Println(lastLine)
				lx := NewLexer(string(data))
				result, lexErrorPull := lx.Tokenize()

				for i := 0; i < len(result); i++ {
					result[i].Line += lastLine - 1
				}

				lastLine = result[len(result)-1].Line - 1

				rq.ErrorsPull.Errors = append(rq.ErrorsPull.Errors, lexErrorPull.Errors...)

				var rightPart []Token

				for _, item := range tokens[idxToInsert+1:] {
					rightPart = append(rightPart, item)
				}

				for i := 0; i < len(rightPart); i++ {
					rightPart[i].Line += lastLine - 1
				}

				temp := append(tokens[:idxToInsert-1], result[:len(result)-1]...)
				temp = append(temp, rightPart...)
				return temp, rq.ErrorsPull
			}
		}
	}

	return tokens, rq.ErrorsPull
}
