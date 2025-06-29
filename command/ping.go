package command

import "github.com/nagaremono/greenis/internal"

type PingCommand struct{}

func (h PingCommand) Handle(c *internal.Context) error {
	output := "+PONG\r\n"

	c.Output.WriteString(output)
	return nil
}
