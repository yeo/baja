package baja

import (
	"os"
	"path/filepath"
	//"github.com/yeo/baja/command"
)

// Initalize a new blog directory
func Setup(name string) error {
	root := filepath.Join(".", name)
	os.MkdirAll(root, os.ModePerm)
	content := filepath.Join(".", name, "content")
	os.MkdirAll(content, os.ModePerm)
	theme := filepath.Join(".", name, "theme")
	os.MkdirAll(theme, os.ModePerm)

	return nil
}
