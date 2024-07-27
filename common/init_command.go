package baja

import (
	"fmt"
	"os"
	"path/filepath"
)

type InitCommand struct {
}

func (cmd *InitCommand) ArgDesc() string {
	return "site-name"
}

func (cmd *InitCommand) Help() string {
	return "Setup skeleton for a new project"
}

func (cmd *InitCommand) Run(s *Site, args []string) int {
	if err := Create(s.Meta.Name); err == nil {
		fmt.Println("Error when creating site", err)
		return 1
	}

	return 0
}

// Initalize a new blog directory
func Create(name string) error {
	path := []string{
		filepath.Join(".", name),
		filepath.Join(".", name, "content"),
		filepath.Join(".", name, "theme/baja"),
		filepath.Join(".", name, "public/asset"),
	}

	for _, p := range path {
		os.MkdirAll(p, os.ModePerm)
	}

	c := NewConfig("./" + name + "/baja.yaml")
	c.WriteFile()
	return nil
}
