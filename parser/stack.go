package parser

type Stack []Node

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val Node) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (Node, bool) {
	if s.IsEmpty() {
		return &Boolean{value: false, Line: 0}, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
