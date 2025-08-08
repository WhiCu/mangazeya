package network

import (
	"time"

	"github.com/WhiCu/mangazeya/internal/core/network"
	"github.com/WhiCu/mangazeya/pkg/chart"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Quit      key.Binding
	StopStart key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.StopStart}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.StopStart},
	}
}

var defaultKeys = keyMap{
	Quit:      key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	StopStart: key.NewBinding(key.WithKeys("s", "ctrl+s"), key.WithHelp("s", "stop/start")),
}

type Model struct {
	charts      map[string]*chart.Chart[uint64]
	Stop        bool
	networkList network.NetworkList
	Rate        time.Duration
}

var _ tea.Model = (*Model)(nil)

func initModel(nl network.NetworkList) *Model {
	charts := make(map[string]*chart.Chart[uint64])
	charts["Беспроводная сеть"] = chart.NewChart[uint64](10, 10, 10)
	charts["Беспроводная сеть"].Add(nl["Беспроводная сеть"].BytesSentRate)
	return &Model{
		charts:      charts,
		networkList: nl,
		Stop:        false,
		Rate:        1 * time.Second,
	}
}

func New(rate time.Duration) *Model {
	return &Model{Rate: rate}
}

type tickMsg struct{}

func tick(t time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(t)
		return tickMsg{}
	}
}
