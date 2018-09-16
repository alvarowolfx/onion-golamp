package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"periph.io/x/periph/host"
)

func main() {
	_, err := host.Init() // Init periph.io
	if err != nil {
		log.Fatal(err)
	}

	led := NewLed("11")
	defer led.Close()
	led.Off() // Turn off led on start

	httpLamp := NewHttpLamp(led)
	httpLamp.Start() // Configure http handlers

	hkLamp := NewHomekitLamp(led)
	go hkLamp.Start() // Homekit server blocks execution

	server := &http.Server{Addr: ":8080"}
	go server.ListenAndServe() // Http server blocks execution

	log.Println("Http server started")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigs // Wait for signal
		log.Println(sig)

		// Shutdown all services
		hkLamp.Stop()
		log.Println("Homekit stopped")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		server.Shutdown(ctx)
		log.Println("Http server stopped")

		done <- true
	}()

	log.Println("Press ctrl+c to stop...")
	<-done // Wait
}
