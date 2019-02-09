package baja

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/yeo/baja/utils"
)

// Build execute template and content to generate our real static conent
func Build() int {
	config := DefaultConfig()

	db := BuildDB(config)
	CompileNodes(db)
	CompileAsset(config)

	return 0
}

func CompileAsset(config *Config) {
	// This should go into a site/path helper
	utils.CopyDir("themes/"+config.Theme+"/static/css", "public/css")
	utils.CopyDir("themes/"+config.Theme+"/static/js", "public/js")
	utils.CopyDir("static", "public")
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
