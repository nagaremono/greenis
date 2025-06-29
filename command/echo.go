package command

import (
	"errors"
	"strconv"

	"github.com/nagaremono/greenis/internal"
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

	argLen := len(string(strArg))
	str := "$" + strconv.Itoa(argLen) + "\r\n" + string(strArg) + "\r\n"

	c.Output.WriteString(str)
	return nil
}
