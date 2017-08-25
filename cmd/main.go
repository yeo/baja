package main

import (
	"github.com/yeo/baja/command"
	"os"
)

func main() {
	command.Register("init", &command.InitCmd{})
	command.Process(os.Args)
}
