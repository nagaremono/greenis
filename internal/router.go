package internal

import (
	"bytes"
	"errors"
)

type CommandRouter struct {
	table map[string]CommandHandler
}

func (c *CommandRouter) Register(name string, handler CommandHandler) {
	c.table[name] = handler
}

func (c *CommandRouter) Handle(name string, args ...Resp) ([]byte, error) {
	h, ok := c.table[name]
	if !ok {
		return nil, errors.New("handler not found")
	}

	ctx := &Context{
		Params: args,
		Output: bytes.NewBuffer([]byte{}),
	}

	err := h.Handle(ctx)
	if err != nil {
		return nil, err
	}

	return ctx.Output.Bytes(), nil
}

func NewRouter() *CommandRouter {
	r := &CommandRouter{
		table: make(map[string]CommandHandler),
	}

	return r
}
