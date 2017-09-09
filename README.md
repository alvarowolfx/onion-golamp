# Onion Omega 2 + Golang + Homekit Demo

Demo project on how to run a Golang program on an embbeded hardware like Onion Omega 2. In this case this project simulates a Homekit device for iOS, making it possible to control it via Home App. Also there is http endpoints to control the LED, showing many ways that we can interface with Onion via Golan.

## Schematic 

Attach a LED to GPIO 11 on Omega 2 to see this project working.

## How to build for Omega 

This command will generate a binary file compatible with Omega architecture.

`make build`

## Copy to Omega

Change your omega address on the Makefile, then run the command: 

`make copy`