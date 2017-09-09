package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	led := NewGpioOutput(11)
	led.Off() // Turn off led on start

	httpLamp := NewHttpLamp(led)
	httpLamp.Start() // Configure http handlers

	hkLamp := NewHomekitLamp(led)
	go hkLamp.Start() // Homekit server blocks execution

	server := &http.Server{Addr: ":8080"}
	go func() {
		server.ListenAndServe()
	}()

	log.Println("Http server started")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs // Wait for signal

		// Shutdown all services
		server.Shutdown(nil)
		hkLamp.Stop()

		log.Println(sig)

		done <- true
	}()

	log.Println("Press ctrl+c to stop...")
	<-done // Wait
}
