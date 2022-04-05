package main

import (
	"github.com/pythrick/dateflix/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort = "8080"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		r := api.NewHTTPServer()

		err := http.ListenAndServe(":"+defaultPort, r)
		if err != nil {
			done <- syscall.SIGTERM
		}
	}()
	<-done
}
