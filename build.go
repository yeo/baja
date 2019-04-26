package baja

import (
	"os"
	"path/filepath"

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
