package command

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type CmdRunner interface {
	Run(args []string) int
}

type CmdHandler func([]string) CmdRunner

var registries map[string]CmdRunner

func init() {
	registries = map[string]CmdRunner{}
}

func Register(runner CmdRunner) {
	fullCmdName := ""
	if t := reflect.TypeOf(runner); t.Kind() == reflect.Ptr {
		fullCmdName = t.Elem().Name()
	} else {
		fullCmdName = t.Name()
	}

	cmd := strings.ToLower(fullCmdName[:len(fullCmdName)-len("Cmd")])
	registries[cmd] = runner
}

func printHelp() {
	fmt.Println("Usage")
	fmt.Println("baja command [param 1]...[param N]")
	fmt.Println("  init name to init a new one")
	fmt.Println("  node path/to/content to create new node")
	fmt.Println("Default to build")
}

func Process(args []string) {
	if len(args) <= 1 {
		printHelp()
		registries["build"].Run(args)

		os.Exit(0)
	}

	command := args[1]

	if runner := registries[command]; runner != nil {

		params := []string{}
		if len(args) >= 3 {
			params = args[2:]
		}
		os.Exit(registries[args[1]].Run(params))
	} else {
		fmt.Println("Unknow command", args)
		os.Exit(255)
	}
}
