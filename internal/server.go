package internal

import (
	"fmt"
	"net"
	"os"

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

func (s *Server) HandleNext() error {
	c := s.NextConn()
	defer c.Close()

	// Start with hard coding the command
	out, err := s.r.Handle("PING")
	if err != nil {
		return err
	}

	_, err = c.Write(out)
	if err != nil {
		return err
	}

	return nil
}
