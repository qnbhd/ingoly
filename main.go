package main

import (
	"fmt"
	"ingoly/utils/parser"
	"ingoly/utils/tokenizer"
)

func main() {
	lx := tokenizer.New(
		`5 + 7 - 6`)
	result := lx.Tokenize()
	var jp parser.Parser
	ps := jp.New(result)
	ast := ps.Parse()
	ast.PrintRecursive()

	for line, instruction := range ast.Tree {
		execResult, _ := instruction.Execute()
		fmt.Print("Line Num: ")
		fmt.Print(line)
		fmt.Println(" Result: ", execResult)
	}
}
