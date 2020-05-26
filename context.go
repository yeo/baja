package baja

import (
	"time"
)

type Context struct {
	Config  *Config
	Current *Current
	Theme   *Theme
}

func NewContext(config *Config) *Context {
	c := Context{
		Config: config,
		Theme:  NewThemeFromConfig(config),
		Current: &Current{
			IsHome: false,
			IsDir:  true,
			IsTag:  false,
			IsList: true,
		},
	}

	return &c
}

// Current is a struct about various current state we pass to template to help us do some business logic depend on a context
type Current struct {
	IsHome bool
	IsDir  bool
	IsTag  bool
	IsList bool

	CompiledAt time.Time
}
