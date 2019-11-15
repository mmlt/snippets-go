package main

import (
	"errors"
	"fmt"
)

// Predefined errors.
var ErrFoo = errors.New("foo error")


// Simple error type.
// usage: return XyzError{Info: "context"}
type XyzError struct {
	Info string
}

func (e XyzError) Error() string {
	return fmt.Sprintf("xyz error: %s", e.Info)
}


// Wrapping another error.
type GenericError struct {
	Info string
	Inner  error
}

func (e GenericError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("generic error: %s: %v", e.Info, e.Inner)
	}
	return fmt.Sprintf("generic error: %s", e.Info)
}

func (e GenericError) Unwrap() error {
	return e.Inner
}


// Wrapping another error with %w
// return fmt.Errorf("preprocessing %s: %w", j.name, err)


// Checking errors with Is() and As()
err := f()
if errors.Is(err, ErrFoo) {
// you know you got an ErrFoo
// respond appropriately
}

var xyz XyzError
if errors.As(err, &xyz) {
// you know you got a XyzError
// xyz's fields are populated
// respond appropriately
}


/*
The Upspin project uses a custom package, upspin.io/errors, to represent error conditions that arise inside the system.
These errors satisfy the standard Go error interface, but are implemented using a custom type, upspin.io/errors.Error,
that has properties that have proven valuable to the project.

type Error struct {
	Path upspin.PathName
	User upspin.UserName
	Op  Op
	Kind Kind
	Err error
}
https://commandcenter.blogspot.nl/2017/12/error-handling-in-upspin.html
*/