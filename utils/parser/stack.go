package parser

type Stack []Value

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val Value) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (Value, bool) {
	if s.IsEmpty() {
		return NumberValue{0}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
