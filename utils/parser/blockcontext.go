package parser

type BlockContext struct {
	Vars map[string]Value
}

func NewBlockContext() *BlockContext {
	return &BlockContext{Vars: make(map[string]Value)}
}
