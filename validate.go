package validator

import (
	"sync"
)

// Validator must be implemented in order to pass the validator object into the Validate function.
type Validator interface {
	Validate(e *Errors)
}

type vfWrapper struct {
	vf func(e *Errors)
}

func (v vfWrapper) Validate(e *Errors) {
	v.vf(e)
}

// Validate runs n number of supplied Validator objects.
func (e *Errors) Validate(validators ...Validator) {

	// declare WaitGroup
	wg := &sync.WaitGroup{}

	// loop-over supplied validators
	for i := range validators {
		// new goroutine busy count
		wg.Add(1)

		go func(wg *sync.WaitGroup, i int) {
			// release goroutines busy count on exit
			defer wg.Done()

			// get current validator object
			validator := validators[i]

			// run validation method
			validator.Validate(e)
		}(wg, i)
	}

	// wait for all concurrent goroutines to finish
	wg.Wait()
}

// Validate wraps Validate method with dot notated JSON disabled.
func Validate(validators ...Validator) *Errors {

	// create errors object with dot notated JSON disabled
	e := NewErrors()

	// run validators
	e.Validate(validators...)

	if !e.HasAny() {
		return nil // return nil when there are no errors
	}

	return e
}

// ValidateP wraps Validate method with dot notated JSON enabled.
func ValidateP(validators ...Validator) *Errors {

	// create errors object with dot notated JSON enabled
	e := NewErrorsP()

	// run validators
	e.Validate(validators...)

	if !e.HasAny() {
		return nil // return nil when there are no errors
	}

	return e
}

// ValidateFunc wraps any function in a 'Validator' for custom validator rules.
func ValidateFunc(fn func(e *Errors)) Validator {
	return vfWrapper{fn}
}
