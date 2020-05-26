package node

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"os"
	"sort"

	"github.com/yeo/baja"
)

// ListPage is an index page, it isn't constructed from a markdown file but from a list of related markdown such as tag or category
type ListPage struct {
	Current   *baja.Current
	Title     string
	Permalink string
	Nodes     []map[string]interface{}
}

type IndexNode struct {
	Dir     string
	Nodes   []*Node
	Current *baja.Current
}

func NewIndex(dir string, nodes []*Node) *IndexNode {
	n := &IndexNode{
		Dir: dir,
		Current: &baja.Current{
			IsHome:     false,
			IsDir:      false,
			IsTag:      false,
			CompiledAt: time.Now(),
		},
		Nodes: nodes,
	}

	if dir == "" {
		n.Current.IsHome = true
	}

	if strings.HasPrefix(dir, "tag/") {
		n.Current.IsTag = true
	} else {
		n.Current.IsDir = true
	}

	return n
}

func (n *IndexNode) Compile(site *baja.Site) {
	theme := site.Theme

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

	tpl, err := template.New("layout").Funcs(baja.FuncMaps()).ParseFiles(theme.LayoutPath("default"))
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
