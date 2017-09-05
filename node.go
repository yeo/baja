package baja

import (
	"fmt"
	"os"
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

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func BuildNodeTree(config *Config) *TreeNode {
	//err := filepath.Walk(directory, visit)
	return nil
}

func (t *TreeNode) Compile() {

}
