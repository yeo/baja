package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type InitCmd struct {
}

func (c *InitCmd) Run(args []string) int {
	fmt.Println("Init", args)
	baja.Setup(args[0])
	return 0
}
