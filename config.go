package baja

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type SiteConfig struct {
	Name   string `"yaml":name`
	Author string `"yaml":author`
}

type Config struct {
	Theme string `"yaml":theme`
	Site  string `"yaml":site`
	path  string
}

func NewConfig(path string) *Config {
	c := Config{path: path}
	return &c
}

func (c *Config) WriteFile() {
	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("Cannot write config file file %v", err)
	}

	ioutil.WriteFile(c.path, d, 0644)
}
