package models

import (
	"fmt"
)

// MissingEnvError is a custom error class that getsthrown when
//an environment variable required by hue-remote-v2 is missing.
type MissingEnvError struct {
	name string
}

func (e MissingEnvError) Error() string {
	if &e == nil {
		return ""
	}
	return fmt.Sprintf("Environment variable %s was not provided", e.name)
}

// NewMissingEnvError returns a new instance of MissingEnvError
func NewMissingEnvError(name string) (err MissingEnvError) {
	return MissingEnvError{
		name: name,
	}
}
