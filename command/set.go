package command

import (
	"fmt"
	"greenis/internal"
	"strconv"
	"time"
)

var CommandName string = "Set"

type SetCommand struct{}

func (h SetCommand) Handle(c *internal.Context) error {
	err := validateArgs(c)
	if err != nil {
		return err
	}

	err = internal.Store.Set(c.Params[0].String(), c.Params[1])
	if err != nil {
		return fmt.Errorf("failed to set value: %w", err)
	}

	if len(c.Params) > 2 {
		dur, ok := c.Params[3].(internal.RespBString)
		if !ok {
			err := c.W.Write(internal.NullBString)
			return &internal.InvalidArgsError{
				Command: CommandName,
				Args:    c.Params,
				Err:     err,
			}
		}

		expDur, err := time.ParseDuration(dur.String() + "ms")
		if err != nil {
			err := c.W.Write(internal.NullBString)
			return &internal.InvalidArgsError{
				Command: CommandName,
				Args:    c.Params,
				Err:     err,
			}
		}
		time.AfterFunc(expDur, func() {
			internal.Store.Delete(c.Params[0].String())
		})
	}

	err = c.W.Write(internal.RespSString("OK"))
	if err != nil {
		return err
	}

	return nil
}

func validateArgs(c *internal.Context) error {
	if len(c.Params) < 2 {
		err := c.W.Write(internal.NullBString)
		return &internal.InvalidArgsError{
			Command: CommandName,
			Args:    c.Params,
			Err:     err,
		}
	}

	_, ok := c.Params[0].(internal.RespBString)
	if !ok {
		err := c.W.Write(internal.NullBString)
		return &internal.InvalidArgsError{
			Command: CommandName,
			Args:    c.Params,
			Err:     err,
		}
	}

	if len(c.Params) > 2 {
		px, ok := c.Params[2].(internal.RespBString)
		if !ok || px != "px" {
			err := c.W.Write(internal.NullBString)
			return &internal.InvalidArgsError{
				Command: CommandName,
				Args:    c.Params,
				Err:     err,
			}
		}

		dur, ok := c.Params[3].(internal.RespBString)
		if !ok {
			err := c.W.Write(internal.NullBString)
			return &internal.InvalidArgsError{
				Command: CommandName,
				Args:    c.Params,
				Err:     err,
			}
		}

		_, err := strconv.ParseInt(string(dur), 10, 0)
		if err != nil {
			err := c.W.Write(internal.NullBString)
			return &internal.InvalidArgsError{
				Command: CommandName,
				Args:    c.Params,
				Err:     err,
			}
		}
	}

	return nil
}
