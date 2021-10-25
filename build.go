//go:build build

//go:generate go build ./cmd/cli/spm.go
//go:generate go build ./cmd/controller/spm-controller.go
//go:generate go build ./cmd/wrapper/spm-wrapper.go

package build