package main

import (
	"github.com/yiwenlong/ServiceUIDemo-desktop/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT)
		for sig := range signalChan {
			log.Printf("Received signal: %d (%s)", sig, sig)
			server.Stop(0)
			os.Exit(0)
		}
	}()
	server.Boot("localhost:8000")
	select {}
}
