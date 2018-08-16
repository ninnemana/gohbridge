package hue

import "github.com/pkg/errors"

var (
	// ErrNotImplemented is a placeholder errors message for missing
	// interface implementations.
	ErrNotImplemented = errors.New("functionality is not implemented")

	ErrNoUser = errors.New("user parameter missing from context")

	ErrNoHost = errors.New("host parameter missing from context")
)
