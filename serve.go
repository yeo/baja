package baja

import (
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"log"
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
	//watcher := NewWatcher(cwd)
	//go watcher.Run()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Cannot watch directory")
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				Build()
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./content")
	if err != nil {
		log.Fatal(err)
	}

	Run(addr, directory)
	return 0
}
