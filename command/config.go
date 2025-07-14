package command

import (
	"errors"
	"greenis/internal"
	"strings"
)

type ConfigCommand struct{}

func (h ConfigCommand) Handle(c *internal.Context) error {
	if len(c.Params) < 2 {
		return errors.New("invalid args count")
	}
	sub, ok := c.Params[0].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}

	switch strings.ToLower(string(sub)) {
	case "get":
		return getConfigCmd.Handle(c)
	default:
		return errors.New("unknown sub command")
	}

	return nil
}
