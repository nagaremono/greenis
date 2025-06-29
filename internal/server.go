package internal

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

type Server struct {
	l net.Listener
	r *CommandRouter
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	s.l = l
}

func (s *Server) Close() {
	s.l.Close()
}

func (s *Server) NextConn() net.Conn {
	c, err := s.l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	return c
}

func (s *Server) HandleNext(c net.Conn) error {
	for {
		r, err := Parse(c)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		arr, ok := r.(RespArray)
		if !ok {
			return errors.New("unhandled type")
		}
		var cmd string

		switch v := arr[0].(type) {
		case RespSString:
			cmd = string(v)
		case RespBString:
			cmd = string(v)
		default:
			return errors.New("unhandled type")
		}

		fmt.Printf("command received: %s, args: %v\n", cmd, arr[1:])

		out, err := s.r.Handle(cmd)
		if err != nil {
			return err
		}

		_, err = c.Write(out)
		if err != nil {
			return err
		}
	}
}

func (s *Server) RegisterRouter(r *CommandRouter) {
	s.r = r
}
