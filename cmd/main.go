package main

import (
	"github.com/yeo/baja/command"
	"os"
)

func main() {
	command.Register("init", &command.InitCmd{})
	command.Register("build", &command.BuildCmd{})
	command.Register("clean", &command.CleanCmd{})
	command.Process(os.Args)
}
