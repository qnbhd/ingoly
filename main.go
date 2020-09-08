package main

import (
	"fmt"
	"ingoly/utils/parser"
	"ingoly/utils/tokenizer"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("example.ig")

	if err != nil {
		fmt.Println(err)
	}

	lx := tokenizer.New(string(data))
	result := lx.Tokenize()
	var jp parser.Parser
	ps := jp.New(result)
	ast := ps.Parse()
	ast.PrintRecursive()

	for line, instruction := range ast.Tree {
		_, err := instruction.Execute()
		if err == nil {
			fmt.Print("Line Num: ")
			fmt.Print(line)
			fmt.Print(" successful executed")
			fmt.Println()
		}
	}

}
