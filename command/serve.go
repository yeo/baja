package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type ServeCmd struct {
}

func (c *ServeCmd) Run(args []string) int {
	fmt.Println("Run server")

	addr := "0.0.0.0"
	if args[0] != nil {
		addr = args[0]
	}
	return baja.Serve(fmt.Sprintf("%s:2803", addr), "./public")
}
