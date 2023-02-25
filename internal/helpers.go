package internal

import (
	"github.com/bgrewell/elgato-tools/internal/types"
	"strings"
)

func contains(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func deviceType(device *types.ServiceEntry) (dt string) {
	for _, info := range device.InfoFields {
		if strings.HasPrefix(info, "md=") {
			if strings.HasPrefix(info, "md=Elgato Key Light") {
				return "key_light"
			} else {
				return "unknown"
			}
		}
	}
	return "unknown"
}
