package main

import (
	"fmt"
	"github.com/yeo/baja/command"
	"os"
)

var GitCommit string
var AppVersion string

func main() {
	fmt.Sprintf("Baja %s %s\n\n", GitCommit, AppVersion)

	command.Register("init", &command.InitCmd{})
	command.Register("build", &command.BuildCmd{})
	command.Register("clean", &command.CleanCmd{})
	command.Register("serve", &command.ServeCmd{})
	command.Register("new", &command.CreateCmd{})
	command.Process(os.Args)
}
