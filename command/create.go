package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type CreateCmd struct {
}

func (c *CreateCmd) Run(args []string) int {
	fmt.Println("Add a draft node", args)
	if baja.CreateNode(args[0], args[1]) == nil {
		return 0
	}

	return 1
}
