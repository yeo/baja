package baja

import (
	"github.com/yeo/baja/utils"
)

func compileAsset(config *Config) {
	// This should go into a site/path helper
	utils.CopyDir("themes/"+config.Theme+"/static/css", "public/css")
	utils.CopyDir("themes/"+config.Theme+"/static/js", "public/js")
	utils.CopyDir("static", "public")
}
