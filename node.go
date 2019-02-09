package baja

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
	//"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	//"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		return "/" + filepath.Base(n.Name)
	} else {
		return "/" + n.BaseDirectory + "/" + filepath.Base(n.Name)
	}
}

func (n *Node) data() map[string]interface{} {
	unsafe := blackfriday.Run([]byte(n.Body))
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	html := unsafe

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
	fmt.Println("Compile", n.Name, "in", directory)

	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)

	log.Println(n.templatePaths)

	tpl := template.New("layout")

	tpl, err = tpl.ParseFiles("themes/baja/layout/default.html", "themes/baja/node.html")
	if err != nil {
		log.Panic(err)
	}

	if err := tpl.Execute(w, n.data()); err != nil {
		log.Println("Fail to render", err)
	}

	fmt.Println("Write ", n.BaseDirectory, n.Name, "to", directory+"/index.html")
	w.Flush()
}

func CompileNodes(db *NodeDB) {
	// Build individual node
	for i, node := range db.NodeList {
		fmt.Printf("Build progress %d/%d. File: %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	// Now build the main index page
	BuildIndex("", db.Publishable(), true)

	// Now build directory inde
	for dir, nodes := range db.ByCategory() {
		fmt.Println("Build category", dir)
		BuildIndex(dir, nodes, false)
	}
	for tag, nodes := range db.ByTag() {
		fmt.Println("Build tag", tag)
		BuildIndex("tag/"+tag, nodes, false)
	}
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
