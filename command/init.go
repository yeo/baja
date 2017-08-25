package command

import (
	"fmt"
)

type InitCmd struct {
}

func (c *InitCmd) Run() int {
	fmt.Println("Init")

	return 0
}
