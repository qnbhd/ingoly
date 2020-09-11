package parser

type BlockContext struct {
	Vars map[string]Node
}

func NewBlockContext() *BlockContext {
	return &BlockContext{Vars: make(map[string]Node)}
}
