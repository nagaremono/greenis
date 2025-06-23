package command

type PingCommand struct{}

func (h PingCommand) Handle(c *Context) error {
	output := "+PONG\r\n"

	c.output.WriteString(output)
	return nil
}
