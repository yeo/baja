package baja

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"sort"
)

func BuildIndex(dir string, nodes []*Node, home bool) {
	if len(nodes) == 0 {
		fmt.Println("targetDirectory", dir, "has no file. Skip")
		return
	}

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

	data := struct {
		Title     string
		Permalink string
		Nodes     []map[string]interface{}
	}{
		dir,
		dir,
		nodeData,
	}

	tpl, err := template.New("layout").ParseFiles("themes/baja/layout/default.html")
	tpl, err = tpl.ParseFiles("themes/baja/index.html")

	if _, err := os.Stat("themes/baja/" + dir + "/index.html"); err == nil {
		tpl, err = tpl.ParseFiles("themes/baja/" + dir + "/index.html")
	}

	if home == true {
		if _, err := os.Stat("themes/baja/home.html"); err == nil {
			tpl, err = tpl.ParseFiles("themes/baja/home.html")
		}
	}

	if err := tpl.Execute(w, data); err != nil {
		fmt.Println("Fail to render", err)
	}
	w.Flush()
}
