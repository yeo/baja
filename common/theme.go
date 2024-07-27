package baja

import (
	"html/template"
	"path/filepath"

	"github.com/yeo/baja/utils"
)

type Theme struct {
	Name string
	path string
}

func NewThemeFromConfig(config *Config) *Theme {
	path, _ := filepath.Abs("themes/" + config.Theme)

	t := Theme{
		Name: config.Theme,
		path: path,
	}

	return &t
}

func (t *Theme) LayoutPath(name string) string {
	return t.path + "/layout/" + name + ".html"
}

func (t *Theme) NodePath(node string) string {
	return t.path + "/" + node + ".html"
}

func (t *Theme) Path() string {
	return t.path + "/"
}

func (t *Theme) SubPath(subpath string) string {
	return t.path + "/" + subpath
}

func FuncMaps() template.FuncMap {
	funcMap := template.FuncMap{
		"asset": utils.GenerateAssetHash,
	}

	return funcMap
}
