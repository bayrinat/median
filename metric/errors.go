package metric

import "errors"

var (
	errNotPositiveCapacity = errors.New("capacity value must be greater than zero")
)
