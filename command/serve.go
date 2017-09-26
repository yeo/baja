package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type ServeCmd struct {
}

func (c *SerbeCmd) Run(args []string) int {
	fmt.Println("Run server")
	return baja.Serve("./public")
}
