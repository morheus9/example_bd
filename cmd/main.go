package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := Execute(); err != nil {
		outputInitError(err)
	}
}

func outputInitError(err error) {
	fmt.Println(`{"error":"` + err.Error() + `"}`)

	os.Exit(1)
}

func waitSignalExit(cancel func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch

	cancel()
}
