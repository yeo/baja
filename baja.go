package baja

import (
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

func (db *NodeDB) AddDirectory(dir string) {
	db.DirectoryList = append(db.DirectoryList, dir)
}

func (db *NodeDB) Append(n *Node) {
	db.NodeList = append(db.NodeList, n)
	db.Total = len(db.NodeList)
}
