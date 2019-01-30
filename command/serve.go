package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type ServeCmd struct {
}

func (c *ServeCmd) Run(args []string) int {
	fmt.Println("Run server")
	return baja.Serve("localhost:2803", "./public")
}
