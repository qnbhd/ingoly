package parser

type BlockContext struct {
	Vars map[string]ScopeVar
}

func NewBlockContext() *BlockContext {
	return &BlockContext{Vars: make(map[string]ScopeVar)}
}
