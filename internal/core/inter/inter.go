package inter

import (
	"regexp"

	"github.com/shirou/gopsutil/v4/net"
)

var (
	regexpIPv6 = regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}(%[0-9a-zA-Z]{1,})?|::(ffff(:0{1,4}){0,1}:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))(/(12[0-8]|1[01][0-9]|[0-9]{1,2}))?$`)
	regexpIPv4 = regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})\.){3}(25[0-5]|2[0-4][0-9]|1?[0-9]{1,2})(/([1-3]?[0-9]))?$`)
)

type Addr struct {
	Type string `json:"type"`
	IP   string `json:"ip"`
}

type Interface struct {
	MTU          int      `json:"mtu"`          // maximum transmission unit
	HardwareAddr string   `json:"hardwareAddr"` // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string `json:"flags"`        // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []Addr   `json:"addrs"`
}

func Interfaces() (map[string]Interface, error) {
	stat, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	interfaces := make(map[string]Interface)
	for _, i := range stat {
		interfaces[i.Name] = Interface{
			MTU:          i.MTU,
			HardwareAddr: i.HardwareAddr,
			Flags:        i.Flags,
			Addrs:        parseAddrs(i.Addrs),
		}
	}

	return interfaces, nil
}

func parseAddrs(interfaceAddrList net.InterfaceAddrList) []Addr {
	addrs := make([]Addr, len(interfaceAddrList))
	for i := range interfaceAddrList {
		addrs[i] = Addr{
			Type: typeAddr(interfaceAddrList[i].Addr),
			IP:   interfaceAddrList[i].Addr,
		}
	}
	return addrs
}

func typeAddr(s string) string {
	switch {
	case regexpIPv6.MatchString(s):
		return "ipv6"
	case regexpIPv4.MatchString(s):
		return "ipv4"
	default:
		return "unknown"
	}
}
