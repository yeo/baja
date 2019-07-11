package baja

import (
	"os"
	"path/filepath"

	"github.com/yeo/baja/cfg"
)

// Initalize a new blog directory
func Setup(name string) error {
	path := []string{
		filepath.Join(".", name),
		filepath.Join(".", name, "content"),
		filepath.Join(".", name, "theme/baja"),
		filepath.Join(".", name, "public/asset"),
	}

	for _, p := range path {
		os.MkdirAll(p, os.ModePerm)
	}

	c := cfg.NewConfig("./" + name + "/baja.yaml")
	c.WriteFile()
	return nil
}
