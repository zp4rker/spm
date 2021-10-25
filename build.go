//go:build build

//go:generate go build ./cmd/cli/jpm.go
//go:generate go build ./cmd/controller/jpm-controller.go
//go:generate go build ./cmd/wrapper/jpm-wrapper.go

package build