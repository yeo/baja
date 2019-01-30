package baja

func compileAsset(config *Config) {
	// This should go into a site/path helper
	CopyDir("themes/"+config.Theme+"/static/css", "public/css")
	CopyDir("themes/"+config.Theme+"/static/js", "public/js")
	CopyDir("static", "public")
}
