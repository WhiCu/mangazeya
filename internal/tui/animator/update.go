package animator

import (
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
		}
	case tea.QuitMsg:
		return nil, tea.Quit
	case tickMsg:
		if m.Stop {
			break
		}
		m.CurrentFrame = (m.CurrentFrame + 1) % len(m.Frames)
		return m, tick(m.Rate)
	}
	return m, nil
}
