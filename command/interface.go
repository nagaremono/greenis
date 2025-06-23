package command

import (
	"bytes"
)

type Context struct {
	params []any
	output *bytes.Buffer
}

type CommandHandler interface {
	Handle(c *Context) error
}
