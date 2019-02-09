package command

import (
	"fmt"
	"os"
)

type CmdRunner interface {
	Run(args []string) int
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
	if len(args) <= 1 {
		fmt.Println("baja init name")
		fmt.Println("baja node path/to/content")
		os.Exit(1)
	}

	if runner := registries[args[1]]; runner != nil {
		if len(args) >= 3 {
			fmt.Println("Run with", args[2:])
			os.Exit(registries[args[1]].Run(args[2:]))
		} else {
			fmt.Println("Run with no argument")
			os.Exit(registries[args[1]].Run([]string{}))
		}
	} else {
		fmt.Println("Unknow command", args)
		os.Exit(255)
	}
}
