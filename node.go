package baja

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/microcosm-cc/bluemonday"
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
	Title string
}

type Node struct {
	Meta *NodeMeta
	Body string

	Params *NodeParams

	Raw           string
	Path          string
	templatePaths []string
}

func NewNode(path string) *Node {
	n := Node{Path: path}

	return &n
}

func (n *Node) data() map[string]interface{} {
	unsafe := blackfriday.Run([]byte(n.Body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return map[string]interface{}{
		"meta": n.Meta,
		"body": html,
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

	n.Body = string(content[2])
}

func (n *Node) FindTheme(c *Config) {
	// Find theme
	pathComponents := strings.Split(n.Path, "/")
	pathComponents[0] = c.Theme
	n.templatePaths = []string{"theme"}
	for _, p := range pathComponents {
		n.templatePaths = append(n.templatePaths, p)
	}
}

func (n *Node) Compile() {
	to := "public/" + n.Path
	dotPosition := strings.LastIndex(to, ".")
	directory := to[0:dotPosition]
	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)

	fmt.Println(n.templatePaths)

	tpl, err := template.ParseFiles(n.templatePaths...)
	if err != nil {
		log.Println("Cannot parse tempalte", n.Path)
	}

	tpl.ExecuteTemplate(w, "layout", n.data())
}

type visitor func(path string, f os.FileInfo, err error) error

func visit(node *TreeNode) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		fmt.Printf("Visited: %s\n", path)

		if f.IsDir() {
			os.MkdirAll("./public/"+path, os.ModePerm)
			return nil
		}

		//Super simple parsing
		n := NewNode(path)
		n.Parse()
		n.FindTheme(DefaultConfig())
		n.Compile()

		return nil
	}
}

func BuildNodeTree(config *Config) *TreeNode {
	n := &TreeNode{}
	_ = filepath.Walk("./content", visit(n))
	return nil
}

func (t *TreeNode) Compile() {

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
