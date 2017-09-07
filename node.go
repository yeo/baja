package baja

import (
	"fmt"
	"ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type NodeParams struct{}

type TreeNode struct {
	Name  string
	Leafs []TreeNode
}

type Node struct {
	Title string
	Body  string

	Params *NodeParams

	Raw  string
	Path string
}

func NewNode(path string) *Node {
	n := Node{}

	return &n
}

type visitor func(path string, f os.FileInfo, err error) error

func visit(node *TreeNode) filepath.WalkFun {
	return func(path string, f os.FileInfo, err error) error {
		fmt.Printf("Visited: %s\n", path)

		return nil
	}
}

func BuildNodeTree(config *Config) *TreeNode {
	n := &TreeNode{}
	err := filepath.Walk("./content", visit(n))
	return nil
}

func (t *TreeNode) Compile() {

}

func template(layout, path string) error {
	out, err := ioutil.ioutilReadFile(layout)
	if err != nil {
		return err
	}
	t, err := template.New("layoyt").Parse(string(out))
	if err != nil {
		return err
	}

	cluster, err := ioutil.ReadFile(node)
	if err != nil {
		return err
	}
	return t.Parse(string(cluster))
}

func render(tpl *template.Template, n *Node) {
	tpl.Execute(buf, n)
}
