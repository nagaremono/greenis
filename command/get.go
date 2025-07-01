package command

import (
	"errors"
	"strconv"

	"github.com/nagaremono/greenis/internal"
)

type GetCommand struct{}

func (h GetCommand) Handle(c *internal.Context) error {
	if len(c.Params) != 1 {
		return errors.New("invalid args count")
	}
	key, ok := c.Params[0].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}

	val, ok := internal.Store.Get(string(key))
	if !ok {
		c.Output.WriteString("$-1\r\n")
		return nil
	}

	var str string
	switch v := val.(type) {
	case internal.RespBString:
		str = string(v)
	case internal.RespSString:
		str = string(v)
	default:
		return errors.New("unsupported type")

	}

	str = "$" + strconv.Itoa(len(str)) + "\r\n" + str + "\r\n"
	c.Output.WriteString(str)

	return nil
}
