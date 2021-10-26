package main

import "github.com/zp4rker/jpm/internal/core"

func main() {
	controller := core.Controller{}

	if err := controller.InitSock(); err != nil {

	}
}