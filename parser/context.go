package parser

type Context struct {
	Vars      map[string]Node
	Functions map[string]func(w Executor, curNode Node, argCount, line int)
}

func NewContext() *Context {
	return &Context{Vars: make(map[string]Node),
		Functions: make(map[string]func(w Executor, s Node, argCount, line int))}
}
