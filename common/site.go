package baja

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type SiteMeta struct {
	Name    string `"yaml":name`
	Author  string `"yaml":author`
	BaseURL string `"yaml":baseURL`
}

type SitePath struct {
	Content string
	Output  string
}

type Site struct {
	Config *Config
	Theme  *Theme

	Meta *SiteMeta
	Path *SitePath
}

func LoadSite(configpath string) *Site {
	path, err := filepath.Abs(configpath)
	if err != nil {
		log.Fatalf("configpath invalid %s", configpath)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	config := &Config{path: path}
	err = yaml.Unmarshal(data, &config)

	outputPath, _ := filepath.Abs("./public")
	contentPath, _ := filepath.Abs("./content")
	site := Site{
		Config: config,
		Theme:  NewThemeFromConfig(config),
		Meta:   &SiteMeta{},
		Path: &SitePath{
			// TODO: Load these from config
			Output:  outputPath,
			Content: contentPath,
		},
	}

	return &site
}
