package main

import (
	"fmt"

	"github.com/zp4rker/jpm/internal/spm"
)

func main() {
	controller := spm.Controller{}

	if err := controller.InitSock(); err != nil {
		panic(err)
	}

	fmt.Println("Starting controller...")
	defer controller.CloseSock()
	if err := controller.Start(); err != nil {
		fmt.Printf("Exited with error %v\n", err)
	}
}
