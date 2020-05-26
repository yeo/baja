package server

import (
	"github.com/yeo/baja"
)

type ServerCommand struct{}

func (cmd *ServerCommand) ArgDesc() string {
	return "[address] [port]"
}

func (cmd *ServerCommand) Help() string {
	return "Run static server. Bind to address. Default is localhost only"
}

func (cmd *ServerCommand) Run(site *baja.Site, args []string) int {
	addr := "127.0.0.1:2803"
	if args != nil && len(args) > 0 && args[0] != "" {
		addr = args[0]

		if len(args) > 1 {
			addr = addr + ":" + args[1]
		} else {
			addr = addr + ":2803"
		}
	}
	return Serve(addr, "./public")
}
