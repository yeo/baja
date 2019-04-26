package baja

import (
	"log"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
)

type Server struct {
	staticPath string
}

func router(e *echo.Echo, s *Server) {
	e.Static("/", s.staticPath)
}

func Run(addr, public string) {
	e := echo.New()
	s := &Server{
		staticPath: public,
	}
	router(e, s)

	e.Logger.Fatal(e.Start(addr))
}

// Build execute template and content to generate our real static conent
func Serve(addr, directory string) int {
	watcher, err := Watch("./content")
	if err != nil {
		log.Fatalf("Cannot watch directory")
	}
	defer watcher.Close()

	// Build our site immediately to serve dev
	go Build()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				color.Yellow("Receive file change event %s", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					color.Yellow("  modified file:", event.Name)
				}
				Build()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				color.Red("Watch error:%s", err)
			}
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	Run(addr, directory)
	return 0
}
