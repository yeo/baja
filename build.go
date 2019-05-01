package baja

import (
	"bufio"
	"fmt"
	"html/template"
	"strings"
	"time"

	"os"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
	"github.com/yeo/baja/utils"
)

// Build executes template and content to generate our real static conent
func Build() int {
	config := DefaultConfig()

	os.RemoveAll("./public")
	CompileAsset(config)

	db := BuildDB(config)
	CompileNodes(db)

	return 0
}

// CompileAsset copy asset from theme or static into public and also generate a hash version of those file
func CompileAsset(config *Config) {
	theme := GetTheme(config)
	utils.CopyDir(theme.SubPath("static/"), "public")
	utils.CopyDir("static", "public")

	// Now generate hash
	err := filepath.Walk("./public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			color.Red("Error while access %q: %v\n", path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		color.Green("Generate hash for %s", path)
		utils.CopyFileWithHash(path)

		return nil
	})

	if err != nil {
		color.Red("error compile asser%v", err)
		return
	}
}

type visitor func(path string, f os.FileInfo, err error) error

func visit(db *NodeDB) filepath.WalkFunc {

	return func(path string, f os.FileInfo, err error) error {
		color.Green("\t%s", path)

		if f.IsDir() {
			return nil
		}

		db.Append(NewNode(path))

		return nil
	}
}

func BuildDB(config *Config) *NodeDB {
	db := &NodeDB{
		NodeList: []*Node{},
	}
	color.Green("Scan content")
	_ = filepath.Walk("./content", visit(db))
	return db
}

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

func CompileNodes(db *NodeDB) {
	// Build individual node
	color.Yellow("Start build html\n  Build individual page")
	for i, node := range db.NodeList {
		color.Yellow("\t%d/%d:  %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	current := &Current{
		IsHome:     false,
		IsDir:      false,
		IsTag:      false,
		CompiledAt: time.Now(),
	}
	// Now build the main index pag
	current.IsHome = true
	BuildIndex("", db.Publishable(), current)

	// Now build directory inde
	color.Cyan("  Build category")
	for dir, nodes := range db.ByCategory() {
		color.Cyan("    %s ", dir)
		current := &Current{
			IsHome:     false,
			IsDir:      true,
			IsTag:      false,
			CompiledAt: time.Now(),
		}

		BuildIndex(dir, nodes, current)
	}

	color.Cyan("  Build tag")
	for tag, nodes := range db.ByTag() {
		color.Cyan("    %s ", tag)
		current := &Current{
			IsHome:     false,
			IsDir:      false,
			IsTag:      true,
			CompiledAt: time.Now(),
		}
		BuildIndex("tag/"+tag, nodes, current)
	}
	color.Green("Done! Enjoy")
}

func CreateNode(dir, title string) error {
	slug := strings.Replace(title, " ", "-", -1)

	file, err := os.Create("content/" + dir + "/" + slug + ".md")
	if err != nil {
		color.Red("Cannot create file in %s. Check directory permission. Err: err", dir, err)
		return err
	}

	defer file.Close()

	content := `+++
date = "%s"
title = "%s"
draft = true

tags = []
+++`
	fmt.Fprintf(file, fmt.Sprintf(content, time.Now().Format(time.RFC3339), title))

	return nil
}
