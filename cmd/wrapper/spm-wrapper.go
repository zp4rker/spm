package main

import (
	"fmt"
	"github.com/zp4rker/jpm/internal/spm"
	"net"
	"os"
	"os/signal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a command string to be executed!")
		os.Exit(1)
	}

	wrapper, err := spm.NewWrapper(os.Args[1])
	if err != nil {
		panic(err)
	}

	if err = wrapper.Start(); err != nil {
		panic(err)
	}

	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, spm.SIGRQFB)

		for range sigchan {
			fmt.Println("Recieved SIGRQFB")
		}
	}()

	go func() {
		conn, err := net.Dial("unix", spm.SockAddr)
		if err != nil {
			return
		}

		_, _ = conn.Write([]byte(fmt.Sprintf("/register %v\n", os.Getpid())))
	}()

	// start go routines here

	if err = wrapper.Wait(); err != nil {
		panic(err)
	}
}