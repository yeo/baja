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
	fmt.Printf("Baja %s.%s\n\n", AppVersion, GitCommit)

	command.Register(&command.InitCmd{})
	command.Register(&command.BuildCmd{})
	command.Register(&command.CleanCmd{})
	command.Register(&command.ServeCmd{})
	command.Register(&command.CreateCmd{})

	command.Process(os.Args)
}
