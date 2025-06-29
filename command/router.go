package command

import "github.com/nagaremono/greenis/internal"

func InitRouter() *internal.CommandRouter {
	r := internal.NewRouter()

	r.Register("PING", &PingCommand{})
	r.Register("ECHO", &EchoCommand{})

	return r
}
