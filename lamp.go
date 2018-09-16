package main

import (
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

type Lamp interface {
	ID() string
	On()
	Off()
	Toggle()
	GetCurrentState() bool
	Close()
}

type Led struct {
	Pin     string
	gpioPin gpio.PinIO
	state   bool
}

func NewLed(pin string) *Led {
	gpioPin := gpioreg.ByName(pin)
	return &Led{
		Pin:     pin,
		gpioPin: gpioPin,
		state:   false,
	}
}

func (g *Led) ID() string {
	return g.Pin
}

func (g *Led) On() {
	g.gpioPin.Out(gpio.High)
	g.state = true
}

func (g *Led) Off() {
	g.gpioPin.Out(gpio.Low)
	g.state = false
}

func (g *Led) Toggle() {
	if g.state {
		g.Off()
	} else {
		g.On()
	}
}

func (g *Led) GetCurrentState() bool {
	return g.state
}

func (g *Led) Close() {
	g.Off()
	g.gpioPin.Halt()
}
