package command

import (
	"fmt"
	"os"
)

type CmdRunner interface {
	Run() int
}

type CmdHandler func([]string) CmdRunner

var registries map[string]CmdRunner

func Register(cmd string, runner CmdRunner) {
	if registries == nil {
		registries = map[string]CmdRunner{}
	}

	registries[cmd] = runner
}

func Process(args []string) {
	if len(args) == 0 {
		fmt.Println("baja init name")
		fmt.Println("baja node path/to/content")
		os.Exit(1)
	}

	if runner := registries[args[1]]; runner != nil {
		os.Exit(registries[args[1]].Run())
	} else {
		fmt.Println("Unknow command")
		os.Exit(255)
	}
}
