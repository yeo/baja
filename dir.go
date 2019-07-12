package baja

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/yeo/baja/cfg"
	"os"
	"sort"
)

type IndexNode struct {
	Dir     string
	Nodes   []*Node
	Current *Current
}

func (n *IndexNode) Compile() {
	current := &Current{
		IsHome:     false,
		IsDir:      false,
		IsTag:      false,
		CompiledAt: time.Now(),
	}

	if n.Dir == "" {
		current.IsHome = true
	}

	if strings.Stat

	theme := GetTheme(cfg.Default())

	targetDirectory := "public/" + n.Dir
	os.MkdirAll(targetDirectory, os.ModePerm)

	f, err := os.Create(targetDirectory + "/index.html")
	if err != nil {
		fmt.Println("Cannot create index.html in", targetDirectory, ". error", err)
	}

	w := bufio.NewWriter(f)

	sort.Slice(n.Nodes, func(i, j int) bool { return n.Nodes[i].Meta.Date.After(n.Nodes[j].Meta.Date) })
	nodeData := make([]map[string]interface{}, len(n.Nodes))

	for i, n := range n.Nodes {
		nodeData[i] = n.data()
	}

	data := ListPage{
		n.Current,
		n.Dir,
		n.Dir,
		nodeData,
	}

	tpl, err := template.New("layout").Funcs(FuncMaps()).ParseFiles(theme.LayoutPath("default"))
	tpl, err = tpl.ParseFiles(theme.NodePath("index"))

	log.Println("Build index", n.Dir, theme.SubPath(n.Dir+".html"))
	if _, err := os.Stat(theme.SubPath(n.Dir + ".html")); err == nil {
		tpl, err = tpl.ParseFiles(theme.SubPath(n.Dir + ".html"))
	}

	if _, err := os.Stat(theme.Path() + n.Dir + "/index.html"); err == nil {
		tpl, err = tpl.ParseFiles(theme.Path() + n.Dir + "/index.html")
	}

	if n.Current.IsHome {
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
