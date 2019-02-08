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
	Type          string
}

type Node struct {
	Meta *NodeMeta
	Body template.HTML
	Name string

	Params *NodeParams

	Raw           string
	Path          string
	BaseDirectory string
	templatePaths []string
}

var NodeDB map[string][]*Node
