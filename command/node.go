package command

import (
	"fmt"
)

type NodeCmd struct {
}

func (c *NodeCmd) Run() int {
	fmt.Println("Create node")

	return 0
}
