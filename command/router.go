package command

import "github.com/nagaremono/greenis/internal"

func InitRouter() *internal.CommandRouter {
	r := internal.NewRouter()

	r.Register("PING", &PingCommand{})
	r.Register("ECHO", &EchoCommand{})
	r.Register("GET", &GetCommand{})
	r.Register("SET", &SetCommand{})

	return r
}
