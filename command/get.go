package command

import (
	"errors"

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
		err := c.W.Write(internal.NullBString)
		if err != nil {
			return err
		}
		return nil
	}

	switch val.(type) {
	case internal.RespBString:
	case internal.RespSString:
	default:
		return errors.New("unsupported type")

	}

	err := c.W.Write(val)
	if err != nil {
		return err
	}

	return nil
}
