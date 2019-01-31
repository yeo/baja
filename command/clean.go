package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type CleanCmd struct {
}

func (c *CleanCmd) Run(args []string) int {
	baja.Clean()
	return 0
}
