package tui

import (
	"time"

	"github.com/WhiCu/mangazeya/internal/tui/animator"
	tea "github.com/charmbracelet/bubbletea"
)

type subModel struct {
	animatator tea.Model
}

func newSubModel() subModel {
	return subModel{animatator: animator.New(animator.StringFrames(frame), time.Second/10)}
}

func (m subModel) Init() tea.Cmd {
	var cmds tea.BatchMsg
	cmds = append(cmds, m.animatator.Init())
	return tea.Batch(cmds...)
}

func (m subModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds tea.BatchMsg
	var cmd tea.Cmd
	m.animatator, cmd = m.animatator.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m subModel) View() string {
	return m.animatator.View()
}
