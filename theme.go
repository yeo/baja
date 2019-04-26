package baja

type Theme struct {
	Name string
	path string
}

func GetTheme(config *Config) *Theme {
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
