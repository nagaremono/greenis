package internal

import (
	"errors"
)

type CommandRouter struct {
	table map[string]CommandHandler
}

func (c *CommandRouter) Register(name string, handler CommandHandler) {
	c.table[name] = handler
}

func (c *CommandRouter) Handle(name string, w *ResponseWriter, args ...Resp) error {
	h, ok := c.table[name]
	if !ok {
		return errors.New("handler not found")
	}

	ctx := &Context{
		Params: args,
		W:      w,
	}

	err := h.Handle(ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewRouter() *CommandRouter {
	r := &CommandRouter{
		table: make(map[string]CommandHandler),
	}

	return r
}
