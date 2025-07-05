package command

import "greenis/internal"

type PingCommand struct{}

func (h PingCommand) Handle(c *internal.Context) error {
	err := c.W.Write(internal.RespSString("PONG"))
	if err != nil {
		return err
	}

	return nil
}
