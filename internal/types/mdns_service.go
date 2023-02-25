package types

import "net"

type MDNSService struct {
	Instance string   `json:"instance" yaml:"instance"`
	Service  string   `json:"service" yaml:"service"`
	Domain   string   `json:"domain" yaml:"domain"`
	HostName string   `json:"host_name" yaml:"host_name"`
	Port     int      `json:"port" yaml:"port"`
	IPs      []net.IP `json:"ip_addresses" yaml:"ip_addresses"`
	TXT      []string `json:"txt" yaml:"txt"`
}
