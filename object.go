package validator

import (
	"sync"
)

// Errors describes custom validator errors object.
type Errors struct {
	// mutex lock for concurent access
	u *sync.RWMutex

	// map of paths to errors
	m map[string][]string

	// use dot notated JSON paths
	dot bool
}

// NewErrors returns a pointer to an initialized Errors object
// with dot notated JSON disabled (no nested struct).
func NewErrors() *Errors {
	return &Errors{
		u:   &sync.RWMutex{},
		m:   make(map[string][]string),
		dot: false,
	}
}

// NewErrorsP returns a pointer to an initialized Errors object
// with dot notated JSON enabled (nested struct).
func NewErrorsP() *Errors {
	return &Errors{
		u:   &sync.RWMutex{},
		m:   make(map[string][]string),
		dot: true,
	}
}
