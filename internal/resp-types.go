package internal

import (
	"fmt"
	"strconv"
)

type Resp interface {
	Encode() ([]byte, error)
	String() string
}

type (
	RespSString     string
	RespBString     string
	RespArray       []Resp
	RespNullBString string
	RespInt         int
	RespBoolean     bool
)

var NullBString RespNullBString = ""

func (r RespBString) Encode() ([]byte, error) {
	s := string(r)
	length := len(s)
	str := fmt.Sprintf("$%d\r\n%s\r\n", length, s)

	return []byte(str), nil
}

func (r RespBString) String() string {
	return string(r)
}

func (r RespSString) Encode() ([]byte, error) {
	// str := "+" + strconv.Itoa(length) + "\r\n" + s + "\r\n"
	str := fmt.Sprintf("+%s\r\n", r)

	return []byte(str), nil
}

func (r RespSString) String() string {
	return string(r)
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

func (r RespArray) String() string {
	var s string
	for _, v := range r {
		s += v.String()
	}

	return s
}

func (r RespNullBString) Encode() ([]byte, error) {
	return []byte("$-1\r\n"), nil
}

func (r RespNullBString) String() string {
	return string(r)
}

func (r RespInt) Encode() ([]byte, error) {
	str := fmt.Sprintf(":%d\r\n", r)
	return []byte(str), nil
}

func (r RespInt) String() string {
	return strconv.Itoa(int(r))
}

func (r RespBoolean) Encode() ([]byte, error) {
	encStr := ""
	if r {
		encStr = "t"
	} else {
		encStr = "f"
	}

	encStr = fmt.Sprintf("#%s\r\n", encStr)
	return []byte(encStr), nil
}

func (r RespBoolean) String() string {
	if r {
		return "true"
	}
	return "false"
}
