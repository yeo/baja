package baja

import (
	"html/template"
	"time"

	"github.com/fatih/color"
)

type Site struct {
	Name    string
	Author  string
	BaseUrl string

	Config *SiteConfig
}

type Current struct {
	IsHome bool
	IsDir  bool
	IsTag  bool
	IsList bool
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

type ListPage struct {
	Current   *Current
	Title     string
	Permalink string
	Nodes     []map[string]interface{}
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
			color.Red("[ignore]%s page/draft", node.Name)
			continue
		}

		nodes = append(nodes, node)
	}

	return nodes
}
