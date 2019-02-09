package baja

import (
	"fmt"
	"html/template"
	"time"
)

type Site struct {
	Name    string
	Author  string
	BaseUrl string
}

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
	Type          string
}

type Node struct {
	Meta *NodeMeta
	Body template.HTML
	Name string // the filename without extension

	Raw           string
	Path          string
	BaseDirectory string // the directory without /content part
	templatePaths []string
}

type NodeDB struct {
	NodeList      []*Node
	DirectoryList []string
	Total         int
}

func (db *NodeDB) Append(n *Node) {
	db.NodeList = append(db.NodeList, n)
	db.Total = len(db.NodeList)
}

func (db *NodeDB) ByCategory() map[string][]*Node {
	categoryNodes := make(map[string][]*Node)

	for _, node := range db.NodeList {
		if node.BaseDirectory == "" {
			// Those are node directly under content/ without any subdirectory
			// they are only appear in / index page and not in subdirectory page
			continue
		}
		if categoryNodes[node.BaseDirectory] == nil {
			categoryNodes[node.BaseDirectory] = []*Node{}
		}
		categoryNodes[node.BaseDirectory] = append(categoryNodes[node.BaseDirectory], node)
	}

	return categoryNodes
}

func (db *NodeDB) ByTag() map[string][]*Node {
	tagsNode := make(map[string][]*Node)
	for _, node := range db.NodeList {
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

	return tagsNode
}

func (db *NodeDB) Publishable() []*Node {
	nodes := []*Node{}

	for _, node := range db.NodeList {
		if node.IsPage() || node.Meta.Draft {
			fmt.Printf("node %s is ignored because it's a draft or page %s %s\n", node.Name, node.Meta.Type, node.Meta.Draft)
			continue
		}

		nodes = append(nodes, node)
	}

	return nodes
}
