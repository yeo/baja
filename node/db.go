package node

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/yeo/baja"
)

// NodeDB is the in-memory database of all the page
type NodeDB struct {
	NodeList      []*Node
	DirectoryList []string
	Total         int
	Site          *baja.Site
}

func (db *NodeDB) Append(n *Node) {
	db.NodeList = append(db.NodeList, n)
	db.Total = len(db.NodeList)
}

func (db *NodeDB) All() []*Node {
	return db.NodeList
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

type visitor func(path string, f os.FileInfo, err error) error

func visit(db *NodeDB) filepath.WalkFunc {

	return func(path string, f os.FileInfo, err error) error {
		color.Green("\t%s", path)

		if f.IsDir() {
			return nil
		}

		db.Append(NewNode(db.Site, path))

		return nil
	}
}

// BuildDB calculate a tree to represent all of node
// This tree can be query/group/filter
func BuildDB(ctx *baja.Context) *NodeDB {
	db := &NodeDB{
		NodeList: []*Node{},
	}
	color.Green("Scan content")
	_ = filepath.Walk("./content", visit(db))
	return db
}
