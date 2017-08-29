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

func BuildNodeTree(directory string) *TreeNode {
	err := filepath.Walk(directory, visit)

}
