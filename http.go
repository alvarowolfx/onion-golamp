package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpLamp struct {
	Lamp Lamp
}

type HttpLampResponse struct {
	Message string `json:"message"`
	State   string `json:"state"`
}

func NewHttpLamp(lamp Lamp) *HttpLamp {
	return &HttpLamp{
		Lamp: lamp,
	}
}

func (hl *HttpLamp) Start() {
	id := hl.Lamp.ID()
	onURL := fmt.Sprintf("/lamp/%s/on", id)
	offURL := fmt.Sprintf("/lamp/%s/off", id)

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
