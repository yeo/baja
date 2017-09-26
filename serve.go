package baja

import (
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
	watcher := NewWatcher(cwd)
	go watcher.Run()
	Run(addr, directory)
	return 0
}
