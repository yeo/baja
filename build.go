package baja

import (
	"time"

	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/yeo/baja/cfg"
	"github.com/yeo/baja/utils"
)

// ListPage is an index page, it isn't constructed from a markdown file but from a list of related markdown such as tag or category
type ListPage struct {
	Current   *Current
	Title     string
	Permalink string
	Nodes     []map[string]interface{}
}

// Build executes template and content to generate our real static conent
func Build() int {
	config := cfg.Default()

	os.RemoveAll("./public")
	db := BuildDB(config)

	CompileAsset(config)
	CompileNodes(db)

	return 0
}

// CompileAsset copy asset from theme or static into public and also generate a hash version of those file
func CompileAsset(config *cfg.Config) {
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

func CompileNodes(db *NodeDB) {
	color.Yellow("Start build html\n  Build individual page")
	for i, node := range db.NodeList {
		color.Yellow("\t%d/%d:  %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	// Now build the main index page
	current.IsHome = true
	indexNode := IndexNode{"", db.Publishable(), current}
	indexNode.Compile()

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

		indexNode := IndexNode{dir, nodes, current}
		indexNode.Compile()
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
		indexNode := IndexNode{"tag/" + tag, nodes, current}
		indexNode.Compile()
	}
	color.Green("Done! Enjoy")
}
