package main

import (
	"fmt"
	"ingoly/parser"
	"ingoly/parser/tokenizer"
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

	lx := tokenizer.NewLexer(string(data))
	result, lexErrorPull := lx.Tokenize()
	rq := tokenizer.NewRequirer()
	result, _ = rq.Require(result)

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

	indexedContext := ast.Index()

	semanticPullError := ast.Execute(indexedContext)
	semanticPullError.Print()
}
