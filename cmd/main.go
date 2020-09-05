package main

import (
	"fmt"
	"os"

	"github.com/yeo/baja"
	"github.com/yeo/baja/cleaner"
	"github.com/yeo/baja/node"
	"github.com/yeo/baja/render"
	"github.com/yeo/baja/server"
)

type CmdRunner interface {
	Run(*baja.Site, []string) int
	ArgDesc() string
	Help() string
}

var (
	GitCommit  string
	AppVersion string
)

func printHelp(registries map[string]CmdRunner) {
	fmt.Println("Usage:")
	fmt.Println("  baja command [param 1]...[param N]")

	fmt.Println("\n\nCommands:")
	for name, cmd := range registries {
		fmt.Printf("  %s %s: %s\n", name, cmd.ArgDesc(), cmd.Help())
	}
	//fmt.Println("  node path/to/content to create new node")
}

func main() {
	fmt.Printf("Baja %s. Rev %s\n\n", AppVersion, GitCommit)

	registries := make(map[string]CmdRunner)
	registries["init"] = &baja.InitCommand{}
	registries["build"] = &render.Command{}
	registries["clean"] = &cleaner.Command{}
	registries["server"] = &server.ServerCommand{}
	registries["serve"] = registries["server"]
	registries["create"] = &node.CreateCommand{}

	os.Exit(process(registries, os.Args[1:]))
}

func process(registries map[string]CmdRunner, args []string) int {
	var command string

	if len(args) >= 1 {
		command = args[0]
	} else {
		printHelp(registries)
		return 1
	}

	runner := registries[command]

	if runner == nil {
		fmt.Println("Unknow command", args)
		printHelp(registries)
		return 255
	}

	params := args[1:]
	site := baja.LoadSite("./baja.yaml")
	return runner.Run(site, params)
}
