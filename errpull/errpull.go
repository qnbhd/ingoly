package errpull

import (
	"fmt"
	"github.com/fatih/color"
)

type InnerError struct {
	Err        error
	SourceLine int
}

func NewInnerError(err error, line int) InnerError {
	return InnerError{err, line}
}

type ErrorsPull struct {
	Errors []InnerError
}

func NewErrorsPull() *ErrorsPull {
	return &ErrorsPull{Errors: []InnerError{}}
}

func (ep *ErrorsPull) Print() {
	for _, err := range ep.Errors {
		errTrace := fmt.Sprintf("%q", err.Err)
		errTrace = errTrace[1 : len(errTrace)-1]
		color.HiRed("[!] Error %s nearby %d line \n", errTrace, err.SourceLine)
	}
}

func (ep *ErrorsPull) GetLength() int {
	return len(ep.Errors)
}

func (ep *ErrorsPull) IsEmpty() bool {
	return ep.GetLength() == 0
}
