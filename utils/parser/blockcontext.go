package parser

type BlockContext struct {
	Vars      map[string]Node
	Functions map[string]func(w Executor, curNode Node, argCount, line int)
}

func NewBlockContext() *BlockContext {
	return &BlockContext{Vars: make(map[string]Node),
		Functions: make(map[string]func(w Executor, s Node, argCount, line int))}
}
