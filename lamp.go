package main

import (
	"github.com/brian-armstrong/gpio"
)

type Lamp interface {
	ID() int
	On()
	Off()
	Toggle()
	GetCurrentState() bool
	Close()
}

type Led struct {
	Pin     int
	gpioPin gpio.Pin
	state   bool
}

func NewLed(pin int) *Led {
	gpioPin := gpio.NewOutput(uint(pin), false)
	return &Led{
		Pin:     pin,
		gpioPin: gpioPin,
		state:   false,
	}
}

func (g *Led) ID() int {
	return g.Pin
}

func (g *Led) On() {
	g.gpioPin.High()
	g.state = true
}

func (g *Led) Off() {
	g.gpioPin.Low()
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
	g.gpioPin.Close()
}
