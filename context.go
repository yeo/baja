package baja

import (
	"time"
)

type Context struct{}

// Current is a struct about various current state we pass to template to help us do some business logic depend on a context
type Current struct {
	IsHome bool
	IsDir  bool
	IsTag  bool
	IsList bool

	CompiledAt time.Time
}
