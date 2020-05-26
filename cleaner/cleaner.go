package cleaner

import (
	"fmt"
	"os"

	"github.com/yeo/baja"
)

type Command struct {
}

func (cmd *Command) ArgDesc() string {
	return ""
}

func (cmd *Command) Help() string {
	return "Clean output directory"
}

func (cmd *Command) Run(site *baja.Site, args []string) int {
	cleans := []string{"public"}

	for _, d := range cleans {
		fmt.Println("Clean", d)
		os.RemoveAll(fmt.Sprintf("./%s", d))
	}

	return 0
}
