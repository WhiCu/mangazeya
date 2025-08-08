package network

import (
	"time"

	"github.com/WhiCu/mangazeya/internal/core/network"
	"github.com/WhiCu/mangazeya/pkg/chart"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Up        key.Binding
	Down      key.Binding
	Quit      key.Binding
	StopStart key.Binding
	Help      key.Binding
	StopRecv  key.Binding
	StopSent  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Help,
		k.Quit,
		k.StopStart,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
		{k.StopStart, k.StopRecv, k.StopSent},
		{k.Up, k.Down},
	}
}

var defaultKeys = keyMap{
	Up:        key.NewBinding(key.WithKeys("k", "w", "up", "right"), key.WithHelp("↑/➝/k/w", "move up")),
	Down:      key.NewBinding(key.WithKeys("j", "s", "down", "left"), key.WithHelp("↓/🠔/j/s", "move down")),
	Quit:      key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	StopStart: key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "stop/start")),
	Help:      key.NewBinding(key.WithKeys("h", "?"), key.WithHelp("h", "toggle help")),
	StopRecv:  key.NewBinding(key.WithKeys("r", "ctrl+r"), key.WithHelp("r", "stop recv")),
	StopSent:  key.NewBinding(key.WithKeys("s", "ctrl+s"), key.WithHelp("s", "stop sent")),
}

type Model struct {
	charts      map[string]*chart.Chart[uint64]
	Stop        bool
	StopRecv    bool
	StopSent    bool
	networkList network.NetworkList
	Rate        time.Duration
	help        help.Model
}

var _ tea.Model = (*Model)(nil)

func initModel(nl network.NetworkList) *Model {
	charts := make(map[string]*chart.Chart[uint64])
	for k := range nl {
		charts[k] = chart.NewChart[uint64](8, 32, 8)
		charts[k].AddLegend("BytesRecvRate", "BytesSentRate")
	}
	return &Model{
		charts:      charts,
		networkList: nl,
		Stop:        false,
		Rate:        1 * time.Second,
		help:        help.New(),
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
