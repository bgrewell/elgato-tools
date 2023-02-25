package internal

import (
	"fmt"
	"github.com/bgrewell/elgato-tools/internal/types"
	"github.com/hashicorp/mdns"
)

func ConvertServiceEntry(in *mdns.ServiceEntry) (out *types.ServiceEntry) {
	out = &types.ServiceEntry{
		Name:       in.Name,
		Host:       in.Host,
		AddrV4:     in.AddrV4,
		AddrV6:     in.AddrV6,
		Port:       in.Port,
		Info:       in.Info,
		InfoFields: in.InfoFields,
	}
	return out
}

func Discover(service string, filter *string) (entries []*types.ServiceEntry, err error) {
	entriesCh := make(chan *mdns.ServiceEntry, 50)
	entries = make([]*types.ServiceEntry, 0)
	go func() {
		for entry := range entriesCh {
			if filter != nil && contains(entry.InfoFields, *filter) {
				fmt.Printf("Got new entry: %v\n", entry)
				entries = append(entries, ConvertServiceEntry(entry))
			}
		}
	}()

	err = mdns.Lookup(service, entriesCh)
	if err != nil {
		return nil, err
	}
	close(entriesCh)
	return entries, nil
}
