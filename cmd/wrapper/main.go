package main

import (
	"fmt"
	"github.com/zp4rker/jpm/internal/wrapper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a command string to be executed!")
		os.Exit(1)
	}

	proc, err := wrapper.NewWrapper(os.Args[1])
	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, syscall.Signal(0x1f))

	go func() {
		for range sigchan {
			fmt.Println("Recieved Signal(0x1f)")
		}
	}()

	// start go routines here

	err = proc.Run()
	if err != nil {
		panic(err)
	}
}