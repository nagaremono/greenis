package internal

import "bytes"

type Context struct {
	Params []Resp
	Output *bytes.Buffer
}

type CommandHandler interface {
	Handle(c *Context) error
}
