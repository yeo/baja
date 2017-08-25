package command

import (
	"fmt"
	"os"
)

type CmdRunner interface {
	Run() bool
}

type CmdHandler func([]string) CmdRunner

var registries map[string]CmdRunner

func Register(cmd string, runner CmdRunner) {
	registries[cmd] = runner
}

func Process(args []string) {
	if len(args) == 0 {
		fmt.Println("baja init name")
		fmt.Println("baja node path/to/content")
		os.Exit(1)
	}
}
