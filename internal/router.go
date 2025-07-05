package internal

import (
	"errors"
	"strings"
)

type CommandRouter struct {
	table map[string]CommandHandler
}

func (c *CommandRouter) Register(name string, handler CommandHandler) {
	cmd := strings.ToLower(name)
	c.table[cmd] = handler
}

func (c *CommandRouter) Handle(name string, w *ResponseWriter, args ...Resp) error {
	cmd := strings.ToLower(name)
	h, ok := c.table[cmd]
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
