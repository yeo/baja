package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type ServeCmd struct {
}

func (c *ServeCmd) Run(args []string) int {
	fmt.Println("Run server")

	addr := "127.0.0.1"
	if args != nil && len(args) > 0 && args[0] != "" {
		addr = args[0]
	}
	return baja.Serve(fmt.Sprintf("%s:2803", addr), "./public")
}
