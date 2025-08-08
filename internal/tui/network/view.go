package network

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	bs := m.charts["Беспроводная сеть"]
	io := m.networkList["Беспроводная сеть"]
	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			bs.View(),
			lipgloss.JoinVertical(
				lipgloss.Left,
				"Беспроводная сеть",
				fmt.Sprintf("BytesRecv: %d\nBytesSent: %d", io.BytesRecv, io.BytesSent),
				fmt.Sprintf("BytesRecvRate: %d\nBytesSentRate: %d", io.BytesRecvRate, io.BytesSentRate),
				fmt.Sprintf("PacketsRecv: %d\nPacketsSent: %d", io.PacketsRecv, io.PacketsSent),
				fmt.Sprintf("PacketsRecvRate: %d\nPacketsSentRate: %d", io.PacketsRecvRate, io.PacketsSentRate),
			),
		),
		m.help.View(defaultKeys),
	)
	// return ""
}
