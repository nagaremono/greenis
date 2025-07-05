package command

import (
	"errors"
	"fmt"
	"time"

	"greenis/internal"
)

type SetCommand struct{}

func (h SetCommand) Handle(c *internal.Context) error {
	if len(c.Params) < 2 {
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

	if len(c.Params) > 2 {
		px, ok := c.Params[2].(internal.RespBString)
		if !ok || px != "px" {
			return fmt.Errorf("unknown args passed: %v", px)
		}

		dur, ok := c.Params[3].(internal.RespBString)
		if !ok {
			return fmt.Errorf("unknown parameter type passed: %v", dur)
		}

		expDur, err := time.ParseDuration(string(dur) + "ms")
		if err != nil {
			return fmt.Errorf("failure to parse expiry duration: %w", err)
		}
		time.AfterFunc(expDur, func() {
			internal.Store.Delete(string(key))
		})
	}

	err = c.W.Write(internal.RespSString("OK"))
	if err != nil {
		return err
	}

	return nil
}
