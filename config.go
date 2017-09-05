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

func DefaultConfig() *Config {
	data, err := ioutil.ReadFile("./baja.yaml")
	if err != nil {
		log.Fatalf("Cannot read config file %v. Did you forget to run init?", err)
	}

	c := Config{path: "./baja.yaml"}
	err = yaml.Unmarshal(data, &c)

	return &c
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
