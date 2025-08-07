package tui

// A simple program that counts down from 5 and then exits.

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var frame = []string{
	`
╔════╤╤╤╤════╗
║    │││ \   ║
║    │││  O  ║
║    OOO     ║
	`,
	`
╔════╤╤╤╤════╗
║    ││││    ║
║    ││││    ║
║    OOOO    ║
	`,
	`
╔════╤╤╤╤════╗
║   / │││    ║
║  O  │││    ║
║     OOO    ║
	`,
	`
╔════╤╤╤╤════╗
║    ││││    ║
║    ││││    ║
║    OOOO    ║
	`,
}

// keyMap defines the key bindings for the application
type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help},
		{k.Quit},
	}
}

var defaultKeys = keyMap{
	Help: key.NewBinding(key.WithKeys("h", "?"), key.WithHelp("h", "toggle help")),
	Quit: key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
}

func NewProgram() *tea.Program {
	return tea.NewProgram(
		initModel(),
		tea.WithAltScreen(),
	)
}

func initModel() tea.Model {
	return model{
		sub:  newSubModel(),
		keys: defaultKeys,
		help: help.New(),
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	// UI dimensions
	width  int
	height int

	// Logic
	sub tea.Model

	// Controls
	keys keyMap

	// Help
	help help.Model
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	var cmds tea.BatchMsg
	cmds = append(cmds, m.sub.Init())
	return tea.Batch(cmds...)
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.BatchMsg
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.help.Width = m.width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}
	m.sub, cmd = m.sub.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		mainWindowStyle.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Width(m.width-3).
			Height(m.height-4).
			Render(
				m.rednerMainFrame(),
			),
		m.help.View(m.keys),
	)

}

func (m model) rednerMainFrame() string {
	return m.sub.View()
}
