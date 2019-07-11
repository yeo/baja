package baja

import (
	"html/template"
	"os"
	"strings"

	"github.com/yeo/baja/cfg"
	"github.com/yeo/baja/utils"
)

type Theme struct {
	Name string
	path string
}

func GetTheme(config *cfg.Config) *Theme {
	t := Theme{
		Name: config.Theme,
		path: "themes/" + config.Theme + "/",
	}

	return &t
}

func (t *Theme) LayoutPath(name string) string {
	return t.path + "layout/" + name + ".html"
}

func (t *Theme) NodePath(node string) string {
	return t.path + node + ".html"
}

func (t *Theme) Path() string {
	return t.path
}

func (t *Theme) SubPath(subpath string) string {
	return t.path + subpath
}

func (t *Theme) FindTheme(n *Node) {
	// Find theme
	pathComponents := strings.Split(n.BaseDirectory, "/")
	n.templatePaths = []string{t.LayoutPath("default")}
	lookupPath := t.Path()
	for _, p := range pathComponents {
		if _, err := os.Stat(lookupPath + "/node.html"); err == nil {
			n.templatePaths = append(n.templatePaths, lookupPath+"/node.html")
		}
		lookupPath = lookupPath + "/" + p
	}
}

func FuncMaps() template.FuncMap {
	funcMap := template.FuncMap{
		"asset": utils.GenerateAssetHash,
	}

	return funcMap
}
