package errs

import "errors"

var (
	PetsNotFound = errors.New("pets not found")
)
