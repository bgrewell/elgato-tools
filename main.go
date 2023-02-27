package main

import (
	"fmt"
	"github.com/bgrewell/elgato-tools/internal"
	"time"
)

func main() {

	// API Endpoints
	// /elgato/lights
	// /elgato/accessory-info
	// /elgato/settings
	// /elgato/identify (POST ONLY)

	engine := internal.ElgatoEngine{}
	engine.Initialize()
	for _, keylight := range engine.KeyLights {
		err := keylight.Update()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		keylight.SetTemperature(7000)
		time.Sleep(2 * time.Second)
		keylight.SetTemperature(2900)
		time.Sleep(2 * time.Second)
		keylight.SetBrightness(10)
		time.Sleep(2 * time.Second)
		keylight.SetBrightness(50)
		time.Sleep(2 * time.Second)
		keylight.SetBrightness(0)
		time.Sleep(2 * time.Second)
		keylight.SetBrightness(100)
		time.Sleep(2 * time.Second)
		keylight.SetTemperature(5000)
		keylight.SetOn(false)
	}
}
