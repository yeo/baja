package baja

import (
	"bufio"
	"html/template"
	"log"
	"os"
)

func BuildIndex(dir string, nodes []*Node, home bool) {
	directory := "public/" + dir

	sort.Slice(nodes, func(i, j int) bool { return allNodes[i].Meta.Date.After(allNodes[j].Meta.Date) })

	if len(nodes) == 0 {
		log.Println("Directory", dir, "has no file. Skip")
		return
	}

	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)

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
		log.Println("Fail to render", err)
	}
	w.Flush()
}
