package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zp4rker/spm/internal/spm"
)

func main() {
	controller := spm.Controller{}

	if err := controller.InitSock(); err != nil {
		panic(err)
	}
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT)

		for range sigchan {
			controller.CloseSock()
			fmt.Println("Exiting controller...")
		}
	}()

	fmt.Println("Starting controller...")
	if err := controller.Start(); err != nil {
		fmt.Printf("Exited with error %v\n", err)
	}
}
