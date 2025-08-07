package tui

import "github.com/charmbracelet/lipgloss"

const (
	padding = 1
	margin  = 1
)

var (
	mainWindowStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#808080")).
		Padding(padding)
)
