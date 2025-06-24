package internal

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/nagaremono/greenis/command"
)

type Server struct {
	l net.Listener
	r *command.CommandRouter
}

func NewServer() *Server {
	return &Server{
		r: command.NewRouter(),
	}
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
	reader := bufio.NewReader(c)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Println("error reading next input")
			return err
		}
		// Start with hard coding the command
		input = strings.TrimSpace(input)
		if input != "PING" {
			continue
		}

		out, err := s.r.Handle(input)
		if err != nil {
			return err
		}

		_, err = c.Write(out)
		if err != nil {
			return err
		}
	}
}
