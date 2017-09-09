package main

import (
	"log"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

type HomekitLamp struct {
	Lamp      *GpioOutput
	light     *accessory.Lightbulb
	transport hc.Transport
	ticker    *time.Ticker
}

func NewHomekitLamp(lamp *GpioOutput) *HomekitLamp {
	return &HomekitLamp{
		Lamp: lamp,
	}
}

func (hkl *HomekitLamp) Start() {
	hkl.light = accessory.NewLightbulb(accessory.Info{
		Name: "GoLamp",
	})

	hkl.light.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("Client changed switch to on")
			hkl.Lamp.On()
		} else {
			log.Println("Client changed switch to off")
			hkl.Lamp.Off()
		}
	})

	// Monitor gpio state
	go func() {
		lastState := hkl.Lamp.state
		hkl.ticker = time.NewTicker(300 * time.Millisecond)
		for _ = range hkl.ticker.C {
			newState := hkl.Lamp.state
			if lastState != newState {
				hkl.light.Lightbulb.On.SetValue(newState)
				lastState = newState
			}
		}
	}()

	config := hc.Config{Pin: "00102003"}
	var err error
	hkl.transport, err = hc.NewIPTransport(config, hkl.light.Accessory)
	if err != nil {
		log.Panic(err)
	}

	hkl.transport.Start()
}

func (hkl *HomekitLamp) Stop() {
	hkl.transport.Stop()
	hkl.ticker.Stop()
}
