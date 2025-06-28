package main

import (
	"fmt"
	"os"

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
				fmt.Println(err)
				os.Exit(1)
			}
		}()
	}
}
