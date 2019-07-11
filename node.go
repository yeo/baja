package baja

import (
	"bufio"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
	"github.com/russross/blackfriday"
	//"github.com/microcosm-cc/bluemonday"
)

// NewNode creates a *Node object from a path
func NewNode(path string) *Node {
	n := Node{Path: path}

	// Remove content from path to get base directory
	n.BaseDirectory = strings.Join(strings.Split(filepath.Dir(path), "/")[1:], "/")

	filename := filepath.Base(path)
	dotPosition := strings.LastIndex(filename, ".")
	n.Name = filename[0:dotPosition]

	n.Parse()
	n.FindTheme(DefaultConfig())

	return &n
}

// Parse reads the markdown and parse metadata and generate html
func (n *Node) Parse() {
	content, err := ioutil.ReadFile(n.Path)
	if err != nil {
		log.Error().Err(err).Str("Node", n.Path).Str("Message", "Cannot read node")

		return
	}

	part := strings.Split(string(content), "+++")
	if len(part) < 3 {
		log.Fatal().Str("path", n.Path).Msg("Not enough header/body")
	}

	n.Meta = &NodeMeta{}
	toml.Decode(string(part[1]), n.Meta)

	n.Meta.DateFormatted = n.Meta.Date.Format("2006 Jan 02")
	n.Meta.Category = n.BaseDirectory

	n.Body = template.HTML(part[2])
}

func (n *Node) IsPage() bool {
	return n.Meta.Type == "page"
}

func (n *Node) Permalink() string {
	if n.BaseDirectory == "" {
		return "/" + filepath.Base(n.Name) + "/"
	} else {
		return "/" + n.BaseDirectory + "/" + filepath.Base(n.Name) + "/"
	}
}

func (n *Node) data() map[string]interface{} {
	html := blackfriday.Run([]byte(n.Body))

	return map[string]interface{}{
		"Meta":      n.Meta,
		"Body":      template.HTML(html),
		"Permalink": n.Permalink(),
	}
}

func (n *Node) FindTheme(c *Config) {
	// Find theme
	theme := GetTheme(DefaultConfig())
	pathComponents := strings.Split(n.BaseDirectory, "/")
	n.templatePaths = []string{"themes/" + c.Theme + "/layout/default.html"}
	lookupPath := "themes/" + c.Theme
	for _, p := range pathComponents {
		if _, err := os.Stat(lookupPath + "/node.html"); err == nil {
			n.templatePaths = append(n.templatePaths, lookupPath+"/node.html")
		}

		if _, err := os.Stat(lookupPath + "/" + n.Name + ".html"); err == nil {
			n.templatePaths = append(n.templatePaths, lookupPath+"/"+n.Name+".html")
		}

		lookupPath = lookupPath + "/" + p
	}

	if n.Meta.Theme != "" {
		n.templatePaths = append(n.templatePaths, theme.NodePath(n.Meta.Theme))
	}
}

func (n *Node) Compile() {
	directory := "public/" + n.BaseDirectory + "/" + n.Name
	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Error().Err(err).Str("Directory", directory).Msg("Cannot create index file in directory")
	}

	w := bufio.NewWriter(f)

	tpl := template.New("layout").Funcs(FuncMaps())
	tpl, err = tpl.ParseFiles(n.templatePaths...)
	if err != nil {
		log.Panic().Err(err)
	}

	if err := tpl.Execute(w, n.data()); err != nil {
		log.Panic().Str("Node", n.Name).Err(err).Msg("Fail to render node")
	}

	w.Flush()
}
