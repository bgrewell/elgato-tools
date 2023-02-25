package internal

import "github.com/bgrewell/elgato-tools/internal/types"

type ElgatoDevice interface {
	Parse(entry *types.ServiceEntry) error
	Sync() error
	Update() error
}
