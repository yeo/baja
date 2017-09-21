package baja

import (
	"bufio"
	"html/template"
	"log"
	"os"
)

func BuildIndex(dir string, nodes []*Node) {
	directory := "public/" + dir
	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)
	data := map[string]interface{}{
		"Permalink": dir,
		"Nodes":     nodes,
	}

	tpl, err := template.New("layout").ParseFiles("themes/baja/layout/default.html", "themes/baja/list.html")

	if err := tpl.Execute(w, data); err != nil {
		log.Println("Fail to render", err)
	}
	w.Flush()
}
