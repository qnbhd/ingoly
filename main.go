package main

import (
	"fmt"
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
	}

	lx := tokenizer.New(string(data))
	result, lexErrorPull := lx.Tokenize()

	if !lexErrorPull.IsEmpty() {
		lexErrorPull.Print()
		return
	}

	fmt.Println(result)

	//
	//var jp parser.Parser
	//ps := jp.New(result)
	//ast := ps.Parse()
	//ast.Print()
	//ast.Execute()
}
