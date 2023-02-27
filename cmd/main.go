package main

import (
	"flag"
	"fmt"
	"github.com/bgrewell/elgato-tools/internal"
	"strings"
)

func main() {

	// keylight on|off
	// keylight --[n]ame <name> --[t]emp <temp> --[b]rightness <brightness> --no-save
	// keylight --[t]emp <temp> --[b]rightness --no-save
	// Define the named flags
	namePtr := flag.String("name", "", "the name of the light")
	tempPtr := flag.Int("temp", -1, "the temperature of the light")
	brightnessPtr := flag.Int("brightness", -1, "the brightness of the light")
	noSavePtr := flag.Bool("no-save", false, "don't save the changes")

	// Define short options for the named flags
	flag.StringVar(namePtr, "n", "", "the name of the light")
	flag.IntVar(tempPtr, "t", 0, "the temperature of the light")
	flag.IntVar(brightnessPtr, "b", 0, "the brightness of the light")
	flag.BoolVar(noSavePtr, "s", false, "don't save the changes")

	// Parse the command line arguments
	flag.Parse()

	// Get the positional argument
	if flag.NArg() != 1 {
		fmt.Println("error: must specify a single action argument (on or off)")
		return
	}
	var action bool
	if flag.Arg(0) == "on" {
		action = true
	} else {
		action = false
	}

	// Print the values of the named flags and positional argument
	fmt.Println("name:", *namePtr)
	fmt.Println("temp:", *tempPtr)
	fmt.Println("brightness:", *brightnessPtr)
	fmt.Println("no-save:", *noSavePtr)
	fmt.Println("action:", action)

	engine := internal.ElgatoEngine{}
	engine.Initialize()

	for _, light := range engine.KeyLights {
		if strings.Contains(strings.ToLower(light.DisplayName), strings.ToLower(*namePtr)) || *namePtr == "" {
			if *brightnessPtr > -1 {
				light.SetBrightness(*brightnessPtr)
			}
			if *tempPtr > -1 {
				light.SetTemperature(*tempPtr)
			}
			light.SetOn(action)

		}
	}
}
