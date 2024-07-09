# ginkgo-gomega-custom-formatter-example

> Demonstrate using a [Gomega custom formatter](https://onsi.github.io/gomega/#adjusting-output) to pretty print different types of errors

## Usage

1. Clone the repository
1. Install Ginkgo cli via `go install -mod=mod github.com/onsi/ginkgo/ginkgo/v2/cmd/ginkgo`
1. Run tests via `gingko run ./tests`

## Examples

- [register custom formatter](./test/test_suite_test.go)
- [test examples of different kinds of errors](./test/example_test.go)
