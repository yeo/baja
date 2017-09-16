package baja

import (
	"os"
)

type NodeList struct {
	Directory string
}

func (l *NodeList) data() []Node {
}

func (l *NodeList) Compile() {
	directory := "public/" + l.Directory
	os.MkdirAll(directory, os.ModePerm)
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Cannot write to file", err, directory)
	}

	w := bufio.NewWriter(f)

	tpl := template.New("layout")
	for i := len(l.templatePaths) - 1; i >= 0; i-- {
		t := l.templatePaths[i]
		if out, err := ioutil.ReadFile(t); err == nil {
			if tpl, err = tpl.Parse(string(out)); err != nil {
				log.Println("Cannot parse", t, err)
			}
		}
	}
	log.Println("Loaded", l.templatePaths)

	if err := tpl.Execute(w, l.data()); err != nil {
		log.Println("Fail to render", err)
	}
	w.Flush()
}
