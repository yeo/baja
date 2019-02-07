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

type NodeParams struct{}

type TreeNode struct {
	Name  string
	Leafs []TreeNode
	Type  string
}

type NodeMeta struct {
	Title         string
	Draft         bool
	Date          time.Time
	DateFormatted string
	Tags          []string
	Category      string
	Hidden        bool
}

type Node struct {
	Meta *NodeMeta
	Body template.HTML

	Params *NodeParams

	Raw           string
	Path          string
	BaseDirectory string
	Name          string
	templatePaths []string
}

func NewNode(path string) *Node {
	n := Node{Path: path}
	n.BaseDirectory = strings.Join(strings.Split(filepath.Dir(path), "/")[1:], "/")

	dotPosition := strings.LastIndex(path, ".")
	n.Name = path[0:dotPosition]

	return &n
}

func (n *Node) Permalink() string {
	return n.BaseDirectory + "/" + filepath.Base(n.Name)
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

func (n *Node) FindTheme(c *Config) {
	// Find theme
	pathComponents := strings.Split(n.Name, "/")
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
	name := strings.Split(n.Name, "/")
	directory := "public/" + strings.Join(name[1:], "/")
	log.Println("Compile", n.Name, directory)

	log.Println("Mkdir", directory)
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

	log.Println("Loaded", n.templatePaths)

	//tpl.Execute(os.Stdout, n.data())
	if err := tpl.Execute(w, n.data()); err != nil {
		log.Println("Fail to render", err)
	}

	log.Println("Write to ", directory+"/index.html")
	w.Flush()
}

type visitor func(path string, f os.FileInfo, err error) error

var NodeDB map[string][]*Node

func visit(node *TreeNode) filepath.WalkFunc {
	NodeDB = make(map[string][]*Node)

	return func(path string, f os.FileInfo, err error) error {
		fmt.Printf("Visited: %s\n", path)

		if f.IsDir() {
			if _, ok := NodeDB[path]; ok {
				NodeDB[path] = []*Node{}
			}
			return nil
		}

		node.Name = f.Name()
		//Super simple parsing
		n := NewNode(path)
		n.Parse()
		n.FindTheme(DefaultConfig())
		n.Compile()
		if n.Meta.Hidden == false {
			NodeDB[n.BaseDirectory] = append(NodeDB[n.BaseDirectory], n)
		}

		return nil
	}
}

func BuildNodeTree(config *Config) *TreeNode {
	n := &TreeNode{}
	_ = filepath.Walk("./content", visit(n))
	return n
}

func (t *TreeNode) Compile() {
	var tagsNode map[string][]*Node
	tagsNode = make(map[string][]*Node)

	allNodes := []*Node{}
	for dir, nodes := range NodeDB {
		BuildIndex(dir, nodes, false)
		allNodes = append(allNodes, nodes...)

		for _, node := range nodes {
			if len(node.Meta.Tags) > 0 {
				for _, tag := range node.Meta.Tags {
					if tagsNode[tag] == nil {
						tagsNode[tag] = []*Node{node}
					} else {
						tagsNode[tag] = append(tagsNode[tag], node)
					}
				}
			}
		}
	}

	BuildIndex("", allNodes, true)

	// TODO: concurent
	for t, nodes := range tagsNode {
		BuildIndex("tag/"+strings.ToLower(t), nodes, false)
	}
}

func _template(layout, path string) error {
	out, err := ioutil.ReadFile(layout)
	if err != nil {
		return err
	}
	t, err := template.New("layout").Parse(string(out))
	if err != nil {
		return err
	}

	cluster, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	t, err = t.Parse(string(cluster))
	return err
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
hidden = false

tags = []
+++`
	fmt.Fprintf(file, fmt.Sprintf(content, time.Now().Format(time.RFC3339), title))

	return nil
}
