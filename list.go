package baja

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"sort"
)

func BuildIndex(dir string, nodes []*Node, current *Current) {
	theme := GetTheme(DefaultConfig())

	targetDirectory := "public/" + dir
	os.MkdirAll(targetDirectory, os.ModePerm)

	f, err := os.Create(targetDirectory + "/index.html")
	if err != nil {
		fmt.Println("Cannot create index.html in", targetDirectory, ". error", err)
	}

	w := bufio.NewWriter(f)

	sort.Slice(nodes, func(i, j int) bool { return nodes[i].Meta.Date.After(nodes[j].Meta.Date) })
	nodeData := make([]map[string]interface{}, len(nodes))

	for i, n := range nodes {
		nodeData[i] = n.data()
	}

	data := ListPage{
		current,
		dir,
		dir,
		nodeData,
	}

	tpl, err := template.New("layout").Funcs(FuncMaps()).ParseFiles(theme.LayoutPath("default"))
	tpl, err = tpl.ParseFiles(theme.NodePath("index"))

	if _, err := os.Stat(theme.Path() + dir + "/index.html"); err == nil {
		tpl, err = tpl.ParseFiles(theme.Path() + dir + "/index.html")
	}

	if current.IsHome {
		if _, err := os.Stat(theme.NodePath("home")); err == nil {
			tpl, err = tpl.ParseFiles(theme.NodePath("home"))
		}
	}

	if tpl == nil {
		fmt.Println("Cannot create template render")
		return
	}

	if err := tpl.Execute(w, data); err != nil {
		fmt.Println("Fail to render. Check your template for syntax, wrong tag", err)
	}
	w.Flush()
}
