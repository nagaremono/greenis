package command

import (
	"errors"

	"greenis/internal"
)

type EchoCommand struct{}

func (h EchoCommand) Handle(c *internal.Context) error {
	if len(c.Params) != 1 {
		return errors.New("invalid args count")
	}
	strArg, ok := c.Params[0].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}

	err := c.W.Write(strArg)
	if err != nil {
		return err
	}

	return nil
}
