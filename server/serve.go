package server

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo"

	"github.com/mholt/archiver"
	"github.com/yeo/baja/utils"
)

type Server struct {
	staticPath string
}

func KeyAuth() {

}

func Deploy(c echo.Context) {
	apikey := c.FormValue("apikey")

	if os.Getenv("APIKEY") != apikey {
		c.String(401, "Unauthrozied")
		return
	}

	file, err := c.FormFile("bundle")
	if err != nil {
		c.String(400, "Error")
		return
	}

	src, err := file.Open()
	if err != nil {
		c.String(400, "Error")
		return
	}

	defer src.Close()

	err = archiver.Unarchive("test.tar.gz", "test")
}

func router(e *echo.Echo, s *Server) {
	//e.Static("/deploy", Deploy)
	e.Static("/", s.staticPath)
}

func Run(addr, public string) {
	e := echo.New()
	s := &Server{
		staticPath: public,
	}
	router(e, s)

	hostname, _ := os.Hostname()
	log.Printf("Listen on http://%s:%d", hostname, 2803)
	e.Logger.Fatal(e.Start(addr))
}

// Build execute template and content to generate our real static conent
func Serve(addr, directory string) int {
	w := utils.Watch([]string{"./content", "./themes"})

	// Build our site immediately to serve dev
	//go Build()

	go func() {
		for {
			select {
			case event := <-w.Event:
				color.Yellow("Receive file change event %s. Rebuild", event)
				//Build()
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
