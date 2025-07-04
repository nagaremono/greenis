package command

import (
	"errors"
	"fmt"

	"github.com/nagaremono/greenis/internal"
)

type SetCommand struct{}

func (h SetCommand) Handle(c *internal.Context) error {
	if len(c.Params) != 2 {
		return errors.New("invalid args count")
	}
	key, ok := c.Params[0].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}

	err := internal.Store.Set(string(key), c.Params[1])
	if err != nil {
		return fmt.Errorf("failed to set value: %w", err)
	}

	err = c.W.Write(internal.RespBString("OK"))
	if err != nil {
		return err
	}

	return nil
}
