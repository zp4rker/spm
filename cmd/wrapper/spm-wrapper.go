package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/zp4rker/spm/internal/spm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a command string to be executed!")
		os.Exit(1)
	}

	fmt.Println("Starting process wrapper...")

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

	conn, err := net.Dial("unix", spm.SockAddr)
	if err != nil {
		return
	}
	defer conn.Write([]byte("TERMINATED\n"))

	_, _ = conn.Write([]byte(fmt.Sprintf("/register %v\n", os.Getpid())))

	// start go routines here
	go wrapper.StartHeartbeat(conn)

	if err = wrapper.Wait(); err != nil {
		fmt.Printf("Exited with error %v\n", err)
	}

	fmt.Println("Exiting process wrapper")
}
