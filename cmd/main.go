package main

import (
	"fmt"
	"github.com/yeo/baja/command"
	"os"
)

var (
	GitCommit  string
	AppVersion string
)

func main() {
	fmt.Printf("Baja %s Reg %s\n\n", AppVersion, GitCommit)

	command.Register("init", &command.InitCmd{})
	command.Register("build", &command.BuildCmd{})
	command.Register("clean", &command.CleanCmd{})
	command.Register("serve", &command.ServeCmd{})
	command.Register("new", &command.CreateCmd{})

	command.Process(os.Args)
}
