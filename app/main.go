package main

import (
	"github.com/nagaremono/greenis/internal"
)

func main() {
	s := internal.NewServer()
	s.Listen()
	defer s.Close()

	for {
		err := s.HandleNext()
		if err != nil {
			panic(err)
		}
	}
}
