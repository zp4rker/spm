//go:build build

//go:generate go build -o jpm ./cmd/cli/main.go
//go:generate go build -o jpm-controller ./cmd/controller/main.go
//go:generate go build -o jpm-wrapper ./cmd/wrapper/main.go

package build