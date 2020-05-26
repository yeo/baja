package render

import (
	"github.com/yeo/baja"
)

type Command struct{}

func (cmd *Command) ArgDesc() string {
	return ""
}

func (cmd *Command) Help() string {
	return "Render markdown into html for deploy. HTML content is written to public directory"
}

func (cmd *Command) Run(site *baja.Site, args []string) int {
	return Build(site)
}
