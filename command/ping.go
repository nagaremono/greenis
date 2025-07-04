package command

import "github.com/nagaremono/greenis/internal"

type PingCommand struct{}

func (h PingCommand) Handle(c *internal.Context) error {
	err := c.W.Write(internal.RespSString("PONG"))
	if err != nil {
		return err
	}

	return nil
}
