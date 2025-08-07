package animator

import (
	"time"

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

type Frame interface {
	String() string
}

type Model struct {
	Stop         bool
	CurrentFrame int
	Rate         time.Duration
	Frames       []Frame
}

var _ tea.Model = (*Model)(nil)

func New(frames []Frame, rate time.Duration) *Model {
	return &Model{
		CurrentFrame: 0,
		Frames:       frames,
		Rate:         rate,
	}
}

type stringFrame string

func (f stringFrame) String() string {
	return string(f)
}

func StringToFrame(s string) Frame {
	return stringFrame(s)
}

func StringFrames(s []string) []Frame {
	frames := make([]Frame, len(s))
	for i, v := range s {
		frames[i] = StringToFrame(v)
	}
	return frames
}

type tickMsg struct{}

func tick(t time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(t)
		return tickMsg{}
	}
}
