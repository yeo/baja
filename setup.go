package baja

import (
	"os"
	"path/filepath"
	//"github.com/yeo/baja/command"
)

// Initalize a new blog directory
func Setup(name string) error {
	path := []string{
		filepath.Join(".", name),
		filepath.Join(".", name, "content"),
		filepath.Join(".", name, "theme"),
		filepath.Join(".", name, "public"),
	}

	for _, p := range path {
		os.MkdirAll(p, os.ModePerm)
	}

	return nil
}
