package baja

import (
	"time"

	"github.com/yeo/baja/cfg"
)

// Site stores information of this static site and various meta data
type Site struct {
	Name    string
	Author  string
	BaseUrl string

	Config *cfg.SiteConfig
}

// Current is a struct about various current state we pass to template to help us do some business logic depend on a context
type Current struct {
	IsHome bool
	IsDir  bool
	IsTag  bool
	IsList bool

	CompiledAt time.Time
}
