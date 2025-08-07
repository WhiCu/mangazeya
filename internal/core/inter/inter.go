package inter

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

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
	MTU          int      `json:"mtu,omitempty"`          // maximum transmission unit
	HardwareAddr string   `json:"hardwareAddr,omitempty"` // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string `json:"flags,omitempty"`        // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []Addr   `json:"addrs,omitempty"`
}

func (i Interface) String() string {
	b := strings.Builder{}

	if i.MTU != 0 {
		fmt.Fprintf(&b, "MTU: %d\n", i.MTU)
	}

	if i.HardwareAddr != "" {
		fmt.Fprintf(&b, "HardwareAddr: %s\n", i.HardwareAddr)
	}

	if len(i.Flags) != 0 {
		fmt.Fprintf(&b, "Flags: %s\n", strings.Join(i.Flags, " "))
	}

	if len(i.Addrs) != 0 {
		fmt.Fprintf(&b, "Addrs: ")
		for _, a := range i.Addrs {
			fmt.Fprintf(&b, "%s ", a)
		}
		fmt.Fprintf(&b, "\n")
	}

	return b.String()
}

func (i Interface) JSON() ([]byte, error) {
	return json.Marshal(i)
}

func (i Interface) CoolJSON() ([]byte, error) {
	return json.MarshalIndent(i, "", "  ")
}

type InterfaceList map[string]Interface

func (l InterfaceList) List() []string {
	list := make([]string, 0, len(l))
	for key := range l {
		list = append(list, key)
	}
	return list
}

func (l InterfaceList) Interface(name string) (Interface, bool) {
	i, ok := l[name]
	return i, ok
}

func (l InterfaceList) Count() int {
	return len(l)
}

func (l InterfaceList) String() string {
	b := strings.Builder{}

	for key, value := range l {
		b.WriteString(fmt.Sprintf("%s:\n%s", key, value))
	}

	return b.String()
}

func (l InterfaceList) JSON() ([]byte, error) {
	return json.Marshal(l)
}

func (l InterfaceList) CoolJSON() ([]byte, error) {

	return json.MarshalIndent(l, "", "  ")
}

func Interfaces() (InterfaceList, error) {
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
