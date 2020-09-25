package parser

import (
	"reflect"
	"strings"
)

type VarWithAnnotation struct {
	Name       string
	Annotation string
}

func typeof(v interface{}) string {
	result := reflect.TypeOf(v).String()
	if strings.HasPrefix(result, "*parser.") {
		result = result[len("*parser."):]
	}
	return result
}
