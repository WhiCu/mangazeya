package animator

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit},
	}
}

var defaultKeys = keyMap{
	Quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultKeys.Quit):
			return m, tea.Quit
		}
	case tea.QuitMsg:
		return nil, tea.Quit
	case tickMsg:
		m.CurrentFrame = (m.CurrentFrame + 1) % len(m.Frames)
		return m, tick(m.Rate)
	}
	return m, nil
}
