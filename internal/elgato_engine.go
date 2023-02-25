package internal

import (
	"fmt"
	"github.com/bgrewell/elgato-tools/internal/types"
	"log"
)

type ElgatoEngine struct {
	DiscoveredDevices []*types.ServiceEntry `json:"discovered_devices" yaml:"discovered_devices"`
	KeyLights         []*ElgatoKeyLight     `json:"key_lights" yaml:"key_lights"`
}

func (ee *ElgatoEngine) Initialize() {
	filter := "mf=Elgato"
	devices, err := Discover("_elg._tcp", &filter)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	ee.DiscoveredDevices = devices
	ee.KeyLights = make([]*ElgatoKeyLight, 0)
	ee.parseDevices()
}

func (ee *ElgatoEngine) parseDevices() {
	for _, device := range ee.DiscoveredDevices {
		info, err := GetInfo(device)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("%v\n", info)

		switch info.ProductName {
		case "Elgato Key Light":
			kl := ElgatoKeyLight{}
			err = kl.Parse(device, info)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
			ee.KeyLights = append(ee.KeyLights, &kl)
		default:
			fmt.Printf("unknown device type: %s\n", info.ProductName)
		}
		// 1: Call accessory-info api and get info on accessory
		// 2: Store that information
		// 3: Call API's based on features of device

		//device_type := deviceType(device)
		//switch device_type {
		//case "key_light":
		//	kl := ElgatoKeyLight{}
		//	err := kl.Parse(device)
		//	if err != nil {
		//		fmt.Printf("error: %v\n", err)
		//	}
		//	ee.KeyLights = append(ee.KeyLights, &kl)
		//default:
		//
		//}
	}
}
