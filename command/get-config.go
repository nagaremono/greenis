package command

import (
	"errors"
	"greenis/internal"
)

var getConfigCmd = &GetConfigCommand{}

type GetConfigCommand struct{}

func (h GetConfigCommand) Handle(c *internal.Context) error {
	if len(c.Params) != 2 {
		return errors.New("invalid args count")
	}
	keyword, ok := c.Params[0].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}
	if string(keyword) != "get" {
		return errors.New("unknown command")
	}
	key, ok := c.Params[1].(internal.RespBString)
	if !ok {
		return errors.New("invalid args type")
	}

	var val string

	switch key {
	case "dir":
		val = internal.GetRDBConfig().GetDir()
	case "dbfilename":
		val = internal.GetRDBConfig().GetFile()
	default:
		return errors.New("unknown config")
	}

	err := c.W.Write(internal.RespArray{
		key,
		internal.RespBString(val),
	})
	if err != nil {
		return err
	}

	return nil
}
