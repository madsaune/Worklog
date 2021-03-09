package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	w := NewWorklogClient(os.Args)
	w.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		w.Stop()
		fmt.Printf("\n\n%s", w)

		os.Exit(0)
	}()

	fmt.Println("Started tracking...")
	fmt.Println()
	fmt.Println("Press CTRL+C to stop.")
	for {
		time.Sleep(10 * time.Minute) // or runtime.Gosched() or similar per @misterbee
	}
}
