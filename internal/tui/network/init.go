package network

import (
	"github.com/WhiCu/mangazeya/internal/core/network"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	return tick(m.Rate)
}

func NewProgram(nl network.NetworkList) *tea.Program {
	return tea.NewProgram(
		initModel(nl),
	)
}
