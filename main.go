package main

import (
	"fmt"
	"ingoly/utils/parser"
	"ingoly/utils/tokenizer"
	"io/ioutil"
	"os"
)

func main() {

	fileInput := os.Args[1:][0]

	if fileInput == "" {
		return
	}

	data, err := ioutil.ReadFile(fileInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	lx := tokenizer.New(string(data))
	result, lexErrorPull := lx.Tokenize()

	if !lexErrorPull.IsEmpty() {
		lexErrorPull.Print()
		return
	}

	var jp parser.Parser
	ps := jp.New(result)
	ast, parserErrorPull := ps.Parse()

	if !parserErrorPull.IsEmpty() {
		parserErrorPull.Print()
		return
	}

	ast.Print()

	semanticPullError := ast.Execute()

	semanticPullError.Print()
}
