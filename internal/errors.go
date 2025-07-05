package internal

import (
	"fmt"
)

var ErrInvalidArgs *InvalidArgsError

type InvalidArgsError struct {
	Command string
	Args    []Resp
	Err     error
}

func (i *InvalidArgsError) Error() string {
	msg := fmt.Sprintf("Invalid args received for command: %s, args: %v ", i.Command, i.Args)
	return msg
}

func (i *InvalidArgsError) Unwrap() error {
	return i.Err
}
