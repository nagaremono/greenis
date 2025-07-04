package internal

import (
	"fmt"
)

type Resp interface {
	Encode() ([]byte, error)
}

type (
	RespSString     string
	RespBString     string
	RespArray       []Resp
	RespNullBString string
)

var NullBString RespNullBString = ""

func (r RespBString) Encode() ([]byte, error) {
	s := string(r)
	length := len(s)
	str := fmt.Sprintf("$%d\r\n%s\r\n", length, s)

	return []byte(str), nil
}

func (r RespSString) Encode() ([]byte, error) {
	// str := "+" + strconv.Itoa(length) + "\r\n" + s + "\r\n"
	str := fmt.Sprintf("+%s\r\n", r)

	return []byte(str), nil
}

func (r RespArray) Encode() ([]byte, error) {
	str := fmt.Sprintf("*%d\r\n", len(r))
	for _, v := range r {
		s, err := v.Encode()
		if err != nil {
			return nil, err
		}
		str += string(s)
	}

	return []byte(str), nil
}

func (r RespNullBString) Encode() ([]byte, error) {
	return []byte("$-1\r\n"), nil
}
