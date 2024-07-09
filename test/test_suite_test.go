package test_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	format.RegisterCustomFormatter(formatter)
	RunSpecs(t, "Test Suite")
}

func formatter(value any) (string, bool) {
	// handle github.com/pkg/errors with a stack
	pkgErr, isPkgError := value.(interface{ Cause() error })
	if isPkgError {
		return fmt.Sprintf("%+v", pkgErr), true
	}

	return "", false
}
