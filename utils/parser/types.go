package parser

type IntNumber struct {
	value int
	Line  int
}

func (in *IntNumber) Walk(v Visitor) {
	if !v.EnterNode(in) {
		return
	}
	v.LeaveNode(in)
}

type FloatNumber struct {
	value float64
	Line  int
}

func (fn *FloatNumber) Walk(v Visitor) {
	if !v.EnterNode(fn) {
		return
	}
	v.LeaveNode(fn)
}

type Boolean struct {
	value bool
	Line  int
}

func (bn *Boolean) Walk(v Visitor) {
	if !v.EnterNode(bn) {
		return
	}
	v.LeaveNode(bn)
}

type String struct {
	value string
	Line  int
}

func (sn *String) Walk(v Visitor) {
	if !v.EnterNode(sn) {
		return
	}
	v.LeaveNode(sn)
}
