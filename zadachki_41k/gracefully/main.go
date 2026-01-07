package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	
	signal := <-stop

	time.Sleep(2 * time.Second)

	fmt.Println("gracefully stopped", signal.String())
}