package internal

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type ParseRespError struct {
	RespType string
	Err      error
}

func (e *ParseRespError) Error() string { return e.Err.Error() }

type ParseFunc func(r *bufio.Reader) (Resp, error)

func Parse(r io.Reader) (Resp, error) {
	reader := bufio.NewReader(r)

	first, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch string(first) {
	case "+":
		return parseString(reader)
	case "$":
		return parseBulkString(reader)
	case "*":
		return parseArray(reader)
	}

	return nil, &ParseRespError{Err: errors.New("unhandled first byte")}
}

func parseString(r io.Reader) (RespSString, error) {
	s := bufio.NewScanner(r)
	ok := s.Scan()
	if !ok {
		return "", &ParseRespError{"sstring", errors.New("eof or unreadable")}
	}

	return RespSString(s.Text()), nil
}

func parseBulkString(r *bufio.Reader) (RespBString, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", &ParseRespError{"bstring", err}
	}
	_, err = strconv.Atoi(
		strings.TrimSuffix(str, "\r\n"),
	)
	if err != nil {
		return "", &ParseRespError{"bstring", err}
	}
	str, err = r.ReadString('\n')
	if err != nil {
		return "", &ParseRespError{"bstring", err}
	}

	return RespBString(strings.TrimSuffix(str, "\r\n")), nil
}

func parseArray(r *bufio.Reader) (RespArray, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return nil, &ParseRespError{"array", err}
	}

	count, err := strconv.Atoi(
		strings.TrimSuffix(str, "\r\n"),
	)
	if err != nil {
		return nil, &ParseRespError{"array", err}
	}

	var arr RespArray
	for range count {
		d, err := Parse(r)
		if err != nil {
			return nil, &ParseRespError{"array", err}
		}

		arr = append(arr, d)
	}

	return arr, nil
}
