package baja

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

// NodeMeta is meta data of a node, usually map directly to node toml metadata section
type NodeMeta struct {
	Title         string
	Draft         bool
	Date          time.Time
	DateFormatted string
	Tags          []string
	Category      string
	Type          string // node type. Eg page or post
	Theme         string // a custom template file inside theme directory without extension
}

// Node hold information of a specifc page we are rendering
type Node struct {
	Meta *NodeMeta
	Body template.HTML

	Raw           string // full absolute path to markdown file
	Path          string
	BaseDirectory string   // the directory without /content part
	Name          string   // the filename without extension
	templatePaths []string // a list of template files that are discovered for this node. These templates are used to render content
}

// ListPage is an index page, it isn't constructed from a markdown file but from a list of related markdown such as tag or category
type ListPage struct {
	Current   *Current
	Title     string
	Permalink string
	Nodes     []map[string]interface{}
}

// NodeDB is the in-memory database of all the page
type NodeDB struct {
	NodeList      []*Node
	DirectoryList []string
	Total         int
}

// BuildDB calculate a tree to represent all of node
// This tree can be query/group/filter
func BuildDB(config *Config) *NodeDB {
	db := &NodeDB{
		NodeList: []*Node{},
	}
	color.Green("Scan content")
	_ = filepath.Walk("./content", visit(db))
	return db
}

func (db *NodeDB) Append(n *Node) {
	db.NodeList = append(db.NodeList, n)
	db.Total = len(db.NodeList)
}

// ByTag category groups node by category(category is the directoy name)
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

// ByTag groups node by tag
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

// Publishable returns a list of node that can be publish, as in non-draft mode or non page
func (db *NodeDB) Publishable() []*Node {
	nodes := []*Node{}

	for _, node := range db.NodeList {
		if node.IsPage() {
			color.Red("\tignore %s because it's a standalone page", node.Name)
			continue
		}

		if node.Meta.Draft {
			color.Red("\tignore %s because it's in draft mode ", node.Name)
			continue
		}

		nodes = append(nodes, node)
	}

	return nodes
}
