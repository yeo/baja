package baja

import (
	"github.com/fatih/color"
	"github.com/radovskyb/watcher"
)

func Watch(dir []string) *watcher.Watcher {
	w := watcher.New()

	w.SetMaxEvents(1)

	for _, d := range dir {
		color.Yellow("Watch to build %s", d)
		if err := w.AddRecursive(d); err != nil {
			color.Red("Cannot watch %v", err)
		}
	}

	return w
}
