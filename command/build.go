package command

import (
	"fmt"
	"github.com/yeo/baja"
)

type BuildCmd struct {
}

func (c *BuildCmd) Run(args []string) int {
	fmt.Println("Build", args)
	return baja.Build()
}
