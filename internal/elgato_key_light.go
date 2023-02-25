package internal

import (
	"fmt"
	"github.com/bgrewell/elgato-tools/internal/types"
	"math"
)

type ElgatoKeyLight struct {
	types.ServiceEntry
	types.AccessoryInfo
	Lights *types.Lights
}

func (kl *ElgatoKeyLight) Parse(entry *types.ServiceEntry, info *types.AccessoryInfo) error {
	kl.ServiceEntry = *entry
	kl.AccessoryInfo = *info
	return kl.Sync()
}

func (kl *ElgatoKeyLight) Sync() error {
	lights, err := GetLights(&kl.ServiceEntry)
	if err != nil {
		return err
	}

	// Update lights
	kl.Lights = lights

	return nil
}

func (kl *ElgatoKeyLight) Update() error {
	lights, err := PutLights(&kl.ServiceEntry, kl.Lights)
	if err != nil {
		return err
	}

	// Update lights
	kl.Lights = lights

	return nil
}

func (kl *ElgatoKeyLight) SetOn(on bool) (err error) {
	for _, light := range kl.Lights.Lights {
		if on {
			light.On = 1
		} else {
			light.On = 0
		}
	}
	return kl.Update()
}

func (kl *ElgatoKeyLight) SetBrightness(percent int) (err error) {
	if percent < 0 || percent > 100 {
		return fmt.Errorf("brightness must be a number between 0 and 100")
	}
	for _, light := range kl.Lights.Lights {
		light.Brightness = percent
		if percent == 0 {
			light.On = 0
		} else {
			light.On = 1
		}
	}
	return kl.Update()
}

func (kl *ElgatoKeyLight) SetTemperature(kelvin int) (err error) {
	if kelvin < 2900 || kelvin > 7000 {
		return fmt.Errorf("temperature must be between 2900 and 7000")
	}
	v := float64(1000000) / float64(kelvin)
	v = math.Round(v/50.0) * 50.0
	for _, light := range kl.Lights.Lights {
		light.Temperature = int(v)
	}
	return kl.Update()
}
