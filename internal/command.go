package internal

type Context struct {
	Params []Resp
	W      *ResponseWriter
}

type CommandHandler interface {
	Handle(c *Context) error
}
