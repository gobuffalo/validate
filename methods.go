package validator

import (
	"encoding/json"
	"sort"

	"github.com/Jeffail/gabs"
)

// https://google.github.io/styleguide/jsoncstyleguide.xml

// Add will add (append) a new error message to the list of errors using the given path.
func (e *Errors) Add(path string, msg string) {

	// check for nil, no errors
	if e == nil {
		return
	}

	e.u.Lock()
	defer e.u.Unlock()

	// initialize error messages slice when needed
	if _, ok := e.m[path]; !ok {
		e.m[path] = []string{}
	}

	// populate internal map with paths mapped to error messages
	e.m[path] = append(e.m[path], msg)
}

// Count returns the number of errors.
func (e *Errors) Count() int {

	var total int

	// check for nil, no errors
	if e == nil {
		return 0
	}

	// loop-over path keys
	for path := range e.m {
		// append to total errors count
		total += len(e.m[path])
	}

	return total
}

// HasAny returns 'true' if errors have occurred, 'false' if no errors occurred.
func (e *Errors) HasAny() bool {
	return e.Count() > 0
}

// Exists returns 'true' if given path exist and has errors, else returns 'false'.
func (e *Errors) Exists(path string) bool {

	// check for nil, no errors
	if e == nil {
		return false
	}

	// check existence of given path
	if _, ok := e.m[path]; !ok {
		return false
	}

	// get number of errors for a given path
	if len(e.m[path]) == 0 {
		return false // skip for non-errors, should never happen
	}

	return true
}

// Get returns an array of error messages for the given path.
func (e *Errors) Get(path string) []string {

	// check existence of given path
	if !e.Exists(path) {
		return []string{}
	}

	return e.m[path]
}

// Keys returns all paths (sorted) which have errors.
func (e *Errors) Keys() []string {

	// check for nil, no errors
	if e == nil {
		return []string{}
	}

	// initialize keys slice with capacity of internal path map
	keys := make([]string, 0, len(e.m))

	for key := range e.m {
		keys = append(keys, key)
	}

	// sort paths in reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	return keys
}

// sync synchronizes internal map with container interface.
func (e *Errors) sync() *gabs.Container {

	// container for dynamic JSON
	c := gabs.New()

	// check for nil, no errors
	if e == nil {
		return c
	}

	// get all paths
	paths := e.Keys()

	e.u.Lock()
	defer e.u.Unlock()

	for _, path := range paths {

		var err error

		// get all errors
		values, ok := e.m[path]
		if !ok {
			panic("internal map error, path does not exist") // WTF?!
		}

		// loop-over errors and append them one-by-one
		for _, val := range values {
			if e.dot { // use JSON dot paths
				err = c.ArrayAppendP(val, path)
			} else { // use plain JSON paths
				err = c.ArrayAppend(val, path)
			}
			if err != nil {
				panic(err) // WTF?!
			}
		}
	}

	return c
}

// Data returns errors as native datatype.
func (e *Errors) Data() interface{} {
	return e.sync().Data()
}

// JSON returns JSON formatted errors (as string).
func (e *Errors) JSON() string {
	return e.sync().String()
}

// Raw returns JSON formatted errors (as json.RawMessage).
func (e *Errors) Raw() json.RawMessage {
	return json.RawMessage(e.sync().Bytes())
}

// Error implements the default golang error interface.
func (e *Errors) Error() string {
	return e.sync().String()
}
