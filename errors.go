package sharding

import (
	"errors"
)

// ErrBadSchemeParameter is returned when a scheme is badly defined.
//
var ErrBadSchemeParameter = errors.New("sharding: bad scheme parameter")

// BadScheme is the panic message for a Scheme that was not created with the
// New function.
//
const BadScheme = "sharding: uninitialised scheme"
