package main

import (
	"fmt"
	"greenis/command"
	"greenis/internal"
	"os"
)

func main() {
	s := internal.NewServer()
	s.RegisterRouter(command.InitRouter())
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
