package baja

import (
	"bufio"
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"time"
	//"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
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
		log.Fatal("Cannot parse", n.Path)
	}

	part := strings.Split(string(content), "+++")
	if len(part) < 3 {
		log.Fatal("Not enough header/body", n.Path)
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
	pathComponents := strings.Split(n.BaseDirectory, "/")
	n.templatePaths = []string{"themes/" + c.Theme + "/layout/default.html"}
	lookupPath := "themes/" + c.Theme
	for _, p := range pathComponents {
		if _, err := os.Stat(lookupPath + "/node.html"); err == nil {
			n.templatePaths = append(n.templatePaths, lookupPath+"/node.html")
		}
		lookupPath = lookupPath + "/" + p
	}
}

func (n *Node) Compile() {
	directory := "public/" + n.BaseDirectory + "/" + n.Name
	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)

	tpl := template.New("layout")

	tpl, err = tpl.ParseFiles("themes/baja/layout/default.html", "themes/baja/node.html")
	if err != nil {
		log.Panic(err)
	}

	if err := tpl.Execute(w, n.data()); err != nil {
		log.Println("Fail to render", err)
	}

	w.Flush()
}

func CompileNodes(db *NodeDB) {
	// Build individual node
	color.Yellow("Start build html\nBuild individual page")
	for i, node := range db.NodeList {
		color.Yellow("\t%d/%d:  %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	current := &Current{
		IsHome: false,
		IsDir:  false,
		IsTag:  false,
	}
	// Now build the main index pag
	current.IsHome = true
	BuildIndex("", db.Publishable(), current)

	// Now build directory inde
	color.Cyan("Build category")
	for dir, nodes := range db.ByCategory() {
		color.Cyan("\t√%s ", dir)
		current := &Current{
			IsHome: false,
			IsDir:  true,
			IsTag:  false,
		}

		BuildIndex(dir, nodes, current)
	}

	color.Cyan("Build tag")
	for tag, nodes := range db.ByTag() {
		color.Cyan("\t√%s ", tag)
		current := &Current{
			IsHome: false,
			IsDir:  false,
			IsTag:  true,
		}
		BuildIndex("tag/"+tag, nodes, current)
	}
	color.Green("Done! Enjoy")
}

func CreateNode(dir, title string) error {
	slug := strings.Replace(title, " ", "-", -1)

	file, err := os.Create("content/" + dir + "/" + slug + ".md")
	if err != nil {
		log.Fatalf("Cannot create file in", dir, ". Check directory permission. Err", err)
	}

	defer file.Close()

	content := `+++
date = "%s"
title = "%s"
draft = true

tags = []
+++`
	fmt.Fprintf(file, fmt.Sprintf(content, time.Now().Format(time.RFC3339), title))

	return nil
}
