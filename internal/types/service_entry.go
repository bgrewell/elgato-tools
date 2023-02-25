package types

import "net"

type ServiceEntry struct {
	Name       string   `json:"name" yaml:"name"`
	Host       string   `json:"host" yaml:"host"`
	AddrV4     net.IP   `json:"addr_v4" yaml:"addr_v4"`
	AddrV6     net.IP   `json:"addr_v6" yaml:"addr_v6"`
	Port       int      `json:"port" yaml:"port"`
	Info       string   `json:"info" yaml:"info"`
	InfoFields []string `json:"info_fields" yaml:"info_fields"`
}
