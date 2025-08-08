package network

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/shirou/gopsutil/v4/net"
)

type Network struct {
	BytesSent       uint64 `json:"bytesSent,omitempty"`
	BytesSentRate   uint64 `json:"bytesSentRate,omitempty"`
	BytesRecv       uint64 `json:"bytesRecv,omitempty"`
	BytesRecvRate   uint64 `json:"bytesRecvRate,omitempty"`
	PacketsSent     uint64 `json:"packetsSent,omitempty"`
	PacketsSentRate uint64 `json:"packetsSentRate,omitempty"`
	PacketsRecv     uint64 `json:"packetsRecv,omitempty"`
	PacketsRecvRate uint64 `json:"packetsRecvRate,omitempty"`
	Errin           uint64 `json:"errin,omitempty"`
	Errout          uint64 `json:"errout,omitempty"`
	Dropin          uint64 `json:"dropin,omitempty"`
	Dropout         uint64 `json:"dropout,omitempty"`
	Fifoin          uint64 `json:"fifoin,omitempty"`
	Fifoout         uint64 `json:"fifoout,omitempty"`
}

func (n Network) String() string {
	b := strings.Builder{}

	fmt.Fprintf(&b, "BytesSent: %d\n", n.BytesSent)
	fmt.Fprintf(&b, "BytesSentRate: %d\n", n.BytesSentRate)
	fmt.Fprintf(&b, "BytesRecv: %d\n", n.BytesRecv)
	fmt.Fprintf(&b, "BytesRecvRate: %d\n", n.BytesRecvRate)
	fmt.Fprintf(&b, "PacketsSent: %d\n", n.PacketsSent)
	fmt.Fprintf(&b, "PacketsSentRate: %d\n", n.PacketsSentRate)
	fmt.Fprintf(&b, "PacketsRecv: %d\n", n.PacketsRecv)
	fmt.Fprintf(&b, "PacketsRecvRate: %d\n", n.PacketsRecvRate)
	fmt.Fprintf(&b, "Errin: %d\n", n.Errin)
	fmt.Fprintf(&b, "Errout: %d\n", n.Errout)
	fmt.Fprintf(&b, "Dropin: %d\n", n.Dropin)
	fmt.Fprintf(&b, "Dropout: %d\n", n.Dropout)
	fmt.Fprintf(&b, "Fifoin: %d\n", n.Fifoin)
	fmt.Fprintf(&b, "Fifoout: %d\n", n.Fifoout)

	return b.String()
}

func (n *Network) JSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Network) CoolJSON() ([]byte, error) {
	return json.MarshalIndent(n, "", "  ")
}

type NetworkList map[string]Network

func Networks() (NetworkList, error) {
	stat, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}

	networks := make(map[string]Network, len(stat))
	for _, value := range stat {
		networks[value.Name] = Network{
			BytesSent:   value.BytesSent,
			BytesRecv:   value.BytesRecv,
			PacketsSent: value.PacketsSent,
			PacketsRecv: value.PacketsRecv,
			Errin:       value.Errin,
			Errout:      value.Errout,
			Dropin:      value.Dropin,
			Dropout:     value.Dropout,
			Fifoin:      value.Fifoin,
			Fifoout:     value.Fifoout,
		}
	}

	return networks, nil
}

func (nl NetworkList) Reboot() error {
	stat, err := net.IOCounters(true)
	if err != nil {
		return err
	}

	for _, value := range stat {
		prev := nl[value.Name]
		nl[value.Name] = Network{
			BytesSent:       value.BytesSent,
			BytesSentRate:   value.BytesSent - prev.BytesSent,
			BytesRecv:       value.BytesRecv,
			BytesRecvRate:   value.BytesRecv - prev.BytesRecv,
			PacketsSent:     value.PacketsSent,
			PacketsSentRate: value.PacketsSent - prev.PacketsSent,
			PacketsRecv:     value.PacketsRecv,
			PacketsRecvRate: value.PacketsRecv - prev.PacketsRecv,
			Errin:           value.Errin,
			Errout:          value.Errout,
			Dropin:          value.Dropin,
			Dropout:         value.Dropout,
			Fifoin:          value.Fifoin,
			Fifoout:         value.Fifoout,
		}
	}

	return nil
}

func (nl NetworkList) Network(name string) (Network, error) {
	network, ok := nl[name]
	if !ok {
		return Network{}, fmt.Errorf("network %s not found", name)
	}

	return network, nil
}

func (nl NetworkList) String() string {
	b := strings.Builder{}

	for key, value := range nl {
		fmt.Fprintf(&b, "%s:\n", key)
		WithTab(&b, &value)
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

func WithTab(w io.Writer, network *Network) {
	fmt.Fprintf(w, "\tBytesSent: %v\n", network.BytesSent)
	fmt.Fprintf(w, "\tBytesSentRate: %v\n", network.BytesSentRate)
	fmt.Fprintf(w, "\tBytesRecv: %v\n", network.BytesRecv)
	fmt.Fprintf(w, "\tPacketsSent: %v\n", network.PacketsSent)
	fmt.Fprintf(w, "\tPacketsRecv: %v\n", network.PacketsRecv)
	fmt.Fprintf(w, "\tErrin: %v\n", network.Errin)
	fmt.Fprintf(w, "\tErrout: %v\n", network.Errout)
	fmt.Fprintf(w, "\tDropin: %v\n", network.Dropin)
	fmt.Fprintf(w, "\tDropout: %v\n", network.Dropout)
	fmt.Fprintf(w, "\tFifoin: %v\n", network.Fifoin)
	fmt.Fprintf(w, "\tFifoout: %v\n", network.Fifoout)
}

func (nl NetworkList) JSON() ([]byte, error) {
	return json.Marshal(nl)
}

func (nl NetworkList) CoolJSON() ([]byte, error) {
	return json.MarshalIndent(nl, "", "  ")
}
