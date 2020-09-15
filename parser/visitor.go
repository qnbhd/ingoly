package parser

type Visitor interface {
	EnterNode(node Node) bool
	LeaveNode(node Node)
}
