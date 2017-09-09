package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpLamp struct {
	Lamp *GpioOutput
}

type HttpLampResponse struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

func NewHttpLamp(lamp *GpioOutput) *HttpLamp {
	return &HttpLamp{
		Lamp: lamp,
	}
}

func (hl *HttpLamp) Start() {
	pin := hl.Lamp.Pin
	onURL := fmt.Sprintf("/lamp/%d/on", pin)
	offURL := fmt.Sprintf("/lamp/%d/off", pin)

	http.HandleFunc(onURL, func(res http.ResponseWriter, req *http.Request) {
		hl.Lamp.On()

		json.NewEncoder(res).Encode(HttpLampResponse{
			Message: "ok",
			State:   "On",
		})
	})

	http.HandleFunc(offURL, func(res http.ResponseWriter, req *http.Request) {
		hl.Lamp.Off()

		json.NewEncoder(res).Encode(HttpLampResponse{
			Message: "ok",
			State:   "Off",
		})
	})
}
