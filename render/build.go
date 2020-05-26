package render

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/yeo/baja"
	"github.com/yeo/baja/node"
	"github.com/yeo/baja/utils"
)

// Build executes template and content to generate our real static conent
func Build(site *baja.Site) int {
	ctx := baja.NewContext(site.Config)

	os.RemoveAll("./public")
	db := node.BuildDB(site, ctx)

	CompileAsset(ctx)
	CompileNodes(db)

	return 0
}

// CompileAsset copy asset from theme or static into public and also generate a hash version of those file
func CompileAsset(ctx *baja.Context) {
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

func CompileNodes(db *node.NodeDB) {
	color.Yellow("Build individual page")
	for i, node := range db.All() {
		color.Yellow("\t%d/%d:  %s\n", i+1, db.Total, node.Path)
		node.Compile()
	}

	indexNode := node.NewIndex("", db.Publishable())
	indexNode.Compile(db.Site.Config)

	color.Cyan("Build category")
	for dir, nodes := range db.ByCategory() {
		color.Cyan("    %s ", dir)
		indexNode := node.NewIndex(dir, nodes)
		indexNode.Compile(db.Site.Config)
	}

	color.Cyan("Build tag")
	for tag, nodes := range db.ByTag() {
		color.Cyan("    %s ", tag)
		indexNode := node.NewIndex("tag/"+tag, nodes)
		indexNode.Compile(db.Site.Config)
	}

	color.Green("💥 Done! Enjoy. 🏖")
}
