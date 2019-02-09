package command

import (
	"github.com/yeo/baja"
)

type BuildCmd struct {
}

func (c *BuildCmd) Run(args []string) int {
	return baja.Build()
}
