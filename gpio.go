package main

import (
	"os/exec"
	"strconv"
)

type GpioOutput struct {
	Pin    int
	state  bool
	pinStr string
}

func NewGpioOutput(pin int) *GpioOutput {
	pinStr := strconv.Itoa(pin)
	exec.Command("gpioctl", "dirout", pinStr).Run()
	return &GpioOutput{
		Pin:    pin,
		state:  false,
		pinStr: pinStr,
	}
}

func (g *GpioOutput) On() {
	exec.Command("gpioctl", "dirout-high", g.pinStr).Run()
	g.state = true
}

func (g *GpioOutput) Off() {
	exec.Command("gpioctl", "dirout-low", g.pinStr).Run()
	g.state = false
}

func (g *GpioOutput) Toggle() {
	if g.state {
		g.Off()
	} else {
		g.On()
	}
}
