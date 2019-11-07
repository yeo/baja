package baja

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo"

	"github.com/yeo/baja/utils"
)

type Server struct {
	staticPath string
}

func router(e *echo.Echo, s *Server) {
	e.Static("/deploy", s.staticPath)
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
	w := utils.Watch([]string{"./content", "./themes"})

	// Build our site immediately to serve dev
	go Build()

	go func() {
		for {
			select {
			case event := <-w.Event:
				color.Yellow("Receive file change event %s. Rebuild", event)
				Build()
			case err := <-w.Error:
				color.Red("Watch error:%s", err)
			case <-w.Closed:
				return
			}
		}
	}()

	go func() {
		// Start the watching process - it'll check for changes every 100ms.
		if err := w.Start(time.Millisecond * 100); err != nil {
			log.Fatalln(err)
		}
	}()

	Run(addr, directory)
	return 0
}
