package baja

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func Watch(dir string) (*fsnotify.Watcher, error) {
	w, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	err = w.Add(dir)

	_ = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			w.Add(path)
		}

		return nil
	})

	return w, nil
}
