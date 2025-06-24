package main

import (
	"github.com/nagaremono/greenis/internal"
)

func main() {
	s := internal.NewServer()
	s.Listen()
	defer s.Close()

	for {
		c := s.NextConn()
		go func() {
			defer c.Close()

			err := s.HandleNext(c)
			if err != nil {
				panic(err)
			}
		}()
	}
}
