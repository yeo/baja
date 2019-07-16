package baja

import (
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
	ctx := NewContext(cfg.Default())

	os.RemoveAll("./public")
	db := BuildDB(ctx)

	CompileAsset(ctx)
	CompileNodes(db)

	return 0
}

// CompileAsset copy asset from theme or static into public and also generate a hash version of those file
func CompileAsset(ctx *Context) {
	utils.CopyDir(ctx.Theme.SubPath("static/"), "public")
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
	color.Yellow("Build individual page")
	for i, node := range db.All() {
		color.Yellow("\t%d/%d:  %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	indexNode := NewIndex("", db.Publishable())
	indexNode.Compile()

	color.Cyan("Build category")
	for dir, nodes := range db.ByCategory() {
		color.Cyan("    %s ", dir)
		indexNode := NewIndex(dir, nodes)
		indexNode.Compile()
	}

	color.Cyan("Build tag")
	for tag, nodes := range db.ByTag() {
		color.Cyan("    %s ", tag)
		indexNode := NewIndex("tag/"+tag, nodes)
		indexNode.Compile()
	}

	color.Green("💥 Done! Enjoy. 🏖")
}
