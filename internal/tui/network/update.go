package network

import (
	"github.com/WhiCu/mangazeya/internal/core/network"
	"github.com/WhiCu/mangazeya/pkg/chart"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultKeys.Quit):
			return m, tea.Quit
		case key.Matches(msg, defaultKeys.StopStart):
			m.Stop = !m.Stop
			if !m.Stop {
				return m, tick(m.Rate)
			}
		case key.Matches(msg, defaultKeys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, defaultKeys.StopRecv):
			// fmt.Println("StopRecv")
			m.StopRecv = !m.StopRecv
		case key.Matches(msg, defaultKeys.StopSent):
			// fmt.Println("StopSent")
			m.StopSent = !m.StopSent
		}
	case tea.QuitMsg:
		return nil, tea.Quit
	case tickMsg:
		m.networkList.Reboot()
		if m.Stop {
			break
		}
		m.updateCharts(m.networkList, m.charts)
		return m, tick(m.Rate)
	}
	return m, nil
}

func (m *Model) updateCharts(networkList network.NetworkList, chart map[string]*chart.Chart[uint64]) {
	for k, v := range networkList {
		m.updateChart(v, chart[k])
	}
}

func (m *Model) updateChart(network network.Network, chart *chart.Chart[uint64]) {
	switch {
	case m.StopRecv && m.StopSent:
		// fmt.Println("StopRecv && StopSent")
		chart.Add(0, 0)
	case m.StopSent:
		// fmt.Println("StopRecv")
		chart.Add(network.BytesRecvRate, 0)
	case m.StopRecv:
		// fmt.Println("StopSent")
		chart.Add(0, network.BytesSentRate)
	default:
		// fmt.Println("Default")
		chart.Add(network.BytesRecvRate, network.BytesSentRate)
	}
}
