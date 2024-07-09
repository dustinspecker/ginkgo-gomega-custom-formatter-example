package internal

import (
	"errors"

	pkgErrors "github.com/pkg/errors"
)

// ReturnsError returns a normal Go error
func ReturnsError() error {
	return errors.New("error from ReturnsError")
}

// ReturnsPkgError returns a wrapped error using github.com/pkg/errors
func ReturnsPkgError() error {
	err := ReturnsError()

	return pkgErrors.Wrapf(err, "error from ReturnsPkgError")
}
